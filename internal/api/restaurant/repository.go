package restaurant

import (
	"context"
	"time"

	"github.com/CesarDelgadoM/api-reports/pkg/database"
	"github.com/CesarDelgadoM/api-reports/pkg/httperrors"
	"go.mongodb.org/mongo-driver/bson"
)

type IRepository interface {
	Create(restaurant *Restaurant) error
	Find(userId uint, name string) (*RestaurantData, error)
	FindAll(userId uint) (*[]RestaurantData, error)
	Update(userId uint, name string, restaurant *RestaurantData) error
	Delete(userId uint, name string) error
}

type Repository struct {
	client *database.MongoDB
}

func NewRepository(client *database.MongoDB) IRepository {
	return &Repository{
		client: client,
	}
}

func (repo *Repository) Create(restaurant *Restaurant) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err := repo.client.CollectionRestaurant().InsertOne(ctx, restaurant)
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) Find(userId uint, name string) (*RestaurantData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{
		"userid": userId,
		"name":   name,
	}

	var restaurant RestaurantData

	err := repo.client.CollectionRestaurant().FindOne(ctx, filter).Decode(&restaurant)
	if err != nil {
		return nil, err
	}

	return &restaurant, err
}

func (repo *Repository) FindAll(userId uint) (*[]RestaurantData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{
		"userid": userId,
	}

	var restaurants []RestaurantData

	cursor, err := repo.client.CollectionRestaurant().Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var restaurant RestaurantData
		if err := cursor.Decode(&restaurant); err != nil {
			return nil, err
		}

		restaurants = append(restaurants, restaurant)
	}

	return &restaurants, err
}

func (repo *Repository) Update(userId uint, name string, restaurant *RestaurantData) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{
		"userid": userId,
		"name":   name,
	}

	update := bson.M{
		"$set": restaurant,
	}

	result, err := repo.client.CollectionRestaurant().UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return httperrors.RestaurantNotFound
	}

	return nil
}

func (repo *Repository) Delete(userId uint, name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{
		"userid": userId,
		"name":   name,
	}

	result, err := repo.client.CollectionRestaurant().DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return httperrors.RestaurantNotFound
	}

	return nil
}
