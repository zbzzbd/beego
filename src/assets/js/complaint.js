function getJob(code, callback) {
    $.ajax({
        url:"/search/job",
        data: {
            id : code
        },                               
        type:"get",
        success:function(data){
            if (data && data.error) {
               alert("查询失败：" + data.error)
            }else { 
              callback(data);

            }
        }
    });
}

function deleteComplaint(id) {
    $('.ui.modal.delete').modal({
        closable  : false,
        onDeny: function(){
        },
        onApprove: function() {
            $.ajax({
                url:"/job/complaint/del/" + id,
                type:"get",
                success:function(data){
                    if (data && data.error) {
                        alert(data.error);
                    } else {
                        window.location.reload();
                    }
                }
            });
        }
    }).modal('show');
}