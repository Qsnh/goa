<div class="container register-box">
    <div class="row justify-content-center">
        <div class="col-sm-4">
            <div class="card center-form-box">
                <div class="card-body">
                    <h3 class="text-center">注册</h3>
                    <form action="" method="post" class="form-horizontal">
                    {{ .xsrfdata }}
                        <div class="form-group">
                            <label>昵称</label>
                            <input type="text" name="nickname" class="form-control form-control-sm" placeholder="昵称" >
                        </div>
                        <div class="form-group">
                            <label>邮箱</label>
                            <input type="text" name="username" class="form-control form-control-sm" placeholder="邮箱" >
                        </div>
                        <div class="form-group">
                            <label>密码</label>
                            <input type="password" name="password" class="form-control form-control-sm" placeholder="密码" >
                        </div>
                        <div class="form-group">
                            <label>再输入一次密码</label>
                            <input type="password" name="password_confirmation" class="form-control form-control-sm" placeholder="再输入一次密码" >
                        </div>
                        <div class="form-group">
                            <button class="btn btn-outline-light btn-block btn-sm" type="submit">注册</button>
                        </div>
                    </form>
                </div>
            </div>
            
            <div class="row lh-46">
                <div class="col-sm text-right">
                    <a href="/login">登录</a>
                </div>
            </div>

        </div>
    </div>
</div>