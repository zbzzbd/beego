<div class="p10">
    <div class="ui card full p30">
        <div class="ui grid">
            <div class="twelve wide column">
                <div style="font-size:30px">作业总数:</div>
            </div>
            <div class="four wide column">
                <div style="font-size:30px">{{.Job.Count}}</div>
            </div>
        </div>
    </div>
</div>


<div class="p10">
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
                    <div class="column">{{.Job.WaitValidCount}}</div>
                </div>
                <div class="row">
                    <div class="column"><a href="/project/job/valid">待审核</a></div>
                </div>
            </div>
        </div>

        <div class="three wide column">
            <div class="ui card p20 full">
                <div class="row">
                    <div class="column">{{.Job.ValidRefuseCount}}</div>
                </div>
                <div class="row">
                    <div class="column">待修改</div>
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
