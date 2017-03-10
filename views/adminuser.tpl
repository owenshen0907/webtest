{{define "rootuser"}}
<div class="btn-toolbar list-toolbar">
    <!---button class="btn btn-primary"><i class="fa fa-plus"></i> 添加用户</button---->
    <button class="btn btn-default">Import导入</button>
    <button class="btn btn-default">Export导出</button>
  <div class="btn-group">
  </div>
</div>
<table class="table">
  <thead>
    <tr>
      <th>ID#</th>
      <th>用户名</th>
      <th>昵称</th>
      <th>权限</th>
      <th style="width: 3.5em;">Modify</th>
    </tr>
  </thead>
  <tbody>

    <tr>
      <td>1</td>
      <td>root</td>
      <td>超级管理员</td>
      <td>激活</td>
      <td>
          <a href="user.html"><i class="fa fa-pencil"></i></a>
		<input name="qqqqq" value="tom" style="display:none">
        <!---  <a href="#myModal" onclick=resetname() type="submit" role="button" data-toggle="modal"><i  class="fa fa-trash-o"></i></a>-->
	   <button href="#myModal" onclick=resetname() type="submit" role="button" data-toggle="modal"><i  class="fa fa-trash-o"></i></button>

      </td>
    </tr>

 	{{str2html .UsrInfo}}
<script>
function resetname(){
	var tmp = document.getElementsByName("qqqqq")[0].value
	document.getElementsByName("use").item(0).textContent = tmp
}
</script>
  </tbody>
</table>

<ul class="pagination">
<!---
  <li><a href="#">&laquo;</a></li>
  <li><a href="#">1</a></li>
  <li><a href="#">2</a></li>
  <li><a href="#">3</a></li>
  <li><a href="#">4</a></li>
  <li><a href="#">5</a></li>
  <li><a href="#">&raquo;</a></li>
--->
{{str2html .Page}}
</ul>

<div class="modal small fade" id="myModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
  <div class="modal-dialog">
    <div class="modal-content">
        <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
            <h3 id="myModalLabel">确认删除</h3>
        </div>
        <div class="modal-body">
            <p class="error-text"><i class="fa fa-warning modal-icon"></i>你确定要删掉该<span name="use"></span>么？<br>该动作无法撤销，请谨慎操作.</p>
        </div>
        <div class="modal-footer">
            <button class="btn btn-default" data-dismiss="modal" aria-hidden="true">取消</button>
            <button class="btn btn-danger" data-dismiss="modal">确认</button>
        </div>
      </div>
    </div>
</div>

{{end}}

{{define "rootuseredit"}}
<ul class="nav nav-tabs">
  <li class="active"><a href="#home" data-toggle="tab">个人资料</a></li>
  <li><a href="#profile" data-toggle="tab">安全设置</a></li>
</ul>

<div class="row">
  <div class="col-md-4">
    <br>
    <div id="myTabContent" class="tab-content">
      <div class="tab-pane active in" id="home">
      <form id="tab" method="post">
        <div class="form-group">
        <label >账户</label>
		<input name="IsUsersEdit" value="1" style="display:none">
		<input name="IsUsersInfo" value="1" style="display:none">
        <input name="uid" type="text" value="{{.IsId}}" class="form-control" readonly>
        </div>
        <div class="form-group">
        <label>昵称</label>
        <input name="UpdataName" type="text" value="{{.IsName}}" class="form-control">
        </div>
        <div class="form-group">
        <label>电话</label>
        <input name="UpdataTel" type="text" value="{{.IsTel}}" class="form-control">
        </div>
        <div class="form-group">
        <label>Email</label>
        <input name="UpdataEmail" type="text" value="{{.IsEmail}}" class="form-control">
        </div>

        <div class="form-group">
          <label>个性签名</label>
          <textarea value="Smith" rows="3" class="form-control">永远自由自我
永远高唱我歌
告别千里</textarea>
        </div>

    <div class="btn-toolbar list-toolbar">
      <button class="btn btn-primary" type="submit"><i class="fa fa-save"></i>保存</button>
    </div>
        </form>
      </div>

      <div class="tab-pane fade" id="profile">

        <form method="post">
          <div class="form-group">
            <label >原密码</label>
			<input name="IsUsersEdit" value="1" style="display:none">
			<input name="IsUsersPsw" value="1" style="display:none">
            <input type= "password" name="Psw"
            {{if .IsPsw}} placeholder="请输入正确的密码" {{end}}  class="form-control">
            <label>新密码</label>
            <input data-val-length-max="10" data-val-length-min="6" type= "password" name="Psw1"
            {{if .IsPsw1}} placeholder="两次密码不一致" {{end}} class="form-control">
            <label>确认密码</label>
            <input data-val-length-max="10" data-val-length-min="6" type= "password" name="Psw2"
            {{if .IsPsw1}} placeholder="两次密码不一致" {{end}} class="form-control">
            <input name="uid" type="text" value="{{.IsId}}" class="form-control" style="display:none">
          </div>
          <div>
			<!--a href="#myModal" data-toggle="modal" class="btn btn-danger">更新</a--->
			 <button type="submit"  data-toggle="modal" class="btn btn-danger" >更新</button>
          </div>
        </form>
      </div>
    </div>
<!----
    <div class="btn-toolbar list-toolbar">
      <button >保存</button>
      <a href="#myModal" data-toggle="modal" class="btn btn-danger">重置</a>
    </div>
--->	
  </div>
</div>

<div class="modal small fade" id="myModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <button type="button" class="close" data-dismiss="modal" aria-hidden="true">×</button>
        <h3 id="myModalLabel">Delete Confirmation</h3>
      </div>
      <div class="modal-body">
        
        <p class="error-text"><i class="fa fa-warning modal-icon"></i>Are you sure you want to delete the user?</p>
      </div>
      <div class="modal-footer">
        <button class="btn btn-default" data-dismiss="modal" aria-hidden="true">取消</button>
        <button class="btn btn-danger" data-dismiss="modal">确认更新密码</button>
      </div>
    </div>
  </div>
</div>

{{end}}