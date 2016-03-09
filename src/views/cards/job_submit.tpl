<div class="ui card p10 full">
    <h4 class="ui teal header">{{.Label}}</h4>
    <div class="ui grid p10">
        <div class="row">
            <label>提交人员：</label>
            {{.ProduceUser.Name}}
        </div>
        <div class="row">
            <label>提交结果：</label>
            {{.Status}}
        </div>
        <div class="row">
            <label>附加说明：</label>
            {{.Remark}}
        </div>
        <div class="row">
            <label>提交附件：</label>
            {{range .Files}}
            <a href="{{.Url}}" class="pr20" download="{{.Name}}" target="_blank">{{.Name}}</a>
            {{end}}
        </div>
    </div>
    <div >
        <label class="ui teal right ribbon label">提交时间：{{.Created}}</label>
    </div>
</div>
