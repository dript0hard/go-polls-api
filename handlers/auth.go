package handlers

import (
    "errors"
    "net/http"

    "github.com/go-chi/chi"
    "github.com/go-chi/render"
    "github.com/dript0hard/pollsapi/models"
    "github.com/dript0hard/pollsapi/config"
    "golang.org/x/crypto/pbkdf2"
)

func AuthRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("admin: index"))
	})

	r.Get("/register", registerHandler)

    return r
}


type AuthUserRequest struct {
	*models.User
}

func (authUser *AuthUserRequest) Bind(r *http.Request) error {

    if authUser.Username == "" {
		return errors.New("Missing username.")
	}

    if authUser.Password == "" {
		return errors.New("Missing password.")
	}

    if authUser.Email == "" {
		return errors.New("Missing email.")
	}

	return nil
}

type AuthUserResponse struct {
    *models.User
	JwtToken string `json:"jwt_token"`
}

func NewAuthUserRequest(user *models.User) *AuthUserResponse {
	return &AuthUserResponse{User: user}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
    data := &AuthUserRequest{}
    if err := render.Bind(r, data); err != nil {
        render.Render(w, r, ErrInvalidRequest(err))
        return
    }
    user := models.User
    passwordHash, err := generatePassword(user.Password)

}

func generatePassword(password string) (string, error) {
    return password, nil
}

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

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}
