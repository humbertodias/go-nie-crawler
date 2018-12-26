package main

import (
	"github.com/humbertodias/go-nie-crawler/model"
	"github.com/humbertodias/go-nie-crawler/nie"
	"testing"
)

func TestProvincias(t *testing.T) {
	provincias := nie.ScrapyProvincias("13")
	if len(provincias) == 0 {
		t.Errorf("Length of provincias was incorrect, got: %d, want: %s.", 0, ">0")
	}
}

func TestTramites(t *testing.T) {
	provincia := model.Provincia{
		ID:   "13",
		Name: "Ciudad Real",
		URL:  "https://sede.administracionespublicas.gob.es/icpplus/citar?p=13",
	}
	provincias := []model.Provincia{provincia}
	tramites := nie.ScrapyTramites(provincias)
	if len(tramites) == 0 {
		t.Failed()
	}
}
