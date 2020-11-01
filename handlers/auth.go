package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/dript0hard/pollsapi/database"
	pollserrors "github.com/dript0hard/pollsapi/errors"
	"github.com/dript0hard/pollsapi/models"
	"github.com/dript0hard/pollsapi/utils/password"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/google/uuid"
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
	Username string
	Email    string
	Password string
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
	CreatedAt time.Time `json:"created_at"`
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	JwtToken  string    `json:"jwt_token"`
}

func NewAuthUserReponse(user *models.User) *AuthUserResponse {
	return &AuthUserResponse{
		CreatedAt: user.CreatedAt,
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
	}
}

func (aur *AuthUserResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	data := &AuthUserRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, pollserrors.ErrInvalidRequest(err))
		return
	}

	user := models.User{
		Username: data.Username,
		Email:    data.Email,
	}
	hash := password.NewPasswordSha512().HashPassword(data.Password)
	user.Password = hash.String()

	db, _ := database.OpenDB()

	err := db.Create(&user).Error
	if err != nil {
		render.Render(w, r, pollserrors.ErrUserAlreadyExists(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewAuthUserReponse(&user))
}

type LoginRequest struct {
	Email    string
	Password string
}

func (lr *LoginRequest) Bind(r *http.Request) error {

	if lr.Password == "" {
		return errors.New("Missing password.")
	}

	if len(lr.Password) < 8 {
		return errors.New("Password must be longer than 8 character.")
	}

	if lr.Email == "" {
		return errors.New("Missing email.")
	}

	return nil
}

type LoginResponse struct {
	JwtToken string `json:"jwt_token"`
}

func NewLoginResponse(token string) *LoginResponse {
	return &LoginResponse{
		JwtToken: token,
	}
}

func (lr *LoginResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func loginHandler(w http.ResponseWriter, r *http.Request) {}
