<div class="card">
    <div class="card-header">
        修改密码
    </div>
    <div class="card-body">
        <div class="row justify-content-center">
            <div class="col-sm-4">
                <form action="" method="post" class="form-horizontal">
                {{.xsrfdata}}
                    <div class="form-group">
                        <label>原密码</label>
                        <input type="password" name="old_password" class="form-control form-control-sm" placeholder="原密码" required>
                    </div>
                    <div class="form-group">
                        <label>新密码</label>
                        <input type="password" name="new_password" class="form-control form-control-sm" placeholder="原密码" required>
                    </div>
                    <div class="form-group">
                        <label>再输入一次新密码</label>
                        <input type="password" name="new_password_confirmation" class="form-control form-control-sm" placeholder="再输入一次新密码" required>
                    </div>
                    <div class="form-group">
                        <button type="submit" class="btn btn-outline-light btn-block">立即修改</button>
                    </div>
                </form>
            </div>
        </div>
    </div>
</div>