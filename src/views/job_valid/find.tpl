<div class="ui card p10 full mb20">
    <form class="ui form find" method="get" onchange="submitForm('.ui.form.find')" >
        <div class="inline fields">
            <div class="field">
                <label>作业编号</label>
                <input name="code" value="{{.code}}" style="width: 194px;">
            </div>

            <div class="field">
                <label>业务单元</label>
                <select class="ui dropdown" name="create_user_id">
                    <option value="">请选择</option>
                    {{range .Employees}}
                    {{if and (eq (printf "%d" .Id) $.create_user_id) (eq .Roles $.BussinessMen)}}
                    <option selected="" value="{{.Id}}" >{{.Name}}</option>
                    {{else if eq .Roles $.BussinessMen}}
                    <option value="{{.Id}}" >{{.Name}}</option>
                    {{end}}
                    {{end}}
                </select>
            </div>

            <div class="field">
                <label>项目名称</label>
                <select class="ui search dropdown" name="project" id="project" >
                    <option value="">请选择</option>
                    {{range .ProjectNames}}
                    {{if eq .Name $.project__name}}
                    <option selected="" value="{{.Name}}">{{.Name}}</option>
                    {{else}}
                    <option value="{{.Name}}">{{.Name}}</option>
                    {{end}}
                    {{end}}
                </select>

            </div>
            <div class="field">
                <label>作业部门</label>
                <select class="ui dropdown" name="department" id="department">
                    <option value="">请选择</option>
                    {{range .Departments}}
                    {{if eq .Department $.department}}
                    <option selected="" value="{{.Department}}" tag="{{.Role}}">{{.Department}}</option>
                    {{else}}
                    <option value="{{.Department}}" tag="{{.Role}}">{{.Department}}</option>
                    {{end}}
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

            <div class="field">
                <label>作业状态</label>
                <select class="ui dropdown" name="status" id ="status">
                <option value="">请选择</option> 
                {{range  $elem:=.Status}}
                {{if eq $elem  $.status}}
                 <option selected ="" value="{{$elem}}"> {{$elem}}</option>
                 {{else}}
                 <option value="{{$elem}}">{{$elem}}</option>
                 {{end}}
                 {{end}}
                </select>
            </div>
        </div>

        <div class="inline fields">
            <div class="field">
                <label>实际完成时间</label>
                <input class="datetimepicker" name="submit_time_from" value="{{.submit_time__gte}}">
                <label>至</label>
                <input class="datetimepicker" name="submit_time_to" value="{{.submit_time__lte}}">
            </div>

            <div class="field">
                <label>要求完成时间</label>
                <input class="datetimepicker" name="finish_time_from" value="{{.finish_time__gte}}">
                <label>至</label>
                <input class="datetimepicker" name="finish_time_to" value="{{.finish_time__lte}}">
            </div>
        </div>



            
        <div class="field text-center">
            <!--<button class="ui submit green button">查找</button>-->
            <a class="ui red button" onclick="clearForm('.ui.form.find')">清空</a>
        </div>

        <div class="ui error message"></div>
    </form>
</div>
