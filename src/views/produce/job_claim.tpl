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
        <form class="ui form" enctype ="multipart/form-data">
            <div class="inline fields">
                <div class=" field">
                    <label>补充说明</label>
                </div>
                <div class="fourteen wide field">
                    <textarea id="claim-remark" rows="5" name="remark"></textarea>
                </div>
            </div>

            <div class="field">
                <div class="text-center">
                    <div class="ui submit green button" onclick="jobClaim({{.jobId}}, 1)">认领任务</div>
                    <div class="ui orange button" onclick="jobClaim({{.jobId}}, 2)">任务无法完成，拒绝认领</div>
                </div>
            </div>

            <div class="ui error message"></div>
        </form>
    </div>
</div>

<script>
    function jobClaim(jobId, accept) {
        $.ajax({
            url:"/produce/job/claim/" + jobId,
            data: {
                status: accept,
                remark: $("#claim-remark").val()
            },
            type:"post",
            success:function(data){
                if (data && data.error) {
                    alert("认领错误：" + data.error)
                }
                else {
                    window.location.href = '/produce/job/claim';
                }
            }
        });
    }
</script>