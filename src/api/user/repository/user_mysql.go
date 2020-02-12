package repository

import (
	"context"
	"database/sql"

	"github.com/andisis/gocrud/src/api/user"
	"github.com/andisis/gocrud/src/helper"
	"github.com/andisis/gocrud/src/model"

	"github.com/sirupsen/logrus"
)

type mysqlUserRepository struct {
	Conn *sql.DB
}

// NewSQLUserRepository will create an object that represent the user.Repository interface
func NewSQLUserRepository(Conn *sql.DB) user.Repository {
	return &mysqlUserRepository{Conn}
}

func (mysql *mysqlUserRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*model.User, error) {
	rows, err := mysql.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer rows.Close()

	result := make([]*model.User, 0)
	for rows.Next() {
		data := new(model.User)

		err = rows.Scan(
			&data.ID,
			&data.Fullname,
			&data.Email,
			&data.Username,
			&data.CreatedAt,
			&data.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, data)
	}

	return result, nil
}

func (mysql *mysqlUserRepository) Fetch(ctx context.Context) ([]*model.User, error) {
	query := `SELECT id, fullname, email, username, created_at, updated_at FROM users ORDER BY fullname ASC`

	return mysql.fetch(ctx, query)
}

func (mysql *mysqlUserRepository) GetByID(ctx context.Context, id int) (res *model.User, err error) {
	query := `SELECT id, fullname, email, username, created_at, updated_at FROM users WHERE id = $1`

	list, err := mysql.fetch(ctx, query, id)

	if len(list) > 0 {
		res = list[0]
	} else {
		return nil, helper.ErrNotFound
	}

	return
}
