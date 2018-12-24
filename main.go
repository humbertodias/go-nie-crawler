package main

import (
	"encoding/json"
	"github.com/humbertodias/go-crawler-demo/nie"
	"io/ioutil"
)

func writeJSON(arr interface{}, filename string) {
	jsonMarshal, err := json.Marshal(arr)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, jsonMarshal, 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	oficinas := nie.ScrapyOficinas()
	writeJSON(oficinas, "oficinas.json")

	tramites := nie.ScrapyTramites(oficinas)
	writeJSON(tramites, "tramites.json")
}
