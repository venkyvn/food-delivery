package userbiz

import (
	"context"
	"go-food-delivery/common"
	"go-food-delivery/modules/user/usermodel"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, conditions map[string]string, moreKeys ...string) (*usermodel.User, error)
	Create(ctx context.Context, user *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registrationBiz struct {
	registerStorage RegisterStorage
	hasher          Hasher
}

func NewRegistrationBiz(store RegisterStorage, hasher Hasher) *registrationBiz {
	return &registrationBiz{
		registerStorage: store,
		hasher:          hasher,
	}
}

func (biz *registrationBiz) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, _ := biz.registerStorage.FindUser(ctx, map[string]string{"email": data.Email})

	if user != nil {
		return usermodel.ErrEmailExisted
	}

	salt := common.GenSalt(50)

	password := biz.hasher.Hash(data.Password + salt)
	data.Password = password
	data.Salt = salt
	data.Role = "user"

	if err := biz.registerStorage.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}
