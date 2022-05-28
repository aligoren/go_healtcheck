package providers

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func CheckPostgresSQLDB() ProviderResult {

	connStr := os.Getenv("POSTGRESQL_DB")

	result := ProviderResult{
		Name:    "POSTGRESQL_DB",
		Status:  true,
		Message: "PostgresSQL is works!",
	}

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		result.Status = false
		result.Error = fmt.Sprintf("Error while connection database %v", err)

		return result
	}

	if err := db.PingContext(context.Background()); err != nil {

		result.Status = false
		result.Error = fmt.Sprintf("Error while pinging database %v", err)

		return result
	}

	return result
}
