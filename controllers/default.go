package controllers

import (
	//	"net/url"
	//	"fmt"
	//	"database/sql"
	//"html/template"
	//"net/http"
	"strconv"
	"webtest/models"

	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/session"
	//	"github.com/astaxie/beego/context"
	"regexp"
	"strings"
)

type MainController struct {
	beego.Controller
}
type LoginController struct {
	beego.Controller
}
type RegisterController struct {
	beego.Controller
}
type TestController struct {
	beego.Controller
}
type HomeController struct {
	beego.Controller
}
type AdminController struct {
	beego.Controller
}
type ListController struct {
	beego.Controller
}
type ContentController struct {
	beego.Controller
}

var IsFlag bool
//var globalSessions *session.Manager
//var value string

//type Logininfo struct {
//	Usrname string
//	Psw     string
//}
//func init() {
//	globalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid", "enableSetCookie,omitempty": true, "gclifetime":3600, "maxLifetime": 3600, "secure": false, "sessionIDHashFunc": "sha1", "sessionIDHashKey": "", "cookieLifeTime": 3600, "providerConfig": ""}`)
//	go globalSessions.GC()
//}
//func login(w http.ResponseWriter, r *http.Request) {
//	sess, _ := globalSessions.SessionStart(w, r)
//	defer sess.SessionRelease(w)
//	username := sess.Get("username")
//	if r.Method == "GET" {
//		t, _ := template.ParseFiles("login.gtpl")
//		t.Execute(w, nil)
//	} else {
//		sess.Set("username", r.Form["username"])
//	}
//}
func (c *LoginController) Get() {

	c.TplName = "login.tpl"

}
func (c *LoginController) Post() {

	//	c.Ctx.WriteString(fmt.Sprint(c.Input()))
	usrname := c.Input().Get("usrname")
	psw := c.Input().Get("psw")
	autoLogin := c.Input().Get("autoLogin") == "on"
	sqlscript := "select uid from usr_info where uid =" + "\"" + usrname + "\"" + ";"
	_, isnull := models.Query(sqlscript)
	if isnull == 1 {
		sqlscript = "select psw from usr_info where uid =" + "\"" + usrname + "\"" + ";"
		sqlpsw, _ := models.Query(sqlscript)
		if psw == sqlpsw {
			IsFlag = true
			maxAge := 0
			if autoLogin {
				maxAge = 1<<31 - 1
			}
			c.Ctx.SetCookie("usrname", usrname, maxAge, "/")
			c.Ctx.SetCookie("psw", psw, maxAge, "/")
		}
	}
	//	if beego.AppConfig.String("usrname") == usrname &&
	//		beego.AppConfig.String("psw") == psw {
	//		maxAge := 0
	//		if autoLogin {
	//			maxAge = 1<<31 - 1
	//		}
	//		c.Ctx.SetCookie("usrname", usrname, maxAge, "/")
	//		c.Ctx.SetCookie("psw", psw, maxAge, "/")
	//	}

	c.Redirect("/home", 301)
	return
}
func (c *RegisterController) Get() {
	c.Data["CheakUsr"] = c.Input().Get("CheakUsr") == "1"
	c.Data["CheakTel"] = c.Input().Get("CheakTel") == "1"
	c.Data["CheakEmail"] = c.Input().Get("CheakEmail") == "1"
	c.TplName = "register.tpl"
}
func (c *RegisterController) Post() {
	url_r := "/register?"
	registerFlag := true
	usrname := c.Input().Get("user")
	email := c.Input().Get("email")
	tel := c.Input().Get("tel")
	psw := c.Input().Get("psw")

	sqlscript := "select uid from usr_info where uid =" + "\"" + usrname + "\"" + ";"
	_, isnull := models.Query(sqlscript)
	if isnull ==1 {

		url_r = url_r+"CheakUsr=1"
		registerFlag = false
	}
	if CheckMobile(tel)==false{
		if strings.Contains(url_r,"="){
			url_r = url_r + "&CheakTel=1"
		}else {
			url_r = url_r + "CheakTel=1"
		}

		//1表示格式错了
		registerFlag = false
	}
	if (strings.ContainsAny(email, ".") &&strings.ContainsAny(email, "@"))==false{
		if strings.Contains(url_r,"="){
			url_r = url_r + "&CheakEmail=1"
		}else {
			url_r = url_r + "CheakEmail=1"
		}
		//1表示格式错了
		registerFlag = false
	}
	if registerFlag == false{
		c.Redirect(url_r,301)
	}else {
		models.Insert(usrname, email, tel, psw)
	}
	//if isnull == 1 {
	//	c.Redirect("/register?CheakUsr=1", 301)
	//} else {
	//
	//}

	c.Redirect("/login", 301)

	return

}
func (c *TestController) Get() {


	c.TplName = "index.tpl"
}
func (c *HomeController) Get() {
	IsExit := c.Input().Get("Exit") == "1"
	if IsExit {
		c.Ctx.SetCookie("usrname", "", -1, "/")
		c.Ctx.SetCookie("psw", "", -1, "/")
		IsFlag = false
		c.Redirect("/home", 301)
		return
	}
	c.Data["Project_name"] = "主页"
	c.Data["DataSql"] = models.DbQuary()
	c.Data["Date1"] = "qbcdaa"
	c.Data["IsHome"] = true

	IsFlag = checkAccount(c)
	c.Data["IsLogin"] = IsFlag



	c.TplName = "home.tpl"
}
func (c *AdminController) Get() {
	//	c.Data["IsLogin"] = checkAccount()
	//	c.Data["IsLogin"] = IsFlag
	var tmp, CurPage string
	var account, pages, lastAccount int

	c.Data["IsPsw"] = c.Input().Get("IsPsw")=="1"
	c.Data["IsPsw1"] = c.Input().Get("IsPsw1")=="1"

	//idNum:默认用户id,perPeople:默认每页数据
	idNum := 0
	perPeople := 4
	tmp1 := 1
	//取出cookie的用户名
	tmp = readUsrName(c)
	//根据url选择相对应的组建
	IsUsers := c.Input().Get("IsUsers") == "1"
	IsDash := c.Input().Get("IsDash") == "1"
	IsUsersEdit := c.Input().Get("IsUsersEdit") == "1"
	//根据id显示用户详情
	IsUsrId := c.Input().Get("IsUsrId")
	if IsUsrId == "" {
		sqlscript := "select id from usr_info where uid =" + "\"" + tmp + "\"" + ";"
		IsUsrId, _ = models.Query(sqlscript)
	}
	script := "SELECT uid,name,tel,email FROM usr_info where id = '" + IsUsrId + "';"
	tmp2 := models.ReadData(script)
	//	if IsUsrId != "" {
	c.Data["IsId"] = tmp2[0][0]
	c.Data["IsName"] = tmp2[0][1]
	c.Data["IsTel"] = tmp2[0][2]
	c.Data["IsEmail"] = tmp2[0][3]
	//	} else {
	//		c.Data["IsId"] = "用户名"
	//		c.Data["IsName"] = "昵称"
	//		c.Data["IsTel"] = "电话"
	//		c.Data["IsEmail"] = "Email"

	//	}

	sqlscript := "select name from usr_info where uid =" + "\"" + tmp + "\"" + ";"
	//取用户名
	c.Data["UsrName"], _ = models.Query(sqlscript)
	c.Data["IsUsers"] = IsUsers
	c.Data["IsDash"] = IsDash
	c.Data["IsUsersEdit"] = IsUsersEdit
	//取用户数量
	tmp, _ = models.Query("select count(*) from usr_info where uid IS NOT null;")
	c.Data["UsrAccount"] = tmp

	//取用户数据，第i+1位用户开始，每页显示j个用户

	if c.Input().Get("page") != "" {
		CurPage = c.Input().Get("page")
	} else {
		CurPage = "1"
	}
	//当前页码
	account, _ = strconv.Atoi(tmp)
	tmp1, _ = strconv.Atoi(CurPage)
	idNum = perPeople * (tmp1 - 1)
	//总数据数，每页数据数，当前页码，原始url

	c.Data["Page"], pages, lastAccount = models.GeneratePage(account, perPeople, tmp1, "/admin?IsUsers=1")
	if tmp1 == pages {
		perPeople = lastAccount
	}
	c.Data["UsrInfo"] = models.GenerateUsrList(idNum, perPeople)

	c.TplName = "admin1.tpl"
}
func (c *AdminController) Post() {
	//IsUserInfo := c.Input().Get("IsUserInfo")
	//IsUserPsw := c.Input().Get("IsUserPsw")
	Uid := c.Input().Get("uid")
	UpdataName := c.Input().Get("UpdataName")
	UpdataTel := c.Input().Get("UpdataTel")
	UpdataEmail := c.Input().Get("UpdataEmail")

	tmpUrl := "/admin?IsUsersEdit=1"
	updataFlag := true
	updataFlag1 := true

	readPsw := c.Input().Get("Psw")
	readPsw1 := c.Input().Get("Psw1")
	readPsw2 := c.Input().Get("Psw2")

	if UpdataEmail ==""||UpdataName==""||UpdataTel==""{
		updataFlag = false
	}

	username:= readUsrName(c)
	sqlscript := "select psw from usr_info where uid =" + "\"" + username + "\"" + ";"
	sourcePsw, _ := models.Query(sqlscript)

	if readPsw != sourcePsw{
		//显示原始密码输入错误
		tmpUrl = tmpUrl+"&IsPsw=1"
		updataFlag1 = false
	}
	if readPsw1 != readPsw2{
		//显示两次如数的重置密码不一致
		tmpUrl = tmpUrl+"&IsPsw1=1"
		updataFlag1 = false
	}
	//UpdataBlog := c.Input().Get("UpdataBlog")
	//	UpdataPsw := c.Input().Get("UpdataPsw")
	//	var tmp string
	//	tmp = "update usr_info set name ='" + UpdataName + "',tel='" + UpdataTel + "',email='" + UpdataEmail + "' where uid='" + Uid + "' "
	if updataFlag{
		models.Update(UpdataName, UpdataTel, UpdataEmail, Uid)
		c.Redirect("/admin?IsUsersEdit=1", 301)
	}
	if updataFlag1{
		models.UpdatePsw(readPsw2,Uid)
		c.Redirect("/admin?IsUsersEdit=1", 301)
	}else {
		c.Redirect(tmpUrl,301)
	}


}
func (c *ListController) Get() {
	c.Data["IsLogin"] = IsFlag

	c.Data["Sqlssss"] = models.DbQuary()
	c.Data["Sqlsss"] = "测试的数据"
	c.Data["Project_name"] = "列表"
	c.Data["IsList"] = true
	c.TplName = "list.tpl"
}
func (c *ContentController) Get() {
	c.Data["IsLogin"] = IsFlag
	c.Data["Project_name"] = "内容"
	c.Data["IsContent"] = true
	c.TplName = "content.tpl"
}
func readUsrName(c *AdminController) string {
	ck, _ := c.Ctx.Request.Cookie("usrname")

	return ck.Value
}

func checkAccount(ctx *HomeController) bool {
	//func checkAccount() bool {
	//ck, err := ctx.Request.Cookie("usrname")
	//	ck, err := beego.Controller.Ctx.Request.Cookie("usrname")
	var login bool
	ck, err := ctx.Ctx.Request.Cookie("usrname")
	if err != nil {
		return false
	}
	usrname := ck.Value
	ck, err = ctx.Ctx.Request.Cookie("psw")
	if err != nil {
		return false
	}
	psw := ck.Value
	sqlscript := "select uid from usr_info where uid =" + "\"" + usrname + "\"" + ";"
	_, isnull := models.Query(sqlscript)
	if isnull == 1 {

		sqlscript = "select psw from usr_info where uid =" + "\"" + usrname + "\"" + ";"
		sqlpsw, _ := models.Query(sqlscript)
		if sqlpsw == psw {
			login = true
		}
		//	return beego.AppConfig.String("usrname") == usrname &&
		//	beego.AppConfig.String("psw") == psw
	}
	return login
}
func CheckMobile(no string) bool {
	reg := regexp.MustCompile(`^(13[0-9]|14[57]|15[012356789]|18[0-9])\d{8}$`)
	return reg.MatchString(no)
}