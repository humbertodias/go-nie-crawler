package main

import (
	"encoding/json"
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
	oficinas := scrapyOficinas()
	writeJSON(oficinas, "oficinas.json")

	tramites := scrapyTramites(oficinas)
	writeJSON(tramites, "tramites.json")
}
