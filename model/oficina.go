package model

import (
	"github.com/gocolly/colly"
	"strings"
)

type Oficina struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (o Oficina) Valid() bool {
	return !strings.Contains(o.ID, "-1")
}

func NewOficina(host string, elem *colly.HTMLElement) Oficina {
	url, _ := elem.DOM.Attr("value")
	ID := ProvinciaExtractId(url)
	name := elem.Text
	return Oficina{
		ID:   ID,
		Name: name,
		URL:  host + url,
	}
}
