package models

import (
    "time"
    "github.com/astaxie/beego/orm"
)

type Questions struct {
    User *Users `orm:"rel(fk)"`
    Category *Categories `orm:"rel(fk)"`
    Id int64
    Title string
    Description string
    ViewNum int
    IsBan int8
    CreatedAt int64
    UpdatedAt int64
}

func CreateQuestion(categoryId int64, title string, description string, user *Users) (int64, error)  {
    category, _ := FindCategoryById(categoryId)
    question := new(Questions)
    question.Category = category
    question.User = user
    question.Title = title
    question.Description = description
    question.ViewNum = 0
    question.IsBan = -1
    question.CreatedAt = time.Now().Unix()
    question.UpdatedAt = time.Now().Unix()

    return orm.NewOrm().Insert(question)
}