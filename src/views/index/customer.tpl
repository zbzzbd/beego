<div class="pb20">
    <div class="ui card full p30">
        <div class="ui grid">
            <div class="twelve wide column">
                <div style="font-size:30px">我发布的作业:</div>
            </div>
            <div class="four wide column">
                <div style="font-size:30px">{{.Job.CreateCount}}</div>
            </div>
        </div>
    </div>
</div>


<div class="pb20">
    <div class="ui grid">
        <div class="three wide column">
            <div class="ui card p20 full">
                <div class="row">
                    <div class="column">{{.Job.FinishCount}}</div>
                </div>
                <div class="row">
                    <div class="column">作业完成</div>
                </div>
            </div>
        </div>

        <div class="three wide column">
            <div class="ui card p20 full">
                <div class="row">
                    <div class="column">{{.Job.ModifyCount}}</div>
                </div>
                <div class="row">
                    <div class="column"><a href="/job/progress">作业修改</a></div>
                </div>
            </div>
        </div>

        <div class="three wide column">
            <div class="ui card p20 full">
                <div class="row">
                    <div class="column">{{.Job.CancelCount}}</div>
                </div>
                <div class="row">
                    <div class="column">作业取消</div>
                </div>
            </div>
        </div>

        <div class="three wide column">
            <div class="ui card p20 full">
                <div class="row">
                    <div class="column">{{.Job.DoingCount}}</div>
                </div>
                <div class="row">
                    <div class="column">进行中</div>
                </div>
            </div>
        </div>

    </div>
</div>

{{ template "index/my_project.tpl" .}}
