package handlers

import (
	"github.com/labstack/echo/v4"

	"pogoREST/db"
	"pogoREST/errors"
	"pogoREST/model"
)

// HandlerInterface defines various database handlers
type HandlerInterface interface {
	CreateUser(ctx echo.Context) error // Insert record to the DB
}

// DB is the map for various database type and its object
var DB = map[string]func(user model.UserModelImpl) (HandlerInterface, error){
	"postgres": NewHandler,
}

// NewDBEngine create connections based on the connections information provided
func NewDBEngine(dbType string) (HandlerInterface, error) {
	if _, ok := DB[dbType]; ok {
		con, err := db.GetConnection(dbType)
		if err != nil {
			return nil, err
		}
		engine, err := db.ConnectDB(con)
		if err != nil {
			return nil, err
		}
		return DB[dbType](model.NewUserModel(engine))
	}
	return nil, nil
}

// newHTTPError creates an errors.HTTPError object reference
func newHTTPError(code int, msg string) *errors.HTTPError {
	return &errors.HTTPError{Code: code, Msg: msg}
}
