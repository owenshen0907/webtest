{{define "class"}}

<h1>这里添加删除分类！！</h1>
<form class="form-inline" method="post">
    <div class="form-group">

        <input type="text" class="form-control" name="addClassName" placeholder="分类名">
        <button type="submit" class="btn btn-default">添加</button>
        <input class="form-control" STYLE="width: 21em" id="disabledInput" type="text" placeholder="已存在的分类无法重复添加，请注意！！！" disabled>
    </div>
</form>

<form class="form-inline" method="post">
        <div class="form-group">
            <select class="form-control" STYLE="width: 14em" name="delClassName" >
                <!---option>分类1</option--->
                {{str2html .ClassList}}
            </select>
            <button type="submit" class="btn btn-default">删除</button>
            <input class="form-control" STYLE="width: 21em" id="disabledInput" type="text" placeholder="已存在内容的分类无法删除，请注意！！！" disabled>

        </div>
</form>
<form class="form-inline" method="post">
    <div class="form-group">
        <input name="IsClass" value="1" style="display:none">
        <select class="form-control" STYLE="width: 14em" name="sourceClassName" >
            <!---option>分类1</option--->
            {{str2html .ClassList}}
        </select>
        <input type="text" class="form-control" name="updateClassName" placeholder="分类名">
        <button type="submit" class="btn btn-default">修改</button>
    </div>
</form>
<h1>分类管理</h1>
<table class="table table-striped">
    <thead>
    <tr>
        <th>分类名称</th>
        <th>最新内容</th>
        <th>内容总数</th>
    </tr>
    </thead>
    <tbody>
    <!---tr>
        <td>基础资料</td>
        <td>论数据结构的重要性</td>
        <td>13</td>
    </tr--->
    {{str2html .ClassListInfo}}
    </tbody>
</table>
{{end}}