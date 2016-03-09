$(function(){
    function removeUsers() {
        $('#employee option').each(function(){
           if ($(this).val()) $(this).remove();
        });
    }

    var users = $('#employee option');
    function setUsers (val) {
        users.each(function(){
            if ($("#department  option:selected").attr("tag")) {
                if ($(this).attr('tag') && $(this).attr('tag').indexOf($("#department  option:selected").attr("tag")) > -1) {
                    $('#employee').append($(this));
                }
            }

            if (!val) {
                $('#employee').closest('.field').addClass('disabled');
            } else {
                $('#employee').closest('.field').removeClass('disabled');
            }
        });
    }

    $('#department').dropdown({
            placeholder: '请选择',
            onChange: function(value, text, $selectedItem) {
                removeUsers();
                setUsers(value);
            }
        })
    ;
    $('.ui.search.dropdown').dropdown({fullTextSearch:true})
    removeUsers();
    setUsers($("#department option:selected").text());

    $('#projectProgress').dropdown({placeholder:'优先级1为最高'});

    $('#projectBussiness').dropdown({placeholder:'请选择业务担当'});
    $('#projectArt').dropdown({placeholder:'请选择美术单元'});
    $('#projectTech').dropdown({placeholder:'请选择技术单元'});
    
    $('select.dropdown').dropdown({placeholder:'请选择'});
    $('.datetimepicker').datetimepicker({
        format:'Y-m-d H:i', step:5
    });
    $.datetimepicker.setLocale('zh');
    $('.datetimepicker').datetimepicker({format:'Y-m-d H:i', step:5});
}());

function selectFiles (id) {
    $('#'+id).click();
}

function setUploadFiles (e, id) {
    var strFiles = '';
    var files = $(e).prop('files');

    for(var i=0,len=files.length; i<len; i++){
        strFiles += '<label style="line-height: 24px;" class="pr20">' + files[i].name + '</label>';
    }

    $('#'+id+' .file.text').html(strFiles);

    if (files.length > 0) {
        $('#'+id).show();
    } else {
        $('#'+id).hide();
    }
}

function clearForm (formSel) {
    $(formSel).form('clear');
    $(formSel).submit();
}

function submitForm (formSel) {
    $(formSel).submit();
}

