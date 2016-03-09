$(function() {
    $('#job_assign_form').form({
        fields: {
            result: {
                identifier: 'to_user',
                rules: [{
                    type: 'empty',
                    prompt: '转发给谁不能为空'
                }]
            }
        },
        onSuccess:function(){
            job_assign()
            return false;
        }
    });
});

function job_assign() {
    $.ajax({
        url:"/produce/job/assign",
        data: $("#job_assign_form").serialize(),
        type:"post",
        success:function(data){
            if (data && data.error) {
                $("#job_assign_form").removeClass("success").addClass("error")
                $("#job_assign_form #error_message").append("<li>"+data.error+"</li>");
            } else {
                window.location.href =  "/produce/job/claim";
            }
        }
    });
}
