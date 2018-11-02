<header class="container-fluid header-navbar">
    <div class="row">
        <div class="col-sm">
            <div class="container">
                <div class="row">
                    <div class="col-sm">
                        <nav class="navbar navbar-expand-lg navbar-light">
                            <a class="navbar-brand" href="#">GOA</a>
                            <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
                                <span class="navbar-toggler-icon"></span>
                            </button>

                            <div class="collapse navbar-collapse" id="navbarSupportedContent">
                                <ul class="navbar-nav mr-auto">
                                    <li class="nav-item active">
                                        <a class="nav-link" href="/">首页</a>
                                    </li>
                                </ul>
                                <form class="form-inline" method="get" action="{{urlfor "IndexController.Index"}}">
                                    <div class="input-group input-group-sm">
                                        <input class="form-control" type="search" name="keywords" placeholder="请输入关键字">
                                        <div class="input-group-append">
                                            <button class="btn btn-dark" type="submit">搜索</button>
                                        </div>
                                    </div>
                                </form>
                                <ul class="header-right-box">
                                {{if .user}}
                                    <li><a href="/member">{{.user.Nickname}}</a></li>
                                {{else}}
                                    <li>
                                        <a href="/login">登录</a> /
                                        <a href="/register">注册</a>
                                    </li>
                                {{end}}
                                </ul>
                            </div>
                        </nav>
                    </div>
                </div>
            </div>
        </div>
    </div>
</header>