<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <link href="https://lib.baomitu.com/highlight.js/9.13.1/styles/atelier-dune-dark.min.css" rel="stylesheet">
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
    <script src="https://lib.baomitu.com/highlight.js/9.13.1/highlight.min.js"></script>
    <script>hljs.initHighlightingOnLoad();</script>
</body>
</html>