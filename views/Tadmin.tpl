{{define "roothead"}}
 <meta charset="utf-8">
    <title>Bootstrap Admin</title>
    <meta content="IE=edge,chrome=1" http-equiv="X-UA-Compatible">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="">
    <meta name="author" content="">

    <!----link href='http://fonts.useso.com/css?family=Open+Sans:400,700' rel='stylesheet' type='text/css'--->
    <link rel="stylesheet" type="text/css" href="static/admin/lib/bootstrap/css/bootstrap.css">
    <link rel="stylesheet" href="static/admin/lib/font-awesome/css/font-awesome.css">

    <script src="static/admin/lib/jquery-1.11.1.min.js" type="text/javascript"></script>
    <script src="static/admin/lib/jQuery-Knob/js/jquery.knob.js" type="text/javascript"></script>
    <script type="text/javascript">
        $(function() {
            $(".knob").knob();
        });
    </script>
    <link rel="stylesheet" type="text/css" href="static/admin/stylesheets/theme.css">
    <link rel="stylesheet" type="text/css" href="static/admin/stylesheets/premium.css">

      <link rel="stylesheet" type="text/css" href="static/dist/css/wangEditor.min.css">

{{end}}

{{define "rootnav"}}
    <div class="navbar navbar-default" role="navigation">
        <div class="navbar-header">
          <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target=".navbar-collapse">
            <span class="sr-only">Toggle navigation</span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </button>
          <a class="" href="index.html"><span class="navbar-brand"><span class="fa fa-paper-plane"></span> 武汉东湖学院</span></a></div>

        <div class="navbar-collapse collapse" style="height: 1px;">
          <ul id="main-menu" class="nav navbar-nav navbar-right">
            <li class="dropdown hidden-xs">
                <a href="#" class="dropdown-toggle" data-toggle="dropdown">
                    <span class="glyphicon glyphicon-user padding-right-small" style="position:relative;top: 3px;"></span> {{.UsrName}}
                    <i class="fa fa-caret-down"></i>
                </a>

              <ul class="dropdown-menu">
                <li><a href="./">用户信息</a></li>
                <li class="divider"></li>
                <li class="dropdown-header">Admin Panel</li>
                <li><a href="./">Users</a></li>
                <li><a href="/">Home</a></li>
                <li><a tabindex="-1" href="./">Payments</a></li>
                <li class="divider"></li>
                <li><a tabindex="-1" href="/">Logout</a></li>
              </ul>
            </li>
          </ul>

        </div>
      </div>
    </div>
    <div class="copyrights">Collect from <a href="#"  title="owenshen">owenshen123@outlook.com</a></div>
{{end}}

{{define "rootsidenav"}}
    <div class="sidebar-nav">
    <ul>
    <li><a href="#" data-target=".dashboard-menu" class="nav-header" data-toggle="collapse"><i class="fa fa-fw fa-dashboard"></i> 控制面板<i class="fa fa-collapse"></i></a></li>
    <li><ul class="dashboard-menu nav nav-list collapse in">
            <li><a href="/admin?IsDash=1"><span class="fa fa-caret-right"></span> 综合</a></li>
            <li ><a href="/admin?IsUsers=1"><span class="fa fa-caret-right"></span> 用户列表</a></li>
            <li ><a href="/admin?IsUsersEdit=1"><span class="fa fa-caret-right"></span> 个人档案</a></li>
            <li ><a href="/admin?IsClass=1"><span class="fa fa-caret-right"></span>资源分类管理器</a></li>
            <li ><a href="/admin?IsResource=1"><span class="fa fa-caret-right"></span>资源管理器</a></li>
            <li ><a href="/admin?IsEditor=1"><span class="fa fa-caret-right"></span>文章内容编辑器</a></li>

    </ul></li>

    <li data-popover="true" data-content="Items in this group require a <strong><a href='http://portnine.com/bootstrap-themes/aircraft' target='blank'>premium license</a><strong>." rel="popover" data-placement="right"><a href="#" data-target=".premium-menu" class="nav-header collapsed" data-toggle="collapse"><i class="fa fa-fw fa-fighter-jet"></i> Premium Features<i class="fa fa-collapse"></i></a></li>
        <li><ul class="premium-menu nav nav-list collapse">
                <li class="visible-xs visible-sm"><a href="#">- Premium features require a license -</a></span>
            <li ><a href="premium-profile.html"><span class="fa fa-caret-right"></span> Enhanced Profile</a></li>
            <li ><a href="premium-blog.html"><span class="fa fa-caret-right"></span> Blog</a></li>

    </ul></li>

        <li><a href="#" data-target=".accounts-menu" class="nav-header collapsed" data-toggle="collapse"><i class="fa fa-fw fa-briefcase"></i> Account <span class="label label-info">+3</span></a></li>
        <li><ul class="accounts-menu nav nav-list collapse">
            <li ><a href="sign-in.html"><span class="fa fa-caret-right"></span> Sign In</a></li>
            <li ><a href="sign-up.html"><span class="fa fa-caret-right"></span> Sign Up</a></li>
            <li ><a href="reset-password.html"><span class="fa fa-caret-right"></span> Reset Password</a></li>
    </ul></li>

        <li><a href="#" data-target=".legal-menu" class="nav-header collapsed" data-toggle="collapse"><i class="fa fa-fw fa-legal"></i> Legal<i class="fa fa-collapse"></i></a></li>
        <li><ul class="legal-menu nav nav-list collapse">
            <li ><a href="privacy-policy.html"><span class="fa fa-caret-right"></span> Privacy Policy</a></li>
            <li ><a href="terms-and-conditions.html"><span class="fa fa-caret-right"></span> Terms and Conditions</a></li>
    </ul></li>

        <li><a href="help.html" class="nav-header"><i class="fa fa-fw fa-question-circle"></i> Help</a></li>
            <li><a href="faq.html" class="nav-header"><i class="fa fa-fw fa-comment"></i> Faq</a></li>
                <li><a href="http://portnine.com/bootstrap-themes/aircraft" class="nav-header" target="blank"><i class="fa fa-fw fa-heart"></i> Get Premium</a></li>
            </ul>
    </div>
{{end}}

