<div class="container">
    <div class="row">
        <div class="col-sm-12">
            <nav aria-label="breadcrumb">
                <ol class="breadcrumb">
                    <li class="breadcrumb-item"><a href="{{urlfor "IndexController.Index"}}">首页</a></li>
                    <li class="breadcrumb-item active" aria-current="page">我要提问</li>
                </ol>
            </nav>
        </div>
        <div class="col-sm-12">
            <div class="alert alert-warning">
                <h3>在本站提问之前，请确保您做了下面的这些工作</h3>
                <ul>
                    <li>1.尝试在你准备提问论坛的历史文档中搜索答案。</li>
                    <li>2.尝试搜索互联网以找到答案。</li>
                    <li>3.尝试请教懂行的朋友以找到答案。</li>
                </ul>
            </div>
        </div>
        <div class="col-sm-12">
            <form action="" method="post" class="form-horizontal">
                {{.xsrfdata}}
                <div class="form-group">
                    <label>分类</label>
                    <select name="category_id" class="form-control">
                        <option value="">无</option>
                        {{range $index, $category := .categories}}
                            <option value="{{$category.Id}}">{{$category.Name}}</option>
                        {{end}}
                    </select>
                </div>
                <div class="form-group">
                    <label>标题</label>
                    <input type="text" name="title" class="form-control" placeholder="标题">
                </div>
                <div class="form-group">
                    <label>详情</label>
                    <editor></editor>
                </div>
                <div class="form-group text-right">
                    <button type="submit" class="btn btn-primary">发起提问</button>
                </div>
            </form>
        </div>
    </div>
</div>