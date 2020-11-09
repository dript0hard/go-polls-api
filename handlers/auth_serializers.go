package handlers

import(
	"net/http"
	"errors"
	"time"

	"github.com/dript0hard/pollsapi/models"
	"github.com/google/uuid"
)

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

