<div class="ui card p10 full">
    <h2 class="text-center">作业列表</h2>
    <div style="overflow: auto" class="mb20"> 
        <table class="ui celled table" style="min-width: 1600px;">
            <thead>
            <tr>
                <th>序号</th>
                <th>作业编号</th>
                <th>项目名称</th>
                <th>作业要求</th>
                <th>作业对象</th>
                <th>作业部门</th>
                <th>作业单元</th>
                <th>业务单元</th>
                <th>业务留言</th>
                <th>作业发布时间</th>
                <th>作业认领时间</th>
                <th>要求完成时间</th>
                <th>任务用时／Min</th>
                <th>任务延时／Min</th>
                <th>操作</th>
            </tr></thead>
            <tbody>
            {{range $index, $elem := .jobs}}
            <tr>
                <td>
                    {{AddInt $index 1}}
                </td>
                <td>
                    <a href="/job/view/{{.Id}}">{{.Code}}</a>
                </td>
                <td>{{.Project.Name}}</td>
                <td>{{.Type}}</td>
                <td>{{.Target}}</td>
                <td>{{.Department}}</td>
                <td>{{.Employee.Name}}</td>
                <td>{{.CreateUser.Name}}</td>
                <td>{{.Message}}</td>
                <td>{{TimeFormat .Created}}</td>
                <td>{{TimeFormat .ClaimTime}}</td>
                <td> {{TimeFormat .FinishTime}} </td>
                <td>{{GetTimeDiff .SubmitTime .ClaimTime}}</td>
                <td>{{GetTimeDiff .SubmitTime .FinishTime}}</td>
                <td style="min-width: 100px;">
                    {{if eq (printf "%d" .SubmitStatus) "0"}}
                    <a class="ui purple button" href="/produce/job/submit/{{.Id}}?status=1">提交</a>
                    {{else if eq (printf "%d" .SubmitStatus) "2"}}
                    <a class="ui purple button" href="/produce/job/submit/{{.Id}}?status=1">再次提交</a>
                    {{end}}
                </td>
            </tr>
            {{end}}
            </tbody>
        </table> 
    </div>

    <div>
        {{template "public/pager.tpl"}}
    </div>
</div>
