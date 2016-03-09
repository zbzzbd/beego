   <div class="ui card p10 full">
        <form class="ui form find" method="get">
            <div class="inline fields">
                <div class="field">
                    <label>作业编号</label>
                    <input name="code" value="{{.job__code}}">
                </div>

                <div class="field">
                <label>项目名称</label>
                <select class="ui search dropdown" name="project" id="project" >
                    <option value="">请选择</option>
                    {{range .Projects}}
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
                             {{if eq .Department  $.job__department}}
                            <option selected="" value="{{.Department}}" tag="{{.Role}}">{{.Department}}</option>
                            {{else}}
                            <option value="{{.Department}}" tag="{{.Role}}">{{.Department}}</option>
                            {{end}}
                        {{end}}
                    </select>
                </div>
                <div class="field">
                    <label>答复需求</label>
                    <select class="ui search dropdown" name="response" >
                        <option value="">请选择</option>
                        {{range $elem := .Response}}
                            {{if eq  $elem "1" }}
                                {{if eq  $.response "1" }}                                
                                    <option selected="" value="{{$elem}}">需要答复</option>
                                {{else}}
                                    <option value="{{$elem}}">需要答复</option>
                                {{end}}                           
                            {{else}}
                                {{if eq  $.response "0" }}                                
                                        <option selected="" value="{{$elem}}">不需要答复</option>
                                    {{else}}
                                        <option value="{{$elem}}">不需要答复</option>
                                    {{end}}     
                                {{end}}
                        {{end}}
                    </select>
                </div>
            </div>

            <div class="inline fields">
                <div class="field disabled">
                    <label>作业单元</label>
                    <select class="ui search dropdown" name="employee_id" id="employee">
                        <option value="">请选择</option>
                        {{range .Employees}}
                            <option value="{{.Id}}" tag="{{.Roles}}">{{.Name}}</option>
                        {{end}}
                    </select>
                </div>

                <div class="field">
                    <label>投诉时间</label>
                    <input class="datetimepicker" id="datetimepicker-end-time-from" name="end_time_from" value="{{.end_time__gte}}">
                    <label>至</label>
                    <input class="datetimepicker" id="datetimepicker-end-time-to" name="end_time_to" value="{{.end_time__lte}}">
                </div>
            </div>

            <div class="field text-center">
                <button class="ui submit green button">查找</button>
                <a class="ui reset button" href="/job">清空</a>
            </div>

            <div class="ui error message"></div>
        </form>
    </div>
