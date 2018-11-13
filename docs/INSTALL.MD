# 安装

### 拉取代码

```
go get github.com/Qsnh/goa
```

### 安装数据表

将 `data` 下面的SQL文件导入到数据库中

### 配置文件

```
cp .env.example .env
```

打开 `.env` 文件，配置相关的信息，如：数据库的链接信息，阿里云的邮件配置，自动备份存储路径等。

### 编译运行

```
cd ~/go/src/github.com/Qsnh/goa
go build -o goa main.go
```

### 后台地址

```
https://youdomain.app/static/backend/dist/index.html
```
