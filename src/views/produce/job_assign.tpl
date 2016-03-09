{{if eq .error ""}}
<div class="pb200">
    {{str2html .cards}}
</div>
{{else}}
<div class="ui p10 full red message">
    错误： {{.error}}
</div>
{{end}}

<div class="fixed bottom">
    <div class="ui card p10 full">
        <form class="ui form" method="post" id="job_assign_form">

            <div class="inline fields">
                <div class="field">
                <label>转发给谁: </label>
                </div>
                <select class="ui search dropdown" name="to_user" >
                    <option value="">请选择</option>
                    {{range .toUsers }}
                        <option value="{{.Id}}">{{.Name}}</option>
                    {{end}}
                </select>
            </div>


            <div class="inline fields">
                <div class=" field">
                    <label>转发说明: </label>
                </div>
                <div class="fourteen wide field">
                    <textarea id="claim-remark" rows="5" name="remark"></textarea>
                </div>
            </div>

            <div class="text-center pb10">
                <input type="hidden" name="job_id" value="{{.jobId}}">
                <button class="ui two wide primary button field" type="submit">提交</button>
            </div>

            <div class="ui error message" id="error_message"></div>

        </form>
    </div>
</div>
