<link href="https://cdn.bootcss.com/simplemde/1.11.2/simplemde.min.css" rel="stylesheet">

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
                    <textarea name="description" id="description"></textarea>
                </div>
                <div class="form-group">
                    <button type="submit" class="btn btn-primary">发起提问</button>
                </div>
            </form>
        </div>
    </div>
</div>

<script src="https://cdn.bootcss.com/simplemde/1.11.2/simplemde.min.js"></script>
<script>
    var simplemde = new SimpleMDE({ element: document.getElementById("description") });
</script>