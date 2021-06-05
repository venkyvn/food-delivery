package restaurantmodel

import (
	"go-food-delivery/common"
	"go-food-delivery/modules/user/usermodel"
	"strings"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string          `json:"name" gorm:"column:name;"`
	UserId          int             `json:"-" gorm:"column:owner_id;"`
	User            *usermodel.User `json:"user" gorm:"preload:false;"`
	Addr            string          `json:"address" gorm:"column:addr;"`
	Logo            *common.Image   `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images  `json:"cover" gorm:"column:cover;"`
	LikedCount      int             `json:"liked_count" gorm:"-"`
}

type RestaurantUpdate struct {
	common.SQLModel `json:",inline"`
	Name            *string        `json:"name" gorm:"column:name;"`
	Addr            *string        `json:"address" gorm:"column:addr;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;"`
	UserId          int            `json:"-" gorm:"column:owner_id;"`
	Addr            string         `json:"address" gorm:"column:addr;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

func (res *RestaurantCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)
	if len(res.Name) == 0 {
		return ErrNameCannotBeEmpTy
	}
	return nil
}

var (
	ErrNameCannotBeEmpTy = common.NewCustomError(nil, "restaurant name cannot be blank", "ErrNameCannotBeEmpty")
)

func (r *Restaurant) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)
}

func (r *RestaurantCreate) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)
}

func (r *RestaurantUpdate) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DbTypeRestaurant)
}
