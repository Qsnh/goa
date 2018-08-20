<div class="container">
    <div class="row">
        <div class="col-sm-9">
            <div class="title">{{.question.Title}}</div>
            <div class="info">
                <p>{{.question.Category.Name}} | {{.question.CreateAt}}</p>
            </div>
            <div class="description">
                {{.question.Description}}
            </div>
        </div>
        <div class="col-sm-3 text-center">
            <p>
                <img src="" width="50" height="50" class="img-circle">
            </p>
            <p>{{.question.User.Nickname}}</p>
        </div>
    </div>
</div>