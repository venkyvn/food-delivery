package userstorage

import (
	"context"
	"go-food-delivery/common"
	"go-food-delivery/modules/user/usermodel"
)

func (s *sqlStore) Create(ctx context.Context, user *usermodel.UserCreate) error {
	db := s.db

	if err := db.Create(&user).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
