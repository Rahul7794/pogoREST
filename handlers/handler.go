package handlers

import (
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"

	"pogoREST/model"
)

// Custom error message with code
var (
	errBadRequest     = newHTTPError(http.StatusBadRequest, "bad request")
	errRecordNotFound = newHTTPError(http.StatusNotFound, "record not found")
)

// handler stores mutex for thread safety and UserModelImpl reference to access various functions
type handler struct {
	HandlerInterface
	mutex *sync.Mutex
	user  model.UserModelImpl
}

// CreateUser godoc
// @Summary Create a model.User record
// @Description Create a model.User record in DB
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.User true "New model.User"
// @Success 201 {object} model.User
// @Failure 400 {object} errors.HTTPError
// @Router /user [post]
// CreateUser creates a model.User record
func (h *handler) CreateUser(ctx echo.Context) error {
	return h.withLockContext(func() error {
		var user model.User
		if err := ctx.Bind(&user); err != nil {
			return errBadRequest
		}
		err := h.user.CreateUser(&user)
		if err != nil {
			return err
		}
		return ctx.JSON(http.StatusCreated, &user)
	})
}

// NewHandler initializes the mutex and return HandlerInterface object
func NewHandler(user model.UserModelImpl) (HandlerInterface, error) {
	return &handler{
		user:  user,
		mutex: new(sync.Mutex),
	}, nil
}

// withLockContext provide lock for thread safety
func (h *handler) withLockContext(fn func() error) error {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	return fn()
}
