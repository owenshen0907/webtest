{{define "editor"}}
<!---script type="text/javascript" src="static/dist/js/lib/jquery-1.10.2.min.js"></script---->
<script type="text/javascript" src="static/dist/js/wangEditor.min.js"></script>
<div style="width:95%; margin:0 auto;">
    <form action="localhost:8088/" method="post">

        <select class="form-control" STYLE="width: 21em" >
            {{str2html .ClassList}}
        </select>
        <input type="text" class="form-control" STYLE="width: 21em"  placeholder="作者">
        <input type="text" class="form-control" STYLE="width: 21em"  placeholder="编辑">
        <input type="text" class="form-control" STYLE="width: 21em"  placeholder="标签">
        <textarea class="form-control" rows="3" placeholder="摘要！"></textarea>

<textarea id="textarea1" name="content" type="text" style="height:400px;max-height:500px;">

        <p>请输入内容...</p>
</textarea>
        <input id="btn1" value="提交" type="submit"/>
    </form>
    <button id="btn5">清空内容</button>
</div>
<script type="text/javascript">
    wangEditor.config.mapAk ='BIspAv95IM97ymtyyE36hmlhjd5n3IvX'
    var textarea = document.getElementById('textarea1');
    var editor = new wangEditor('textarea1');
    $('#btn1').click(function () {
        // 获取编辑器区域完整html代码
        var html = editor.$txt.html();
        // 获取编辑器纯文本内容
        // var text = editor.$txt.text();
        // 获取格式化后的纯文本
        // var formatText = editor.$txt.formatText();
    });
    $('#btn2').click(function () {
        // 销毁编辑器
        editor.destroy();
    });
    $('#btn3').click(function () {
        // 恢复编辑器
        editor.undestroy();
    });
    $('#btn4').click(function () {
        editor.$txt.append('<p>新追加的内容</p>');
    });
    $('#btn5').click(function () {
        // 清空内容。
        // 不能传入空字符串，而必须传入如下参数
        editor.$txt.html('<p><br></p>');
    });
    $('#btn6').click(function () {
        // 切换成英文
        editor.config.lang = wangEditor.langs['en'];
    });

    editor.config.uploadImgUrl = '/';
    editor.config.uploadImgFileName = 'myfile'
    editor.config.uploadHeaders = {
        'Accept' : 'text/x-json'
    };
    // editor.config.hideLinkImg = true;
    editor.config.emotions = {
        // 支持多组表情

        // 第一组，id叫做 'default'
        'default': {
            title: '默认',  // 组名称
            data: 'http://www.wangeditor.com/wangEditor/test/emotions.data'  // 服务器的一个json文件url
        },
        // 第二组，id叫做'weibo'
        // 下面还可以继续，第三组、第四组、、、
    };
    editor.onchange = function () {
        // 编辑区域内容变化时，实时打印出当前内容
        console.log(this.$txt.html());
    };
    editor.create();
</script>
{{end}}