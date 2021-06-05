package common

const (
	DbTypeUser       = 1
	DbTypeRestaurant = 2
)

const CurrentUser = "current_user"

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
