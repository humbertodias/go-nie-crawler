package main

import (
	"github.com/humbertodias/go-nie-crawler/helper"
	"github.com/humbertodias/go-nie-crawler/nie"
)

func main() {
	provincias := nie.ScrapyProvincias()
	helper.WriteJSON(provincias, "provincias.json")

	tramites := nie.ScrapyTramites(provincias)
	helper.WriteJSON(tramites, "tramites.json")

	oficinas := nie.ScrapyOficinas(tramites)
	helper.WriteJSON(oficinas, "oficinas.json")
}
