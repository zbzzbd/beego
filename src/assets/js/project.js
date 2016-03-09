;
$(function() {
    $('#project_search').search({
        source: window.projectNames,
        searchFullText: true
    });

    $('#project_create_form')
        .form({
            fields: {
                started: {
                    identifier: 'started',
                    rules: [{
                        type: 'empty',
                        prompt: '请选择启动时间'
                    }]
                },
                name: {
                    identifier: 'name',
                    rules: [{
                        type: 'empty',
                        prompt: '请填写项目名称'
                    }]
                },
                bussiness_user: {
                    identifier: 'bussiness_user',
                    rules: [{
                        type: 'empty',
                        prompt: '请选择业务担当'
                    }]
                },
                art_user: {
                    identifier: 'art_user',
                    rules: [{
                        type: 'empty',
                        prompt: '请选择美术单元'
                    }]
                },
                tech_user: {
                    identifier: 'tech_user',
                    rules: [{
                        type: 'empty',
                        prompt: '请选择技术单元'
                    }]
                },
            },
            onSuccess: function() {
                $.ajax({
                    url: "/project/create",
                    data: new FormData($("#project_create_form")[0]),
                    async: false,
                    cache: false,
                    contentType: false,
                    processData: false,
                    type: "post",
                    success: function(data) {
                        if (data && data.result != 0) {
                            $(".ui.error.message").html(data.error);
                        } else if (data.id) {
                            $(".ui.error.message").html("");
                            window.location.href = '/project/list';
                        } else {
                            $(".ui.error.message").html("未知错误");
                        }
                    }
                });
                return false;
            }
        });
});
 
function delProject(id) {
    $('.ui.modal.delete').modal({
        closable  : false,
        onDeny: function(){
        },
        onApprove: function() {
            $.ajax({
                url:"/project/del/" + id,
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
