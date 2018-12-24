package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gocolly/colly"
)

const HOST = "https://sede.administracionespublicas.gob.es"
const HOST_START_SCRAPPING = HOST + "/icpplus/"

type Oficina struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (o Oficina) String() string {
	return fmt.Sprintf("%d: %s", o.ID, o.Name)
}

func (o Oficina) Valid() bool {
	return strings.Contains(o.URL, "?p=") && !strings.Contains(o.ID, "-1")
}

func NewOficina(elem *colly.HTMLElement) Oficina {
	url, _ := elem.DOM.Attr("value")
	parts := strings.Split(url, "=")
	var ID = "-1"
	if len(parts) > 1 {
		ID = parts[1]
	}
	name := elem.Text

	return Oficina{
		ID:   ID,
		Name: name,
		URL:  HOST + url,
	}

}

type Tramite struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (o Tramite) String() string {
	return fmt.Sprintf("%s %s %s", o.ID, o.Name, o.URL)
}

func (o Tramite) Valid() bool {
	return !strings.Contains(o.ID, "-1")
}

func NewTramite(url string, elem *colly.HTMLElement) Tramite {
	ID, _ := elem.DOM.Attr("value")
	name := elem.Text

	return Tramite{
		ID:   ID,
		URL:  url,
		Name: name,
	}
}

func scrapyOficinas() []Oficina {
	var oficinas []Oficina

	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("#form", func(e *colly.HTMLElement) {

		e.ForEach("option", func(index int, elem *colly.HTMLElement) {

			oficina := NewOficina(elem)
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
