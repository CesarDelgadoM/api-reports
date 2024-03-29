package user

import "golang.org/x/crypto/bcrypt"

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Register struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
}

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  []byte `json:"-"`
}

func NewUser(firstName, lastName, email, password string) *User {
	user := User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
	user.SetHashedPassword(password)

	return &user
}

func (user *User) SetHashedPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
