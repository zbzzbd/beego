<div class="ui card p10 full mb20">
    <h2 class="text-center">编辑用户</h2>
    <form class="ui form user-edit" >
        <input name="id" value="{{.user.Id}}" type="hidden" id="user-id">
        <div class="inline fields">
            <div class="field">
                <label>公司名称</label>
                <select class="ui dropdown" name="company" >
                    <option value="">请选择</option>
                    {{range .companies}}
                    {{if eq .Id $.user.Company.Id}}
                    <option selected="" value="{{.Id}}">{{.Name}}</option>
                    {{else}}
                    <option value="{{.Id}}">{{.Name}}</option>
                    {{end}}
                    {{end}}
                </select>
            </div>

            <div class="field">
                <label>用户角色</label>
                <select class="ui dropdown" name="role" id="department">
                    <option value="">请选择</option>
                    {{range .roles}}
                    {{if eq .Role $.user.Roles}}
                    <option selected="" value="{{.Role}}" tag="{{.Role}}">{{.Department}}</option>
                    {{else}}
                    <option value="{{.Role}}" tag="{{.Role}}">{{.Department}}</option>
                    {{end}}
                    {{end}}
                </select>
            </div>
        </div>

        <div class="inline fields">
            <div class="field">
                <label>用户名称</label>
                <input name="name" value="{{.user.Name}}" style="width: 194px;">
            </div>

            <div class="field">
                <label>用户邮箱</label>
                <input type="email" name="email" value="{{.user.Email}}" style="width: 194px;">
            </div>
        </div>

        <div class="inline fields">
            <div class="field">
                <label>用户手机</label>
                <input name="mobile" value="{{.user.Mobile}}" style="width: 194px;">
            </div>
        </div>

        <div class="ui error message">

        </div>

        <div class="field text-center">
            <button class="ui green button" type="submit">保存</button>
        </div>
    </form>
</div>
