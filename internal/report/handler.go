package report

import (
	"github.com/CesarDelgadoM/api-reports/config"
	"github.com/CesarDelgadoM/api-reports/internal/middlewares"
	"github.com/CesarDelgadoM/api-reports/pkg/logger/zap"
	"github.com/gofiber/fiber/v2"
)

type handler struct {
	app    *fiber.App
	config config.ServicesConf
}

func NewHandler(app *fiber.App, config config.ServicesConf) {
	h := handler{
		app:    app,
		config: config,
	}

	h.InitRouters()
}

func (h *handler) InitRouters() {
	api := h.app.Group("/report/generate")

	api.Use(middlewares.IsAuthenticated).Post("/branches", h.reportBranches)
}

// Make a post request to producer_report service
// Start generating a branches report
func (h *handler) reportBranches(ctx *fiber.Ctx) error {
	agent := fiber.Post(h.config.Producer.Url)
	agent.ContentType(h.config.Producer.ContentType)
	agent.Body(ctx.Body())

	_, response, errors := agent.Bytes()
	if len(errors) > 0 {
		zap.Logger.Error("Failed to request producer report service: ", errors)
		return errors[0]
	}

	return ctx.JSON(string(response))
}
