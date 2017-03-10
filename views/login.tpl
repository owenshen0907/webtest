<!DOCTYPE html>
<html lang="zh-cn">
{{template "head1"}}


<body>
<p>{{.usrname}}</p>
<p>{{.psw}}</p>

<div class="container">
		<h1><a href="/">主页</a></h1>
      <form class="form-signin" role="form" method="POST" action="/login">
        <h2 class="form-signin-heading">Please sign in</h2>
        <input type="text" class="form-control" name="usrname" placeholder="accountID" required autofocus>
        <input type="password" class="form-control" name="psw" placeholder="Password" required>
        <div class="checkbox">
          <label>
            <input type="checkbox" value="on" name="autoLogin"> Remember me
          </label>
        </div>
        <button class="btn btn-lg btn-primary btn-block" type="submit" >登录</button>
		<a class="btn btn-lg btn-primary btn-block" href="/register">注册</a>
      </form>

    </div> <!-- /container -->
</body>
</html>