package handlers

import (
	"errors"
	"net/http"

	"github.com/dript0hard/pollsapi/database"
	pollserrors "github.com/dript0hard/pollsapi/errors"
	"github.com/dript0hard/pollsapi/models"
	"github.com/dript0hard/pollsapi/utils/password"
	"github.com/go-chi/render"
	"gorm.io/gorm"
)

var (
	authDb, _ = database.OpenDB()
)

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
