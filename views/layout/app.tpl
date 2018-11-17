<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <link rel="shortcut icon" href="/favicon.ico" />
    <meta name="keywords" content="{{.PageKeywords}}">
    <meta name="author" content="GOA">
    <meta name="description" content="{{.PageDescription}}">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no">
    <link crossorigin="anonymous" integrity="sha384-nyXucUVWQWGGNq3mSgG/FGxW4Jmsv+0uCbHZ0VpHRsGZt/PlUT31hz9sUek822eI" href="https://lib.baomitu.com/highlight.js/9.13.1/styles/atom-one-dark.min.css" rel="stylesheet">
    <link rel="stylesheet" href="/static/assets/main.css">
    <title>{{.PageTitle}} - {{.AppName}}</title>
</head>
<body>
    <div id="app">
    {{template "components/header.tpl" .}}

    {{template "components/message.tpl" .}}

    {{if not .IsActive}}
    <div class="container">
        <div class="row">
            <div class="col-sm">
                <div class="alert alert-warning">
                    您的账户还未激活，请先 <a class="color-primary" href="{{urlfor "MemberController.SendActiveMail"}}">激活</a>
                </div>
            </div>
        </div>
    </div>
    {{end}}

    {{.LayoutContent}}

    {{template "components/footer.tpl" .}}
    </div>

    <script src="/static/assets/bundle.js"></script>
    <script src="https://lib.baomitu.com/highlight.js/9.13.1/highlight.min.js"></script>
    <script>hljs.initHighlightingOnLoad();</script>
</body>
</html>