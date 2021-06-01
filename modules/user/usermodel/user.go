package usermodel

import (
	"errors"
	"go-food-delivery/common"
)

const EntityName = "user"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email;"`
	Password        string        `json:"password" gorm:"column:password;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	Phone           string        `json:"phone" gorm:"column:phone"`
	Role            string        `json:"-" gorm:"column:role"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"avatar;"`
}

func (User) TableName() string {
	return "users"
}

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email;"`
	Password        string        `json:"password" gorm:"column:password;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	Phone           string        `json:"phone" gorm:"column:phone"`
	Role            string        `json:"-" gorm:"column:role"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"avatar;"`
}

func (UserCreate) TableName() string {
	return "users"
}

var (
	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)
)
