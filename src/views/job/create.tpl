<div class="ui card p10 full">
    <form class="ui form" enctype ="multipart/form-data" id="job_create_form">
        <div class="inline fields">
            <div class="field">
                <label>项目名称</label>
                <select class="ui search dropdown" name="project_id">
                    <option value="">请选择</option>
                    {{range .ProjectNames}}
                    {{if $.is_edit}}
                    {{if eq .Id $.Job.Project.Id}}
                    <option selected="" value="{{.Id}}">{{.Name}}</option>
                    {{else}}
                    <option value="{{.Id}}">{{.Name}}</option>
                    {{end}}
                    {{else}}
                    <option value="{{.Id}}">{{.Name}}</option>
                    {{end}}
                    {{end}}
                </select>
            </div>
        </div>

        <div class="inline fields">
            <div class="field">
                <label>作业要求</label>
                <select class="ui search dropdown" name="type">
                    <option value="">请选择</option>
                    {{range $elem := .Types}}
                    {{if $.is_edit}}
                    {{if eq $elem $.Job.Type}}
                    <option selected="" value="{{$elem}}">{{$elem}}</option>
                    {{else}}
                    <option value="{{$elem}}">{{$elem}}</option>
                    {{end}}
                    {{else}}
                    <option value="{{$elem}}">{{$elem}}</option>
                    {{end}}
                    {{end}}
                </select>
            </div>

            <div class="field">
                <label>作业部门</label>
                <select class="ui dropdown" name="department" id="department">
                    <option value="">请选择</option>
                    {{range .Departments}}
                    {{if $.is_edit}}
                    {{if eq .Department $.Job.Department}}
                    <option selected="" value="{{.Department}}" tag="{{.Role}}">{{.Department}}</option>
                    {{else}}
                    <option value="{{.Department}}" tag="{{.Role}}">{{.Department}}</option>
                    {{end}}
                    {{else}}
                    <option value="{{.Department}}" tag="{{.Role}}">{{.Department}}</option>
                    {{end}}
                    {{end}}
                </select>
            </div>
        </div>

        <div class="inline fields">
            <div class="field">
                <label>作业对象</label>
                <input type="text" name="target" placeholder="例如：官网首页 > LOGO更新" style="min-width: 350px" value="{{.Job.Target}}">
            </div>


            <div class="field disabled">
                <label>作业单元</label>
                <select class="ui search dropdown" name="employee_id" id="employee">
                    <option value="">请选择</option>
                    {{range .Employees}}
                    {{if $.is_edit}}
                        {{if eq .Id $.Job.Employee.Id}}
                        <option selected="" value="{{.Id}}" tag="{{.Roles}}">{{.Name}}</option>
                        {{else}}
                        <option value="{{.Id}}" tag="{{.Roles}}">{{.Name}}</option>
                        {{end}}
                    {{else}}
                         <option value="{{.Id}}" tag="{{.Roles}}">{{.Name}}</option>
                    {{end}}
                    {{end}}
                </select>
            </div>
        </div>

        <div class="inline fields">
            <div class="field">
                <label>修改网址</label>
                <input type="text" name="target_url" placeholder="http://www.chinarun.com (必须带http://或https://)" style="min-width: 350px" value="{{.Job.TargetUrl}}">
            </div>

            <div class="field">
                <label>验收时间</label>
                <input class="datetimepicker" id="datetimepicker-finish-time" name="finish_time" value="{{TimeFormat .Job.FinishTime}}">
            </div>
        </div>

        <div class="inline fields">
            <div class=" field">
                <label>作业描述</label>
            </div>
            <div class="fourteen wide field">
                <textarea rows="5" name="desc">{{.Job.Desc}}</textarea>
            </div>
        </div>

        <div class="inline field">
            <label>上传附件:(<span style="color:red;">附件总大小不能大于10M</span>)</label>

            <div class="field">
                    <input type="file" name="files[]" onchange="upload_files(this)" id="fileToUpload0">
                    <button class="ui primary button" id="add_files">添加附件</button>
            </div>

                 </div>

        <div class="inline fields">
            <div class="field">
                <label style="color: transparent">已经上传</label>
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
                <label>业务留言</label>
                <input type="text" name="message" style="min-width: 400px" placeholder="限16字" value="{{.Job.Message}}">
            </div>
        </div>


        <div class="field pb50">
            <div class="ui submit green button fr">提交</div>
            <div class="ui reset button fr">清空</div>
        </div>

        <div class="ui error message"></div>
    </form>
</div>
<script>       
    window.onload=function(){     
        $(function(){     
            $('#job_create_form')     
                    .form({       
                        fields: {     
                           project_id: {     
                              identifier: 'project_id',     
                               rules: [      
                                    {     
                                         type   : 'empty',     
                                         prompt : '项目名称不能为空'       
                                     }     
                                ]     
                             },        
                             type: {       
                                 identifier: 'type',       
                                 rules: [      
                                     {     
                                         type   : 'empty',     
                                         prompt : '请选择作业要求'        
                                     }     
                                 ]     
                             },        
                             department: {     
                                 identifier: 'department',     
                                 rules: [      
                                     {     
                                         type   : 'empty',     
                                         prompt : '请选择部门单元'        
                                    }     
                                 ]     
                             },        
                             employee_id: {        
                                 identifier: 'employee_id',        
                                 rules: [      
                                     {     
                                         type   : 'empty',     
                                         prompt : '请选择作业单元'        
                                     }     
                                 ]     
                             },        
                            finish_time: {        
                                 identifier: 'finish_time',        
                                 rules: [      
                                     {     
                                         type   : 'empty',     
                                         prompt : '请填写验收时间'        
                                     }     
                                 ]     
                             },        
                             target: {     
                                 identifier: 'target',     
                                 rules: [      
                                     {     
                                         type   : 'empty',     
                                         prompt : '作业对象不能为空'       
                                     }     
                                 ]     
                             },        
                             message: {        
                                 identifier: 'message',        
                                 rules: [      
                                     {     
                                         type   : 'maxLength[16]',     
                                         prompt : '业务留言最多16字'      
                                     }     
                                 ]     
                             }     
                         },        
                         onSuccess: function() {       
                             $.ajax({      
                                 url:"{{.post_url}}",      
                                 data: new FormData($("#job_create_form")[0]),     
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
                                         window.location.href = '/job/view/' + data.id;        
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