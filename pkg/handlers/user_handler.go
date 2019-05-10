package handlers

import (
	"github.com/phuwn/fakeFB/pkg/models"
	"github.com/phuwn/fakeFB/pkg/server"
)

// StoreUserHandler return all approver
func StoreUserHandler(user *models.User) error {
	srv := server.GetSrvConfig()
	db := srv.DB()

	err := srv.Store().User.Create(db, user)
	if err != nil {
		srv.Logger().Errorf("[StoreUserHandler].User.Create err: %s", err.Error())
	}
	return err
}
