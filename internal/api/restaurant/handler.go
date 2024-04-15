package restaurant

import (
	"strconv"

	"github.com/CesarDelgadoM/api-reports/internal/middlewares"
	"github.com/CesarDelgadoM/api-reports/pkg/httperrors"
	"github.com/CesarDelgadoM/api-reports/pkg/logger/zap"
	"github.com/CesarDelgadoM/api-reports/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

const (
	operationName = "restaurant: "
)

type handler struct {
	app     *fiber.App
	service IService
}

func NewHandler(app *fiber.App, service IService) {
	handler := &handler{
		app:     app,
		service: service,
	}

	handler.InitRouters()
}

func (handler *handler) InitRouters() {
	api := handler.app.Group("/report/restaurant")

	api.Use(middlewares.IsAuthenticated).Post("/create", handler.createRestaurant)
	api.Use(middlewares.IsAuthenticated).Get("/get/userid::userid/name::name", handler.getRestaurant)
	api.Use(middlewares.IsAuthenticated).Get("/getall/userid::userid", handler.getRestaurants)
	api.Use(middlewares.IsAuthenticated).Put("/update", handler.updateRestaurant)
	api.Use(middlewares.IsAuthenticated).Delete("delete/userid::userid/name::name", handler.deleteRestaurant)
}

func (handler *handler) createRestaurant(ctx *fiber.Ctx) error {
	zap.Log.Info(operationName, utils.GetRequestURI(ctx))

	var request Request

	err := ctx.BodyParser(&request)
	if err != nil {
		zap.Log.Error(httperrors.ErrBodyParser, err)
		return httperrors.BadRequest
	}

	if request.UserId == 0 {
		return httperrors.UserIdCannotBeZero
	}

	request.Restaurant.UserId = request.UserId

	err = handler.service.CreateRestaurant(&request.Restaurant)
	if err != nil {
		return err
	}

	return ctx.JSON("restaurant created")
}

func (handler *handler) getRestaurant(ctx *fiber.Ctx) error {
	zap.Log.Info(operationName, utils.GetRequestURI(ctx))

	name := ctx.Params("name")
	userId, err := strconv.Atoi(ctx.Params("userid"))
	if err != nil {
		return err
	}

	restaurant, err := handler.service.FindRestaurant(uint(userId), name)
	if err != nil {
		return err
	}

	return ctx.JSON(restaurant)
}

func (handler *handler) getRestaurants(ctx *fiber.Ctx) error {
	zap.Log.Info(operationName, utils.GetRequestURI(ctx))

	userId, err := strconv.Atoi(ctx.Params("userid"))
	if err != nil {
		return err
	}

	restaurant, err := handler.service.FindAllRestaurants(uint(userId))
	if err != nil {
		return err
	}

	return ctx.JSON(restaurant)
}

func (handler *handler) updateRestaurant(ctx *fiber.Ctx) error {
	zap.Log.Info(operationName, utils.GetRequestURI(ctx))

	var request Request

	err := ctx.BodyParser(&request)
	if err != nil {
		zap.Log.Error(httperrors.ErrBodyParser, err)
		return httperrors.BadRequest
	}

	restaurantData := request.Restaurant.MapToRestaurantData()

	err = handler.service.UpdateRestaurant(request.UserId, request.Name, &restaurantData)
	if err != nil {
		return err
	}

	return ctx.JSON("restaurant updated")
}

func (handler *handler) deleteRestaurant(ctx *fiber.Ctx) error {
	zap.Log.Info(operationName, utils.GetRequestURI(ctx))

	name := ctx.Params("name")
	userId, err := strconv.Atoi(ctx.Params("userid"))
	if err != nil {
		return err
	}

	err = handler.service.DeleteRestaurant(uint(userId), name)
	if err != nil {
		return err
	}

	return ctx.JSON("restaurant deleted")
}
