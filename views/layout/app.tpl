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

    {{.LayoutContent}}

    {{template "components/footer.tpl" .}}

    <script src="/static/bundle.js"></script>
</body>
</html>