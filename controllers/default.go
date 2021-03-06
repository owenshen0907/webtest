package controllers

import (
	//	"net/url"
	//	"fmt"
	//	"database/sql"
	//"html/template"
	//"net/http"
	"strconv"
	"webtest/models"
	"time"
	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/session"
	//	"github.com/astaxie/beego/context"
	"regexp"
	"strings"
	"fmt"
	"os"
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
	//根据url选择相对应的组建--获取模块开关
	IsUsers := c.Input().Get("IsUsers") == "1"
	IsDash := c.Input().Get("IsDash") == "1"
	IsUsersEdit := c.Input().Get("IsUsersEdit") == "1"
	IsEditor := c.Input().Get("IsEditor") == "1"
	IsClass := c.Input().Get("IsClass") == "1"
	IsResource := c.Input().Get("IsResource") == "1"

	var USERNAME,USERACCOUNT, CurPage string
	var account, pages, lastAccount int
	//****************8取出cookie的用户名*************
	USERNAME = readUsrName(c)
	sqlscript := "select name from usr_info where uid =" + "\"" + USERNAME + "\"" + ";"
	//取用户名
	c.Data["UsrName"], _ = models.Query(sqlscript)
	//****************8取出用户姓名完毕****************

	c.Data["IsPsw"] = c.Input().Get("IsPsw")=="1"
	c.Data["IsPsw1"] = c.Input().Get("IsPsw1")=="1"

	//***********************************根据id显示用户详情*************************
	IsUsrId := c.Input().Get("IsUsrId")
	if IsUsrId == "" {
		sqlscript := "select id from usr_info where uid =" + "\"" + USERNAME + "\"" + ";"
		IsUsrId, _ = models.Query(sqlscript)
	}
	script := "SELECT uid,name,tel,email FROM usr_info where id = '" + IsUsrId + "';"
	tmp := models.ReadData(script)
	//	if IsUsrId != "" {
	c.Data["IsId"] = tmp[0][0]
	c.Data["IsName"] = tmp[0][1]
	c.Data["IsTel"] = tmp[0][2]
	c.Data["IsEmail"] = tmp[0][3]
	//***********************************显示用户详情完毕***************************
	//******************************************生成用户列表信息*************************************
	//idNum:默认用户id,perPeople:默认每页数据,CurPages当前页码
	idNum := 0
	perPeople := 4
	CurPages := 1
	//取用户数量
	USERACCOUNT, _ = models.Query("select count(*) from usr_info where uid IS NOT null;")
	c.Data["UsrAccount"] = USERACCOUNT
	//取用户数据，第i+1位用户开始，每页显示j个用户
	if c.Input().Get("page") != "" {
		CurPage = c.Input().Get("page")
	} else {
		CurPage = "1"
	}
	//当前页码
	account, _ = strconv.Atoi(USERACCOUNT)
	CurPages, _ = strconv.Atoi(CurPage)
	idNum = perPeople * (CurPages - 1)
	//总数据数，每页数据数，当前页码，原始url
	c.Data["Page"], pages, lastAccount = models.GeneratePage(account, perPeople, CurPages, "/admin?IsUsers=1")
	if CurPages == pages {
		perPeople = lastAccount
	}
	c.Data["UsrInfo"] = models.GenerateUsrList(idNum, perPeople)
	//**********************************************用户列表信息生成完毕**************************************
	//**************************显示分类&标签表*************************************
	c.Data["ClassList"] = models.GenerateClass()
	c.Data["TagList"] = models.GenerateTag()
	c.Data["ClassListInfo"] = models.GenerateClassList()
	//**************************显示分类表结束**********************************

	//模块开关
	c.Data["IsUsers"] = IsUsers
	c.Data["IsDash"] = IsDash
	c.Data["IsUsersEdit"] = IsUsersEdit
	c.Data["IsEditor"] = IsEditor
	c.Data["IsClass"] = IsClass
	c.Data["IsResource"] = IsResource

	c.TplName = "admin1.tpl"
}
func (c *AdminController) Post() {
	//IsUserInfo := c.Input().Get("IsUserInfo")
	//IsUserPsw := c.Input().Get("IsUserPsw")
	//*************修改用户信息********************
	Uid := c.Input().Get("uid")
	UpdataName := c.Input().Get("UpdataName")
	UpdataTel := c.Input().Get("UpdataTel")
	UpdataEmail := c.Input().Get("UpdataEmail")

	tmpUrl := "/admin?IsUsersEdit=1"
	updataFlag := true

	if UpdataEmail ==""||UpdataName==""||UpdataTel==""{
		updataFlag = false
	}
	if updataFlag{
		models.Update(UpdataName, UpdataTel, UpdataEmail, Uid)
		c.Redirect("/admin?IsUsersEdit=1", 301)
	}
	//*************修改用户信息结束********************
	//*********************8****修改用户密码************************************
	readPsw := c.Input().Get("Psw")
	readPsw1 := c.Input().Get("Psw1")
	readPsw2 := c.Input().Get("Psw2")
	updataFlag1 := true
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

	if updataFlag1{
		models.UpdatePsw(readPsw2,Uid)
		c.Redirect("/admin?IsUsersEdit=1", 301)
	}else {
		c.Redirect(tmpUrl,301)
	}
	//*****************************修改用户密码结束***********************************8

	//************添加类****************
	addClassName :=c.Input().Get("addClassName")
	fmt.Println("addClassName："+addClassName)
	if addClassName !=""{
		tmpFlag := cheakClass(addClassName)
		if tmpFlag{
			models.InsertOne("class","cls_content",addClassName)
			c.Redirect("/admin?IsClass=1",301)
		}else {
			fmt.Println("已有类无法重复添加")
			c.Redirect("/admin?IsClass=1",301)
			//c.Ctx.WriteString("已有类无法重复添加。")
		}
	}

	//************添加类完毕*************
	//************删除类****************
	delClassName :=c.Input().Get("delClassName")
	fmt.Println("delClassName："+delClassName)
	if delClassName != ""{
		tmpFlag1 := cheakDelClass(delClassName)
		if tmpFlag1{
			models.Remove("class","cls_content",delClassName)
			c.Redirect("/admin?IsClass=1",301)
			fmt.Println("可以删除改类")
		}else {
			fmt.Println("已有数据类无法删除")
			c.Redirect("/admin?IsClass=1",301)
			//c.Ctx.WriteString("已有类无法重复添加。")
		}
	}
	//************删除类完毕*************
	//************修改类****************
	updateClassName :=c.Input().Get("updateClassName")
	sourceClassName :=c.Input().Get("sourceClassName")
	fmt.Println("updateClassName："+updateClassName)
	fmt.Println("sourceClassName："+sourceClassName)
	if sourceClassName != ""{
		tmpFlag2 := cheakClass(sourceClassName)
		if !tmpFlag2{
			models.UpdateOne("class","cls_content",updateClassName,"cls_content",sourceClassName)
			//c.Redirect("/admin?IsClass=1",301)
			fmt.Println("已经修改类")
		}else {
			fmt.Println("未知错误")
			c.Redirect("/admin?IsClass=1",301)
			//c.Ctx.WriteString("已有类无法重复添加。")
		}
	}
	//************修改类完毕*************
	//************上传文件****************
	//选择文件类型
	getFiles := c.Input().Get("Options")
	if getFiles != ""{
		f, h, err1 := c.GetFile("myfile")
		if err1 == nil {
			var path string
			datePath := time.Now().Format("20060102")
			dir, _ := os.Getwd()
			dir = dir+"/upload/"+getFiles+"/"
			path = dir+datePath
			URL := "/upload/"+getFiles+"/"+datePath
			//当前的目录
			if !PathExists(path){
				err := os.Mkdir(path, os.ModePerm)  //在当前目录下生成md目录
				if err != nil {
					fmt.Println(err)
				}
			}
			// 设置保存目录
			//dirPath := "./upload/" + datePath
			// 设置保存文件名
			FileName := h.Filename
			c.SaveToFile("myfile",path+"/"+FileName)
			fmt.Println(URL+"/"+FileName)
			URL = URL +"/"+FileName
			classname := c.Input().Get("selectClassName")
			class_id_script := "select class_id from class where cls_content = '"+classname+"'"
			class_id := models.ReadDataOne(class_id_script)
			tagname := c.Input().Get("selectTagName")
			tag_id_script := "select tag_id from tag where tag_content = '"+tagname+"'"
			tag_id := models.ReadDataOne(tag_id_script)
			title := c.Input().Get("resourceTitle")
			fmt.Println(title+class_id[0],tag_id[0])
			models.InsertFour("resource","title",title,"class_id",class_id[0],"tag_id",tag_id[0],"url",URL)
			//c.Ctx.WriteString(URL+"/"+FileName)
		}
		//content := c.Input().Get("content")


		//c.Redirect("/", 302)
		defer f.Close()
	}
	//************上传文件结束*************

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

func cheakClass(insertClass string)bool{
	//存在的话返回false
	script := "SELECT * FROM class"
	tmp := models.ReadDataOne(script)
	tmpFlag := true
	for _,v := range tmp{
		if insertClass == v{
			tmpFlag = false
		}
	}
	return tmpFlag
}

func cheakDelClass(delClass string)bool{
	script := "SELECT max(c.cls_content),  count(a.art_title) count_title FROM class c LEFT JOIN article a ON c.class_id = a.class_id GROUP BY c.class_id"
	tmp := models.ReadData(script)
	fmt.Println(tmp)
	tmpFlag := false
	for i,_ := range tmp{
		fmt.Println(tmp[i][0])
		if delClass == tmp[i][0]{
			if tmp[i][1] == "0"{
				fmt.Println(tmp[i][0],tmp[i][1])
				tmpFlag = true
			}
		}
	}
	return tmpFlag
}
func checkAccount(ctx *HomeController) bool {
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
	}
	return login
}
func CheckMobile(no string) bool {
	reg := regexp.MustCompile(`^(13[0-9]|14[57]|15[012356789]|18[0-9])\d{8}$`)
	return reg.MatchString(no)
}
func PathExists(path string) (bool) {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}