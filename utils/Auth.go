package utils

import (
	"github.com/astaxie/beego"
	"strconv"
)

func AuthSign(Id int, Email string, Password string) string  {
	s := strconv.Itoa(Id) + Email + beego.Substr(Password, 0, 13)
	return SHA256Encode(s)
}

func AuthSignCheck(Id int, Email string, Password string, sign string) bool  {
	s := AuthSign(Id, Email, Password)
	return s == sign
}