package models

import (
	"bytes"
	//	"fmt"
	"math"
	//	"os"
	"database/sql"
	"strconv"
	"strings"
	"fmt"
)

func GeneratePage(num int, perpage int, curpage int, url string) (string, int, int) {
	var (
		mpurl       string
		from        int
		to          int
		lastAccount int
		pages       int
		//ret   *bytes.Buffer = new(bytes.Buffer)
		ret bytes.Buffer = bytes.Buffer{}
	)
	//添加URL
	if strings.Index(url, "?") != -1 {
		mpurl = url + "&"
	} else {
		mpurl = url + "?"
	}
	//判断页数
	if num > perpage {
		page := 5
		offset := 5
		//		总页数=数据／每页的数据数
		lastAccount = num % perpage
		pages = int(math.Ceil(float64(num / perpage)))
		if lastAccount != 0 {
			pages++
		}
		if page > pages {
			from = 1
			to = pages
		} else {
			from = curpage - offset
			to = curpage + page + offset - 1
			if from < 1 {
				to = curpage + 1 - from
				from = 1
				if (to-from) < page && (to-from) < pages {
					to = page
				}
			} else if to > pages {
				from = curpage - pages + to
				to = pages
				if (to-from) < page && (to-from) < pages {
					from = pages - page + 1
				}
			}
		}
		//生成html
		if (curpage-offset) > 1 && pages > page {
			ret.WriteString("<li>" + "<a href=\"" + mpurl + "page=1\" >«</a>" + "</li>")
		}

		if curpage > 1 {
			ret.WriteString("<li>" + "<a href=\"" + mpurl + "page=" + strconv.Itoa(curpage-1) + "\" >‹</a>" + "</li>")
		}

		for i := from; i <= to; i++ {
			if i == curpage {
				ret.WriteString("<li ><a>" + strconv.Itoa(i) + "</a></li>")
			} else {
				ret.WriteString("<li>" + "<a href=\"" + mpurl + "page=" + strconv.Itoa(i) + "\" >" + strconv.Itoa(i) + "</a>" + "</li>")
			}
		}

		if curpage < pages {
			ret.WriteString("<li>" + "<a href=\"" + mpurl + "page=" + strconv.Itoa(curpage+1) + "\" >›</a>" + "</li>")
		}

		if to < pages {
			ret.WriteString("<li>" + "<a href=\"" + mpurl + "page=" + strconv.Itoa(pages) + "\" >»</a>" + "</li>")
		}

		if ret.Len() > 0 {
			//			return "<li class=\"p_bar\"><span >Records:" + strconv.Itoa(num) + "</span>" + ret.String() + "</li>"

			return ret.String(), pages, lastAccount
		}

	}
	return "", pages, lastAccount
}

func GenerateUsrList(idNum int, perPage int) string {
	script := "SELECT * FROM usr_info limit " + strconv.Itoa(idNum) + "," + strconv.Itoa(perPage) + ";"
	tmp := ReadData(script)
	var ret bytes.Buffer = bytes.Buffer{}
	for i := 0; i < perPage; i++ {
		ret.WriteString("<tr><td>" + tmp[i][0] + "</td><td id=\"" + tmp[i][1] + "\">" + tmp[i][1] + "</td><td>" + tmp[i][2] + "</td><td>" + "激活</td>" + "<td><a href=\"/admin?IsUsersEdit=1&IsUsrId=" + tmp[i][0] + "\">" + "<i class=\"fa fa-pencil\"></i></a>" + "<a href=\"#myModal\" role=\"button\" data-toggle=\"modal\"><i class=\"fa fa-trash-o\"></i></a></td></tr>")
	}
	return ret.String()
}
func GenerateUsrInfo(IsUsrId string) [][]string {
	script := "SELECT uid,name,tel,email FROM usr_info where id = '" + IsUsrId + "';"
	tmp := ReadData(script)
	return tmp
}
func GenerateClassList()string{
	script := "select max(c.cls_content),max(a.art_title),count(a.art_title) count_title FROM class c left join article a on c.class_id = a.class_id group by c.class_id"
	tmp := ReadData(script)
	//<tr>
	//<td>基础资料{{.ClassListInfo}}</td>
	//<td>论数据结构的重要性</td>
	//<td>13</td>
	//<td>
	//<a href="#"><i  class="fa fa-trash-o"></i></a>
	//</td>
	//</tr>
	var ret bytes.Buffer = bytes.Buffer{}
	for i := 0;i<len(tmp);i++{
		ret.WriteString("<tr><td>"+tmp[i][0]+"</td><td>"+tmp[i][1]+"</td><td>"+tmp[i][2]+"</td><tr>")
	}
	fmt.Println(tmp)
	return ret.String()

}
func GenerateClass()string{
	script := "SELECT cls_content FROM class"
	tmp := ReadDataOne(script)
	var ret bytes.Buffer = bytes.Buffer{}
	////<option>分类1</option>

	for i:=0 ; i< len(tmp);{
		ret.WriteString("<option>"+tmp[i]+"</option>")
		i++
	}
	//fmt.Println(tmp)

	return ret.String()
}
func ReadData(script string) [][]string { //返回多列数据
	var DBdata [][]string
	rows, _ := db.Query(script)
	columns, _ := rows.Columns()
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		var tmpArr []string
		_ = rows.Scan(scanArgs...)
		for _, col := range values {
			if col == nil {
				tmpArr = append(tmpArr, "NULL")
			} else {
				tmpArr = append(tmpArr, string(col))
			}
		}
		DBdata = append(DBdata, tmpArr)
	}

	defer rows.Close()
	return DBdata
}
func ReadDataOne(script string) []string { //返回一列数据
	var tmpArr []string
	rows, _ := db.Query(script)
	columns, _ := rows.Columns()
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {

		_ = rows.Scan(scanArgs...)
		for _, col := range values {
			if col == nil {
				tmpArr = append(tmpArr, "NULL")
			} else {
				tmpArr = append(tmpArr, string(col))
			}
		}
		//DBdata = append(DBdata, tmpArr)
	}

	defer rows.Close()
	return tmpArr
}
