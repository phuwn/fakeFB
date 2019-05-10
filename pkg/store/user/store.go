package user

import (
	"github.com/jinzhu/gorm"

	"github.com/phuwn/fakeFB/pkg/models"
)

// Store of User
type Store interface {
	Create(db *gorm.DB, user *models.User) error
}
