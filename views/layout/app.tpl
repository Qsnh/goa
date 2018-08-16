<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <link href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
    <link rel="stylesheet" href="/static/css/app.css">
    <title>GOA</title>
</head>
<body>
    {{template "components/header.tpl" .}}

    {{template "components/message.tpl" .}}

    {{.LayoutContent}}

    {{template "components/footer.tpl" .}}
</body>
</html>