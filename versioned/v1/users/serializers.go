package users

import (
	"time"

	"github.com/gin-gonic/gin"
)

type UserSerializer struct {
	C *gin.Context
	UserModel
}

type UsersSerializer struct {
	C     *gin.Context
	Users []UserModel
}

type UserResponse struct {
	ID       int    `db:"id, primarykey, autoincrement" json:"id"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"-"`
	Name     string `db:"name" json:"name"`
	Phone    string `db:"phone" json:"phone"`
	Address  string `db:"address" json:"address"`
	Country  string `db:"country" json:"country"`
	Zipcode  string `db:"zipcode" json:"zipcode"`
}

type UserCreateSerializer struct {
	ID        int       `db:"id, primarykey, autoincrement" json:"id"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"-"`
	Name      string    `db:"name" json:"name"`
	Phone     string    `db:"phone" json:"phone"`
	Address   string    `db:"address" json:"address"`
	Country   string    `db:"country" json:"country"`
	Zipcode   string    `db:"zipcode" json:"zipcode"`
	UpdatedAt time.Time `db:"updated_at" json:"-"`
	CreatedAt time.Time `db:"created_at" json:"-"`
}

func (s *UserSerializer) Response() UserResponse {
	response := UserResponse{
		ID:      s.ID,
		Email:   s.Email,
		Name:    s.Name,
		Phone:   s.Phone,
		Address: s.Address,
		Country: s.Country,
		Zipcode: s.Zipcode,
	}
	return response
}

func (s *UsersSerializer) Response() []UserResponse {
	response := []UserResponse{}
	for _, user := range s.Users {
		serializer := UserSerializer{s.C, user}
		response = append(response, serializer.Response())
	}
	return response
}
