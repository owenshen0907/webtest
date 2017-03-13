{{define "resource"}}
<h1>资源上传管理</h1>


<h3>请选择文件类型：</h3>
<label class="radio-inline">
    <input type="radio" name="inlineRadioOptions" id="inlineRadio1" value="option1"> PDF
</label>
<label class="radio-inline">
    <input type="radio" name="inlineRadioOptions" id="inlineRadio2" value="option2"> PPT
</label>
<label class="radio-inline">
    <input type="radio" name="inlineRadioOptions" id="inlineRadio3" value="option3"> WORLD
</label>
<label class="radio-inline">
    <input type="radio" name="inlineRadioOptions" id="inlineRadio3" value="option3"> MP4
</label>

<form id="fform" method="POST" enctype="multipart/form-data">
    <input id="myfile" name="myfile" type="file" />
    <input type="submit" value="保存"  />
</form>
{{end}}