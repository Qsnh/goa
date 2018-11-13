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

在访问后台地址前，你需要配置下跨域配置，这个在 `.env` 文件中的 `CORS_ORIGINAL` 中。

另外你需要到 `static/backend` 下面这样操作：

```
npm install
npm build
```

之后你可以访问下面的地址了：

```
https://youdomain.app/static/backend/dist/index.html
```

管理员的账户和密码你可以在 `.env` 文件中配置：

```
ADMIN_USER=
ADMIN_PASS=
```
