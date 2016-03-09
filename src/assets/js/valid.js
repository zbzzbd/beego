$(function() {
    function is_job_needed_done() {
        return $('#valid_form input:radio[name="result"]:checked').val() == 1
    }

    function strDateToDate(value){
        if(value.indexOf("T") != -1) {
            return new Date(value)
        }
        var ft = value.split(" ").join("T")
        return new Date(ft)
    }

    $.fn.form.settings.rules.finishTime = function(value) {
        if (!is_job_needed_done()) {
            return true
        }

        if (value == ""){
            return true
        }

        var required_finish_time = strDateToDate($('#job_required_finish_time').val())
        var r = strDateToDate(value) <= required_finish_time
        return r
    }

    $.fn.form.settings.rules.validTime = function(value) {
        if (is_job_needed_done()) {
            if ($('.ui.form .datetimepicker').val() === '') {
                return false
            }
        }
        return true
    }

    $('#valid_form').form({
        fields: {
            result: {
                identifier: 'result',
                rules: [{
                    type: 'checked',
                    prompt: '审核结果不能为空'
                }]
            },
            finish_time: {
                identifier: 'finish_time',
                rules: [{
                    type: 'validTime',
                    prompt: '任务完成时间不能为空'
                }, {
                    type: 'finishTime',
                    prompt: '要求完成时间必须早于验收时间'
                }]
            }
        },
        onSuccess:function(){
            jobValid();
            return false;
        }
    });
});

function jobValid() {
    $.ajax({
    url:"/project/job/valid",
    data: $("#valid_form").serialize(),
    type:"post",
    success:function(data){
        console.log(data)
        if (data && data.error) {
            $("#valid_form").removeClass("success").addClass("error")
            $("#valid_error_msg ul li").remove()
            $("#valid_error_msg ul").append("<li>"+data.error+"</li>");
        } else {
            window.location.href =  "/project/job/valid";
        }
    }
    });
}
