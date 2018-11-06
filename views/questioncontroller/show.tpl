<div class="container">
    <div class="row">
        <div class="col-sm-12">
            <nav aria-label="breadcrumb">
                <ol class="breadcrumb">
                    <li class="breadcrumb-item"><a href="{{urlfor "IndexController.Index"}}">首页</a></li>
                    <li class="breadcrumb-item">
                        <a href="{{(urlquery (urlfor "IndexController.Index") "category_id" .question.Category.Id)}}">{{.question.Category.Name}}</a>
                    </li>
                    <li class="breadcrumb-item active" aria-current="page">{{.question.Title}}</li>
                </ol>
            </nav>
        </div>
        <div class="col-sm-9">
            <div class="card">
                <div class="card-header">
                <b>问题：</b>{{.question.Title}}
                </div>
                <div class="card-body">
                    <div class="question-description">
                    {{str2html .question.Description}}
                    </div>
                </div>
                <div class="card-footer">
                    <small>{{.question.Category.Name}} | {{.question.User.Nickname}} | {{.question.CreatedAt}}</small>
                </div>
            </div>
        </div>
        <div class="col-sm-3">
            <ul class="list-group member-left-box text-center">
                <li class="list-group-item member-info-box">
                    <p>
                        <a href="{{urlfor "MemberController.Index"}}">
                            <img src="{{.question.User.Avatar}}" width="80" height="80" class="rounded-circle">
                        </a>
                    </p>
                    <p class="nickname">{{.question.User.Nickname}}</p>
                    {{if .question.User.Website}}
                        <p><a class="btn btn-primary mt-2 mb-2 btn-sm" href="{{.question.User.Website}}">个人站点</a></p>
                    {{end}}
                    {{if .question.User.Weibo}}
                    <p>微博：{{.question.User.Weibo}}</p>
                    {{end}}
                    {{if .question.User.Wechat}}
                        <p>微信：{{.question.User.Wechat}}</p>
                    {{end}}
                </li>
            </ul>
        </div>
    </div>
</div>

<div class="container mt-5">
    <div class="row">
        <div class="col-sm-9">

            <div class="card">
                <div class="card-header">
                    我要回答
                </div>
                <div class="card-body">
                    <form action="{{urlfor "QuestionController.AnswerHandler" ":id" .question.Id}}" class="form-horizontal" method="post">
                        {{ .xsrfdata }}
                        <div class="form-group">
                            <editor></editor>
                        </div>
                        <div class="form-group text-right">
                            <button class="btn btn-primary">回答TA</button>
                        </div>
                    </form>
                </div>
            </div>

            <div class="card mt-15">
                <div class="card-header">
                    小伙伴们的回答
                </div>
                <div class="card-body">
                    <table class="table table-hover table-borderless">
                        <tbody>
                        {{if .Answers}}
                        {{range $index, $answer := .Answers}}
                        <tr>
                            <td class="text-center" width="100">
                                <p><img src="{{$answer.User.Avatar}}" class="rounded-circle" width="40" height="40"></p>
                                <p><small>{{$answer.User.Nickname}}</small></p>
                            </td>
                            <td>
                            <div class="answer-content-box">
                            {{str2html $answer.Content}}
                            </div>
                            </td>
                        </tr>
                        {{end}}
                        {{else}}
                        <tr>
                            <td colspan="2" class="text-center">暂无回答</td>
                        </tr>
                        {{end}}
                        </tbody>
                    </table>
                </div>
            </div>

            <div class="mt-3 mb-3">
            {{str2html .Paginator}}
            </div>
        </div>
        <div class="col-sm-3">
            <div class="card">
                <div class="card-header">
                    广告
                </div>
                <div class="card-body">

                </div>
            </div>
        </div>
    </div>
</div>

<script src="https://lib.baomitu.com/highlight.js/9.13.1/highlight.min.js"></script>
<script>hljs.initHighlightingOnLoad();</script>