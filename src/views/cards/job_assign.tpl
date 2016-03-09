
<div class="ui card p10 full">
    <h4 class="ui teal header">作业转发</h4>
    <div class="ui grid p10">
        <div class="row">
            <label>任务类型：</label>
            {{.Job.Type}}
        </div>

        <div class="row">
            <label>转发人员：</label>
            {{.FromUser.Name}}
        </div>
        <div class="row">
            <label>转发到：</label>
            {{.ToUser.Name}}
        </div>
        <div class="row">
            <label>备注：</label>
            {{.Remark}}
        </div>
    </div>
    <div >
        <label class="ui teal right ribbon label">转发时间：{{.Created}}</label>
    </div>
</div>
