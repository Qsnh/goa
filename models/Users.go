package models

import (
    "github.com/astaxie/beego/orm"
    "github.com/astaxie/beego"
    "time"
    "goa/libs"
)

type Users struct {
    Id int
    Nickname string
    Email string
    Password string
    IsLock int
    CreatedAt int64
    UpdatedAt int64
    Questions []*Question `orm:"reverse(many)"`
}

func UserNicknameExists(nickname string) bool  {
    user := Users{}
    db := orm.NewOrm()
    err := db.QueryTable(user).Filter("nickname", nickname).One(&user)
    if err == orm.ErrNoRows {
        return false
    } else {
        return true
    }
}

func UserEmailExists(email string) bool {
    user := Users{Email:email}
    db := orm.NewOrm()
    err := db.QueryTable(user).Filter("email", email).One(&user)
    if err == orm.ErrNoRows {
        return false
    } else {
        return true
    }
}

func CreateUser(nickname string, email string, password string) (int64, error)  {
    db := orm.NewOrm()

    isLock, _ := beego.AppConfig.Int("user_register_lock_status")

    user := Users{
        Nickname: nickname,
        Email: email,
        Password: libs.SHA256Encode(password),
        IsLock: isLock,
        CreatedAt: time.Now().Unix(),
        UpdatedAt: time.Now().Unix(),
    }

    id, err := db.Insert(&user)
    return id, err
}