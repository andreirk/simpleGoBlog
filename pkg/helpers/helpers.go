package helpers

import (
	"log"
	"net/http"
	"web3/pkg/config"
)

// 28 Need to access AppConfig for
// session info
var app config.AppConfig

func ErrorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// 28 Checks if a user is authenticated
func IsAuthenticated(r *http.Request) bool {
	// Check if user_id exists
	exists := app.Session.Exists(r.Context(), "user_id")
	return exists
}
