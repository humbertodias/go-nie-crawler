get:
	go get -u github.com/gocolly/colly/...

run:
	go run main.go

init:
	mkdir dist

build-mac:  init
	GOOS=darwin GOARCH=amd64 go build -o dist/nie-crawler-mac

build-lin:  init
	GOOS=linux GOARCH=amd64 go build -o dist/nie-crawler-lin

build-win:  init
	GOOS=windows GOARCH=amd64 go build  -o dist/nie-crawler-win

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
