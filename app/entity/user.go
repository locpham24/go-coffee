package entity

import (
	"github.com/locpham24/go-coffee/app/form"
	"github.com/locpham24/go-coffee/app/model"
	"github.com/locpham24/go-coffee/app/orm"
)

type UserEntity struct{}

type IUserEntity interface {
	Create(input form.RegisterPhoneNumber) (*model.User, error)
}

func (e *UserEntity) Create(input form.RegisterPhoneNumber) (model.User, error) {
	//TODO: normalize phone number

	//TODO: check existed phone number

	//TODO: hash password

	user := model.User{
		PhoneNumber: input.PhoneNumber,
		Password:    input.Password,
	}

	err := orm.User.Create(&user)

	return user, err
}
