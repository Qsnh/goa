package models

import "github.com/astaxie/beego/orm"

type Categories struct {
    Id int64
    Name string
    CreatedAt int64
    Questions []*Questions `orm:"reverse(many)"`
}

func FindCategoryById(id int64) (*Categories, error)  {
    category := new(Categories)
    db := orm.NewOrm()
    err := db.QueryTable(category).Filter("id", id).One(category)
    if err != nil {
        return nil, err
    }
    return category, nil
}

func AllCategories() ([]*Categories, error)  {
    var categories []*Categories
    _, err := orm.NewOrm().QueryTable("categories").All(&categories)
    if err != nil {
        return nil, err
    }
    return categories, nil
}