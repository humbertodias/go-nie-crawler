package model

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

type Provincia struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (o Provincia) String() string {
	return fmt.Sprintf("%d: %s", o.ID, o.Name)
}

func (o Provincia) Valid() bool {
	return strings.Contains(o.URL, "?p=") && !strings.Contains(o.ID, "-1")
}

func NewProvincia(host string, elem *colly.HTMLElement) Provincia {
	url, _ := elem.DOM.Attr("value")
	ID := ProvinciaExtractId(url)
	name := elem.Text
	return Provincia{
		ID:   ID,
		Name: name,
		URL:  host + url,
	}

}

func ProvinciaExtractId(url string) string {
	var id = "-1"
	parts := strings.Split(url, "=")
	if len(parts) > 1 {
		id = parts[1]
	}
	return id
}
