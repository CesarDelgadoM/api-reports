package branch

import (
	"strconv"

	"github.com/CesarDelgadoM/api-reports/internal/middlewares"
	"github.com/CesarDelgadoM/api-reports/pkg/httperrors"
	"github.com/CesarDelgadoM/api-reports/pkg/logger/zap"
	"github.com/CesarDelgadoM/api-reports/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

const (
	operationName = "branch: "
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
	api := handler.app.Group("/report/restaurant/branch")

	api.Use(middlewares.IsAuthenticated).Post("/create", handler.createBranch)
	api.Use(middlewares.IsAuthenticated).Get("/get/userid::userid/name::name/manager::manager", handler.getBranch)
	api.Use(middlewares.IsAuthenticated).Get("/getall/userid::userid/name::name", handler.getBranchs)
	api.Use(middlewares.IsAuthenticated).Put("/update", handler.updateBranch)
	api.Use(middlewares.IsAuthenticated).Delete("delete/userid::userid/name::name/manager::manager", handler.deleteBranch)
}

func (handler *handler) createBranch(ctx *fiber.Ctx) error {
	zap.Logger.Info(operationName, utils.GetRequestURI(ctx))

	var request Request

	err := ctx.BodyParser(&request)
	if err != nil {
		zap.Logger.Error(httperrors.ErrBodyParser, err)
		return httperrors.BadRequest
	}

	err = handler.service.CreateBranch(request.UserId, request.Name, &request.Branch)
	if err != nil {
		return err
	}

	return ctx.JSON("branch created")
}

func (handler *handler) getBranch(ctx *fiber.Ctx) error {
	zap.Logger.Info(operationName, utils.GetRequestURI(ctx))

	userId, _ := strconv.Atoi(ctx.Params("userid"))
	name := ctx.Params("name")
	manager := ctx.Params("manager")

	branch, err := handler.service.FindBranch(uint(userId), name, manager)
	if err != nil {
		return err
	}

	return ctx.JSON(branch)
}

func (handler *handler) getBranchs(ctx *fiber.Ctx) error {
	zap.Logger.Info(operationName, utils.GetRequestURI(ctx))

	userId, _ := strconv.Atoi(ctx.Params("userid"))
	name := ctx.Params("name")

	branchs, err := handler.service.FindBranchs(uint(userId), name)
	if err != nil {
		return err
	}

	return ctx.JSON(branchs)
}

func (handler *handler) updateBranch(ctx *fiber.Ctx) error {
	zap.Logger.Info(operationName, utils.GetRequestURI(ctx))

	var request Request

	err := ctx.BodyParser(&request)
	if err != nil {
		zap.Logger.Error(httperrors.ErrBodyParser, err)
		return httperrors.BadRequest
	}

	err = handler.service.UpdateBranch(request.UserId, request.Name, request.Manager, &request.Branch)
	if err != nil {
		return err
	}

	return ctx.JSON("branch updated")
}

func (handler *handler) deleteBranch(ctx *fiber.Ctx) error {
	zap.Logger.Info(operationName, utils.GetRequestURI(ctx))

	userId, _ := strconv.Atoi(ctx.Params("userid"))
	name := ctx.Params("name")
	manager := ctx.Params("manager")

	err := handler.service.DeleteBranch(uint(userId), name, manager)
	if err != nil {
		return err
	}

	return ctx.JSON("branch delete")
}
