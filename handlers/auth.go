package handlers

import (
    "errors"
    "net/http"

    "github.com/go-chi/chi"
    "github.com/go-chi/render"
    "github.com/dript0hard/pollsapi/models"
    pollserrors "github.com/dript0hard/pollsapi/errors"
    "github.com/dript0hard/pollsapi/utils/password"
    "github.com/dript0hard/pollsapi/database"
)

func AuthRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("admin: index"))
	})

	r.Post("/register", registerHandler)

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

func NewAuthUserReponse(user *models.User) *AuthUserResponse {
	return &AuthUserResponse{User: user}
}

func (aur *AuthUserResponse) Render(w http.ResponseWriter, r *http.Request) error {
    // delete(aur.Password, string)
	return nil
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
    data := &AuthUserRequest{}
    if err := render.Bind(r, data); err != nil {
        render.Render(w, r, pollserrors.ErrInvalidRequest(err))
        return
    }
    user := data.User
    hashres := password.NewPasswordSha512().HashPassword(user.Password)
    user.Password = hashres.String()

    db, _ :=  database.OpenDB()
    db.Create(user)

    render.Status(r, http.StatusCreated)
    render.Render(w, r, NewAuthUserReponse(user))
}
