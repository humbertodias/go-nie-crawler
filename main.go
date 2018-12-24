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
	provincias := nie.ScrapyProvincias()
	writeJSON(provincias, "provincias.json")

	tramites := nie.ScrapyTramites(provincias)
	writeJSON(tramites, "tramites.json")

	oficinas := nie.ScrapyOficinas(tramites)
	writeJSON(oficinas, "oficinas.json")
}
