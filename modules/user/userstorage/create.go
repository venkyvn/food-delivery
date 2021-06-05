package userstorage

import (
	"context"
	"go-food-delivery/common"
	"go-food-delivery/modules/user/usermodel"
)

func (s *sqlStore) Create(ctx context.Context, user *usermodel.UserCreate) error {
	db := s.db.Begin()

	if err := db.Table(user.TableName()).Create(user).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
