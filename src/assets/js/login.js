$(function() {
    $('#login_form').form({
        on: 'blur',
        fields: {
            email: {
                identifier  : 'email',
                rules: [
                {
                    type   : 'empty',
                    prompt : '请输入邮箱'
                },
                {
                    type: 'email',
                    prompt: "邮箱格式不正确"
                }
                ]
            },
            password: {
                identifier  : 'password',
                rules: [
                {
                    type   : 'empty',
                    prompt : '请输入密码'
                },
                {
                    type   : 'length[6]',
                    prompt : '密码至少6位'
                }
                ]
            }
        },
        onSuccess:function(){
            login();
            return false;
        }
    });
});

function login() {
    $.ajax({
        url:"/login",
        data: $("#login_form").serialize(),
        type:"post",
        success:function(data){
            if (data && data.error) {
                $("#login_form").removeClass("success").addClass("error")
                $("#login_error_msg ul").append("<li>"+data.error+"</li>");
            } else {
                window.location.href =  "/";
            }
        }
    });
}
