<div class="ui card p10 full">
    <form class="ui form" method="get">
        <div class="inline fields">
            <div class="field">
                <lable>项目名称:</lable>
                <select class="ui search dropdown" name="project_name">
                    {{if .project_default}}
                    <option value={{.project_default.Id}}>{{.project_default.Name}}</option>
                    <option value="">全部</option>
                    {{else}}
                    <option value="">全部</option>
                    {{end}}
                    {{range .AllProjects}}
                    <option value="{{.Id}}">{{.Name}}</option>
                    {{end}}
                </select>
            </div>
            <span style="width:50px;"></span>
            <div class="field">
                <lable>业务担当:</lable>
                <select class="ui dropdown" name="bussiness_user">
                    {{if .bussiness_user_default}}
                    <option value={{.bussiness_user_default.Id}}>{{.bussiness_user_default.Name}}</option>
                    <option value="">全部</option>
                    {{else}}
                    <option value="">全部</option>
                    {{end}}
                    {{range .BussinessUser}}
                    <option value="{{.Id}}">{{.Name}}</option>
                    {{end}}
                </select>
            </div>
        </div>

        <div class="inline fields">
            <div class="field">
                <lable>项目启动日期:</lable>
                <input class="datetimepicker" type=text readonly value="{{.start_date_default}}" name="start_date">
            </div>
            <div class="field">
                <lable>至</lable>
            </div>
            <div class="field">
                <input class="datetimepicker" type=text readonly value="{{.end_date_default}}" name="end_date">
            </div>

            <span style="width:50px;"></span>

            <div class="field">
                <lable>进程现况:</lable>
                <select class="ui dropdown" name="progress">

                    {{if .progress_default}}
                    <option value={{.progress_default.Id}}>{{.progress_default.Name}}</option>
                    <option value="">全部</option>
                    {{else}}
                    <option value="">全部</option>
                    {{end}}
                    {{range .Progress}}
                    <option value="{{.Id}}">{{.Name}}</option>
                    {{end}}
                </select>
            </div>
        </div>

        <div class="field text-center">
            <button class="ui submit green button">搜索</button>
        </div>

        <div class="ui error message"></div>
    </form>
</div>

<div class="ui card p10 full">
    <h2 class="text-center">项目列表
      <a  id="project_export" class="ui teal button fr"  href="">导出</a>
    </h2>
    
    <div class="mb10" style="overflow-x: auto;overflow-y: hidden;">
        <table class="ui striped table" style="min-width: 1200px;">
            <thead>
            <th>序号</th>
            <th>启动时间</th>
            <th>项目名称</th>
            <th>服务项目</th>
            <th>合同编号</th>
            <th>赛事规模</th>
            <th>客户名称</th>
            <th>业务担当</th>
            <th>比赛日期</th>
            <th>开通报名时间</th>
            <th>关闭报名时间</th>
            <th>进程现况</th>
            <th>美术单元</th>
            <th>技术单元</th>
            <th>优先级别</th>
            <th>登记人</th>
            <th>操作</th>
            </thead>
            <tbody>
            {{range $index, $p := .Projects}}
            <tr>
                <td>{{AddInt $index 1}}</td>
                <td>
                    {{TimeFormat $p.Started}}
                </td>
                <td>{{$p.Name}}</td>
                <td>{{$p.ServiceItem}}</td>
                <td>{{$p.ContractNo}}</td>
                <td>{{$p.Scale}}</td>
                <td>{{$p.ClientName}}</td>
                <td>{{$p.BussinessUser.Name}}</td>
                <td>{{date $p.GameDate "Y-m-d"}}</td>
                <td>{{TimeFormat $p.RegStartDate}}</td>
                <td>{{TimeFormat $p.RegCloseDate}}</td>
                <td>{{$p.Progress.Name}}</td>
                <td>{{$p.ArtUser.Name}}</td>
                <td>{{$p.TechUser.Name}}</td>
                <td>{{$p.Priority}}</td>
                <td>{{$p.Registrant.Name}}</td>
                {{if eq $.user_info.Id $p.Registrant.Id}}
                <td><a class="ui blue button" href="/project/edit/{{$p.Id}}">编辑</a> ｜<a onclick="delProject({{$p.Id}})">删除</a></td>

                {{else if eq $.user_info.Id $p.BussinessUser.Id}}
                <td><a class="ui blue button" href="/project/edit/{{$p.Id}}">编辑</a></td>
                {{else}}
                <td><a class="ui teal button" href="/project/detail/{{$p.Id}}">查看</a></td>
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

<div class="ui modal delete">
    <div class="header">项目删除</div>
    <div class="image content"> 
        <div class="description">
            <p>您确定删除此项目？</p>
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
    window.onload =function(){
        $(function(){
            var str = window.location.href;
            if (str.indexOf('?') >0)
            $("#project_export").attr('href','/project/export' + str.substring(str.indexOf('?')))
            else {
                $("#project_export").attr('href','/project/export')
            }
        });
    }
</script>
