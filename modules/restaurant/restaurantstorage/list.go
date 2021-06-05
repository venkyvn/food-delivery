package restaurantstorage

import (
	"context"
	"go-food-delivery/common"
	"go-food-delivery/modules/restaurant/restaurantmodel"
)

func (s *sqlStore) ListDataByCondition(ctx context.Context,
	conditions map[string]interface{},
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	db := s.db

	db = db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where(conditions).
		Where("status in (1)")

	if v := filter; v != nil {
		if v.CityId > 0 {
			db = db.Where("city_id =? ", v.CityId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	//using preload to load user data.
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	// if next currsor != nil => db.find().where id > next cursor . limit ()
	// json.next cursor = result [result.size > 0 ? result.[size-1].id]

	var result []restaurantmodel.Restaurant

	if paging.FakeCursor != "" {
		//order by id desc => where id < uid.
		if uid, err := common.FromBase58(paging.FakeCursor); err == nil {
			db.Where("id < ?", uid.GetLocalID())
		} else {
			db.Offset((paging.Page - 1) * paging.Limit)
		}
	} else {
		db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id DESC").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
