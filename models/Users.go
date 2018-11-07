package models

import (
	"github.com/Qsnh/goa/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"os"
	"strconv"
	"strings"
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
	return user, err
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
		Password:  utils.SHA256Encode(password),
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
	err := db.QueryTable(user).Filter("email", email).Filter("password", utils.SHA256Encode(password)).One(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// 生成密码重置地址
func (user *Users) GeneratePasswordResetUrl() string {
	hashString := ""
	idString := strconv.Itoa(user.Id)
	time := time.Now().Unix()
	timeString := strconv.FormatInt(time, 10)
	hashString += idString
	hashString += timeString
	hashString += beego.Substr(user.Password, 1, 10)
	hashString += idString
	hashed := utils.SHA256Encode(hashString)

	baseUrl := beego.URLFor("UserController.PasswordReset")
	url := utils.Url(baseUrl, "id", user.Id, "time", time, "sign", hashed)
	url = strings.TrimRight(os.Getenv("APP_URL"), "/") + url
	return url
}

// 验证密码重置特征值
func (user *Users) CheckPasswordResetHash(giveHashed string, timeString string) bool {
	hashString := ""
	idString := strconv.Itoa(user.Id)
	hashString += idString
	hashString += timeString
	hashString += beego.Substr(user.Password, 1, 10)
	hashString += idString
	hashed := utils.SHA256Encode(hashString)
	return hashed == giveHashed
}
