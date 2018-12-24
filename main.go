package main

import (
	"encoding/json"
	"io/ioutil"

	. "./model"
	"github.com/gocolly/colly"
)

const HOST = "https://sede.administracionespublicas.gob.es"
const HOST_START_SCRAPPING = HOST + "/icpplus/"

func scrapyOficinas() []Oficina {
	var oficinas []Oficina

	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("#form", func(e *colly.HTMLElement) {
		e.ForEach("option", func(index int, elem *colly.HTMLElement) {

			oficina := NewOficina(HOST, elem)
			if oficina.Valid() {
				c.Visit(oficina.URL)
				oficinas = append(oficinas, oficina)
			}

		})

	})

	// Start scraping
	c.Visit(HOST_START_SCRAPPING)

	return oficinas
}

func scrapyTramites(oficinas []Oficina) []Tramite {
	var tramites []Tramite
	c := colly.NewCollector()

	c.OnHTML("#tramite", func(e *colly.HTMLElement) {

		url := e.Request.URL.String()
		e.ForEach("option", func(_ int, elem *colly.HTMLElement) {
			tramite := NewTramite(url, elem)
			if tramite.Valid() {
				tramites = append(tramites, tramite)
			}
		})

	})

	for _, oficina := range oficinas {
		c.Visit(oficina.URL)
	}

	return tramites
}

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
