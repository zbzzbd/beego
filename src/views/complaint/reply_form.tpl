
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
                    <label>回复</label>
                </div>
                <div class="fourteen wide field">
                    <textarea id="claim-remark" rows="5" name="desc"></textarea>
                </div>
            </div>

            <div class="field pb50">
                <div class="text-center">
                    <button class="ui two wide primary button field" type="submit">提交</button>

                </div>
            </div>

            <div class="ui error message"></div>
        </form>
    </div>
    </div>
<script>
    window.onload=function(){
        $(function(){
            $('.ui.form')
                    .form({
                        fields: {
                             
                             
                            message: {
                                identifier: 'desc',
                                rules: [
                                    {
                                        type   : 'maxLength[100]',
                                        prompt : '业务留言最多100字'
                                    }
                                ]
                            }
                        },
                        onSuccess: function() {
                            $.ajax({
                                url:"/produce/complaint/reply/{{.Id}}",
                                data: new FormData($(".ui.form")[0]),
                                cache: false,
                                contentType: false,
                                processData: false,
                                type:"post",
                                success:function(data){
                                    if (data && data.error) {
                                        $(".ui.error.message").html(data.error);
                                        $(".ui.error.message").show();
                                    }
                                    else if (data.id) {
                                        $(".ui.error.message").html("");
                                        window.location.href = '/produce/complaint/view';
                                    }
                                    else {
                                        $(".ui.error.message").html("未知错误");
                                    }
                                }
                            });


                            return false;
                        }
                    })
            ;
        });
    }
</script>
