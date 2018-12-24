package model

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

type Tramite struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	OficinaID string `json:"oficinaId"`
}

func (o Tramite) String() string {
	return fmt.Sprintf("%s %s %s %s", o.ID, o.Name, o.URL, o.OficinaID)
}

func (o Tramite) Valid() bool {
	return !strings.Contains(o.ID, "-1")
}

func NewTramite(url string, elem *colly.HTMLElement) Tramite {
	ID, _ := elem.DOM.Attr("value")
	name := elem.Text
	oficinaId := OficinaExtractId(url)

	return Tramite{
		ID:        ID,
		URL:       url,
		Name:      name,
		OficinaID: oficinaId,
	}
}