{{define "rootdash"}}
<div class="panel panel-default">
      			  <a href="#page-stats" class="panel-heading" data-toggle="collapse">Latest Stats</a>
      			  <div id="page-stats" class="panel-collapse panel-body collapse in">

                    <div class="row">
                        <div class="col-md-3 col-sm-6">
                            <div class="knob-container">
                                <input class="knob" data-width="200" data-min="0" data-max="3000" data-displayPrevious="true" value="2500" data-fgColor="#92A3C2" data-readOnly=true;>
                                <h3 class="text-muted text-center">Accounts</h3>
                            </div>
                        </div>
                        <div class="col-md-3 col-sm-6">
                            <div class="knob-container">
                                <input class="knob" data-width="200" data-min="0" data-max="4500" data-displayPrevious="true" value="3299" data-fgColor="#92A3C2" data-readOnly=true;>
                                <h3 class="text-muted text-center">Subscribers</h3>
                            </div>
                        </div>
                        <div class="col-md-3 col-sm-6">
                            <div class="knob-container">
                                <input class="knob" data-width="200" data-min="0" data-max="2700" data-displayPrevious="true" value="1840" data-fgColor="#92A3C2" data-readOnly=true;>
                                <h3 class="text-muted text-center">Pending</h3>
                            </div>
                        </div>
                        <div class="col-md-3 col-sm-6">
                            <div class="knob-container">
                                <input class="knob" data-width="200" data-min="0" data-max="15000" data-displayPrevious="true" value="10067" data-fgColor="#92A3C2" data-readOnly=true;>
                                <h3 class="text-muted text-center">Completed</h3>
                            </div>
                        </div>
                    </div>
       			 </div>
  			  </div>
