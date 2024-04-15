package user

import (
	"github.com/CesarDelgadoM/api-reports/internal/middlewares"
	"github.com/CesarDelgadoM/api-reports/pkg/httperrors"
	"github.com/CesarDelgadoM/api-reports/pkg/logger/zap"
	"github.com/CesarDelgadoM/api-reports/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

const (
	operationName = "user: "
)

type handler struct {
	app     *fiber.App
	service IService
}

func NewHandler(app *fiber.App, service IService) {
	handler := handler{
		app:     app,
		service: service,
	}

	handler.InitRouters()
}

func (handler *handler) InitRouters() {
	api := handler.app.Group("/user")

	api.Post("/register", handler.register)
	api.Post("/login", handler.login)

	api.Use(middlewares.IsAuthenticated).Get("/authenticate", handler.authenticate)
	api.Use(middlewares.IsAuthenticated).Patch("/logout", handler.logout)
}

func (handler *handler) register(ctx *fiber.Ctx) error {
	zap.Log.Info(operationName, utils.GetRequestURI(ctx))

	var register Register

	err := ctx.BodyParser(&register)
	if err != nil {
		zap.Log.Error(httperrors.ErrBodyParser, err)
		return httperrors.BadRequest
	}

	err = handler.service.Register(ctx, &register)
	if err != nil {
		return err
	}

	return ctx.JSON("register success")
}

func (handler *handler) login(ctx *fiber.Ctx) error {
	zap.Log.Info(operationName, utils.GetRequestURI(ctx))

	var credentials Credentials

	err := ctx.BodyParser(&credentials)
	if err != nil {
		zap.Log.Error(httperrors.ErrBodyParser, err)
		return httperrors.BadRequest
	}

	err = handler.service.Login(ctx, &credentials)
	if err != nil {
		return err
	}

	return ctx.JSON("login success")
}

func (handler *handler) authenticate(ctx *fiber.Ctx) error {
	zap.Log.Info(operationName, utils.GetRequestURI(ctx))

	user, err := handler.service.Authenticate(ctx)
	if err != nil {
		return err
	}

	return ctx.JSON(user)
}

func (handler *handler) logout(ctx *fiber.Ctx) error {
	zap.Log.Info(operationName, utils.GetRequestURI(ctx))

	handler.service.Logout(ctx)

	return ctx.JSON("logout success")
}
