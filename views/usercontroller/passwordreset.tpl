<div class="container login-box">
    <div class="row justify-content-center">
        <div class="col-sm-4">
            <div class="card center-form-box">
                <div class="card-body">
                    <h3 class="text-center">重置密码</h3>
                    <form action="" method="post" class="form-horizontal">
                        {{ .xsrfdata }}
                        <input type="hidden" name="id" value="{{.user.Id}}">
                        <input type="hidden" name="sign" value="{{.sign}}">
                        <input type="hidden" name="time" value="{{.time}}">
                        <div class="col-sm-12">
                            <div class="form-group">
                                <label>邮箱</label>
                                <input type="text" name="username" class="form-control form-control-sm" value="{{.user.Email}}" disabled>
                            </div>
                            <div class="form-group">
                                <label>新密码</label>
                                <input type="password" name="password" class="form-control form-control-sm" placeholder="密码" >
                            </div>
                            <div class="form-group">
                                <label>再输入一次密码</label>
                                <input type="password" name="password_confirmation" class="form-control form-control-sm" placeholder="再输入一次密码" >
                            </div>
                            <div class="form-group">
                                <button class="btn btn-outline-light btn-block btn-sm" type="submit">立即重置</button>
                            </div>
                        </div>
                    </form>
                </div>
            </div>
            <div class="row lh-46">
                <div class="col-sm-6">
                    <a href="{{urlfor "UserController.Register"}}">注册</a>
                </div>
                <div class="col-sm-6 text-right">
                    <a href="{{urlfor "UserController.Login"}}">登录</a>
                </div>
            </div>

        </div>
    </div>
</div>