function addRequire() {
    $('.ui.modal.edit').modal({
        closable  : false,
        onDeny: function(){
        },
        onApprove: function() {
            $.ajax({
                url:"/require/create",
                data: {name: $("#modal-input").val()},
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

function editRequire(name) {
    $("#modal-input").val(name);

    $('.ui.modal.edit').modal({
        closable  : false,
        onDeny: function(){
        },
        onApprove: function() {
            $.ajax({
                url:"/require/edit",
                data: {name: name, newName: $("#modal-input").val()},
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


function deleteRequire(name) {
    $('.ui.modal.delete').modal({
        closable  : false,
        onDeny: function(){
        },
        onApprove: function() {
            $.ajax({
                url:"/require/delete",
                data: {name: name},
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

function filterRequire(obj) {
    var txtFind = $(obj).val().toLowerCase();

    $("#require-list>tr").each(function(){
        var txt = $(this).find(".name").text().toLowerCase();
        if (txt.indexOf(txtFind) < 0) {
            $(this).hide();
        } else {
            $(this).show();
        }
    });
}
