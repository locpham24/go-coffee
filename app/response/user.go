package response

import (
	"github.com/locpham24/go-coffee/app/model"
	"time"
)

type User struct {
	Id          int    `json:"id"`
	PhoneNumber string `json:"phoneNumber"`
	CreatedAt   string `json:"createdAt"`
}

func PopulateUserView(user model.User) (User, error) {
	var userView User
	createdAt := time.Time(user.CreatedAt).String()

	userView = User{
		Id:          user.Id,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   createdAt,
	}

	return userView, nil
}
