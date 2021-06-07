package restaurantmodel

type Filter struct {
	CityId int `json:"-" form:"city_id"`
}
