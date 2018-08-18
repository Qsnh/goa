<div class="row">
    <div class="col-sm-4 col-sm-offset-4">
        <form action="" method="post" class="form-horizontal">
            {{.xsrfdata}}
            <div class="form-group">
                <label>原密码</label>
                <input type="password" name="old_password" class="form-control" placeholder="原密码" required>
            </div>
                <div class="form-group">
                    <label>新密码</label>
                    <input type="password" name="new_password" class="form-control" placeholder="原密码" required>
                </div>
                <div class="form-group">
                    <label>再输入一次新密码</label>
                    <input type="password" name="new_password_confirmation" class="form-control" placeholder="再输入一次新密码" required>
                </div>
            <div class="form-group">
                <button type="submit" class="btn btn-primary btn-block">点我保存</button>
            </div>
        </form>
    </div>
</div>