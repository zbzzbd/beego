        <div class="ui card p10 full">
            <form class="ui form" enctype ="multipart/form-data">
                <div class="inline fields">
                    <div class="field">
                        <label>启动时间</label> 
                        <input class="datetimepicker" id="datetimepicker-start-time" type=text readonly value={{dateformat .Project.Started "2006-01-02 15:04"}} name="started">
                    </div>
                </div>

                <div class="inline fields">
                    <div class="field">
                        <label>项目名称</label>
                        <input type="text" style="width:360px;" value={{.Project.Name}} name="name"> 
                    </div>
                </div>
                <div class="inline fields">
                    <div class="field"> 
                    <label>服务项目</label>
                    <input type="text" name="service_item" value={{.Project.ServiceItem}}>
                    </div>
                </div>
                <div class="inline fields">
            <div class="field">
                <label>合同编号</label>
                <input type="text" name="contract_no" value={{.Project.ContractNo}}>
            </div>
        </div>
                <div class="inline fields">
                    <div class="field">
                        <label>赛事规模</label>
                        <input type="text" value={{.Project.Scale}} name="scale">
                    </div>
                    <label>/人</label>
                    <span style="width:50px;"></span>
                    <div class="field">
                        <label>优先级别</label>
                        <select class="ui dropdown"  name="priority">
                            {{range .Priority}}
                                {{if  compare $.Project.Priority  .Priority}}
                                    <option selected="selected" value="{{.Priority}}">{{.Priority}}</option>
                                {{else}}
                                    <option value="{{.Priority}}">{{.Priority}}</option>
                                {{end}}
                            {{end}}
                        </select>
                    </div>
                </div>
                
                <div class="inline fields">
                    <div class="field">
                        <label>客户名称</label>
                         <input type="text" name="client_name" value={{.Project.ClientName}}>
                    </div>
                </div>

                <div class="inline fields">
                    <div class="field">
                        <label>业务担当</label>
                        <select class="ui dropdown" name="bussiness_user">
                            {{range .BussinessUser}}
                                {{if compare $.Project.BussinessUser.Id .Id}}
                                    <option selected="selected" value="{{.Id}}">{{.Name}}</option>
                                {{else}}
                                    <option value="{{.Id}}">{{.Name}}</option>
                                {{end}}
                            {{end}}
                        </select>
                    </div>
                </div>

                
                <div class="inline fields">
                    <div class="field">
                        <label>项目进程</label>
                        <select class="ui dropdown" name="progress">
                            {{range .Progress}}
                                {{if compare $.Project.Progress.Id .Id}}
                                    <option selected="selected" value="{{.Id}}">{{.Name}}</option>
                                {{else}}
                                    <option value="{{.Id}}">{{.Name}}</option>
                                {{end}}
                            {{end}}
                            </select>
                    </div>
                   <span style="width:50px;"></span>
                    <div class="field">
                        <label>开通报名时间</label> 
                        <input class="datetimepicker" id="datetimepicker-reg-start-time" type=text readonly value={{dateformat .Project.RegStartDate "2006-01-02 15:04"}} name="reg_start_date" > 
                    </div>
                </div>


                <div class="inline fields">
                    <div class="field">
                        <label>比赛日期</label> 
                        <input class="datetimepicker" id="datetimepicker-game-start-time" type=text readonly value={{dateformat .Project.GameDate "2006-01-02 15:04"}} name="game_date" >
                    </div>
                   <span style="width:50px;"></span>
                    <div class="field">
                        <label>关闭报名时间</label> 
                        <input class="datetimepicker" id="datetimepicker-reg-close-time" type=text readonly value={{dateformat .Project.RegCloseDate "2006-01-02 15:04"}} name="reg_close_date" >
                    </div>
                </div>


                <div class="inline fields">
                    <div class="field">
                        <label>美术单元</label>
                        <select class="ui dropdown" id="projectArt" name="art_user">
                            {{if $.Project.ArtUser}}
                                {{range .ArtUser}}
                                    {{if compare $.Project.ArtUser.Name .Name}}
                                        <option selected="selected" value="{{.Id}}">{{.Name}}</option>
                                    {{else}}
                                        <option value="{{.Id}}">{{.Name}}</option>
                                    {{end}}
                                {{end}}
                           {{else}}
                                <option selected="" value="">请选择美术单元</option>option>
                                {{range .ArtUser}}
                                        <option value="{{.Id}}">{{.Name}}</option>
                                {{end}}
                            {{end}}
                        </select>
                    </div>
                </div>


                <div class="inline fields">
                    <div class="field">
                        <label>技术单元</label>
                        <select class="ui dropdown" id="projectTech" name="tech_user">
                            {{if $.Project.TechUser}}
                                {{range .TechUser}}
                                    {{if compare $.Project.TechUser.Name .Name}}
                                        <option selected="selected" value="{{.Id}}">{{.Name}}</option>
                                    {{else}}
                                        <option value="{{.Id}}">{{.Name}}</option>
                                    {{end}}
                                {{end}}
                            {{else}}
                                <option selected="" value="">请选择技术单元</option>option>
                                 {{range .TechUser}}
                                        <option value="{{.Id}}">{{.Name}}</option>
                                {{end}}
                             {{end}}
                        </select>
                    </div>
                </div>

                <div class="field">
                    <div class="ui submit green button fr">提交</div>
                    <div class="ui reset button fr">重置</div>
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
                            },
                        onSuccess: function() {
                            $.ajax({
                                url: {{Strcat "/project/edit/" .Id}},
                                data: new FormData($(".ui.form")[0]),
                                async: false,
                                cache: false,
                                contentType: false,
                                processData: false,
                                type:"post",
                                success:function(data){
                                    if (data && data.result != 0) {
                                        $(".ui.error.message").html(data.error);
                                    }
                                    else {
                                        $(".ui.error.message").html("");
                                        window.location.href = '/project/list';
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
