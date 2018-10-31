package models

type Answers struct {
	User      *Users     `orm:"rel(fk)"`
	Question  *Questions `orm:"rel(fk)"`
	Id        int64
	Content   string
	CreatedAt int64
	UpdatedAt int64
}
