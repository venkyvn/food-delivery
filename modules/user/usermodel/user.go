package usermodel

import (
	"errors"
	"go-food-delivery/common"
)

const EntityName = "user"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email;"`
	Password        string        `json:"-" gorm:"column:password;"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	Phone           string        `json:"phone" gorm:"column:phone"`
	Role            string        `json:"-" gorm:"column:role"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"avatar;"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email;"`
	Password        string        `json:"-" gorm:"column:password;"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	Phone           string        `json:"phone" gorm:"column:phone"`
	Role            string        `json:"-" gorm:"column:role"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"avatar;"`
}

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email;"`
	Password string `json:"password" form:"password"`
}

func (u *User) Mask(isAdminOrOwner bool) {
	u.GenUID(common.DbTypeUser)
}

func (u *UserCreate) Mask(isAdminOrOwner bool) {
	u.GenUID(common.DbTypeUser)
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

var (
	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)

	ErrUsernameOrPasswordInvalid = common.NewCustomError(
		errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)
)
