compile: go.mod go.sum lebron.go
	go build -ldflags -H=windowsgui lebron.go 
setup:
	go mod init go-sqlite3
	go get github.com/mattn/go-sqlite3
	go get github.com/billgraziano/dpapi
