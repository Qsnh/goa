<div class="container login-box">
    <div class="row justify-content-center">
        <div class="col-sm-4">
            <div class="card center-form-box">
                <div class="card-body">
                    <h3 class="text-center">找回密码</h3>
                    <form action="" method="post" class="form-horizontal">
                        {{ .xsrfdata }}
                        <div class="col-sm-12">
                            <div class="form-group">
                                <label>邮箱</label>
                                <input type="text" name="username" class="form-control form-control-sm" placeholder="请输入注册时的邮箱" required>
                            </div>
                            <div class="form-group">
                                <button class="btn btn-outline-light btn-block btn-sm" type="submit">发送找回密码链接到邮箱</button>
                            </div>
                        </div>
                    </form>
                </div>
            </div>
            <div class="row lh-46">
                <div class="col-sm-6">
                    <a href="{{urlfor "UserController.Login"}}">登录</a>
                </div>
                <div class="col-sm-6 text-right">
                    <a href="{{urlfor "UserController.Register"}}">注册</a>
                </div>
            </div>
        </div>
    </div>
</div>