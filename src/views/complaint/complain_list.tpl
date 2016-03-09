{{template "complaint/complain_find.tpl" .}}
<div class="ui card p10 full">
    <h2 class="text-center">投诉列表
    <a id="job_complaint_export" class="ui teal button fr" href="">导出</a>
    </h2>

    <div style="overflow: auto"  class="mb20" >
    <table class="ui celled table">
        <thead>
        <tr>
            <th>序号</th>
            <th>投诉日期</th>
            <th>投诉来源</th>
            <th>作业编号</th>
            <th>项目名称</th>
            <th>业务单元</th>
            <th>作业部门</th>
            <th>作业单元</th>
            <th>客诉事项</th>
            <th>客诉描述</th> 
            <th>答复需求</th>
            <th>答复情况</th>
            <th>操作</th>
        </tr></thead>
        <tbody>
        {{range $index, $elem := .complains}}
        <tr>
            <td>{{AddInt $index 1}}</td>
            <td>{{date .Created "Y-m-d"}}</td>
            <td>{{.Source}}</td>
            <td><a href="/job/view/{{.Job.Id}}">{{.Job.Code}}</a></td>
            <td>{{.Project.Name}}</td>
            <td>{{.User.Name}}</td>
            <td>{{.Job.Department}}</td>
            <td>{{.Employee.Name}}</td>
            <td>
                {{if eq (printf "%s" .Type) "1"}}
                    延时
                {{else if eq (printf "%s" .Type) "2"}}
                    逻辑错误
                {{else if eq (printf "%s" .Type) "3"}}
                    其它重大错误
                {{end}}
            </td>
            <td>{{.Complain}}</td>
             {{if .Response}}
                <td>需要回复</td>
                 {{if .ReplyStatus}}
                    <td><label class="ui reset button" >已答复</label></td> 
                 {{else}}
                    <td> <label class="ui reset button">未答复</label></td>
                {{end}}
            {{else}}
                <td>不需要答复</td>
                <td></td>
            {{end}} 
            <td>
            {{if eq $.user_info.Id .User.Id}}
            <a class="ui blue button" href="/job/complaint/new/{{.Job.Id}}?editstatus=1">修改</a>
            <a class="ui red button"  onclick="deleteComplaint('{{.Id}}')">删除</a>
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

<div class="ui modal delete">
    <div class="header">
        投诉删除
    </div>
    <div class="image content">
        <div class="description">
           <p>确定删除此投诉?</p>
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
    window.onload = function() {
        $(function(){
            var str = window.location.href;
            if (str.indexOf('?') >0)
                $("#job_complaint_export").attr('href','/job/complaint/export' + str.substring(str.indexOf('?')))
            else {
                $("#job_complaint_export").attr('href','/job/complaint/export')
            }
        });
    }
</script>