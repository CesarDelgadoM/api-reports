package restaurant

import (
	"errors"

	"github.com/CesarDelgadoM/api-reports/pkg/httperrors"
	"github.com/CesarDelgadoM/api-reports/pkg/logger/zap"
)

type IService interface {
	CreateRestaurant(restaurant *Restaurant) error
	FindRestaurant(userId uint, name string) (*RestaurantData, error)
	FindAllRestaurants(userId uint) (*[]RestaurantData, error)
	UpdateRestaurant(userId uint, name string, restaurant *RestaurantData) error
	DeleteRestaurant(userId uint, name string) error
}

type Service struct {
	repo IRepository
}

func NewService(repo IRepository) IService {
	return &Service{
		repo: repo,
	}
}

func (service *Service) CreateRestaurant(restaurant *Restaurant) error {
	err := service.repo.Create(restaurant)
	if err != nil {
		zap.Logger.Error(httperrors.ErrRestaurantCreate, err)
		return httperrors.RestaurantCreateFailed
	}

	zap.Logger.Info("Restaurant created: ", restaurant.Name)
	return nil
}

func (service *Service) FindRestaurant(userId uint, name string) (*RestaurantData, error) {
	restaurant, err := service.repo.Find(userId, name)
	if err != nil {
		zap.Logger.Error(httperrors.ErrRestaurantNotFound, err)
		return nil, httperrors.RestaurantNotFound
	}

	zap.Logger.Info("Restaurant found: ", restaurant.Name)
	return restaurant, nil
}

func (service *Service) FindAllRestaurants(userId uint) (*[]RestaurantData, error) {
	restaurant, err := service.repo.FindAll(userId)
	if err != nil {
		zap.Logger.Error(httperrors.ErrRestaurantNotFound, err)
		return nil, httperrors.RestaurantNotFound
	}

	zap.Logger.Info("Restaurants found: ", len(*restaurant))
	return restaurant, nil
}

func (service *Service) UpdateRestaurant(userId uint, name string, restaurant *RestaurantData) error {
	err := service.repo.Update(userId, name, restaurant)
	if err != nil {
		switch {
		case errors.Is(err, httperrors.RestaurantNotFound):
			zap.Logger.Error(httperrors.ErrRestaurantNotFound, err)
			return httperrors.RestaurantNotFound

		default:
			zap.Logger.Error(httperrors.ErrRestaurantUpdated, err)
			return httperrors.RestaurantUpdateFailed
		}
	}

	zap.Logger.Info("Restaurant updated: ", restaurant.Name)
	return nil
}

func (service *Service) DeleteRestaurant(userId uint, name string) error {
	err := service.repo.Delete(userId, name)
	if err != nil {
		switch {
		case errors.Is(err, httperrors.RestaurantNotFound):
			zap.Logger.Error(httperrors.ErrRestaurantNotFound, err)
			return httperrors.RestaurantNotFound

		default:
			zap.Logger.Error(httperrors.ErrRestaurantDeleted, err)
			return httperrors.RestaurantDeleteFailed
		}
	}

	zap.Logger.Info("Restaurant deleted: ", name)
	return nil
}
