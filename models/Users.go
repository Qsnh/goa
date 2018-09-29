package models

import (
    "github.com/Qsnh/goa/libs"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "time"
)

type Users struct {
    Id        int
    Nickname  string
    Email     string
    Password  string
    IsLock    int
    CreatedAt int64
    UpdatedAt int64
    Questions []*Questions `orm:"reverse(many)"`
}

func FindUserById(id int) (*Users, error) {
    user := new(Users)
    db := orm.NewOrm()
    err := db.QueryTable(user).Filter("id", id).One(user)
    if err == nil {
        return user, nil
    }
    return nil, err
}

func UserNicknameExists(nickname string) bool {
    user := new(Users)
    db := orm.NewOrm()
    err := db.QueryTable(user).Filter("nickname", nickname).One(user)
    if err != nil {
        return false
    } else {
        return true
    }
}

func UserEmailExists(email string) bool {
    user := new(Users)
    db := orm.NewOrm()
    err := db.QueryTable(user).Filter("email", email).One(user)
    if err != nil {
        return false
    } else {
        return true
    }
}

func CreateUser(nickname string, email string, password string) (int64, error) {
    db := orm.NewOrm()

    isLock, _ := beego.AppConfig.Int("user_register_lock_status")

    user := Users{
        Nickname:  nickname,
        Email:     email,
        Password:  libs.SHA256Encode(password),
        IsLock:    isLock,
        CreatedAt: time.Now().Unix(),
        UpdatedAt: time.Now().Unix(),
    }

    id, err := db.Insert(&user)
    return id, err
}

func UserExistsByEmailAndPassword(email string, password string) (*Users, error) {
    user := new(Users)
    db := orm.NewOrm()
    err := db.QueryTable(user).Filter("email", email).Filter("password", libs.SHA256Encode(password)).One(user)
    if err != nil {
        return nil, err
    }
    return user, nil
}
