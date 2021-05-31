package restaurantlikemodel

import "time"

type RestaurantLike struct {
	restaurantId int        `json:"restaurant_id" gorm:"column:restaurant_id;"`
	userId       int        `json:"user_id" gorm:"column:user_id;"`
	createdAt    *time.Time `json:"created_at" gorm:"column:created_at;"`
}

func (RestaurantLike) TableName() string {
	return "restaurant_likes"
}
