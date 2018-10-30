<div class="card">
    <div class="card-header">
        我的资料
    </div>
    <div class="card-body">
        <div class="row justify-content-center">
            <div class="col-sm-6">
                <form action="" method="post" class="form-horizontal">
                    {{.xsrfdata}}
                    <div class="form-group">
                        <label>公司</label>
                        <input type="text" name="company" value="{{.user.Company}}" class="form-control form-control-sm" placeholder="公司">
                    </div>
                        <div class="form-group">
                            <label>职业</label>
                            <input type="text" name="profession" value="{{.user.Profession}}" class="form-control form-control-sm" placeholder="职业">
                        </div>
                        <div class="form-group">
                            <label>年龄</label>
                            <input type="number" name="age" value="{{.user.Age}}" class="form-control form-control-sm" placeholder="年龄">
                        </div>
                        <div class="form-group">
                            <label>我的主页</label>
                            <input type="text" name="website" value="{{.user.Website}}" class="form-control form-control-sm" placeholder="我的主页">
                        </div>
                        <div class="form-group">
                            <label>微博</label>
                            <input type="text" name="weibo" value="{{.user.Weibo}}" class="form-control form-control-sm" placeholder="年龄">
                        </div>
                        <div class="form-group">
                            <label>微信号</label>
                            <input type="text" name="wechat" value="{{.user.Wechat}}" class="form-control form-control-sm" placeholder="微信号">
                        </div>
                    <div class="form-group">
                        <button type="submit" class="btn btn-outline-light btn-block btn-sm">保存</button>
                    </div>
                </form>
            </div>
        </div>
    </div>
</div>