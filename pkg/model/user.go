// gormに依存するstruct
package model

import (
	"github.com/Tomoya113/go-mvc-api-server/pkg/database"

	"github.com/google/uuid"
)

type User struct {
	ID    int
	Name  string
	Email string
}

type CreateUserParams struct {
	Name  string
	Email string
}

type UserModel struct{}

func NewUserModel() UserModel {
	model := UserModel{}
	return model
}

func (m UserModel) GetUsers() ([]User, error) {
	users := []User{}
	if err := database.Get().Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (m UserModel) CreateUser(params CreateUserParams) (id int, err error) {
	id = int(uuid.New().ID())
	user := User{
		ID:    id,
		Name:  params.Name,
		Email: params.Email,
	}
	if err = database.Get().Create(&user).Error; err != nil {
		return id, err
	}

	return id, nil
}

func (m UserModel) GetUser(id int) (User, error) {
	user := User{}
	if err := database.Get().First(&user, id).Error; err != nil {
		return user, err
	}

	return user, nil
}
