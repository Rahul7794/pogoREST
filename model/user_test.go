package model

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestUserModel_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected while opening a stub db con", err)
	}
	defer db.Close()
	reader := sqlx.NewDb(db, "sqlmock")
	created := time.Date(2020, 01, 01, 01, 00, 00, 00, time.UTC)
	tests := []struct {
		name string
	}{
		{
			name: "Success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &User{
				ID:        1,
				FirstName: "rahul",
				LastName:  "singh",
				City:      "abc",
				CreatedAt: created,
			}
			mock.ExpectQuery(`INSERT INTO v1.users (.+) RETURNING`).WithArgs(user.FirstName, user.LastName, user.City).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			userModel := NewUserModel(reader)
			err = userModel.CreateUser(user)
			require.NoError(t, mock.ExpectationsWereMet())
			mock.ExpectClose()
		})
	}

}
