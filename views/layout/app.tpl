<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <link rel="stylesheet" href="/static/assets/main.css">
    <title>GOA</title>
</head>
<body>
    <div id="app">
    {{template "components/header.tpl" .}}

    {{template "components/message.tpl" .}}

    {{.LayoutContent}}

    {{template "components/footer.tpl" .}}
    </div>

    <script src="/static/assets/bundle.js"></script>
</body>
</html>