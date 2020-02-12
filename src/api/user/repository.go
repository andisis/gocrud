package user

import (
	"context"

	"github.com/andisis/gocrud/src/model"
)

// Repository represent the user's repository contract
type Repository interface {
	Fetch(ctx context.Context) ([]*model.User, error)
	GetByID(ctx context.Context, id int) (*model.User, error)
}
