package orm

import "github.com/locpham24/go-coffee/app/model"
import "github.com/locpham24/go-coffee/infra"

type userOrm struct{}

type IUser interface {
	Create(user *model.User) (err error)
	GetByPhoneNumber(phoneNumber string) (user *model.User, err error)
}

var User IUser

func init() {
	User = &userOrm{}
}

func (o *userOrm) Create(user *model.User) (err error) {
	result := infra.GetDB().Create(user)
	return result.Error
}

func (o *userOrm) GetByPhoneNumber(phoneNumber string) (user *model.User, err error) {
	result := infra.GetDB().
		Where("phone_number = ?", phoneNumber).
		Limit(1).
		Order("id DESC").
		Find(&user)

	return user, result.Error
}
