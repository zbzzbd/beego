
<div class="fixed bottom">
    <div class="ui card p10 full">
        <form class="ui form" enctype ="multipart/form-data">  
         <div class="inline fields">
            <div class="field">
                <label>投诉部门: {{str2html RequiredStar}}</label>
            </div>
             <div class="field" >
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

            <div class="field">
                 <label>投诉单元: {{str2html RequiredStar}} </label>
             </div>
             <div class="field">
               <select class="ui search dropdown" name="employee_id" id="employee">
                    <option value="">请选择 </option>
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
           

         </div>   
         <div class="inline fields">  
            <div class="field">
                <label>客诉事项：{{str2html RequiredStar}}</label>
            </div>   
            <div class="four wide field">  
                    <select class="ui dropdown selection" name ="Type" style="min-width: 220px">
                         <option value="">请选择</option>
                         <option value="1">延时</option>
                         <option value="2">逻辑错误</option>
                         <option value="3">其它重大错误</option>
                    </select>  
            </div> 

        </div>
        <div class="inline fields">
            <div class="field">
                <label>投诉来源 {{str2html RequiredStar}}</label>
            </div>
            <div class="four wide field">
                <select class="ui dropdown selection" name="source" style="min-width: 220px">
                    <option value="">请选择</option>
                    <option value="客户">客户</option>
                    <option value="选手">选手</option>
                    <option value="公司内部">公司内部</option>
                </select>
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
                        <label>是否需要回复：</label>
                    </div>
                     <div class=" field"> 
                        <input  type ="radio" name ="reply"  value="1" /> 需要
                        <input  type ="radio" name ="reply"  value="0" /> 不需要
                    </div>
                </div> 

            <div class="text-center pb10">
                <input type="hidden" name="job_id" value="{{.jobId}}">
                <button class="ui two wide primary button field" type="submit">提交</button>
            </div>

        </form>
    </div>
</div>

<script>
    window.onload=function(){
        $(function(){
             var str = window.location.href;
             var site= "";
             if (str.indexOf('?') >0)
                site= "/job/complaint/create"+str.substring(str.indexOf('?'))
             else {
                site= "/job/complaint/create"
             }
             $('.ui.form')
                    .form({
                        fields: {
                             department:{
                                identifier: 'department',
                                rules:[
                                    {
                                        type  : 'empty',
                                        prompt : '投诉部门不能为空'
                                    }
                                ]
                             },
                             employee_id:{
                                identifier : 'employee_id',
                                rules:[
                                    {
                                        type: 'empty',
                                        prompt: '投诉单元不能为空'
                                    }
                                ]
                             },
                             Type :{
                                identifier : 'Type',
                                rules:[
                                    {
                                        type: 'empty',
                                        prompt: '客诉事项不能为空'
                                    }
                                ]
                             },
                             source :{
                                identifier: 'source',
                                rules: [
                                     {
                                        type: 'empty',
                                        prompt: '投诉来源不能为空'
                                     }
                                ]
                             },
                            message: {
                                identifier: 'Desc',
                                rules: [
                                    {
                                        type   : 'maxLength[2000]',
                                        prompt : '业务留言最多2000字'
                                    }
                                ]
                            }
                        },
                        onSuccess: function() {
                            $.ajax({ 
                                url: site,
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
</script>

