<div class="container">
    <div class="row">
        <div class="col-sm-12">
            {{if .flash.notice}}
            <div class="alert alert-success">
                <p>{{.flash.notice}}</p>
            </div>
            {{end}}

            {{if .flash.error}}
            <div class="alert alert-danger">
                <p>{{.flash.error}}</p>
            </div>
            {{end}}

            {{if .flash.warning}}
            <div class="alert alert-warning">
                <p>{{.flash.warning}}</p>
            </div>
            {{end}}
        </div>
    </div>
</div>