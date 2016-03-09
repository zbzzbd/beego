<div class="ui card p10 full">
    <form class="ui form" enctype ="multipart/form-data">
        <div class="inline fields">

            <div class="field">
                <label>作业编号</label>
                <input type="text" name="Jobcode" style="min-width: 220px" onkeyup="getJob(this.value, setJob)">
            </div>


            <div class="field">
                <label>客诉事项</label>
                <select class="ui fluid dropdown" name ="Type" style="min-width: 220px">
                    <option value="1">延时</option>
                    <option value="2">逻辑错误</option>
                    <option value="3">其它重大错误</option>
                </select>

            </div>

        </div>

        <div class="inline fields">
            <div class="field">
                <label>项目名称:</label>
                <lable id="project-name"> </lable>
            </div>

        </div>

        <div class="ui grid">
            <div class="8 wide column">
                <div class="inline fields">
                    <label>作业要求:</label>
                    <label id="job-type"></label>

                </div>
                <div class="inline fields">
                    <label>作业部门:</label>
                    <label id="department-code"> </label>
                </div>
            </div>
        </div>

        <div class="inline fields">
            <div class="field">
                <label>作业对象：</label>
                <label id ="job-target"> </label>
            </div>

            <div class="field">
                <label>作业单元（即执行人）：</label>
                <label  id ="employee-name"> </label>
            </div>
        </div>


        <div class="inline fields">
            <div class="field">
                <label>修改网址：</label>
                <label id="url">  <a href="http://{{.Job.TargetUrl}}">{{.Job.TargetUrl}}</a> </label>
            </div>

            <div class="field">
                <label>验收时间：</label>
                <label id="job-endtime"> {{.Job.FinishTime}} </label>
            </div>
        </div>

        <div class="inline fields">
            <div class=" field">
                <label>作业描述：</label>
                <label id="job-desc">{{.Job.Desc}}</label>
            </div>
        </div>
        <div class="inline fields">
            <div class=" field">
                <label>附件：</label>
                <lable id="job-files"></lable>
                {{range .JobFiles}}
                <a href="{{.Url}}" class="pr20" download="{{.Name}}" target="_blank">{{.Name}}</a>
                {{end}}
            </div>
        </div>

        <div class="inline fields">
            <div class=" field">
                <label>业务留言：</label>
                <label id="job-message"></label>

            </div>
        </div>
        <div class="inline fields">
            <div class=" field">
                <label>客诉描述：</label>
            </div>
            <div class="fourteen wide field">
                <textarea rows="5" name="Desc"  placeholder="此处填写投诉详细描述"></textarea>
            </div>
        </div>

        <div class="inline fields">
            <div class=" field">
                <label>是否需要回复</label>
            </div>
            <div class=" field">
                <input  type ="radio" name ="reply"  value="1" /> 需要
                <input  type ="radio" name ="reply"  value="0" /> 不需要
            </div>
        </div>

        <div class="field">
            <div class="ui submit green button fr" onclick="">提交</div>
            <div class="ui reset button fr">清空</div>
        </div>

        <div class="ui error message"></div>
    </form>
</div>

<script>
    window.onload=function(){
        $(function(){
            $('.ui.form')
                    .form({
                        fields: { 
                                                 
                            Type: {
                                identifier: 'Type',
                                rules: [
                                    {
                                        type   : 'empty',
                                        prompt : '请选择客诉事项'
                                    }
                                ]
                            }

                        },
                        onSuccess: function() {
                            $.ajax({
                                url:"/job/complaint/create",
                                data: new FormData($(".ui.form")[0]),
                                async: false,
                                cache: false,
                                contentType: false,
                                processData: false,
                                type:"post",
                                success:function(data){                                   
                                    if (data && data.error) {
                                        $(".ui.error.message").html(data.error);
                                    }
                                    else if (data.id) {
                                        $(".ui.error.message").html("");
                                        window.location.href = '/job/complaint/view';
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

    function setJob(data) {
        console.info( $("#project-name"), data )
        $("#project-name").html(data.Job.Project.Name)
        $("#job-type").html(data.Job.Type)
        $("#department-code").html(data.Job.Department)
        $("#job-target").html(data.Job.Target)
        $("#employee-name").html(data.Job.Employee.Name)
        $("#url").html(data.Job.TargetUrl)
        $("#job-endtime").html(data.Job.Updated)
        $("#job-desc").html(data.Job.Desc)
        $("#job-message").html(data.Job.Message)
        $("#job-files").html(data.JobFiles[0].Name)
    }
</script>