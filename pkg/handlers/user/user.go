package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/phuwn/fakeFB/pkg/handlers"
	"github.com/phuwn/fakeFB/pkg/models"
	"github.com/phuwn/fakeFB/pkg/server"
	"github.com/phuwn/fakeFB/pkg/util/errs"
)

// StoreUser help store user info
func StoreUser(w http.ResponseWriter, r *http.Request) {
	srv := server.GetSrvConfig()

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		srv.Logger().Errorf("[StoreUser].ReadAll() err: %s", err.Error())
		errs.ResponseError(w, errs.Invalid)
		return
	}

	user := models.User{}

	if err = json.Unmarshal(data, &user); err != nil {
		srv.Logger().Errorf("[StoreUser].Unmarshal() err: %s", err.Error())
		errs.ResponseError(w, errs.Invalid)
		return
	}

	if err = user.Validate(); err != nil {
		srv.Logger().Errorf("[StoreUser].Validate err: %s", err.Error())
		errs.ResponseError(w, err)
		return
	}

	err = handlers.StoreUserHandler(&user)
	if err != nil {
		errs.ResponseError(w, err)
		return
	}

	resJSON, err := json.Marshal(models.Response{Success: "ok"})
	if err != nil {
		srv.Logger().Errorf("[StoreUser].Marshal err: %s", err.Error())
		errs.ResponseError(w, err)
		return
	}

	srv.Logger().Logf("[StoreUser] successful: %s", time.Now())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resJSON)
}
