<div class="card">
    <div class="card-header">
        修改密码
    </div>
    <div class="card-body">
        <div class="row justify-content-center">
            <div class="col-sm-6">
                <form action="" method="post" class="form-horizontal" enctype="multipart/form-data">
                    {{.xsrfdata}}
                    <div class="form-group">
                        <input type="file" class="custom-file-input">
                        <label class="custom-file-label" for="customFile">请选择jpg,png,gif格式的图片文件</label>
                    </div>
                    <div class="form-group">
                        <button type="submit" class="btn btn-outline-light btn-block">立即修改</button>
                    </div>
                </form>
            </div>
        </div>
    </div>
</div>