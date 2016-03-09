<div class="ui card p10 full">
    <h4 class="ui teal header">作业审核</h4>
    <div class="ui grid p10">
        <div class="row">
            <div class="nine wide column">
                <label>回复人：</label>
                {{.UserName}}
            </div>
        </div>

        <div class="row">
            <div class="nine wide column">
                <label>审核结果：</label>
                {{.Result}}
            </div>
        </div>
        
        <div class="row">
            <div class="nice wide column">
                <label>要求完成时间: </label>
                {{.FinishTime}}
            </div>
        </div>

        <div class="row">
            <div class="five wide column">
                <label>审核说明：</label>
                {{.Message}}
            </div>
        </div>
    </div>
    <div >
        <label class="ui teal right ribbon label">审核时间：{{.Created}}</label>
    </div>
</div>
