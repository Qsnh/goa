<div class="container">
    <div class="row">
        <div class="col-sm-12">
            <nav aria-label="breadcrumb">
                <ol class="breadcrumb">
                    <li class="breadcrumb-item"><a href="{{urlfor "IndexController.Index"}}">首页</a></li>
                    <li class="breadcrumb-item active" aria-current="page">{{.User.Nickname}}的主页</li>
                </ol>
            </nav>
        </div>
        <div class="col-sm-3">
            <ul class="list-group member-left-box text-center">
                <li class="list-group-item member-info-box">
                    <p>
                        <img src="{{.User.Avatar}}" width="80" height="80" class="rounded-circle">
                    </p>
                    <p class="nickname">{{.User.Nickname}}</p>
                    {{if .User.Website}}
                        <p><a class="btn btn-primary mt-2 mb-2 btn-sm" href="{{.User.Website}}">个人站点</a></p>
                    {{end}}
                    {{if .User.Weibo}}
                        <p>微博：{{.User.Weibo}}</p>
                    {{end}}
                    {{if .User.Wechat}}
                        <p>微信：{{.User.Wechat}}</p>
                    {{end}}
                </li>
                <li class="list-group-item active">
                    TA的问题
                </li>
                <li class="list-group-item">
                    <a href="{{urlfor "DashboardController.MemberAnswers" ":user_id" .User.Id}}">TA的回答</a>
                </li>
            </ul>
        </div>
        <div class="col-sm-9">
            <div class="card">
                <div class="card-header">
                    TA的问题
                </div>
                <div class="card-body">
                    <table class="table table-hover table-borderless">
                        <thead>
                        <tr>
                            <th class="text-center">标题</th>
                            <th class="text-center">时间</th>
                        </tr>
                        </thead>
                        <tbody>
                        {{if .Questions}}
                        {{range $index, $question := .Questions}}
                        <tr class="text-center">
                            <td><a href="{{urlfor "QuestionController.Show" ":id" $question.Id}}">{{$question.Title}}</a></td>
                            <td>{{timeforhumnas $question.CreatedAt}}</td>
                        </tr>
                        {{end}}
                        {{else}}
                        <tr>
                            <td colspan="2" class="text-center">暂无数据</td>
                        </tr>
                        {{end}}
                        </tbody>
                    </table>

                    <div class="mt-3 mb-3">
                    {{str2html .Paginator}}
                    </div>

                </div>
            </div>
        </div>
    </div>
</div>