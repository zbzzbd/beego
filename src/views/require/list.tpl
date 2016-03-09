{{template "require/find.tpl" .}}
<div>
    <div class="ui card p10 full">
        <h2 class="text-center">
            作业要求列表
            <a onclick="addRequire()" class="ui teal button" style="width: 130px!important;float: right;">新建作业要求</a>
        </h2>

        <div style="overflow-x: auto;overflow-y: hidden" class="mb20">
            <table class="ui celled table">
                <thead>
                <tr>
                    <th>序号</th>
                    <th>名称</th>
                    <th>操作</th>
                </tr></thead>
                <tbody id="require-list">
                {{range $index, $elem := .requires}}
                <tr>
                    <td>
                        {{AddInt $index 1}}
                    </td>
                    <td class="name">{{$elem}}</td>
                    <td>
                        <a class="ui blue button" onclick="editRequire('{{$elem}}')">编辑</a>
                        <a class="ui red button" onclick="deleteRequire('{{$elem}}')">删除</a>
                    </td>
                </tr>
                {{end}}
                </tbody>
            </table>
        </div>
    </div>
</div>


<div class="ui modal delete">
    <div class="header">
        作业要求删除
    </div>
    <div class="image content">
        <div class="description">
            <p>确定删除此作业要求?</p>
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

<div class="ui modal edit">
    <div class="header">
        作业要求
    </div>
    <div class="content">
        <div class="ui input">
            <label style="margin-top: 10px;margin-right: 10px;">新作业要求名称: </label>
            <input id="modal-input">
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