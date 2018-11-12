<p align="center"><img src="static/goa.png" height=300/></p>

<p align="center">
<a href="https://travis-ci.org/Qsnh/goa"><img src="https://travis-ci.org/Qsnh/goa.svg?branch=master" alt="Build Status"></a>
<a><img src="https://img.shields.io/github/last-commit/Qsnh/goa.svg"/></a>
<a><img src="https://img.shields.io/github/issues/Qsnh/goa.svg"/></a>
<a><img src="https://img.shields.io/github/issues-closed/Qsnh/goa.svg"/></a>
<a><img src="https://img.shields.io/github/issues-pr/Qsnh/goa.svg"/></a>
<a><img src="https://img.shields.io/github/issues-pr-closed/Qsnh/goa.svg"/></a>
</p>

## 介绍

基于 Beego + Vue 开发的在线问答系统。

## 功能

+ [x] 邮箱注册
+ [x] 邮件密码找回
+ [x] 会员邮件激活
+ [x] markdown内容提问和回答
+ [x] XSS安全过滤
+ [x] Vue前端小组件
+ [x] 完善的会员功能体系
+ [x] 后台前后端分离
+ [x] API接口
+ [ ] 程序/数据库定时备份
+ [x] ENV环境配置
+ [ ] 单元测试

## 安装

```
# 框架
go get github.com/astaxie/beego
go get -u github.com/beego/bee
# MySQL Driver
go get github.com/go-sql-driver/mysql
# Markdown解析库
go get github.com/russross/blackfriday
# env
go get github.com/joho/godotenv
# 图片验证码
go get github.com/dchest/captcha
```

## 后台地址：

```
https://youdomain.app/static/backend/dist/index.html
```