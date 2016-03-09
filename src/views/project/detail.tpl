<div class="ui card p10 full">
    <form class="ui form" enctype ="multipart/form-data">
        <div class="inline fields">
            <div class="field">
                <label>启动时间</label>
                <input disabled="disabled" id="datetimepicker-start-time" value={{dateformat .Project.Started "2006-01-02 15:04"}} name="started">
            </div>
        </div>

        <div class="inline fields">
            <div class="field">
                <label>项目名称</label>
                <input disabled="disabled" type="text" style="width:360px;" value={{.Project.Name}} name="name">
            </div>
        </div>

        <div class="inline fields">
            <div class="field">
                <label>赛事规模</label>
                <input disabled="disabled" type="text" value={{.Project.Scale}} name="scale">
            </div>
            <label>/人</label>
            <span style="width:50px;"></span>
            <div class="field">
                <label>优先级别</label>
                <input disabled="disabled" type="text" value={{.Project.Priority}} name="priority">
            </div>
        </div>

        <div class="inline fields">
            <div class="field">
                <label>客户名称</label>
                <input disabled="disabled" type="text" name="client_name" value={{.Project.ClientName}}>
            </div>
        </div>

        <div class="inline fields">
            <div class="field">
                <label>业务担当</label>
                <input disabled="disabled" type="text" value={{.Project.BussinessUser.Name}} name="bussiness_user">
            </div>
        </div>


        <div class="inline fields">
            <div class="field">
                <label>项目进程</label>
                <input disabled="disabled" type="text" value={{.Project.Progress.Name}} name="progress">
            </div>
            <span style="width:50px;"></span>
            <div class="field">
                <label>开通报名时间</label>
                <input disabled="disabled" id="datetimepicker-reg-start-time" value={{dateformat .Project.RegStartDate "2006-01-02 15:04"}} name="reg_start_date" >
            </div>
        </div>


        <div class="inline fields">
            <div class="field">
                <label>比赛日期</label>
                <input disabled="disabled" id="datetimepicker-game-start-time" value={{dateformat .Project.GameDate "2006-01-02 15:04"}} name="game_date" >
            </div>
            <span style="width:50px;"></span>
            <div class="field">
                <label>关闭报名时间</label>
                <input disabled="disabled" id="datetimepicker-reg-close-time" value={{dateformat .Project.RegCloseDate "2006-01-02 15:04"}} name="reg_close_date" >
            </div>
        </div>


        <div class="inline fields">
            <div class="field">
                <label>美术单元</label>
                {{if .Project.ArtUser}}
                    <input disabled="disabled" type="text" value={{.Project.ArtUser.Name}} name="art_user">
                {{else}}
                    <input disabled="disabled" type="text" value="" name="art_user">
                {{end}}

            </div>
        </div>


        <div class="inline fields">
            <div class="field">
                <label>技术单元</label>
                {{if .Project.TechUser}}
                    <input disabled="disabled" type="text" value={{.Project.TechUser.Name}} name="tech_user">
                {{else}}
                    <input disabled="disabled" type="text" value="" name="tech_user">
                {{end}}
            </div>
        </div>

        <div class="field">
            <a class="ui green button fr" href="/project/list">Return</a>
        </div>

        <div class="ui error message"></div>
    </form>
</div>
