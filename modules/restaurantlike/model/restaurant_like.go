package restaurantlikemodel

import (
	"fmt"
	"go-food-delivery/common"
	"time"
)

const EntityName = "RestaurantLike"

type RestaurantLike struct {
	RestaurantId int                `json:"restaurant_id" gorm:"column:restaurant_id;"`
	UserId       int                `json:"user_id" gorm:"column:user_id;"`
	CreatedAt    *time.Time         `json:"created_at" gorm:"column:created_at;"`
	User         *common.SimpleUser `json:"user" gorm:"preload:false"`
}

func (RestaurantLike) TableName() string {
	return "restaurant_likes"
}

func (l *RestaurantLike) GetRestaurantId() int {
	return l.RestaurantId
}

func (l *RestaurantLike) GetUserId() int {
	return l.UserId
}

func ErrUserAlreadyLikeThisRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("You already liked this restaurant"),
		fmt.Sprintf("ErrUserAlreadyLikeThisRestaurant"),
	)
}

func ErrUserCannotLikeThisRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("You cannot like this restaurant due to some error occured"),
		fmt.Sprintf("ErrUserCannotLikeThisRestaurant"),
	)
}

func ErrUserCannotUnLikeThisRestaurant(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("You cannot like this restaurant due to some error occured"),
		fmt.Sprintf("ErrUserCannotUnLikeThisRestaurant"),
	)
}

func ErrUserDoesntLikeThisRestaurantYet(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("You cannot unlike this restaurant due to you does not like this one yet"),
		fmt.Sprintf("ErrUserDoesntLikeThisRestaurantYet"),
	)
}
