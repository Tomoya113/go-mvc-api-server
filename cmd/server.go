package main

import (
	"net/http"

	"github.com/Tomoya113/go-mvc-api-server/pkg/database"
)

func main() {
	r := InitializeRouter()
	error := database.Init(nil)
	if error != nil {
		panic(error)
	}
	http.ListenAndServe(":3333", r)
}
