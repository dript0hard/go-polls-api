package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dript0hard/pollsapi/database"
	"github.com/dript0hard/pollsapi/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

var (
	udb, err = database.OpenDB()
)

func Test() chi.Router {
	r := chi.NewRouter()
	r.Get("/", listUsers)
	return r
}

type UserResp struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (ul *UserResp) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewUserResp(user *models.User) *UserResp {
	return &UserResp{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

type UserList []*UserResp

type UserListResponse struct {
	UserList UserList `json:"user_list"`
}

func (ul *UserListResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// ul.Random = "random string"
	return nil
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	ulr := UserListResponse{}
	udb.Table("users").Select("username", "id", "email", "created_at", "updated_at", "password").Find(&ulr.UserList)
	fmt.Printf("USER: %#v", ulr.UserList[0])
	render.Status(r, http.StatusOK)
	render.Render(w, r, &ulr)
}
