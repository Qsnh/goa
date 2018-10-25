<div class="container login-box">
    <div class="row justify-content-center">
        <div class="col-sm-4">
            <div class="card center-form-box">
                <div class="card-body">
                    <h3 class="text-center">登录</h3>
                    <form action="" method="post" class="form-horizontal">
                        {{ .xsrfdata }}
                        <div class="col-sm-12">
                            <div class="form-group">
                                <label>邮箱</label>
                                <input type="text" name="username" class="form-control form-control-sm" placeholder="邮箱" required>
                            </div>
                            <div class="form-group">
                                <label>密码</label>
                                <input type="password" name="password" class="form-control form-control-sm" placeholder="密码" required>
                            </div>
                            <div class="form-group">
                                <button class="btn btn-outline-light btn-block btn-sm" type="submit">登录</button>
                            </div>
                        </div>
                    </form>
                </div>
            </div>
            <div class="row lh-46">
                <div class="col-sm-6">
                    <a href="/register">注册</a>
                </div>
                <div class="col-sm-6 text-right">
                    <a href="/">找回密码</a>
                </div>
            </div>

        </div>
    </div>
</div>