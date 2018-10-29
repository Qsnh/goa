<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <link rel="stylesheet" href="/static/main.css">
    <title>GOA</title>
</head>
<body>
{{template "components/header.tpl" .}}

    {{template "components/message.tpl" .}}

    <div class="container">
        <div class="row">
            <div class="col-sm-3">
                <ul class="list-group member-left-box text-center">
                    <li class="list-group-item member-info-box">
                        <p>
                            <img src="https://ps.ssl.qhimg.com/t015b383a4233d7e9f4.jpg" width="80" height="80" class="rounded-circle">
                        </p>
                        <p class="nickname">{{.user.Nickname}}</p>
                    </li>
                    <li class="list-group-item">
                        <a href="/member">会员中心</a>
                    </li>
                    <li class="list-group-item">
                        <a href="/member/questions">我的问题</a>
                    </li>
                    <li class="list-group-item">
                        <a href="/member/answers">我的回答</a>
                    </li>
                    <li class="list-group-item">
                        <a href="/member/profile">我的资料</a>
                    </li>
                    <li class="list-group-item">
                        <a href="/member/password">修改密码</a>
                    </li>
                    <li class="list-group-item">
                        <a href="/member/logout">安全退出</a>
                    </li>
                </ul>
            </div>
            <div class="col-sm-9">
            {{.LayoutContent}}
            </div>
        </div>
    </div>

    {{template "components/footer.tpl" .}}
</body>
</html>