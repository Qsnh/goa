package models

import (
    "github.com/astaxie/beego/orm"
)

type Question struct {
    Id int
    User_id int
    Category_id int
    Title string
    Description string
    View_num int
    Is_ban int8
    Created_at int
    Updated_at int
    Category *Category `orm:"rel(fk)"`
}

func init() {
    orm.RegisterModel(new(Question))
}