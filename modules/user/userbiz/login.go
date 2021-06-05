package userbiz

import (
	"context"
	"go-food-delivery/common"
	"go-food-delivery/component/tokenprovider"
	"go-food-delivery/modules/user/usermodel"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreKeys ...string) (*usermodel.User, error)
}

type loginBiz struct {
	storeUser     LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBiz(
	storeUser LoginStorage,
	tokenProvider tokenprovider.Provider,
	hasher Hasher,
	expiry int) *loginBiz {
	return &loginBiz{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

/*
Login flow:
1. find user by email
2. compare hash password with salt
3. issue token (access, refresh)
4. return token
*/

func (biz *loginBiz) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {

	user, err := biz.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	passHashed := biz.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}
	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return accessToken, nil
}
