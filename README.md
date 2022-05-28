# Health Check with Golang and Fiber Framework

This is an example project to check services such as PostgresSQL, websites, .etc

## Config

There are two default config keys and they should be placed in the .env file

```dotenv
PORT=":8080"
ENDPOINT="/healthcheck"
```

**PORT:** The port to run Fiber

**ENDPOINT:** The endpoint to show health check data.

## Creating Custom Provider

Create a file under `cmd/providers` folder. For example, we want to create a provider called postgresql.

```go
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

```

That's all. When you create a provider, you just need to call it in the handlers file.

```go
package routes

import (
	"github.com/gofiber/fiber/v2"
	"go_healtcheck/cmd/providers"
	"net/http"
)

func GetHealthCheckData(ctx *fiber.Ctx) error {

	provider := providers.Provider{}

	providers.Init()

	provider.AddProvider(providers.CheckPostgresSQLDB)
	provider.AddProvider(providers.CheckGoogle)

	return ctx.Status(http.StatusOK).JSON(providers.GetProviders())
}
```

That's all :)