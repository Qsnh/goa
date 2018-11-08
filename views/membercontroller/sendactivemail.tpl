<div class="card">
    <div class="card-header">
        会员中心
    </div>
    <div class="card-body">
        <div class="row justify-content-center">
            <div class="col-sm-6">
                <form action="" method="post">
                {{.xsrfdata}}
                    <div class="form-row">
                        <div class="form-group col-sm-12">
                            <label>邮箱</label>
                            <input class="form-control" type="text" value="{{.user.Email}}" disabled>
                        </div>
                        <div class="form-group col-sm-12">
                            <label>验证码</label>
                            <div class="input-group">
                                <input type="text" class="form-control" placeholder="验证码" name="captcha_code" value="">
                            </div>
                        </div>
                        <div class="col-sm-12 mt-3 mb-3">
                            <captcha></captcha>
                        </div>
                        <div class="form-group col-sm-12">
                            <button class="btn btn-block btn-primary">发送激活邮件</button>
                        </div>
                    </div>
                </form>
            </div>
        </div>
    </div>
</div>