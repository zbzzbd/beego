{{template "job_valid/find.tpl" .}}

<div class="pb20">
    <div class="ui card p10 full">
    <h2 class="text-center">
        删除历史列表
        <a id="export" class="ui teal button fr"  href="/project/job/export?{{.path}}">导出</a>
    </h2>

        <div style="overflow: auto"  class="mb20" >
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
               <th>作业认领时间</th>
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
               <td>{{TimeFormat .ClaimTime}}</td>
               <td>{{TimeFormat .SubmitTime}}</td>
               <td>{{GetTimeDiff .SubmitTime .ClaimTime}}</td>
               <td>{{GetTimeDiff .SubmitTime .FinishTime}}</td>
               <td><a class="ui teal button" onclick="recoverJob({{.Id}})">恢复</a></td> 
           </tr>
           {{end}}
           </tbody>
       </table>
    </div>

    <div>
        {{template "public/pager.tpl" .}}
    </div>
</div>

<div class="ui modal delete">
    <div class="header">作业删除</div>
    
    <div class="image content">
      <div class="description">
        <p>确定恢复此作业?</p>
      </div> 
    </div>
    <div class="actions">
        <div class="ui green approve button">
          确定
        </div>
        <div class="ui red deny button">
          取消
        </div>
    </div>
</div>
<script>
    window.onload=function(){
        $(function(){
          var str = window.location.href;
          if (str.indexOf('?')>0)
          $("#export").attr('href', '/project/job/export' + str.substring(str.indexOf('?')) +'&delstatus=1');
          else {
            $("#export").attr('href', '/project/job/export'+'?delstatus=1');
          }           
        });
    }
</script>