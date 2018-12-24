package nie

import (
	"fmt"
	"github.com/gocolly/colly"
	. "github.com/humbertodias/go-crawler-demo/model"
)

const HOST = "https://sede.administracionespublicas.gob.es"
const HOST_START_SCRAPPING = HOST + "/icpplus/"

func ScrapyProvincias() []Provincia {
	var provincias []Provincia

	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("#form", func(e *colly.HTMLElement) {
		e.ForEach("option", func(index int, elem *colly.HTMLElement) {

			oficina := NewProvincia(HOST, elem)
			if oficina.Valid() {
				c.Visit(oficina.URL)
				provincias = append(provincias, oficina)
			}

		})

	})

	// Start scraping
	c.Visit(HOST_START_SCRAPPING)

	return provincias
}

func ScrapyTramites(provincias []Provincia) []Tramite {
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

	for _, oficina := range provincias {
		c.Visit(oficina.URL)
	}

	return tramites
}

// https://sede.administracionespublicas.gob.es/icpplus/acCitar
func ScrapyOficinas(tramites []Tramite) []Oficina {
	var oficinas []Oficina
	c := colly.NewCollector()

	//	c.OnHTML("#idSede", func(e *colly.HTMLElement) {
	c.OnHTML("html", func(e *colly.HTMLElement) {

		url := e.Request.URL.String()
		e.ForEach("option", func(_ int, elem *colly.HTMLElement) {
			oficina := NewOficina(url, elem)
			if oficina.Valid() {
				oficinas = append(oficinas, oficina)
			}
		})

	})

	//
	//	URL := "https://sede.administracionespublicas.gob.es/icpplus/acInfo"
	URL := "https://sede.administracionespublicas.gob.es/icpplus/acCitar"

	for _, tramite := range tramites {
		params := map[string]string{
			"tramite": tramite.ID,
		}
		err := c.Post(URL, params)
		if err != nil {
			fmt.Println(err)
		}
		c.Wait()
	}

	return oficinas
}
