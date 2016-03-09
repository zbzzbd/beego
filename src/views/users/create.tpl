<div class="ui card p10 full mb20">
    <h2 class="text-center">创建用户</h2>
    <form class="ui form user" >
        <div class="inline fields">
            <div class="field">
                <label>公司名称</label>
                <select class="ui dropdown" name="company" >
                    <option value="">请选择</option>
                    {{range .companies}}
                        <option value="{{.Id}}">{{.Name}}</option>
                    {{end}}
                </select>
            </div>

            <div class="field">
                <label>用户角色</label>
                <select class="ui dropdown" name="role" id="department">
                    <option value="">请选择</option>
                    {{range .roles}}
                    <option value="{{.Role}}" tag="{{.Role}}">{{.Department}}</option>
                    {{end}}
                </select>
            </div>
        </div>

        <div class="inline fields">
            <div class="field">
                <label>用户名称</label>
                <input name="name" style="width: 194px;">
            </div>

            <div class="field">
                <label>用户邮箱</label>
                <input type="email" name="email" style="width: 194px;">
            </div>
        </div>

        <div class="inline fields">
            <div class="field">
                <label>用户手机</label>
                <input name="mobile" style="width: 194px;">
            </div>
        </div>

        <div class="ui error message">

        </div>

        <div class="field text-center">
            <button class="ui green button" type="submit">创建</button>
        </div>
    </form>
</div>
