package extractors

import "fmt"

type PostgresExtractor struct {
	Extractor
}

func (postgresExtractor *PostgresExtractor) Extract() {
	fmt.Println("Postgres Extract: " + postgresExtractor.Name)
}
