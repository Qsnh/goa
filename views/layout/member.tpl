<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <link rel="stylesheet" href="/static/main.css">
    <title>GOA</title>
</head>
<body>
{{template "components/header.tpl" .}}

    {{template "components/message.tpl" .}}

    <div class="container">
        <div class="row">
            <div class="col-sm-3 text-center">
                <p>
                    <img src="https://ps.ssl.qhimg.com/t013658e41e8c191970.jpg" class="img-circle" width="50" height="50">
                </p>
                <p>
                {{.user.Nickname}}
                </p>
            </div>
            <div class="col-sm-9">
            {{.LayoutContent}}
            </div>
        </div>
    </div>

    {{template "components/footer.tpl" .}}
</body>
</html>