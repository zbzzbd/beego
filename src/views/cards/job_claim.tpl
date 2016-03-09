<div class="ui card p10 full">
    <h4 class="ui teal header">作业认领</h4>
    <div class="ui grid p10">
        <div class="row">
            <label>认领人员：</label>
            {{.ClaimUser.Name}}
        </div>
        <div class="row">
            <label>认领结果：</label>
            {{.Status}}
        </div>
        <div class="row">
            <label>附加说明：</label>
            {{.Remark}}
        </div>
    </div>
    <div >
        <label class="ui teal right ribbon label">认领时间：{{.Created}}</label>
    </div>
</div>
