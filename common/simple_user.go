package common

type SimpleUser struct {
	SQLModel  `json:",inline"`
	FirstName string `json:"first_name" gorm:"column:first_name;"`
	LastName  string `json:"last_name" gorm:"column:last_name;"`
	Avatar    *Image `json:"avatar,omitempty" gorm:"avatar;"`
}

func (SimpleUser) TableName() string {
	return "users"
}

func (u *SimpleUser) Mask(isAdminOrOwner bool) {
	u.GenUID(DbTypeUser)
}
