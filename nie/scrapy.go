package nie

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	. "github.com/humbertodias/go-nie-crawler/model"
)

const HOST = "https://sede.administracionespublicas.gob.es"
const HOST_START_SCRAPPING = HOST + "/icpplus/"

func ScrapyProvincias(ids ...string) []Provincia {
	var provincias []Provincia

	// Instantiate default collector
	c := NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("#form", func(e *colly.HTMLElement) {
		e.ForEach("option", func(index int, elem *colly.HTMLElement) {

			oficina := NewProvincia(HOST, elem)
			if oficina.Valid() && CanAddID(ids, oficina.ID) {
				c.Visit(oficina.URL)
				provincias = append(provincias, oficina)
			}

		})

	})

	// Start scraping
	c.Visit(HOST_START_SCRAPPING)

	return provincias
}
func CanAddID(ids []string, id string) bool {
	return len(ids) == 0 || Contains(ids, id)
}

func Contains(ar []string, x string) bool {
	for _, n := range ar {
		if x == n {
			return true
		}
	}
	return false
}

func ScrapyTramites(provincias []Provincia) []Tramite {
	var tramites []Tramite
	c := NewCollector()

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
	c := NewCollector()

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

func NewCollector() *colly.Collector {
	cd := colly.Debugger(&debug.LogDebugger{})
	c := colly.NewCollector(cd)

	c.WithTransport(&http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Proxy:           http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 5 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	return c
}
