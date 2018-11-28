package utils

import (
	"crypto"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func SHA256Encode(s string) string {
	sha256 := crypto.SHA256.New()
	sha256.Write([]byte(s))
	return hex.EncodeToString(sha256.Sum(nil))
}

func Url(url string, params ...interface{}) string {
	queryString := ""
	for index, item := range params {
		if index%2 == 0 {
			queryString += item.(string) + "="
		} else {
			queryString += ToString(item) + "&"
		}
	}
	if url != "/" {
		url = strings.TrimRight(url, "/")
	}
	queryString = strings.TrimRight(queryString, "&")
	return url + "?" + queryString
}

func ToString(i interface{}) string {
	switch i.(type) {
	case string:
		return i.(string)
	case int:
		return strconv.Itoa(i.(int))
	case int64:
		return strconv.FormatInt(i.(int64), 10)
	}
	return ""
}

// 将时间转换为人类可阅读的格式
func TimeDiffForHumans(t time.Time) string {
	unix := t.Unix()
	now := time.Now().Unix()
	b := now - unix
	if b < 0 {
		return t.Format("2006-01-01 15:04:05")
	}
	if b < 60 {
		return fmt.Sprintf("%d秒前", b)
	}
	// 单位：分钟
	if b < 3600 {
		b = b / 60
		return fmt.Sprintf("%d分钟前", b)
	}
	// 单位：小时
	b = b / 3600
	if b < 24 {
		return fmt.Sprintf("%d个小时前", b)
	}
	// 单位：天
	b = b / 24
	if b < 30 {
		return fmt.Sprintf("%d天前", b)
	}
	// 单位：月
	b = b / 30
	if b < 12 {
		return fmt.Sprintf("%d个月前", b)
	}
	// 单位：年
	b = b / 12
	if b > 3 {
		return t.Format("2006-01-01 15:04:05")
	}
	return fmt.Sprintf("%d年钱", b)
}

// 获取当前工作目录
func Pwd() string {
	return filepath.Dir(os.Args[0])
}

// 传入开始时间，计算结束时间
func ComputedHandlerSeconds(startTime int64) float64 {
	return float64(time.Now().UnixNano()-startTime) / 1e9
}
