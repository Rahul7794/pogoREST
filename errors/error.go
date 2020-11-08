package errors

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HTTPError is the custom error struct for this project
type HTTPError struct {
	Code int    // Code is the HTTPError code
	Msg  string // Msg is the HTTPError message
}

// Error return string format of error message
func (e *HTTPError) Error() string {
	return e.Msg
}

// MarshalJSON marshals error message
func (e *HTTPError) MarshalJSON() ([]byte, error) {
	return json.Marshal(echo.Map{
		"message": e.Error(),
	})
}

// ErrorHandler returns Error handling object
func ErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if e, ok := err.(*HTTPError); ok {
		code = e.Code
	}

	_ = c.JSON(code, echo.Map{
		"message": err.Error(),
	})

}
