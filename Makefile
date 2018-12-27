TAG_NAME=$(shell git describe --abbrev=0 --tags)
APP_NAME=nie-crawler
MAC_EXE=$(APP_NAME)-darwin-amd64-$(TAG_NAME).app
LIN_EXE=$(APP_NAME)-linux-amd64-$(TAG_NAME)
WIN_EXE=$(APP_NAME)-windows-amd64-$(TAG_NAME).exe

get:
	go get -u github.com/gocolly/colly/...

run:
	go run main.go

init:
	mkdir -p dist

build-mac:  init
	GOOS=darwin GOARCH=amd64 go build -o dist/$(MAC_EXE)
	chmod +x dist/$(MAC_EXE)
	tar zcvf dist/$(MAC_EXE).tar.gz -C dist $(MAC_EXE)

build-lin:  init
	GOOS=linux GOARCH=amd64 go build -o dist/$(LIN_EXE)
	chmod +x dist/$(LIN_EXE)
	tar zcvf dist/$(LIN_EXE).tar.gz -C dist $(LIN_EXE)

build-win:  init
	GOOS=windows GOARCH=amd64 go build -o dist/$(WIN_EXE)
	chmod +x dist/$(WIN_EXE)
	zip -9 -D -r dist/$(WIN_EXE).zip dist/$(WIN_EXE)

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
