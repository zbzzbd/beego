
<div class="fixed bottom">
    <div class="ui card p10 full">
        <form id="valid_form" class="ui form" method="post" action="/project/job/valid">
            <div class="four fields">
                <div class="two wide field">
                    <label><i style="color:red">*</i>审核结果:</labeL>
                </div>

                <div class="ui radio checkbox pr20">
                    <input type="radio" name="result" value="1" onclick="$('#finish-time-required').show();">
                    <label>内容明确, 作业排产</label>
                </div>

                <div class="ui radio checkbox pr20">
                    <input type="radio" name="result" value="2" onclick="$('#finish-time-required').hide();">
                    <label>作业有误, 请重新修改</label>
                </div>

                <div class="ui radio checkbox pr20">
                    <input type="radio" name="result" value="3" onclick="$('#finish-time-required').hide();">
                    <label>任务取消</label>
                </div>
            </div>

            <div class="two fields pb10" >
                <div class="two wide field">
                    <label><i style="color:red;" id="finish-time-required">*</i>要求完成时间:</label>
                </div>

                <div class="filed pl0" >
                    <input class="datetimepicker" name="finish_time" >
                    <input type="hidden" name="job_required_finish_time" value="{{TimeFormat .job.FinishTime}}" id="job_required_finish_time">
                </div>
            </div>

            <div class="two fields">
                <div class="two wide field">
                    <label>补充说明:</label>
                </div>

                <div class="field pl0">
                    <textarea rows="2" name="message" style="color: rgba(0,0,0,.87);border: 1px solid rgba(34,36,38,.15);background-color: #fff;"></textarea>
                </div>
            </div>

            <div class="text-center pb10">
                <input type="hidden" name="job_id" value="{{.jobId}}">
                <button class="ui two wide primary button field" type="submit">提交</button>
            </div>

            <div class="ui error message" id="valid_error_msg">
            <ul>
            </ul>
            </div>
        </form>
    </div>
</div>

