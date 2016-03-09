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
              <div class="inline field">
            <label>上传附件:(<span style="color:red;">附件总大小不能大于10M</span>)</label>

            <div class="field">
                    <input type="file" name="files[]" onchange="upload_files(this)" id="fileToUpload0">
                    <button class="ui primary button" id="add_files">添加附件</button>
            </div>

                 </div>

        <div class="inline fields">
            <div class="field">
                <label style="color: transparent">已经上传</label>
            </div>
            <div class=" field">
                {{range .JobFiles}}
                        <span class="pr20">
                            <a href="{{.Url}}"  download="{{.Name}}" target="_blank">{{.Name}}</a>
                            <i onclick="delJobFile({{.Id}})" class="remove icon"></i>
                        </span>
                {{end}}
            </div>
        </div>
            <div class="inline fields">
                <div class=" field">
                    <label>附加说明</label>
                </div>
                <div class="fourteen wide field">
                    <textarea id="submit-remark" rows="5" name="remark"></textarea>
                </div>
            </div>

            <input type="hidden" name="status" id="submit-status">
            <div class="field">
                {{if eq .status 1}}
                <div class="text-center">
                    <div class="ui submit green button" onclick="jobSubmit({{.jobId}}, 1)">完成任务，提交</div>
                </div>
                {{else if eq .status 2}}
                <div class="text-center">
                    <div class="ui submit green button" onclick="jobSubmit({{.jobId}}, 3)">Good，验收通过</div>
                    <div class="ui submit red button" onclick="jobSubmit({{.jobId}}, 2)">验收不通过，请重新制作</div>

                </div>
                {{end}}
            </div>

            <div class="ui error message"></div>
        </form>
    </div>
</div>


<script>
    function jobSubmit(jobId, accept) {
        $("#submit-status").val(accept)
        $.ajax({
            url:"/produce/job/submit/" + jobId,
            data: new FormData($(".ui.form")[0]),
            async: false,
            cache: false,
            contentType: false,
            processData: false,
            type:"post",
            success:function(data){
                if (data && data.error) {
                    alert("提交作业错误：" + data.error)
                }
                else {
                    if (accept > 1) {
                        window.location.href = '/job/view/' + jobId;
                    } else {
                        window.location.href = '/produce/job/submit';
                    }
                }
            }
        });
    }
</script>
