package user

import (
	"context"

	"github.com/andisis/gocrud/src/model"
)

// Usecase represent the user's usecases
type Usecase interface {
	Fetch(ctx context.Context) ([]*model.User, error)
	GetByID(ctx context.Context, id int) (*model.User, error)
}
