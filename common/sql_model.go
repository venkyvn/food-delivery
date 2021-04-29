package common

import "time"

type SQLModel struct {
	Id        int
	Status    int
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
