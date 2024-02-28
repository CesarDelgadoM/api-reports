package user

import (
	"github.com/CesarDelgadoM/api-reports/pkg/database"
)

type IRepository interface {
	Create(user *User) error
	FindByEmail(email string) (*User, error)
	FindById(id uint) (*User, error)
}

type Repository struct {
	db *database.PostgresDB
}

func NewRepository(db *database.PostgresDB) IRepository {
	return &Repository{
		db,
	}
}

func (repo *Repository) Create(user *User) error {
	err := repo.db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) FindByEmail(email string) (*User, error) {
	var user User

	err := repo.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *Repository) FindById(id uint) (*User, error) {
	var user User

	err := repo.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
