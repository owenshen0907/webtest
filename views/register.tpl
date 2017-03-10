<!DOCTYPE html>
<html lang="zh-cn">
<head>
 {{template "head1"}}
</head>
<body>
<p>{{.usrname}}</p>
<p>{{.psw}}</p>


<div class="container">
<div class="jumbotron">
	<h1><a href="/">主页</a></h1>
      
	<form action="/register" class="form-horizontal" method="post" role="form">
	<h4>创建一个新账户</h4>
    <hr />
<div class="validation-summary-valid text-danger" data-valmsg-summary="true">
	<ul>
		<li style="display:none"></li>
	</ul>
</div>    
	<div class="col-lg-6 ">
        <div class="form-group {{if .CheakUsr}}has-warning has-feedback{{end}} ">
            <label class="control-label" for="UserName">用户名</label>
            <input class="form-control" data-val="true" id="FirstName" name="user" type="text" {{if .CheakUsr}} placeholder="用户名已存在，请尝试更换。"{{else}}placeholder="用户名／邮箱／手机" {{end}}/>
        </div>
        <div class="form-group {{if .CheakEmail}}has-error has-feedback{{end}}">
            <label class="control-label" for="Email">邮箱</label>
                <input class="form-control" data-val="true" id="LastName" name="email" type="text" {{if .CheakEmail}} placeholder="请按正确格式填写邮箱"{{else}}placeholder="eg:somebody@gmail.com" {{end}}/>
                <span class="glyphicon glyphicon-remove form-control-feedback" aria-hidden="true"></span>
               <span id="inputError2Status" class="sr-only">(error)</span>

        </div>
        <div class="form-group {{if .CheakTel}}has-error has-feedback{{end}}">
            <label class="control-label" for="Email">手机号</label>
                <input class="form-control" data-val="true" id="Email" name="tel" type="text" {{if .CheakTel}}id="checkboxWarning" placeholder="请填写真实手机号"{{else}}placeholder="13888888888" {{end}}/>
                <span class="glyphicon glyphicon-remove form-control-feedback" aria-hidden="true"></span>
               <span id="inputError2Status" class="sr-only">(error)</span>

        </div>
        <div class="form-group">
            <label class="control-label" for="Password">密码</label>
            <div>
                <input class="form-control" data-val="true" data-val-length="The Password must be at least 6 characters long." data-val-length-max="100" data-val-length-min="6" data-val-required="The Password field is required." id="Password" name="psw" type="password" />
            </div>
        </div>
        <div class="form-group">
            <div class="g-recaptcha" data-sitekey="6LfMrBQUAAAAAIrQhFoGbKs2vzf9ta6KRn_nHryR"></div>
        </div>
        <div class="form-group">
            <div class="">
                <input type="submit" class="btn btn-lg btn-success btn-block" value="提交" />
            </div>
        </div>
    </div>
    
</form>
</div>
    </div> <!-- /container -->
<script src='https://www.google.com/recaptcha/api.js'></script>

</body>
</html>