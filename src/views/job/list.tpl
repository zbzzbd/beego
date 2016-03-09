<div class="ui card p10 full">
    <h2 class="text-center">作业列表</h2> 
    <div style="overflow-x: auto;overflow-y: hidden" class="mb20">
        <table class="ui celled table">
           <thead>
           <tr>
               <th>序号</th>
               <th>发布时间</th>
               <th>作业编号</th>
               <th>业务单元</th>
               <th>项目名称</th> 
               <th>作业要求</th>
               <th>作业对象</th>
               <th>作业部门</th>
               <th>作业单元</th>
               <th>要求完成时间</th>
               <th>实际完成时间</th>
 
               <th>业务留言</th>
               <th>操作</th> 
           </tr></thead>
           <tbody>
           {{range $index, $elem := .jobs}}
           <tr>
               <td>
                  {{AddInt $index 1}}
               </td>
               <td>{{TimeFormat .Created}}</td>
               <td>
                   <a href="/job/view/{{.Id}}">{{.Code}}</a>
               </td>
               <td>{{.CreateUser.Name}}</td>
               <td>{{.Project.Name}}</td> 
               <td>{{.Type}}</td>
               <td>{{.Target}}</td>
               <td>{{.Department}}</td>
               <td>{{.Employee.Name}}</td>
               <td> {{TimeFormat .FinishTime}}</td>
               <td>{{TimeFormat .SubmitTime}}</td>
               <td>{{.Message}}</td>
               <td>
                   {{if and (eq .CreateUser.Id $.user_info.Id) }}
                        {{if eq (printf "%d" .SubmitStatus) "1"}}
                        <a class="ui purple button" href="/produce/job/submit/{{.Id}}?status=2">验收</a>
                       {{else if eq (printf "%d" .ValidStatus) "2"}}
                            <a class="ui orange button" href="/job/edit/{{.Id}}">审核不通过 － 修改</a>
                       {{else if eq (printf "%d" .ClaimStatus) "2"}}
                            <a class="ui orange button" href="/job/edit/{{.Id}}">拒绝认领 － 编辑</a>
                        {{else if eq (printf "%d" .ValidStatus) "0"}}
                            <a class="ui blue button" href="/job/edit/{{.Id}}">编辑</a>
                       {{end}}
                   {{end}}
               </td> 
           </tr>
           {{end}}
           </tbody>
       </table>
    </div>

    <div>
        {{template "public/pager.tpl" .}}
    </div>
</div>
