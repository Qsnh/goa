package models

import (
	"github.com/Qsnh/goa/libs"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/russross/blackfriday"
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

func AnswerPaginate(questionId string, page int64, pageSize int64) ([]Answers, *libs.BootstrapPaginator, error) {
	db := orm.NewOrm()
	answers := []Answers{}

	total, err := db.QueryTable("answers").Filter("question_id", questionId).Count()
	if err != nil {
		return answers, nil, err
	}

	paginator := new(libs.BootstrapPaginator)
	paginator.Instance(total, page, pageSize, beego.URLFor("QuestionController.show", ":id", questionId))

	if page > paginator.TotalPage {
		return answers, paginator, nil
	}

	var startPosition int64
	if page > 0 {
		startPosition = (page - 1) * pageSize
	}
	_, err = db.QueryTable("answers").Filter("question_id", questionId).RelatedSel().OrderBy("created_at", "id").Limit(pageSize, startPosition).All(&answers)
	if err != nil {
		return answers, paginator, err
	}

	for index, item := range answers {
		answers[index].Content = string(blackfriday.MarkdownCommon([]byte(item.Content)))
	}

	return answers, paginator, nil
}
