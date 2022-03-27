// gormに依存するstruct
package model

import (
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserParams struct {
	Name  string
	Email string
}

type IUserModel interface {
	GetUsers() []User
	CreateUser(params CreateUserParams)
}

type UserModel struct {
	db *gorm.DB
}

func NewUserModel() UserModel {
	model := UserModel{}

	const dsn = "root:password@tcp(127.0.0.1:3306)/go_mvc_api_server"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	model.db = db
	return model
}

func (m UserModel) GetUsers() []User {
	users := []User{}
	m.db.Find(&users)
	return users
}

func (m UserModel) CreateUser(params CreateUserParams) {
	id := uuid.New().ID()
	user := User{
		Id:    int(id),
		Name:  params.Name,
		Email: params.Email,
	}
	m.db.Create(&user)
}
