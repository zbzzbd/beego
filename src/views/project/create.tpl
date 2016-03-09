<div class="ui card p10 full">
    <form class="ui form" enctype ="multipart/form-data" id="project_create_form">
        <div class="inline fields">
            <div class="field">
                
                <label>启动时间{{str2html RequiredStar}}</label>
                <input class="datetimepicker" id="datetimepicker-start-time" type=text readonly value={{dateformat .Now "2006-01-02 15:04"}} name="started" >
            </div>
        </div>
        <div class="inline fields">
            <div class="field">
                <label>项目名称{{str2html RequiredStar}}</label>

                <div class="ui search" style="display:inline-block;" id="project_search">
                    <div class="ui input">
                        <input  class="prompt" type="text" placeholder="项目..." name="name" style="width:350px;">
                    </div>
                    <div class="results"></div>
                </div>

            </div>
        </div>
        <div class="inline fields">
            <div class="field">
                <label>合同编号</label>
                <input type="text" name="contract_no">
            </div>
        </div>
         <div class="inline fields">
            <div class="field">
                <label>服务项目</label>
                <input type="text" name="service_item">
            </div>
        </div>
        <div class="inline fields">
            <div class="field">
                <label>赛事规模</label>
                <input type="text" name="scale">
            </div>
            <label>/人</label>
            <span style="width:50px;"></span>
            <div class="field">
                <label>优先级别</label>
                <select class="ui dropdown" id="projectProgress" name="priority">
                    <option  selected="" value="">优先级1为最高</option>
                    {{range .Priority}}
                    <option value="{{.Priority}}">{{.Priority}}</option>
                    {{end}}
                </select>
            </div>
        </div>
        <div class="inline fields">
            <div class="field">
                <label>客户名称</label>
                <input type="text" name="client_name">
            </div>
        </div>
        <div class="inline fields">
            <div class="field">
                <label>业务担当{{str2html RequiredStar}}</label>
                <select class="ui dropdown" id="projectBussiness" name="bussiness_user">
                    <option selected="" value="">请选择业务担当</option>
                    {{range .BussinessUser}}
                    <option value="{{.Id}}">{{.Name}}</option>
                    {{end}}
                </select>
            </div>
        </div>
        <div class="inline fields">
            <div class="field">
                <label>项目进程</label>
                <select class="ui dropdown" name="progress">
                    {{range .Progress}}
                    <option value="{{.Id}}">{{.Name}}</option>
                    {{end}}
                </select>
            </div>
            <span style="width:50px;"></span>
            <div class="field">
                <label>开通报名时间</label>
                <input class="datetimepicker" id="datetimepicker-reg-start-time" type=text readonly name="reg_start_date" >
            </div>
        </div>
        <div class="inline fields">
            <div class="field">
                <label>比赛日期</label>
                <input class="datetimepicker" id="datetimepicker-game-start-time" type=text readonly name="game_date" >
            </div>
            <span style="width:50px;"></span>
            <div class="field">
                <label>关闭报名时间</label>
                <input class="datetimepicker" id="datetimepicker-reg-close-time" type=text readonly name="reg_close_date" >
            </div>
        </div>
        <div class="inline fields">
            <div class="field">
                <label>美术单元{{str2html RequiredStar}}</label>
                <select class="ui dropdown" id="projectArt" name="art_user">
                    <option selected="" value="">请选择美术单元</option>
                    {{range .ArtUser}}
                    <option value="{{.Id}}">{{.Name}}</option>
                    {{end}}
                </select>
            </div>
        </div>
        <div class="inline fields">
            <div class="field">
                <label>技术单元{{str2html RequiredStar}}</label>
                <select class="ui dropdown" id="projectTech" name="tech_user">
                    <option value="">请选择技术单元</option>
                    {{range .TechUser}}
                    <option value="{{.Id}}">{{.Name}}</option>
                    {{end}}
                </select>
            </div>
        </div>
        <div class="field">
            <div class="ui submit green button fr">提交</div>
            <div class="ui reset button fr">清空</div>
        </div>
        <div class="ui error message"></div>
    </form>
</div>
<script>
window.projectNames = [];
{{range .ProjectNames}}
projectNames.push({"title": {{.Name}}})
{{end}}

console.log(projectNames)
</script>
