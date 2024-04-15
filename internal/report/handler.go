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

	api.Use(middlewares.IsAuthenticated).Post("/branches", h.branchReport)
}

// Make a post request to extractor_reports service
// Start generating a branches report
func (h *handler) branchReport(ctx *fiber.Ctx) error {
	agent := fiber.Post(h.config.Extractor.Url)
	agent.ContentType(h.config.Extractor.ContentType)
	agent.Body(ctx.Body())

	_, response, errors := agent.Bytes()
	if len(errors) > 0 {
		zap.Log.Error("Failed to request producer report service: ", errors)
		return errors[0]
	}

	return ctx.JSON(string(response))
}
