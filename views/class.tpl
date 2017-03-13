{{define "class"}}

<h1>这里添加删除分类！！</h1>
<form class="form-inline">
    <div class="form-group">
        <input name="IsClass" value="1" style="display:none">
        <input type="text" class="form-control" name="addClassName" placeholder="分类名">
    <button type="submit" class="btn btn-default">添加</button>
        <input class="form-control" STYLE="width: 21em" id="disabledInput" type="text" placeholder="已存在的分类无法重复添加，请注意！！！" disabled>
    </div>
</form>

<form class="form-inline">
        <div class="form-group">
            <input name="IsClass" value="1" style="display:none">
            <select class="form-control" STYLE="width: 21em" name="delClassName" >
                <option>分类1</option>
                <option>分类2</option>
                <option>分类3</option>
                <option>分类4</option>
                <option>分类5</option>
            </select>
            <button type="submit" class="btn btn-default">删除</button>
            <input class="form-control" STYLE="width: 21em" id="disabledInput" type="text" placeholder="已存在内容的分类无法删除，请注意！！！" disabled>
        </div>
</form>
<h1>分类管理</h1>
<table class="table table-striped">
    <thead>
    <tr>
        <th>分类名称</th>
        <th>最新内容</th>
        <th>内容总数</th>
        <th>删除分类</th>
    </tr>
    </thead>
    <tbody>
    <tr>
        <td>基础资料</td>
        <td>论数据结构的重要性</td>
        <td>13</td>
        <td>
            <a href="#myModal" onclick=resetname() type="submit" role="button" data-toggle="modal"><i  class="fa fa-trash-o"></i></a>
        </td>
    </tr>
    <tr>
        <td>基础资料</td>
        <td>论数据结构的重要性</td>
        <td>13</td>
        <td>
            <a href="#myModal" onclick=resetname() type="submit" role="button" data-toggle="modal"><i  class="fa fa-trash-o"></i></a>
        </td>
    </tr>
    <tr>
        <td>基础资料</td>
        <td>论数据结构的重要性</td>
        <td>13</td>
        <td>
            <a href="#myModal" onclick=resetname() type="submit" role="button" data-toggle="modal"><i  class="fa fa-trash-o"></i></a>
        </td>
    </tr>
    <tr>
        <td>基础资料</td>
        <td>论数据结构的重要性</td>
        <td>13</td>
        <td>
            <a href="#myModal" onclick=resetname() type="submit" role="button" data-toggle="modal"><i  class="fa fa-trash-o"></i></a>
        </td>
    </tr>
    <tr>
        <td>基础资料</td>
        <td>论数据结构的重要性</td>
        <td>13</td>
        <td>
            <a href="#myModal" onclick=resetname() type="submit" role="button" data-toggle="modal"><i  class="fa fa-trash-o"></i></a>
        </td>
    </tr>
    <tr>
        <td>基础资料</td>
        <td>论数据结构的重要性</td>
        <td>13</td>
        <td>
            <a href="#myModal" onclick=resetname() type="submit" role="button" data-toggle="modal"><i  class="fa fa-trash-o"></i></a>
        </td>
    </tr>
    </tbody>
</table>
{{end}}