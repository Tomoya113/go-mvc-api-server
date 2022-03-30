// jsonに変換できるstruct
package view

import (
	"encoding/json"

	"github.com/Tomoya113/go-mvc-api-server/pkg/model"
)

type UserResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type IUserView interface {
	ConvertUsersToJson(users []model.User) ([]byte, error)
	ConvertUserToJson(users model.User) ([]byte, error)
}

type UserView struct{}

func NewUserView() UserView {
	view := UserView{}
	return view
}

func (v UserView) ConvertUsersToJson(users []model.User) ([]byte, error) {
	var result = []UserResponse{}
	for _, user := range users {
		userResponse := UserResponse{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		}
		result = append(result, userResponse)
	}
	b, err := json.Marshal(result)
	return b, err
}

func (v UserView) ConvertUserToJson(user model.User) ([]byte, error) {
	userResponse := UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
	b, err := json.Marshal(userResponse)
	return b, err
}
