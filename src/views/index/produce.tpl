<div class="pb20">
    <div class="ui card full p30">
        <div class="ui grid">
            <div class="twelve wide column">
                <div style="font-size:30px">我完成的作业:</div>
            </div>
            <div class="four wide column">
                <div style="font-size:30px">{{.Job.FinishCount}}</div>
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
                    <div class="column">{{.Job.WaitClaimCount}}</div>
                </div>
                <div class="row">
                    <div class="column"><a href="/produce/job/claim">待认领</a></div>
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

        <div class="three wide column">
            <div class="ui card p20 full">
                <div class="row">
                    <div class="column">{{.Job.WaitFinishCount}}</div>
                </div>
                <div class="row">
                    <div class="column">待验收</div>
                </div>
            </div>
        </div>

    </div>
</div>

{{ template "index/my_project.tpl" .}}
