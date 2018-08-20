<div class="container">
    <div class="row">
        <div class="col-sm-9 question-box">
            {{range $index, $question := .questions}}
            <div class="question-item">
                <a href="{{urlfor "QuestionController.Show" ":id" $question.Id}}">
                    <img src="https://ps.ssl.qhimg.com/t013658e41e8c191970.jpg" class="img-circle" width="50" height="50" alt="nickname">
                    <span>{{$question.Title}}</span>
                </a>
            </div>
            {{end}}

            <div>
            {{str2html .paginator}}
            </div>
        </div>
        <div class="col-sm-3 right-section">
            <div class="row">
                <div class="col-sm-12 mt-10">
                    <a href="{{urlfor "QuestionController.Create"}}" class="btn btn-block btn-success">我要提问</a>
                </div>
                <div class="col-sm-12 data-statistic">
                    <p>问题总量：<b>121212</b>个</p>
                    <p>已解决：<b>12545</b>个</p>
                    <p>会员总量：<b>21155</b>位</p>
                </div>
                <div class="col-sm-12 mt-10 register-member">
                    <h5>最新注册</h5>
                    <ul>
                        <li><img src="https://ps.ssl.qhimg.com/t013658e41e8c191970.jpg" class="img-circle" width="50" height="50" alt="nickname"></li>
                        <li><img src="https://ps.ssl.qhimg.com/t013658e41e8c191970.jpg" class="img-circle" width="50" height="50" alt="nickname"></li>
                        <li><img src="https://ps.ssl.qhimg.com/t013658e41e8c191970.jpg" class="img-circle" width="50" height="50" alt="nickname"></li>
                        <li><img src="https://ps.ssl.qhimg.com/t013658e41e8c191970.jpg" class="img-circle" width="50" height="50" alt="nickname"></li>
                        <li><img src="https://ps.ssl.qhimg.com/t013658e41e8c191970.jpg" class="img-circle" width="50" height="50" alt="nickname"></li>
                    </ul>
                </div>
            </div>
        </div>
    </div>
</div>