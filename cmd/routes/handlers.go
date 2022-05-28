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
