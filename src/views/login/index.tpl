<div class="ui center aligned right grid" style="margin: auto;width: 450px;display: table;vertical-align: middle;height: 100%;">
  <form id="login_form" class="ui large form" action="/login" method="post" style="display: table-cell; vertical-align: middle;width: 450px;">
    <div class="text-center" style="margin-top: -70px;">
      <img style="width: 190px;" class="mb10" src="/img/chinarun.png">
      <div style="font-size: 32px; color: #28ad75;">项目管理系统</div>
    </div>
    <div class="ui stacked segment">
      <div class="field">
        <div class="ui left icon input">
          <i class="mail icon"></i>
          <input type="text" name="email" placeholder="邮箱">
        </div>
      </div>
      <div class="field">
        <div class="ui left icon input">
          <i class="lock icon"></i>
          <input type="password" name="password" placeholder="密码">
        </div>
      </div>
      <input type="submit" class="ui fluid large submit button" value="登录" style="background-color: #28ad75;color:#fff;">
    </div>

    <div class="ui error message" id="login_error_msg">
      <ul class="list">
      </ul>
    </div>
  </form>
</div>

