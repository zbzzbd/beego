<div class="ui card p10 full">
    <h2 class="text-center">作业列表</h2>
    <div style="overflow: auto" class="mb20"> 
        <table class="ui celled table" >
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
               <th>要求完成时间</th> 
               {{if not .IsReadonly}}
               <th>操作</th>
               {{end}}
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
               <td> {{TimeFormat .FinishTime}}</td> 
               {{if not $.IsReadonly}}
               <td class="text-center" style="min-width: 100px;">
                   <a class="ui blue button" href="/produce/job/claim/{{.Id}}">认领</a><a class="ui green button" href="/produce/job/assign?id={{.Id}}">转发</a>
               </td>
               {{end}}
           </tr>
           {{end}}
           </tbody>
       </table>
    </div>

    <div>
        {{template "public/pager.tpl"}}
    </div>
</div>
