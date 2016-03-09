<!doctype html>
<html lang="zh-cn">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>{{.title}}</title>
    <meta name="Copyright" content="CHINARUN" />
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="shortcut icon" href="/img/favicon.ico" />
    <link rel="stylesheet" href="/public/css/semantic.min.css" />
    <link rel="stylesheet" href="/public/css/main.min.css" />
</head>
<body>
<div class="body" style="height: 100%;">
    <div class="container">
        {{if .NeedTemplate}}
            {{.LayoutMenu}}
            <div class="page">
                <div class="header">
                    {{template "components/user_menu.tpl" .}}
                </div>
                <div class="p30">
                    {{if .RoleError}}
                        {{template "role/error.tpl" .}}
                    {{else}}
                        {{.LayoutContent}}
                    {{end}}
                </div>
            </div>

            {{.Footer}}
        {{else}}
            {{.LayoutContent}}
        {{end}}
    </div>
</div>

    <script src="/public/js/lib.min.js"></script>
    <script src="/public/js/semantic.min.js"></script>

    {{if .IsDev}}
    <script src="/public/js/app.js"></script>
    {{else}}
    <script src="/public/js/app.min.js"></script>
    {{end}}

    {{if .IsDev}}
    <script src="/public/js/dev.js"></script>
    {{end}}
</body>
</html>
