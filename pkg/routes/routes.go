package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/phuwn/fakeFB/pkg/handlers/user"
	"github.com/phuwn/fakeFB/pkg/models"
	"github.com/phuwn/fakeFB/pkg/util/errs"
)

// HealthCheck of server backend
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	res := models.Response{Success: "ok"}

	data, err := json.Marshal(res)
	if err != nil {
		errs.ResponseError(w, errs.Invalid)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// NewRoutes calling API from routes.
func NewRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", HealthCheck).Methods("GET")
	r.HandleFunc("/healthcheck", HealthCheck).Methods("GET")
	user.Routes(r)

	return r
}
