VERSION       := 0.0.1
REVISION      := `git rev-parse --short HEAD`
FLAG := -ldflags='-X main.version=$(VERSION) -X main.revision='$(REVISION)' -s -w -extldflags="-static" -buildid=' -a -tags netgo -installsuffix -trimpath 

all:
	cat ./makefile
build:
	go build
release:
	go build $(FLAG)
	make upx 
	@echo Success!
upx:
	upx --lzma ./*.exe

