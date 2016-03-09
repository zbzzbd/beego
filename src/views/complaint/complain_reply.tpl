<div class="pb20" >
    {{if eq .error ""}}
    <div class="ui card p10 full">
        <h4 class="ui teal header">客户投诉</h4>
        <div class="ui grid p10">
            <div class="row">
                <div class="nine wide column">
                    <label>业务担当：</label>
                    {{.complain.User.Name}}
                </div>
                <div class="five wide column">
                    <label>客诉事项：</label>
                    {{if eq  (printf  "%d"  .complain.Type)  "1"}}
                    <label>延时</label>
                    {{else if eq (printf "%d" .complain.Type) "2" }}
                    <label>逻辑错误</label>
                    {{else}}
                    <label>其它重大错误</label>
                    {{end}}
                </div>
            </div> 
            <div class="row">
                
                <div class="nine wide column">
                    <label>客诉描述：</label>
                    {{.complain.Complain}}
                </div>

            </div>
        </div>
        <div>
            <label class="ui teal right ribbon label">发表于{{TimeFormat .complain.Updated}}</label>
        </div>
    </div>
    {{else}}
    <div class="ui p10 full red message">
        错误： {{.error}}
    </div>
    {{end}}
</div>

<div class="pb20" >
    {{if eq .error ""}}
    <div class="ui card p10 full">
        <h4 class="ui teal header">作业描述</h4>
        <div class="ui grid p10">
            <div class="row">
                <div class="nine wide column">
                    <label>项目名称：</label>
                    {{.complain.Project.Name}}
                </div>
                <div class="five wide column">
                    <label>作业编号：</label>
                    {{.complain.Job.Code}}
                </div>
            </div>
            <div class="row">
                <div class="nine wide column">
                    <label>作业要求：</label>
                    {{.complain.Job.Type}}
                </div>
                <div class="five wide column">
                    <label>作业部门：</label>
                    {{.complain.Job.Department}}
                </div>
            </div>
            <div class="row">
                <div class="nine wide column">
                    <label>作业对象：</label>
                    {{.complain.Job.Target}}
                </div>
                <div class="five wide column">
                    <label>作业单元：</label>
                    {{.complain.Job.Employee.Id}}
                </div>
            </div>
            <div class="row">
                <div class="nine wide column">
                    <label>修改地址：</label>
                    {{.complain.Job.TargetUrl}}
                </div>
                <div class="five wide column">
                    <label>验收时间：</label>
                    {{.complain.Job.FinishTime}}
                </div>
            </div>
            <div class="row">
                <div class="sixteen wide column pr0" style="width: 85px!important;">
                    <label>作业描述：</label>
                    {{.complain.Job.Desc}}
                </div>
                <div class="thirteen wide column pl0">
                    <pre class="m0"> </pre>
                </div>
            </div>
            <div class="row">
                <div class="sixteen wide column">
                    <label>业务留言：</label>
                    {{.complain.Job.Message}}
                </div>
            </div>

            <div class="row">
                <div class="sixteen wide column">
                    <label>作业附件：</label>
                    {{range .JobFiles}}
                    <a href="{{.Url}}" class="pr20" download="{{.Name}}" target="_blank">{{.Name}} </a>
                    {{end}}
                </div>
            </div>
        </div>
        <div >
            <label class="ui teal right ribbon label">发表于 {{TimeFormat  .complain.Job.Created}} </label>
        </div>
    </div>
    {{else}}
    <div class="ui p10 full red message">
        错误： {{.error}}
    </div>
    {{end}}
</div>

<div class="pb20" >
    <div class="ui card p10 full">
        <form class="ui form" enctype ="multipart/form-data">
            <div class="inline fields">
                <div class=" field">
                    <label>上传附件</label>
                    <input type="file" name="files[]" multiple="multiple">
                </div>
                <div class=" field">
                    {{range .JobFiles}}
                        <span class="pr20">
                            <a href="{{.Url}}"  download="{{.Name}}" target="_blank">{{.Name}}</a>
                            <i onclick="delJobFile({{.Id}})" class="remove icon"></i>
                        </span>
                    {{end}}
                </div>
            </div>

            <div class="inline fields">
                <div class=" field">
                    <label>回复</label>
                </div>
                <div class="fourteen wide field">
                    <textarea id="claim-remark" rows="5" name="desc"></textarea>
                </div>
            </div>

            <div class="field pb50">
                <div class="text-center">
                    <div class="ui submit green button" onclick="jobClaim({{.Job.Id}}, 1)">提交</div>

                </div>
            </div>

            <div class="ui error message"></div>
        </form>
    </div>
</div>
<script>
    window.onload=function(){
        $(function(){
            $('.ui.form')
                    .form({
                        fields: {
                             
                             
                            message: {
                                identifier: 'desc',
                                rules: [
                                    {
                                        type   : 'maxLength[100]',
                                        prompt : '业务留言最多100字'
                                    }
                                ]
                            }
                        },
                        onSuccess: function() {
                            $.ajax({
                                url:"/produce/complaint/reply/{{.complain.Id}}",
                                data: new FormData($(".ui.form")[0]),
                                async: false,
                                cache: false,
                                contentType: false,
                                processData: false,
                                type:"post",
                                success:function(data){
                                    if (data && data.error) {
                                        $(".ui.error.message").html(data.error);
                                        $(".ui.error.message").show();
                                    }
                                    else if (data.id) {
                                        $(".ui.error.message").html("");
                                        window.location.href = '/produce/complaint/view';
                                    }
                                    else {
                                        $(".ui.error.message").html("未知错误");
                                    }
                                }
                            });


                            return false;
                        }
                    })
            ;
        });
    }
</script>


