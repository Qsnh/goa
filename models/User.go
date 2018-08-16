package models

import (
    "github.com/astaxie/beego/orm"
)

type User struct {
    Id int
    Nickname string
    Email string
    Password string
    Is_lock int8
    Created_at int
    Updated_at int
    Questions []*Question `orm:"reverse(many)"`
}

func init() {
    orm.RegisterModel(new(User))
}