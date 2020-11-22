package errors

import (
	"github.com/go-chi/render"
	"net/http"
)

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func ErrInternalServerErr(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		StatusText:     "Something went wrong.",
		ErrorText:      err.Error(),
	}
}

var (

    ErrNotFound = &ErrResponse{
        HTTPStatusCode: 404, StatusText: "Resource not found."}
    ErrPollDoesNotExist = &ErrResponse{
                    HTTPStatusCode: http.StatusNotFound,
                    StatusText: "Poll does not exist."}
    ErrChoiceDoesNotExist = &ErrResponse{
                    HTTPStatusCode: http.StatusNotFound,
                    StatusText: "Choice does not exist."}
    ErrUserDoesNotExist = &ErrResponse{
                    HTTPStatusCode: http.StatusUnauthorized,
                    StatusText: "User does not exist."}
    ErrWrongPassword = &ErrResponse{
                    HTTPStatusCode: http.StatusUnauthorized,
                    StatusText: "Wrong Password"}
    ErrUserAlreadyExists = &ErrResponse{
                    HTTPStatusCode: http.StatusBadRequest,
                    StatusText: "User already exists"}
)
