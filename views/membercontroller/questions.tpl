<div class="card">
    <div class="card-header">
        我的问题
    </div>
    <div class="card-body">
        <table class="table table-hover table-borderless">
            <thead>
            <tr>
                <th class="text-center">时间</th>
                <th class="text-center">标题</th>
            </tr>
            </thead>
            <tbody>
            {{if .Questions}}
            {{range $index, $question := .Questions}}
            <tr class="text-center">
                <td>{{$question.CreatedAt}}</td>
                <td><a href="{{urlfor "QuestionController.Show" ":id" $question.Id}}">{{$question.Title}}</a></td>
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