package user

import (
	"github.com/gorilla/mux"
)

// Routes for User
func Routes(r *mux.Router) {
	r.HandleFunc("/login", StoreUser).Methods("POST")
}
