<div class="pb20">
    <div class="ui card p10 full">
        <h2 class="text-center">我的项目</h2>
        <div style="overflow: auto"  class="mb20" >
        <table class="ui celled table">
            <thead>
            <tr>
                <th>序号</th>
                <th>项目名称</th>
                <th>项目开始时间</th>
                <th>客户名称</th>
                <th>项目进程</th>
                <th>业务担当</th>
                <th>美术担当</th>
                <th>技术担当</th>
                <th>登记人</th>
                <th>作业数</th>
            </tr></thead>
            <tbody>
            {{range .Projects}}
            <tr>
                <td>
                    <a href="/project/detail/{{.Id}}">{{.Id}}</a>
                </td>
                <td>{{.Name}}</td>
                <td>{{dateformat .Started "2006-01-02 15:04:05"}}</td>
                <td>{{.ClientName}}</td>
                <td>{{.Progress.Name}}</td>
                <td>{{.BussinessUser.Name}}</td>
                <td>{{.ArtUser.Name}}</td>
                <td>{{.TechUser.Name}}</td>
                <td>{{.Registrant.Name}}</td>
                <td>{{.JobsNum}}</td>
            </tr>
            {{end}}
            </tbody>
            <tfoot>
            <tr>
                <th colspan="100">
                    {{template "public/pager.tpl" .}}
                </th>
            </tr></tfoot>
            
        </table>
        </div>
    </div>

            {{if .error}}
               <h3> 服务器异常</h3>
            {{end}}
</div>
