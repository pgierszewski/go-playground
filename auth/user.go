package main

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int64  `db:"id" json:"id"`
	Email    string `db:"hour" json:"email"`
	password string `db:"password"`
}

type UserDTO struct {
	Email    string `"json:"email"`
	Password string `"json:"password"`
}

type UserService struct {
	r RepositoryInterface
}

func NewUserService(r RepositoryInterface) UserService {
	us := UserService{r}

	return us
}

func (u *UserService) NewUser(userDTO UserDTO) (User, error) {
	hashedPassword, _ := u.hashPassword(userDTO.Password)
	user := User{Email: userDTO.Email, password: hashedPassword}
	user, err := u.r.createUser(user)
	if err != nil {
		return user, err
	}

	return user, err
}

func (u *UserService) LogIn(email string, password string) (User, error) {
	hashedPassword, err := u.hashPassword(password)
	p := User{1, email, hashedPassword}

	return p, err
}

func (u *UserService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}
