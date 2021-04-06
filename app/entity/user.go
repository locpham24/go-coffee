package entity

import (
	"fmt"
	"github.com/locpham24/go-coffee/app/form"
	"github.com/locpham24/go-coffee/app/model"
	"github.com/locpham24/go-coffee/app/orm"
	"github.com/locpham24/go-coffee/utils"
)

type UserEntity struct{}

type IUserEntity interface {
	Create(input form.RegisterPhoneNumber) (*model.User, error)
}

func (e *UserEntity) Create(input form.RegisterPhoneNumber) (model.User, error) {
	var user model.User
	//TODO: normalize phone number

	existedUser, err := orm.User.GetByPhoneNumber(input.PhoneNumber)
	if err != nil {
		return user, err
	}

	if existedUser != nil {
		return user, fmt.Errorf("phone number is existed")
	}

	password := utils.HashPassword(input.Password)
	if len(password) == 0 {
		return model.User{}, nil
	}

	user = model.User{
		PhoneNumber: input.PhoneNumber,
		Password:    password,
	}

	err = orm.User.Create(&user)

	return user, err
}
