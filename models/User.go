package models

import (
    "github.com/astaxie/beego/orm"
    "github.com/astaxie/beego/validation"
)

type User struct {
    Id int `form:"-"`
    Nickname string `form:"nickname"`
    Email string `form:"email"`
    Password string `form:"password"`
    Is_lock int8 `form:"-"`
    Created_at int `form:"-"`
    Updated_at int `form:"-"`
    Questions []*Question `orm:"reverse(many)"`
}

func init() {
    orm.RegisterModel(new(User))
}