package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Categories struct {
	Id        int64        `json:"id"`
	Name      string       `json:"name"`
	CreatedAt time.Time    `json:"created_at"`
	Questions []*Questions `orm:"reverse(many)"`
}

func FindCategoryById(id int64) (*Categories, error) {
	category := new(Categories)
	db := orm.NewOrm()
	err := db.QueryTable(category).Filter("id", id).One(category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func AllCategories() ([]*Categories, error) {
	var categories []*Categories
	_, err := orm.NewOrm().QueryTable("categories").All(&categories)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
