get:
	go get -u github.com/gocolly/colly/...

run:
	go run main.go

build:
	go build

test:
	go test

lint:
	golint

format:
	go fmt

fix:
	go fix

clean:
	rm -rf *.json
