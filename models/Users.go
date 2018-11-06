package models

import (
	"github.com/Qsnh/goa/libs"
	"github.com/astaxie/beego/orm"
	"os"
	"strconv"
	"time"
)

type Users struct {
	Id         int
	Nickname   string
	Avatar     string
	Email      string
	Password   string
	IsLock     int
	Company    string
	Age        int64
	Profession string
	Website    string
	Weibo      string
	Wechat     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Questions  []*Questions `orm:"reverse(many)"`
	Answers    []*Answers   `orm:"reverse(many)"`
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

	isLock, _ := strconv.Atoi(os.Getenv("USER_REGISTER_LOCK_STATUS"))

	user := Users{
		Nickname:  nickname,
		Email:     email,
		Password:  libs.SHA256Encode(password),
		IsLock:    isLock,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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
