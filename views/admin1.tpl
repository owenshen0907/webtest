<!doctype html>
<html lang="en">
<head>
{{template "roothead"}}

</head>
<body class=" theme-blue">

    <!-- Demo page code -->

    <script type="text/javascript">
        $(function() {
            var match = document.cookie.match(new RegExp('color=([^;]+)'));
            if(match) var color = match[1];
            if(color) {
                $('body').removeClass(function (index, css) {
                    return (css.match (/\btheme-\S+/g) || []).join(' ')
                })
                $('body').addClass('theme-' + color);
            }

            $('[data-popover="true"]').popover({html: true});
            
        });
    </script>
    <style type="text/css">
        #line-chart {
            height:300px;
            width:800px;
            margin: 0px auto;
            margin-top: 1em;
        }
        .navbar-default .navbar-brand, .navbar-default .navbar-brand:hover { 
            color: #fff;
        }
    </style>

    <script type="text/javascript">
        $(function() {
            var uls = $('.sidebar-nav > ul > *').clone();
            uls.addClass('visible-xs');
            $('#main-menu').append(uls.clone());
        });
    </script>

    <!-- Le HTML5 shim, for IE6-8 support of HTML5 elements -->
    <!--[if lt IE 9]>
      <script src="http://html5shim.googlecode.com/svn/trunk/html5.js"></script>
    <![endif]-->

    <!-- Le fav and touch icons -->
    <link rel="shortcut icon" href="../assets/ico/favicon.ico">
    <link rel="apple-touch-icon-precomposed" sizes="144x144" href="../assets/ico/apple-touch-icon-144-precomposed.png">
    <link rel="apple-touch-icon-precomposed" sizes="114x114" href="../assets/ico/apple-touch-icon-114-precomposed.png">
    <link rel="apple-touch-icon-precomposed" sizes="72x72" href="../assets/ico/apple-touch-icon-72-precomposed.png">
    <link rel="apple-touch-icon-precomposed" href="../assets/ico/apple-touch-icon-57-precomposed.png">
  

  <!--[if lt IE 7 ]> <body class="ie ie6"> <![endif]-->
  <!--[if IE 7 ]> <body class="ie ie7 "> <![endif]-->
  <!--[if IE 8 ]> <body class="ie ie8 "> <![endif]-->
  <!--[if IE 9 ]> <body class="ie ie9 "> <![endif]-->
  <!--[if (gt IE 9)|!(IE)]><!--> 
   
  <!--<![endif]-->

{{template "rootnav" .}}

{{template "rootsidenav" .}}

<div class="content">
        <div class="header">
            <div class="stats">
   				 <p class="stat"><span class="label label-info">{{.UsrAccount}}</span> 在线人数</p>
   				 <p class="stat"><span class="label label-success">27</span> 资源数</p>
    			 <p class="stat"><span class="label label-danger">15</span> 文章数</p>
			</div>

            <h1 class="page-title">Dashboard</h1>
            <ul class="breadcrumb">
           		 <li><a href="/">Home</a> </li>
           		 <li class="active">控制面板</li>
       		 </ul>

        </div>
        <div class="main-content">
            {{if .IsUsers}}{{template "rootuser" .}}{{end}}
			{{if .IsDash}}{{template "rootdash" .}}{{end}}
			{{if .IsUsersEdit}}{{template "rootuseredit" .}}{{end}}
			{{if .IsHelp}}{{end}}
			{{if .IsContent}}{{end}}
			{{if .IsArticle}}{{end}}
			{{if .IsResource}}{{end}}
			{{if .IsSlide}}{{end}}
			{{if .IsSql}}{{end}}
   			


			{{template "rootfooter"}}
    </div>
 </div>


    <script src="static/admin/lib/bootstrap/js/bootstrap.js"></script>
    <script type="text/javascript">
        $("[rel=tooltip]").tooltip();
        $(function() {
            $('.demo-cancel-click').click(function(){return false;});
        });
    </script>
    
  
</body></html>
