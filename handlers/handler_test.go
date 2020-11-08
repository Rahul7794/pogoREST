package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"pogoREST/model"
)

type UserModelStub struct {
}

func (u *UserModelStub) CreateUser(user *model.User) error {
	user.ID = 1
	user.CreatedAt = time.Date(2020, 01, 01, 01, 00, 00, 00, time.UTC)
	user.FirstName = "Rahul"
	user.LastName = "Singh"
	user.City = "s"
	return nil
}

func TestHandler_CreateUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/user")
	u := &UserModelStub{}
	h, _ := NewHandler(u)
	var userJSON = `{"first_name":"Rahul","last_name":"Singh","city":"s"}`
	if assert.NoError(t, h.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, userJSON, strings.TrimSuffix(rec.Body.String(), "\n"))
	}
}
