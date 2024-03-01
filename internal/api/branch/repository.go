package branch

import (
	"context"
	"errors"
	"time"

	"github.com/CesarDelgadoM/api-reports/pkg/database"
	"github.com/CesarDelgadoM/api-reports/pkg/httperrors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IRepository interface {
	Create(userId uint, name string, branch *Branch) error
	Find(userId uint, namer, nameb string) (*Branch, error)
	FindAll(userId uint, name string) (*[]Branch, error)
	Update(userId uint, name, nameb string, branch *Branch) error
	Delete(userId uint, name, nameb string) error
}

type Repository struct {
	client *database.MongoDB
}

func NewRepository(client *database.MongoDB) IRepository {
	return &Repository{
		client: client,
	}
}

func (repo *Repository) Create(userId uint, name string, branch *Branch) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{
		"userid": userId,
		"name":   name,
	}

	update := bson.M{
		"$push": bson.M{
			"branches": branch,
		},
	}

	result, err := repo.client.CollectionRestaurant().UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return httperrors.BranchNotFound
	}

	return nil
}

func (repo *Repository) Find(userId uint, namer, nameb string) (*Branch, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{
		"userid": userId,
		"name":   namer,
	}

	project := bson.M{
		"_id": 0,
		"branches": bson.M{
			"$elemMatch": bson.M{
				"name": nameb,
			},
		},
	}

	opts := options.Find().SetProjection(project)
	cursor, err := repo.client.CollectionRestaurant().Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	result := struct {
		Branchs []Branch `json:"branches"`
	}{
		Branchs: []Branch{},
	}

	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
	}

	if len(result.Branchs) == 0 {
		return nil, errors.New("array branchs empty")
	}

	return &result.Branchs[0], nil
}

func (repo *Repository) FindAll(userId uint, name string) (*[]Branch, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{
		"userid": userId,
		"name":   name,
	}

	project := bson.M{
		"branches": 1,
		"_id":      0,
	}

	opts := options.Find().SetProjection(project)
	cursor, err := repo.client.CollectionRestaurant().Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	// Struct for mapping the query result
	result := struct {
		Branches []Branch `json:"branches"`
	}{}

	if cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
	}

	return &result.Branches, nil
}

func (repo *Repository) Update(userId uint, namer, nameb string, branch *Branch) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{
		"userid": userId,
		"name":   namer,
		"branches": bson.M{
			"$elemMatch": bson.M{
				"name": nameb,
			},
		},
	}

	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{
					Key:   "branches.$",
					Value: branch,
				},
			},
		},
	}

	result, err := repo.client.CollectionRestaurant().UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return httperrors.BranchNotFound
	}

	return nil
}

func (repo *Repository) Delete(userId uint, namer, nameb string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	filter := bson.M{
		"userid": userId,
		"name":   namer,
	}

	update := bson.M{
		"$pull": bson.M{
			"branches": bson.M{
				"name": nameb,
			},
		},
	}

	result, err := repo.client.CollectionRestaurant().UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return httperrors.BranchNotFound
	}

	return nil
}
