$(function() {
    $('.ui.form.user').form({
        on: 'blur',
        fields: {
            company: {
                identifier  : 'company',
                rules: [
                {
                    type   : 'empty',
                    prompt : '请选择公司名称'
                }
                ]
            },
            role: {
                identifier  : 'role',
                rules: [
                    {
                        type   : 'empty',
                        prompt : '请选择角色'
                    }
                ]
            },
            name: {
                identifier  : 'name',
                rules: [
                    {
                        type   : 'empty',
                        prompt : '请用户名称'
                    }
                ]
            },
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
        },
        onSuccess:function(){
            createUser();
            return false;
        }
    });

    $('.ui.form.user-edit').form({
        on: 'blur',
        fields: {
            company: {
                identifier  : 'company',
                rules: [
                    {
                        type   : 'empty',
                        prompt : '请选择公司名称'
                    }
                ]
            },
            role: {
                identifier  : 'role',
                rules: [
                    {
                        type   : 'empty',
                        prompt : '请选择角色'
                    }
                ]
            },
            name: {
                identifier  : 'name',
                rules: [
                    {
                        type   : 'empty',
                        prompt : '请用户名称'
                    }
                ]
            },
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
        },
        onSuccess:function(){
            editUser();
            return false;
        }
    });
});

function createUser() {
    $.ajax({
        url:"/user/create",
        data: $(".ui.form.user").serialize(),
        type:"post",
        success:function(data){
            if (data && data.error) {
                $(".ui.form.user").removeClass("success").addClass("error");
                $(".ui.error.message").html('<ul><li>' + data.error + '</li></ul>');
            } else {
                window.location.href =  "/user/list";
            }
        }
    });
}


function editUser() {
    $.ajax({
        url:"/user/edit/" + $("#user-id").val(),
        data: $(".ui.form.user-edit").serialize(),
        type:"post",
        success:function(data){
            if (data && data.error) {
                $(".ui.form.user").removeClass("success").addClass("error");
                $(".ui.error.message").html('<ul><li>' + data.error + '</li></ul>');
            } else {
                window.location.href =  "/user/list";
            }
        }
    });
}

function deleteUser(id) {
    $('.ui.modal.delete').modal({
        closable  : false,
        onDeny: function(){
        },
        onApprove: function() {
            $.ajax({
                url:"/user/delete/" + id,
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

function restoreUser(id) {
    $.ajax({
        url:"/user/restore/" + id,
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