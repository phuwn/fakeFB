package user

import (
	"github.com/jinzhu/gorm"

	"github.com/phuwn/fakeFB/pkg/models"
)

type pg struct{}

// NewPGService of User
func NewPGService() Store {
	return &pg{}
}

func (s pg) Create(db *gorm.DB, user *models.User) error {
	err := db.
		Create(user).Error
	return err
}
