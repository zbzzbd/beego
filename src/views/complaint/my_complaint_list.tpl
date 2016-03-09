<!--{{template "complaint/complain_find.tpl" .}}-->
<div class="ui card p10 full">
    <h2 class="text-center">投诉列表</h2>
    <div style="overflow: auto" class="mb20">
    <table class="ui celled table">
        <thead>
        <tr>
            <th>序号</th>
            <th>项目名称</th>
            <th>作业编号</th>
            <th>作业要求</th>
            <th>作业部门</th>
            <th>登记人</th>
            <th>投诉来源</th>
            <th>投诉事项</th>
            <th>投诉内容</th>
            <th>投诉时间</th>
            <th>答复需求</th>
            {{if not .IsReadonly}}
            <th>操作</th>
            {{end}}
        </tr></thead>
        <tbody>
        {{range .complains}}
        <tr>
            <td>{{.Id}}</td>
            <td>{{.Project.Name}}</td>
            <td><a href="/job/view/{{.Job.Id}}">{{.Job.Code}}</a></td>
            <td>{{.Job.Type}}</td>
            <td>{{.Job.Department}}</td>
            <td>{{.User.Name}}</td>
            <td>{{.Source}}</td>
            {{if eq (printf "%s"  .Type) "1"}}
            <td>延时</td>
            {{else if eq (printf "%s"  .Type) "2"}}
            <td>逻辑错误</td>
            {{else if eq (printf "%s"  .Type) "3"}}
            <td>其它重大错误</td>
            {{end}}
            <td>{{.Complain}}</td>
            <td>{{TimeFormat .Created}}</td>
             
         
            {{if .Response}}
            <td>需要回复</td> 
            {{if not $.IsReadonly}}
            <td>{{if .ReplyStatus}}  <label class="ui reset button">已回复</label>  {{else}}  <div class="pb10"><a class="ui teal button" href="/produce/complaint/reply/{{.Id}}?jobid={{.Job.Id}}">回复</a></div> {{end}} </td>
            {{end}}
            {{else}}
            <td>不需要回复</td>
            <td></td>
            {{end}}
        
        </tr>

        {{end}}
        </tbody>
    </table>
    </div>
    <div>
        {{template "public/pager.tpl" .}}
    </div>
</div>
