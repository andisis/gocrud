package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v2"

	"github.com/andisis/gocrud/src/api/user/repository"
)

func TestFetch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "fullname", "email", "username", "created_at", "updated_at"}).
		AddRow(1, "Andi Siswanto", "andisis92@gmail.com", "andisis", time.Now(), time.Now()).
		AddRow(2, "Ammar Syafiq Siswanto", "ammarsyafiq@gmail.com", "ammarsyafiq", time.Now(), time.Now())
	query := `SELECT id, fullname, email, username, created_at, updated_at FROM users
	ORDER BY fullname ASC`

	mock.ExpectQuery(query).WillReturnRows(rows)
	u := repository.NewSQLUserRepository(db)
	list, err := u.Fetch(context.TODO())

	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "fullname", "email", "username", "created_at", "updated_at"}).
		AddRow(1, "Andi Siswanto", "andisis92@gmail.com", "andisis", time.Now(), time.Now())

	query := `SELECT id, fullname, email, username, created_at, updated_at FROM users
	WHERE id = ?`

	mock.ExpectQuery(query).WillReturnRows(rows)
	u := repository.NewSQLUserRepository(db)
	id := int(5)
	user, err := u.GetByID(context.TODO(), id)

	assert.NoError(t, err)
	assert.NotNil(t, user)
}

func TestGetError404(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := `SELECT id, fullname, email, username, created_at, updated_at FROM users
	WHERE id = ?`

	mock.ExpectQuery(query).WillReturnError(err)
	u := repository.NewSQLUserRepository(db)
	id := int(1)
	user, err := u.GetByID(context.TODO(), id)

	assert.Error(t, err)
	assert.Nil(t, user)
}

func TestGetErrorScanRow(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "fullname", "email", "username", "created_at"}).
		AddRow(1, "Andi Siswanto", "andisis92@gmail.com", "andisis", time.Now())

	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	u := repository.NewSQLUserRepository(db)

	_, err = u.Fetch(context.TODO())

	assert.Error(t, err)
}
