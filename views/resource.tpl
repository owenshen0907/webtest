{{define "resource"}}
<h1>资源上传管理</h1>

<h3>请选择文件类型：</h3>

<form method="POST" enctype="multipart/form-data">
    <select class="form-control" STYLE="width: 14em" name="selectClassName" >
        <!---option>分类1</option--->
        {{str2html .ClassList}}
    </select>
    <select class="form-control" STYLE="width: 14em" name="selectTagName" >
        <!---option>分类1</option--->
        {{str2html .TagList}}
    </select>
    <label class="radio-inline">
        <input type="radio" name="Options" id="inlineRadio1" value="PDF"> PDF
    </label>
    <label class="radio-inline">
        <input type="radio" name="Options" id="inlineRadio2" value="PPT"> PPT
    </label>
    <label class="radio-inline">
        <input type="radio" name="Options" id="inlineRadio3" value="WORLD"> WORLD
    </label>
    <label class="radio-inline">
        <input type="radio" name="Options" id="inlineRadio4" value="MP4"> MP4
    </label>
    <input id="myfile" name="myfile" type="file" />
    <input type="text" class="form-control" name="resourceTitle" placeholder="资源名">
    <input type="submit" value="保存"  />
</form>
{{end}}