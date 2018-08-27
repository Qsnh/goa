<header class="header">
    <div class="container">
        <div class="row">
            <div class="col-sm-3">
                <a href="/"><b class="logo-text">GOA</b></a>
            </div>
            <div class="col-sm-9">
                <ul>
                    {{if .user}}
                    <li><a href="/member">{{.user.Nickname}}</a></li>
                    {{else}}
                    <li>
                        <a href="/login">登录</a> /
                        <a href="/register">注册</a>
                    </li>
                    <li><a href="/">首页</a></li>
                    {{end}}
                </ul>
            </div>
        </div>
    </div>
</header>