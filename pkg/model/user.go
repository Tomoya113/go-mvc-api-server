// gormに依存するstruct
package model

import (
	"github.com/Tomoya113/go-mvc-api-server/pkg/database"

	"github.com/google/uuid"
)

type User struct {
	Id    int
	Name  string
	Email string
}

type CreateUserParams struct {
	Name  string
	Email string
}

type IUserModel interface {
	GetUsers() ([]User, error)
	CreateUser(params CreateUserParams) (int, error)
	GetUser(id int) (User, error)
}

type UserModel struct{}

func NewUserModel() UserModel {
	model := UserModel{}
	return model
}

func (m UserModel) GetUsers() ([]User, error) {
	users := []User{}
	err := database.Get().Find(&users).Error
	if err == nil {
		database.Get().Find(&users)
	}

	return users, err
}

func (m UserModel) CreateUser(params CreateUserParams) (id int, err error) {
	id = int(uuid.New().ID())
	user := User{
		Id:    id,
		Name:  params.Name,
		Email: params.Email,
	}
	err = database.Get().Create(&user).Error

	// err != nilにするとreturnを2回書く必要が出てきそうだったのでerr == nilにしました
	if err == nil {
		database.Get().Create(&user)
	}
	return id, err
}

func (m UserModel) GetUser(id int) (User, error) {
	user := User{}
	err := database.Get().First(&user, id).Error
	database.Get().First(&user, id)
	return user, err
}
