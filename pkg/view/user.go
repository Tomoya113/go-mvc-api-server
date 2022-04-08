// jsonに変換できるstruct
package view

import (
	"encoding/json"

	"github.com/Tomoya113/go-mvc-api-server/pkg/model"
)

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ConvertUsersToJson(users []model.User) ([]byte, error) {
	var result = []UserResponse{}
	for _, user := range users {
		userResponse := UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}
		result = append(result, userResponse)
	}
	b, err := json.Marshal(result)
	return b, err
}

func ConvertUserToJson(user model.User) ([]byte, error) {
	userResponse := UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	b, err := json.Marshal(userResponse)
	return b, err
}
