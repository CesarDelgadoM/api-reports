package httperrors

import "github.com/gofiber/fiber/v2"

// Local errors
const (
	// Authentication
	ErrComparePasswrod = "Compare password failed: "
	ErrBodyParser      = "Body parser failed: "

	// Email
	ErrFindByEmail = "Find by email failed: "

	// User
	ErrUserNotFound = "User not found: "
	ErrUserCreate   = "Create user failed: "

	// Restaurant
	ErrRestaurantNotFound = "Restaurant/restaurants not found: "
	ErrRestaurantCreate   = "Create restaurant failed: "
	ErrRestaurantUpdated  = "Restaurant updated failed: "
	ErrRestaurantDeleted  = "Restaurant deleted failed: "

	// Branch
	ErrBranchNotFound = "Branch/branches not found: "
	ErrBranchCreate   = "Branch created failed: "
	ErrBranchUpdate   = "Branch update failed: "
	ErrBranchDelete   = "Branch delete failed: "
)

// Http errors
var (
	// Authentication
	PasswordDoNotMatch = fiber.NewError(fiber.StatusBadRequest, "password do not match")
	InvalidCredentials = fiber.NewError(fiber.StatusBadRequest, "invalid credentials")
	BadRequest         = fiber.NewError(fiber.StatusBadRequest, "Bad request")

	// Email
	EmailFindFailed    = fiber.NewError(fiber.StatusInternalServerError, "failed to find email")
	EmailAlreadyExists = fiber.NewError(fiber.StatusBadRequest, "email already exists")
	EmailNotExists     = fiber.NewError(fiber.StatusBadRequest, "email not exist")

	// User
	UserNotFound       = fiber.NewError(fiber.StatusNotFound, "user not found")
	UserCreateFailed   = fiber.NewError(fiber.StatusInternalServerError, "failed to create user")
	UserIdCannotBeZero = fiber.NewError(fiber.StatusBadRequest, "userid cannot be 0")

	// Restaurant
	RestaurantNotFound     = fiber.NewError(fiber.StatusNotFound, "restaurant/restaurants not found")
	RestaurantCreateFailed = fiber.NewError(fiber.StatusInternalServerError, "failed to create restaurant")
	RestaurantUpdateFailed = fiber.NewError(fiber.StatusInternalServerError, "failed to update restaurant")
	RestaurantDeleteFailed = fiber.NewError(fiber.StatusInternalServerError, "failed to delete restaurant")

	// Branch
	BranchNotFound     = fiber.NewError(fiber.StatusNotFound, "branch/branches not found")
	BranchArrayEmpty   = fiber.NewError(fiber.StatusNotFound, "branch array empty")
	BranchCreateFailed = fiber.NewError(fiber.StatusInternalServerError, "failed to create branch")
	BranchUpdateFailed = fiber.NewError(fiber.StatusInternalServerError, "failed to update branch")
	BranchDeleteFailed = fiber.NewError(fiber.StatusInternalServerError, "failed to delete branch")
)
