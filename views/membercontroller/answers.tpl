<div class="card">
    <div class="card-header">
        我的回答
    </div>
    <div class="card-body">
        <table class="table table-hover table-borderless">
            <thead>
            <tr>
                <th class="text-center">问题</th>
                <th class="text-center">回答时间</th>
            </tr>
            </thead>
            <tbody>
            {{if .Answers}}
            {{range $index, $answer := .Answers}}
            <tr class="text-center">
                <td>{{timeforhumnas $answer.CreatedAt}}</td>
                <td><a href="{{urlfor "QuestionController.Show" ":id" $answer.Question.Id}}">{{$answer.Question.Title}}</a></td>
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