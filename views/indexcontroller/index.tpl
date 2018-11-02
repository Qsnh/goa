<div class="container">
    <div class="row">
        <div class="col-sm-12 mt-15 mb-15">
            <a href="{{urlfor "QuestionController.Create"}}" class="btn btn-primary">我要提问</a>
        </div>
        <div class="col-sm-12">
            {{if .Keywords}}
            <a href="{{urlfor "IndexController.Index"}}" title="点击取消" class="badge badge-info">
                关键字过滤：{{.Keywords}}
            </a>
            {{end}}
            {{if .Category.Name}}
                <a href="{{urlfor "IndexController.Index"}}" title="点击取消" class="badge badge-secondary">
                    分类过滤：{{.Category.Name}}
                </a>
            {{end}}
        </div>
        <div class="col-sm-12 question-box">
            <table class="table table-hover table-borderless question-box">
                <tbody>
                {{range $index, $question := .Questions}}
                <tr>
                    <td class="align-middle" width="60px">
                        <img src="{{$question.User.Avatar}}" class="rounded-circle" width="50" height="50">
                    </td>
                    <td>
                        <h6><a href="{{urlfor "QuestionController.Show" ":id" $question.Id}}">{{$question.Title}}</a></h6>
                        <p class="font-08">
                            <a href="javascript:void(0)"><span class="badge badge-info">{{$question.Category.Name}}</span></a>&nbsp;
                            /&nbsp;<a href="javascript:void(0)">{{$question.User.Nickname}}</a>&nbsp;
                            /&nbsp;{{$question.CreatedAt}}&nbsp;
                            /&nbsp;最新回复&nbsp;
                            {{if $question.AnswerUser}}
                            <a href="javascript:void(0)">{{$question.AnswerUser.Nickname}}</a>&nbsp;{{$question.AnswerAt}}
                            {{end}}
                        </p>
                    </td>
                    <td class="align-middle" width="20">
                        <span class="badge badge-secondary">{{$question.AnswerCount}}</span>
                    </td>
                </tr>
                {{end}}
                </tbody>
            </table>

            <div class="text-right">
            {{str2html .Paginator}}
            </div>
        </div>
    </div>
</div>