package models

import (
    "github.com/astaxie/beego/orm"
)

type Category struct {
    Id int
    Name string
    Created_at int
    Questions []*Question `orm:"reverse(many)"`
}

func init() {
    orm.RegisterModel(new(Category))
}