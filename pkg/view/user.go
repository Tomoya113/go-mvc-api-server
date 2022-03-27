// jsonに変換できるstruct
package view

import (
	"encoding/json"

	"go-mvc-api-server/pkg/model"
)

type IUserView interface {
	ConvertUsersToJson(users []model.User) ([]byte, error)
}

type UserView struct{}

func NewUserView() UserView {
	view := UserView{}
	return view
}

func (v UserView) ConvertUsersToJson(users []model.User) ([]byte, error) {
	b, err := json.Marshal(users)
	return b, err
}
