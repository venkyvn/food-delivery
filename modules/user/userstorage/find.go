package userstorage

import (
	"context"
	"go-food-delivery/common"
	"go-food-delivery/modules/user/usermodel"
)

func (s *sqlStore) FindUser(ctx context.Context, conditions map[string]string, moreKeys ...string) (*usermodel.User, error) {
	db := s.db.Table(usermodel.User{}.TableName())

	for i := range moreKeys {
		db.Preload(moreKeys[i])
	}

	var user usermodel.User

	if err := db.Where(conditions).First(&user).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &user, nil
}
