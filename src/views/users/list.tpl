{{template "users/find.tpl" .}}
<div>
    <div class="ui card p10 full">
        <h2 class="text-center">
            用户列表
            <a href="/user/create" class="ui teal button" style="width: 100px!important;float: right;">新建用户</a>
        </h2>

        <div style="overflow-x: auto;overflow-y: hidden" class="mb20">
            <table class="ui celled table">
                <thead>
                <tr>
                    <th>序号</th>
                    <th>公司名称</th>
                    <th>名称</th>
                    <th>角色</th>
                    <th>邮箱</th>
                    <th>手机</th>
                    <th>操作</th>
                </tr></thead>
                <tbody>
                {{range $index, $elem := .users}}
                <tr>
                    <td>
                        {{AddInt $index 1}}
                    </td>
                    <td>{{.Company.Name}}</td>
                    <td>{{.Name}}</td>
                    <td>{{GetRoleDesc .Roles}}</td>
                    <td>{{.Email}}</td>
                    <td>{{.Mobile}}</td>
                    <td>
                        {{if TimeIsZero .Deleted }}
                            <a class="ui blue button" href="/user/edit/{{.Id}}">编辑</a>
                            <a class="ui red button" onclick="deleteUser('{{.Id}}')">删除</a>
                        {{else}}
                            <a class="ui green button" onclick="restoreUser('{{.Id}}')">恢复</a>
                        {{end}}

                    </td>
                </tr>
                {{end}}
                </tbody>
            </table>
        </div>
    </div>

    <div class="pb50">
        {{template "public/pager.tpl" .}}
    </div>
</div>

<div class="ui modal delete">
    <div class="header">
        用户删除
    </div>
    <div class="image content">
        <div class="description">
           <p>确定删除此用户?</p>
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