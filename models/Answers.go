package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Answers struct {
	User      *Users     `orm:"rel(fk)"`
	Question  *Questions `orm:"rel(fk)"`
	Id        int64
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func AnswerCreate(user *Users, question *Questions, content string, orm *orm.Ormer) (int64, error) {
	answer := new(Answers)
	answer.User = user
	answer.Question = question
	answer.Content = content
	answer.CreatedAt = time.Now()
	answer.UpdatedAt = time.Now()

	return (*orm).Insert(answer)
}