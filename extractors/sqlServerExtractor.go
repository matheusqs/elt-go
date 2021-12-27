package extractors

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

type SqlServerDatabaseConfig struct {
	server   string
	port     int
	user     string
	password string
}

type DatabaseSchema struct {
	id   int
	name string
}

type SqlServerExtractor struct {
	Extractor
}

func (sqlServerExtractor *SqlServerExtractor) Extract() {
	config := SqlServerDatabaseConfig{
		server:   "localhost",
		port:     1433,
		user:     "sa",
		password: "Pas123456789",
	}

	db := openConnection(config)

	rows := getRows(db)

	for _, row := range rows {
		fmt.Println(row.id, row.name)
	}

	defer db.Close()
}

func getRows(db *sql.DB) []DatabaseSchema {
	rowsCursor, _ := db.QueryContext(context.Background(), getSql())

	var rows []DatabaseSchema

	defer rowsCursor.Close()

	for rowsCursor.Next() {
		var name string
		var id int

		rowsCursor.Scan(&id, &name)

		rows = append(rows, DatabaseSchema{id, name})
	}

	return rows
}

func openConnection(config SqlServerDatabaseConfig) *sql.DB {
	connectionString := getConnectionString(config)

	db, err := sql.Open("sqlserver", connectionString)

	if err != nil {
		log.Fatal("Could not connect to database.")
	}

	return db
}

func getConnectionString(config SqlServerDatabaseConfig) string {
	return fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", config.server, config.user, config.password, config.port)
}

func getSql() string {
	return "SELECT id, name FROM master.dbo.Records;"
}
