package controller

import "net/http"

// RegisterControllers is used to register a controller
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
