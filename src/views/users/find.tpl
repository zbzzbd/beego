<div class="ui card p10 full mb20">
    <form class="ui form find" method="get" onchange="submitForm('.ui.form.find')" >
        <div class="inline fields">
            <div class="field">
                <label>公司名称</label>
                <select class="ui dropdown" name="company" >
                    <option value="">请选择</option>
                    {{range .companies}}
                    {{if eq .Code $.company}}
                    <option selected="" value="{{.Code}}">{{.Name}}</option>
                    {{else}}
                    <option value="{{.Code}}">{{.Name}}</option>
                    {{end}}
                    {{end}}
                </select>
            </div>

            <div class="field">
                <label>用户角色</label>
                <select class="ui dropdown" name="role" >
                    <option value="">请选择</option>
                    {{range .roles}}
                    {{if eq .Role $.role}}
                    <option selected="" value="{{.Role}}" tag="{{.Role}}">{{.Department}}</option>
                    {{else}}
                    <option value="{{.Role}}" tag="{{.Role}}">{{.Department}}</option>
                    {{end}}
                    {{end}}
                </select>
            </div>
            <div class="field">
                <label>用户名称</label>
                <select class="ui search dropdown" name="user_id" >
                    <option value="">请选择</option>
                    {{range .allUsers}}
                    {{if eq (printf "%d" .Id) $.user_id}}
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
                <label>用户邮箱</label>
                <input name="email" value="{{.email}}" style="width: 194px;">
            </div>
            <div class="field">
                <label>用户手机</label>
                <input name="mobile" value="{{.mobile}}" style="width: 194px;">
            </div>
        </div>

        <div class="field text-center">
            <a class="ui red button" onclick="clearForm('.ui.form.find')">清空</a>
        </div>
    </form>
</div>
