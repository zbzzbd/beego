<div class="ui card p10 full">
    <h2 class="text-center">作业列表</h2> 
    <div style="overflow-x: auto;overflow-y: hidden" class="mb20">
        <table class="ui celled table" style="min-width: 1600px;">
            <thead>
            <tr>
                <th>序号</th>
                <th>作业状态</th>
                <th>作业编号</th>
                <th>项目名称</th>
                <th>业务单元</th>
                <th>作业要求</th>
                <th>作业对象</th>
                <th>作业部门</th>
                <th>作业单元</th>
                <th>发布时间</th>
                <th>审核时间</th>
                <th>要求完成时间</th>
                <th>实际完成时间</th>
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
                <td class="{{GetStatusColor $elem}}">{{GetStatusDesc $elem}}</td>
                <td>
                    <a href="/job/view/{{.Id}}">{{.Code}}</a>
                </td>
                <td>{{.Project.Name}}</td>
                <td>{{.CreateUser.Name}}</td>
                <td>{{.Type}}</td>
                <td>{{.Target}}</td>
                <td>{{.Department}}</td>
                <td>{{.Employee.Name}}</td>
                <td>{{TimeFormat .Created}}</td>
                <td>{{TimeFormat .ValidTime}}</td>
                <td> {{TimeFormat .FinishTime}} </td>
                <td>{{TimeFormat .SubmitTime}}</td>
                <td>{{GetTimeDiff .SubmitTime .ClaimTime}}</td>
                <td>{{GetTimeDiff .SubmitTime .FinishTime}}</td>
                <td style="min-width: 100px;">
                    <div><a class="ui red button" href="/job/complaint/new/{{.Id}}">投诉</a></div>
                </td>
            </tr>
            {{end}}
            </tbody>
        </table>
    </div>

    <div>
        <div class="ui right floated pagination menu">
            <a class="icon item" href="{{.paginator.PageLinkPrev}}">
                <i class="left chevron icon"></i>
            </a>

            {{range $index, $page := .paginator.Pages}}
            {{if $.paginator.IsActive .}}
            <a class="item active" href="{{$.paginator.PageLink $page}}">{{$page}}</a>
            {{else}}
            <a class="item" href="{{$.paginator.PageLink $page}}">{{$page}}</a>
            {{end}}
            {{end}} 
            <a class="icon item" href="{{.paginator.PageLinkNext}}">
                <i class="right chevron icon"></i>
            </a>
        </div> 
    </div>
</div>
