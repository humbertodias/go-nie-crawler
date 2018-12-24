package model

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

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

func NewOficina(host string, elem *colly.HTMLElement) Oficina {
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
		URL:  host + url,
	}

}
