<div class="container">
    <div class="row">
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
                    <textarea name="description" class="form-control" rows="5" placeholder="详情"></textarea>
                </div>
                <div class="form-group">
                    <button type="submit" class="btn btn-primary">发起提问</button>
                </div>
            </form>
        </div>
    </div>
</div>