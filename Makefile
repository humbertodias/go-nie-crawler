TAG_NAME=$(shell git describe --abbrev=0 --tags)
APP_NAME=nie-crawler
MAC_DIST=dist/$(APP_NAME)-darwin-amd64-$(TAG_NAME).app
LIN_DIST=dist/$(APP_NAME)-linux-amd64-$(TAG_NAME)
WIN_DIST=dist/$(APP_NAME)-windows-amd64-$(TAG_NAME).exe

get:
	go get -u github.com/gocolly/colly/...

run:
	go run main.go

init:
	mkdir -p dist

build-mac:  init
	GOOS=darwin GOARCH=amd64 go build -o $(MAC_DIST)
	chmod +x $(MAC_DIST)
	tar zcvf $(MAC_DIST).tar.gz $(MAC_DIST)

build-lin:  init
	GOOS=linux GOARCH=amd64 go build -o $(LIN_DIST)
	chmod +x $(LIN_DIST)
	tar zcvf $(LIN_DIST).tar.gz $(LIN_DIST)

build-win:  init
	GOOS=windows GOARCH=amd64 go build -o $(WIN_DIST)
	chmod +x $(WIN_DIST)
	zip -9 -r $(WIN_DIST).zip $(WIN_DIST)

build: build-mac    build-lin   build-win

test:
	go test -v

lint:
	golint

format:
	go fmt

fix:
	go fix

clean-dist:
	rm -rf dist

clean:  clean-dist
	rm -rf *.json
