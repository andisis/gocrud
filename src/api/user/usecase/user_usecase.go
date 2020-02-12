package usecase

import (
	"context"
	"time"

	"github.com/andisis/gocrud/src/api/user"
	"github.com/andisis/gocrud/src/model"

	"github.com/sirupsen/logrus"
)

type userUsecase struct {
	userRepo       user.Repository
	contextTimeout time.Duration
}

// NewUserUsecase will create new an userUsecase object representation of user.Usecase interface
func NewUserUsecase(usr user.Repository, timeout time.Duration) user.Usecase {
	return &userUsecase{
		userRepo:       usr,
		contextTimeout: timeout,
	}
}

func (usr *userUsecase) Fetch(c context.Context) ([]*model.User, error) {
	ctx, cancel := context.WithTimeout(c, usr.contextTimeout)
	defer cancel()

	listUser, err := usr.userRepo.Fetch(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return listUser, nil
}

func (usr *userUsecase) GetByID(c context.Context, id int) (*model.User, error) {
	ctx, cancel := context.WithTimeout(c, usr.contextTimeout)
	defer cancel()

	res, err := usr.userRepo.GetByID(ctx, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return res, nil
}
