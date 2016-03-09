<div class="ui card p10 full">
    <form class="ui form find" method="get">
        <div class="inline fields">
            <div class="field">
                <label>作业编号</label>
                <input name="code" value="{{.code}}">
            </div>
        </div>

        <div class="inline fields">
            <div class="field">
                <label>项目名称</label>
                <input name="project"  value="{{.project__name}}">
            </div>
            <div class="field">
                <label>作业部门</label>
                <select class="ui dropdown" name="department" id="department">
                    <option value="">请选择</option>
                    {{range .Departments}}
                    <option value="{{.Department}}" tag="{{.Role}}">{{.Department}}</option>
                    {{end}}
                </select>
            </div>
        </div>

        <div class="inline fields">
            <div class="field">
                <label>作业要求</label>
                <select class="ui search dropdown" name="type" >
                    <option value="">请选择</option>
                    {{range $elem := .Types}}
                    {{if eq $elem $.type}}
                    <option selected="" value="{{$elem}}">{{$elem}}</option>
                    {{else}}
                    <option value="{{$elem}}">{{$elem}}</option>
                    {{end}}
                    {{end}}
                </select>
            </div>

            <div class="field disabled">
                <label>作业单元</label>
                <select class="ui search dropdown" name="employee_id" id="employee">
                    <option value="">请选择</option>
                    {{range .Employees}}
                    {{if eq (printf "%d" .Id) $.employee_id}}
                    <option selected="" value="{{.Id}}" tag="{{.Roles}}">{{.Name}}</option>
                    {{else}}
                    <option value="{{.Id}}" tag="{{.Roles}}">{{.Name}}</option>
                    {{end}}
                    {{end}}
                </select>
            </div>
        </div>

        <div class="inline fields">
            <div class="field">
                <label>完成时间</label>
                <input id="datetimepicker-end-time-from" name="finish_time_from" value="{{.end_time__gte}}">
                <label>至</label>
                <input id="datetimepicker-end-time-to" name="end_time_to" value="{{.end_time__lte}}">
            </div>
        </div>

        <div class="field text-center">
            <button class="ui submit green button">查找</button>
            <a class="ui red button" href="{{.ClearUrl}}">清空</a>
        </div>

        <div class="ui error message"></div>
    </form>
</div>

<script>
    window.onload=function(){
        $(function(){
            var datatimeOpt = {format:'Y-m-d'};
            $('#datetimepicker-end-time-from').datetimepicker(datatimeOpt);
            $('#datetimepicker-end-time-to').datetimepicker(datatimeOpt);
        });
    }
</script>