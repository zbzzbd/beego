<div class="ui card p10 full">
    <h4 class="ui teal header">{{.Label}}</h4>
    <div class="ui grid p10">
        <div class="row">
            <div class="nine wide column">
                <label>项目名称：</label>
                {{.Project.Name}}
            </div>
            <div class="five wide column">
                <label>作业编号：</label>
                {{.Code}}
            </div>
        </div>
        <div class="row">
            <div class="nine wide column">
                <label>作业要求：</label>
                {{.Type}}
            </div>
            <div class="five wide column">
                <label>作业部门：</label>
                {{.Department}}
            </div>
        </div>
        <div class="row">
            <div class="nine wide column">
                <label>作业对象：</label>
                {{.Target}}
            </div>
            <div class="five wide column">
                <label>作业单元：</label>
                {{.Employee.Name}}
            </div>
        </div>
        <div class="row">
            <div class="nine wide column">
                <label>修改地址：</label>
                <a href="{{.TargetUrl}}" target="_blank">{{.TargetUrl}}</a>
            </div>
            <div class="five wide column">
                <label>验收时间：</label>
                {{.FinishTime}}
            </div>
        </div>
        <div class="row">
            <div class="three wide column pr0" style="width: 90px!important;">
                <label>作业描述：</label>
            </div>
            <div class="thirteen wide column pl0">
                <pre class="m0">{{.Desc}}</pre>
            </div>
        </div>
        <div class="row">
            <div class="sixteen wide column">
                <label>业务留言：</label>
                {{.Message}}
            </div>
        </div>

        <div class="row">
            <div class="sixteen wide column">
                <label>作业附件：</label>
                {{range .Files}}
                <a href="{{.Url}}" class="pr20" download="{{.Name}}" target="_blank">{{.Name}}</a>
                {{end}}
            </div>
        </div>

    </div>
    <div >
        {{if .IsCreate}}
        <label class="ui teal right ribbon label">创建时间：{{.Updated}}</label>
        {{else}}
        <label class="ui teal right ribbon label">修改时间：{{.Updated}}</label>
        {{end}}
    </div>
</div>
