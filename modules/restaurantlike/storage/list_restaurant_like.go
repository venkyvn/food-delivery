package restaurantlikestorage

import (
	"context"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"go-food-delivery/common"
	"go-food-delivery/modules/restaurantlike/model"
	"time"
)

const timeLayout = "2006-01-02T15:04:05.999999"

func (s *sqlStore) FetchRestaurantLike(ctx context.Context, ids []int) (map[int]int, error) {

	result := make(map[int]int)

	type sqlData struct {
		RestaurantId int `gorm:"column:restaurant_id;"`
		Count        int `gorm:"column:count;"`
	}

	var likeResult []sqlData

	if err := s.db.
		Table(restaurantlikemodel.RestaurantLike{}.TableName()).
		Select("restaurant_id, count(restaurant_id) as count").
		Where("restaurant_id in (?)", ids).
		Group("restaurant_id").
		Find(&likeResult).
		Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, item := range likeResult {
		result[item.RestaurantId] = item.Count
	}

	return result, nil
}

func (s *sqlStore) GetUserLikedRestaurantById(
	ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantlikemodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]common.SimpleUser, error) {

	db := s.db

	db = db.Table(restaurantlikemodel.RestaurantLike{}.TableName()).Where(conditions)

	if v := filter; v != nil {
		if v.RestaurantId > 0 {
			db.Where("restaurant_id = ?", v.RestaurantId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	db.Preload("User")
	if v := paging.FakeCursor; v != "" {
		timeCreated, err := time.Parse(timeLayout, string(base58.Decode(v)))

		if err != nil {
			return nil, common.ErrDB(err)
		}

		db.Where("created_at < ?", timeCreated.Format("2006-01-02 15:04:05"))
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	var listRestLikeResult []restaurantlikemodel.RestaurantLike

	if err := db.Limit(paging.Limit).
		Order("created_at DESC").
		Find(&listRestLikeResult).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	var listUserLikeRestaurant = make([]common.SimpleUser, len(listRestLikeResult))

	for i, item := range listRestLikeResult {
		listUserLikeRestaurant[i] = *item.User
		listUserLikeRestaurant[i].CreatedAt = item.CreatedAt
		listUserLikeRestaurant[i].UpdatedAt = nil

		if i == len(listUserLikeRestaurant)-1 {
			cursorStr := base58.Encode([]byte(fmt.Sprintf("%v", item.CreatedAt.Format(timeLayout))))
			paging.NextCursor = cursorStr
		}
	}

	return listUserLikeRestaurant, nil
}
