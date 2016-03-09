function jobAssign(jobId, callback) {
    $.ajax({
        url: "/job/assign",
        data: {
            id: jobId,
            employee_id: $.cookie('uid')
        },
        type: "get",
        success: function(data) {
            if (data && data.error) {
                alert("认领错误：" + data.error)
            } else {
                window.location.reload();
            }
        }
    });
}

function delJobFile(fileId, callback) {
    $.ajax({
        url: "/job/file/del",
        data: {
            id: fileId
        },
        type: "get",
        success: function(data) {
            if (data && data.error) {
                alert("错误：" + data.error)
            } else {
                window.location.reload();
            }
        }
    });
}
var totalSize = 0;
var i =0;
function del_files(obj,id) {
    if (document.getElementById(id).files && document.getElementById(id).files[0]!=null) { 
         size= document.getElementById(id).files[0].size
        if (totalSize >0){
            totalSize -= size;  
        }
      } 
    $(obj).parent(".field").remove()
     
    return false
}

function upload_files(obj) {
    
    if (document.getElementById('fileToUpload'+i).files) {
        size= document.getElementById('fileToUpload'+i).files[0].size
        totalSize += size;  
        if (totalSize >= 1024*1024*10) {
            alert("文件已经大于10MB，请重新选择")
            totalSize -= size;
        } 
      } else {
        alert("上传文件获取为空")
      }  
}

$(function() {
    $("#add_files").click(function(){
        i=i+1; 
        s='fileToUpload'+i 
        var template = $('<div class="field"> <input type="file" name="files[]" onchange="upload_files(this)" id="'+s+'"> <button class="ui red button" onclick="del_files(this,s)">删除附件</button> </div>')  
        $(this).parent(".field").after(template) 
        
        return false
    }); 
});
   
function delJob(id) {
    $('.ui.modal.delete').modal({closable: false, onDeny: function(){}, onApprove: function(){
        $.ajax({
            url:"/job/delete/" +id,
            type: "get",
            success: function(data){
                if (data && data.error) {
                    alert(data.error);
                }else {
                    window.location.reload();
                }
            }
        });

    }}).modal('show');
}

function recoverJob(id) {
    $('.ui.modal.delete').modal({closable: false, onDeny: function(){}, onApprove: function(){
        $.ajax({
            url:"/job/recover/" +id,
            type: "get",
            success: function(data){
                if (data && data.error) {
                    alert(data.error);
                }else {
                    window.location.reload();
                }
            }
        });

    }}).modal('show');
}