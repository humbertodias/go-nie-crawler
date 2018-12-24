get:
	go get -u github.com/gocolly/colly/...

run:
	go run main.go

lint:
	golint

format:
	go fmt

fix:
	go fix

clean:
	rm -rf oficinas.json tramites.json 