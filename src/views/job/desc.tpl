<div class="ui card p10 full mb20">
    <h4 class="ui teal header">作业描述</h4>
    <div class="ui grid p10">
        <div class="row">
            <div class="nine wide column">
                <label>项目名称：</label>
                {{.Job.Project.Name}}
            </div>
            <div class="five wide column">
                <label>作业编号：</label>
                {{.Job.Code}}
            </div>
        </div>
        <div class="row">
            <div class="nine wide column">
                <label>作业要求：</label>
                {{.Job.Type}}
            </div>
            <div class="five wide column">
                <label>作业部门：</label>
                {{.Job.Department}}
            </div>
        </div>
        <div class="row">
            <div class="nine wide column">
                <label>作业对象：</label>
                {{.Job.Target}}
            </div>
            <div class="five wide column">
                <label>作业单元：</label>
                {{.Job.Employee.Name}}
            </div>
        </div>
        <div class="row">
            <div class="nine wide column">
                <label>修改地址：</label>
                {{.Job.TargetUrl}}
            </div>
            <div class="five wide column">
                <label>验收时间：</label>
                {{TimeFormat .Job.FinishTime}}
            </div>
        </div>
        <div class="row">
            <div class="three wide column pr0" style="width: 85px!important;">
                <label>作业描述：</label>
            </div>
            <div class="thirteen wide column pl0">
                <pre class="m0">{{.Job.Desc}}</pre>
            </div>
        </div>
        <div class="row">
            <div class="sixteen wide column">
                <label>业务留言：</label>
                {{.Job.Message}}
            </div>
        </div>

        <div class="row">
            <div class="sixteen wide column">
                <label>作业附件：</label>
                {{range .JobFiles}}
                <a href="{{.Url}}" class="pr20" download="{{.Name}}" target="_blank">{{.Name}}</a>
                {{end}}
            </div>
        </div>
    </div>
    <div >
        <label class="ui teal right ribbon label">发表于{{TimeFormat .Job.Updated}}</label>
    </div>
</div>
