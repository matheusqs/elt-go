package main

import (
	"dito/elt/extractors"
	"dito/elt/factories"
)

func main() {
	dbTypes := [4]string{"sqlServer", "postgres"}

	for _, dbType := range dbTypes {
		extractor, _ := factories.GetExtractor(dbType)
		printDetails(extractor)
	}
}

func printDetails(g extractors.IExtractor) {
	g.Extract()
}
