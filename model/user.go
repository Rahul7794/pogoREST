package model

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Required for side effects
)

// UserModelImpl is the interface with various DB operation function
type UserModelImpl interface {
	CreateUser(user *User) error
}

// User stores user details
type User struct {
	ID        uint32    `json:"-"`          // ID is the primary key
	FirstName string    `json:"first_name"` // FirstName of User
	LastName  string    `json:"last_name"`  // LastName of User
	City      string    `json:"city"`       // City of User
	CreatedAt time.Time `json:"-"`          // CreatedAt is the time User record was created in DB
}

// UserModel stores *sqlx.DB reference
type UserModel struct {
	db *sqlx.DB
}

// NewUserModel creates an object
func NewUserModel(db *sqlx.DB) UserModelImpl {
	return &UserModel{
		db: db,
	}
}

// CreateUser creates a new User Record in DB
func (u *UserModel) CreateUser(user *User) error {
	result, err := u.db.NamedQuery(`
		INSERT INTO v1.users (first_name, last_name, city)
		VALUES (:first_name, :last_name, :city)
		RETURNING id, created_at`,
		map[string]interface{}{
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"city":       user.City,
		})

	if err != nil {
		return err
	}
	defer result.Close()
	_ = result.Next()
	var id uint32
	var createdAt time.Time
	err = result.Scan(&id, &createdAt)
	if err != nil {
		return err
	}
	user.CreatedAt = createdAt
	user.ID = id
	return nil
}
