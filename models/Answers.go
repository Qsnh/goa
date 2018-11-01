package models

import "time"

type Answers struct {
	User      *Users     `orm:"rel(fk)"`
	Question  *Questions `orm:"rel(fk)"`
	Id        int64
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