<!--***********************--->
<div class="row">
    <div class="col-sm-6 col-md-6">
        <div class="panel panel-default">
            <div class="panel-heading no-collapse">Not Collapsible<span class="label label-warning">+10</span></div>
            <table class="table table-bordered table-striped">
              <thead>
                <tr>
                  <th>First Name</th>
                  <th>Last Name</th>
                  <th>Username</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td>Mark</td>
                  <td>Tompson</td>
                  <td>the_mark7</td>
                </tr>
                <tr>
                  <td>Ashley</td>
                  <td>Jacobs</td>
                  <td>ash11927</td>
                </tr>
                <tr>
                  <td>Audrey</td>
                  <td>Ann</td>
                  <td>audann84</td>
                </tr>
                <tr>
                  <td>John</td>
                  <td>Robinson</td>
                  <td>jr5527</td>
                </tr>
                <tr>
                  <td>Aaron</td>
                  <td>Butler</td>
                  <td>aaron_butler</td>
                </tr>
                <tr>
                  <td>Chris</td>
                  <td>Albert</td>
                  <td>cab79</td>
                </tr>
              </tbody>
            </table>
        </div>
    </div>
    <div class="col-sm-6 col-md-6">
        <div class="panel panel-default">
            <a href="#widget1container" class="panel-heading" data-toggle="collapse">Collapsible </a>
            <div id="widget1container" class="panel-body collapse in">
                <h2>Here's a Tip</h2>
                <p>This template was developed with <a href="http://middlemanapp.com/" target="_blank">Middleman</a> and includes .erb layouts and views.</p>
                <p>All of the views you see here (sign in, sign up, users, etc) are already split up so you don't have to waste your time doing it yourself!</p>
                <p>The layout.erb file includes the header, footer, and side navigation and all of the views are broken out into their own files.</p>
                <p>If you aren't using Ruby, there is also a set of plain HTML files for each page, just like you would expect.</p>
            </div>
        </div>
    </div>
</div>

<div class="row">
    <div class="col-sm-6 col-md-6">
        <div class="panel panel-default"> 
            <div class="panel-heading no-collapse">
                <span class="panel-icon pull-right">
                    <a href="#" class="demo-cancel-click" rel="tooltip" title="Click to refresh"><i class="fa fa-refresh"></i></a>
                </span>

                Needed to Close
            </div>
            <table class="table list">
              <tbody>
                  <tr>
                      <td>
                          <a href="#"><p class="title">Care Hospital</p></a>
                          <p class="info">Sales Rating: 86%</p>
                      </td>
                      <td>
                          <p>Date: 7/19/2012</p>
                          <a href="#">View Transaction</a>
                      </td>
                      <td>
                          <p class="text-danger h3 pull-right" style="margin-top: 12px;">$20,500</p>
                      </td>
                  </tr>
                  <tr>
                      <td>
                          <a href="#"><p class="title">Custom Eyesight</p></a>
                          <p class="info">Sales Rating: 58%</p>
                      </td>
                      <td>
                          <p>Date: 7/19/2012</p>
                          <a href="#">View Transaction</a>
                      </td>
                      <td>
                          <p class="text-danger h3 pull-right" style="margin-top: 12px;">$12,600</p>
                      </td>
                  </tr>
                  <tr>
                      <td>
                          <a href="#"><p class="title">Clear Dental</p></a>
                          <p class="info">Sales Rating: 76%</p>
                      </td>
                      <td>
                          <p>Date: 7/19/2012</p>
                          <a href="#">View Transaction</a>
                      </td>
                      <td>
                          <p class="text-danger h3 pull-right" style="margin-top: 12px;">$2,500</p>
                      </td>
                  </tr>
                  <tr>
                      <td>
                          <a href="#"><p class="title">Safe Insurance</p></a>
                          <p class="info">Sales Rating: 82%</p>
                      </td>
                      <td>
                          <p>Date: 7/19/2012</p>
                          <a href="#">View Transaction</a>
                      </td>
                      <td>
                          <p class="text-danger h3 pull-right" style="margin-top: 12px;">$22,400</p>
                      </td>
                  </tr>
                    
              </tbody>
            </table>
        </div>
    </div>
    <div class="col-sm-6 col-md-6">
        <div class="panel panel-default">
            <a href="#widget2container" class="panel-heading" data-toggle="collapse">Collapsible </a>
            <div id="widget2container" class="panel-body collapse in">
                <h2>Built with Less</h2>
                <p>The CSS is built with Less. There is a compiled version included if you prefer plain CSS.</p>
                <p>Fava bean jícama seakale beetroot courgette shallot amaranth pea garbanzo carrot radicchio peanut leek pea sprouts arugula brussels sprout green bean. Spring onion broccoli chicory shallot winter purslane pumpkin gumbo cabbage squash beet greens lettuce celery. Gram zucchini swiss chard mustard burdock radish brussels sprout groundnut. Asparagus horseradish beet greens broccoli brussels.</p>
                <p><a class="btn btn-primary">Learn more »</a></p>
            </div>
        </div>
    </div>
</div>
{{end}}

{{define "rootfooter"}}
            <footer>
                <hr>                
                <p class="pull-right">Collect from <a href="#" title="owen" target="_blank">owenshen123@outlook.com</a></p>
                <p>© 2017 <a href="#" target="_blank">owen</a></p>
            </footer>
{{end}}