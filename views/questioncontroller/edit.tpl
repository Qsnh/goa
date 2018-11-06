<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <link href="https://lib.baomitu.com/highlight.js/9.13.1/styles/atelier-dune-dark.min.css" rel="stylesheet">
    <link rel="stylesheet" href="/static/assets/main.css">
    <title>GOA</title>
    <script>
        window.text = '{{.Question.Description}}';
    </script>
</head>
<body>
<div id="app">
    {{template "components/header.tpl" .}}

    {{template "components/message.tpl" .}}

    <div class="container">
        <div class="row">
            <div class="col-sm-12">
                <nav aria-label="breadcrumb">
                    <ol class="breadcrumb">
                        <li class="breadcrumb-item"><a href="{{urlfor "IndexController.Index"}}">首页</a></li>
                        <li class="breadcrumb-item active" aria-current="page">问题编辑</li>
                    </ol>
                </nav>
            </div>
            <div class="col-sm-12">
                <form action="" method="post" class="form-horizontal">
                {{.xsrfdata}}
                    <div class="form-group">
                        <label>分类</label>
                        <select name="category_id" class="form-control">
                            <option value="">无</option>
                        {{range $index, $category := .categories}}
                            <option value="{{$category.Id}}"
                            {{if eq $category.Id $.Question.Category.Id}}
                                    selected
                            {{end}}
                            >{{$category.Name}}</option>
                        {{end}}
                        </select>
                    </div>
                    <div class="form-group">
                        <label>标题</label>
                        <input type="text" name="title" value="{{.Question.Title}}" class="form-control" placeholder="标题">
                    </div>
                    <div class="form-group">
                        <label>详情</label>
                        <editor></editor>
                    </div>
                    <div class="form-group text-right">
                        <button type="submit" class="btn btn-primary">保存</button>
                    </div>
                </form>
            </div>
        </div>
    </div>

    {{template "components/footer.tpl" .}}
</div>

<script src="/static/assets/bundle.js"></script>
<script src="https://lib.baomitu.com/highlight.js/9.13.1/highlight.min.js"></script>
<script>hljs.initHighlightingOnLoad();</script>
</body>
</html>