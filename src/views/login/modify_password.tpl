<div class="ui center aligned right grid" style="margin: auto; width: 450px">
  <div class="column">
    <h2 class="ui teal image header">
      <div class="content">
         
        <h1>修改密码</h1>
      </div>
    </h2>
    <form class="ui large form" id="modifyPassword_form">
      <div class="ui stacked segment">
        <div class="field">
          <div class="ui left icon input">
            <i class="lock icon"></i>
             <input type="password" name="oldPassword" placeholder="输入旧密码" onchange="">
          </div>
        </div>
        <div class="field">
          <div class="ui left icon input">
            <i class="lock icon"></i>
             <input type="password" name="newPassword" placeholder="输入新密码">
          </div>
          <div class="ui left icon input">
            <i class="lock icon"></i>
             <input type="password" name="newPassword_confirm" placeholder="确认新密码">
          </div>
        </div>
        <input type="submit" class="ui fluid large submit button" style="background-color: #28ad75;color:#fff;" value="确定">
      </div>
    
      <div class="ui error message"> </div>

    </form>

  </div>
</div>

<script>
    window.onload=function(){
        $(function(){ 
            $('.ui.large.form')
                    .form({
                        fields: { 
                                                 
                            oldPassword: {
                                identifier: 'oldPassword',
                                rules: [
                                    {
                                        type   : 'empty',
                                        prompt : '请输入旧密码'
                                    }
                                ]
                            },
                             newPassword: {
                              identifier: 'newPassword',
                                rules :[
                                 {
                                   type : 'empty',
                                   prompt : '请输入新密码'
                                 }
                              ]
                            },

                             newPassword_confirm: {
                                identifier: 'newPassword_confirm',
                                rules: [
                                    {
                                        type   : 'empty',
                                        prompt : '请输入确认新密码'
                                    }
                                ]
                            }
                           


                        },
                        onSuccess: function() {
                            $.ajax({
                                url:"/modify_password",
                                data: new FormData($(".ui.large.form")[0]),
                                async: false,
                                cache: false,
                                contentType: false,
                                processData: false,
                                type:"post",
                                success:function(data){                                   
                                    if (data && data.error) {
                                        $(".ui.error.message").html(data.error);
                                        $(".ui.error.message").show()
                                    }
                                    else if (data.id) {
                                        $(".ui.error.message").html("修改密码成功"); 
                                        $('.ui.large.form').form('clear')
                                        $(".ui.error.message").show();
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