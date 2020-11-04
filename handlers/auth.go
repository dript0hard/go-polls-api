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
	"gorm.io/gorm"
)

var (
	authDb, _ = database.OpenDB()
)

func AuthRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/login", loginHandler)
	r.Post("/register", registerHandler)

	return r
}

type AuthUserRequest struct {
	Username string
	Email    string
	Password string
}

func (authUser *AuthUserRequest) Bind(r *http.Request) error {

	//TODO(dript0hard): Check if email is valid addr.

	if authUser.Username == "" {
		return errors.New("Missing username.")
	}

	if authUser.Password == "" {
		return errors.New("Missing password.")
	}

	if len(authUser.Password) < 8 {
		return errors.New("Password must be longer than 8 character.")
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
	JwtToken  string    `json:"jwt_token,omitempty"`
}

func NewAuthUserResponse(user *models.User) *AuthUserResponse {
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

	hash := password.NewPBKDF2PasswordSha512().HashPassword(data.Password)
	user.Password = hash.String()

	err := authDb.Create(&user).Error
	if err != nil {
		render.Render(w, r, pollserrors.ErrUserAlreadyExists)
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewAuthUserResponse(&user))
}

type LoginRequest struct {
	Email    string
	Password string
}

func (lr *LoginRequest) Bind(r *http.Request) error {

	if lr.Password == "" {
		return errors.New("Missing password.")
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

func loginHandler(w http.ResponseWriter, r *http.Request) {
	data := &LoginRequest{}
	// check input.
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, pollserrors.ErrInvalidRequest(err))
		return
	}

	// Verify if email is not valid or if exists in the db.
	user := models.User{}

	dbErr := authDb.Where("email = ?", data.Email).First(&user).Error
	if dbErr != nil {
		if errors.Is(dbErr, gorm.ErrRecordNotFound) {
			render.Render(w, r, pollserrors.ErrUserDoesNotExist)
			return
		}
		render.Render(w, r, pollserrors.ErrInternalServerErr(dbErr))
		return
	}

	// Verify password with the db hash.
	hr := password.NewHashResult(user.Password)
	ok := password.NewPBKDF2PasswordSha512().ValidatePassword(data.Password, hr)
	if ok {
		lr := NewLoginResponse("")
		render.Status(r, http.StatusOK)
		render.Render(w, r, lr)
		return
	}
	// else you are not whu you want to become :P
	render.Status(r, http.StatusUnauthorized)
	render.Render(w, r, pollserrors.ErrWrongPassword)
}
