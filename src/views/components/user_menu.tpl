 
<div class="clearfix text-center" style="font-size: 14px;">
    {{if .return}}
        <div class="fl">
        <a onclick="history.back();">返回</a>
        </div>
    {{end}}  
    <div class="fr">
        {{if .no_login}}
        <a href="/login">登录</a>
        {{else}}

        <label class="pr20">{{.user_info.Company.Name}}</label>
        <label class="pr20" style="color: #20ad6e;">{{.user_info.Name}}</label>
        <a href="/modify_password">修改密码</a>
        <a href="/logout">退出</a>
        {{end}}
    </div>
</div>
