package main

import (
	"net/http"

	"github.com/Tomoya113/go-mvc-api-server/pkg/database"
)

func main() {
	r := InitializeRouter()
	err := database.Init(nil)
	if err != nil {
		panic(err)
	}
	if err := http.ListenAndServe(":3333", r); err != nil {
		panic(err)
	}
}
