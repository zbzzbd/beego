{{template "job/find.tpl" .}}

<div class="ui card p10 full">
    <h2 class="text-center">作业列表</h2>
    <div style="overflow: auto"  class="mb20" >
    <table class="ui celled table">
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

            {{if not .IsReadonly}}
            <th class="text-center">操作</th>
            {{else}}
            <th class="text-center">状态</th>
            {{end}}
        </tr></thead>
        <tbody>
        {{range $index, $elem := .jobs}}
        <tr>
            <td>
                {{AddInt $index 1}}
            </td>
            <td> <a href="/job/view/{{.Id}}">{{.Code}}</a></td>
            <td>{{.Project.Name}}</td>
            <td>{{.Type}}</td>
            <td>{{.Target}}</td>
            <td>{{.Department}}</td>
            <td>{{.Employee.Name}}</td>
            <td>{{.CreateUser.Name}}</td>
            <td>{{.Message}}</td>

            <td class="text-center">
                <a class="ui blue button" href="/project/job/valid/{{.Id}}">审核</a>
            </td>
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
