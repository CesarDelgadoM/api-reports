package user

import (
	"errors"
	"strconv"
	"time"

	"github.com/CesarDelgadoM/api-reports/internal/middlewares"
	"github.com/CesarDelgadoM/api-reports/pkg/httperrors"
	"github.com/CesarDelgadoM/api-reports/pkg/logger/zap"
	"github.com/CesarDelgadoM/api-reports/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type IService interface {
	Register(ctx *fiber.Ctx, register *Register) error
	Login(ctx *fiber.Ctx, credentials *Credentials) error
	Authenticate(ctx *fiber.Ctx) (*User, error)
	Logout(ctx *fiber.Ctx)
}

type Service struct {
	repo IRepository
}

func NewUserService(repo IRepository) IService {
	return &Service{
		repo: repo,
	}
}

func (service *Service) Register(ctx *fiber.Ctx, register *Register) error {

	if register.Password != register.PasswordConfirm {
		return httperrors.PasswordDoNotMatch
	}

	user, _ := service.repo.FindByEmail(register.Email)
	if user != nil {
		return httperrors.EmailAlreadyExists
	}

	user = NewUser(register.FirstName, register.LastName, register.Email, register.Password)

	err := service.repo.Create(user)
	if err != nil {
		zap.Log.Error(httperrors.ErrUserCreate, err)
		return httperrors.UserCreateFailed
	}

	zap.Log.Info("User register success: ", user.Email)
	return nil
}

func (service *Service) Login(ctx *fiber.Ctx, credentials *Credentials) error {
	user, err := service.repo.FindByEmail(credentials.Email)
	if err != nil {

		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			return httperrors.EmailNotExists

		default:
			zap.Log.Error(httperrors.ErrFindByEmail, err)
			return httperrors.EmailFindFailed
		}
	}

	err = user.ComparePassword(credentials.Password)
	if err != nil {
		zap.Log.Error(httperrors.ErrComparePasswrod, err)
		return httperrors.InvalidCredentials
	}

	var payload = jwt.StandardClaims{
		Subject:   strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte(utils.SecretKey))
	if err != nil {
		return httperrors.InvalidCredentials
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 1),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	zap.Log.Info("Login success: ", user.Email)
	return nil
}

func (service *Service) Authenticate(ctx *fiber.Ctx) (*User, error) {
	id, _ := middlewares.GetUserId(ctx)

	user, err := service.repo.FindById(id)
	if err != nil {
		zap.Log.Error(httperrors.UserNotFound, err)
		return nil, httperrors.UserNotFound
	}

	zap.Log.Info("Authentication success: ", user.Email)
	return user, nil
}

func (service *Service) Logout(ctx *fiber.Ctx) {

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now(),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	zap.Log.Info("Logout success")
}
