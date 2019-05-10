package store

import "github.com/phuwn/fakeFB/pkg/store/user"

// Store of server
type Store struct {
	User user.Store
}
