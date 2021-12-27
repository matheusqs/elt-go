package factories

import (
	"dito/elt/extractors"
	"fmt"
)

func GetExtractor(extractorType string) (extractors.IExtractor, error) {
	if extractorType == "sqlServer" {
		return newSqlServer(), nil
	}

	if extractorType == "postgres" {
		return newPostgres(), nil
	}

	return nil, fmt.Errorf("Wrong extractor type passed")
}

func newSqlServer() extractors.IExtractor {
	return &extractors.SqlServerExtractor{
		Extractor: extractors.Extractor{
			Name: "SQL Server",
		},
	}
}

func newPostgres() extractors.IExtractor {
	return &extractors.PostgresExtractor{
		Extractor: extractors.Extractor{
			Name: "Postgres",
		},
	}
}
