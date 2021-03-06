package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	dblogin := beego.AppConfig.String("dbuser")+":"+beego.AppConfig.String("dbpsw")+
	"@tcp("+beego.AppConfig.String("dbhost")+":"+beego.AppConfig.String("dbport")+
	")/"+beego.AppConfig.String("dbschema")+"?charset=utf8"
	fmt.Println("加载数据库信息："+dblogin)
	//db, _ = sql.Open("mysql", "root:owenshen123@tcp(127.0.0.1:3306)/test?charset=utf8")
	db,_ = sql.Open("mysql",dblogin)
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.Ping()
}

//插入demo
func Insert(usr string, email string, tel string, psw string) {
	//	stmt, err := db.Prepare(`INSERT user_info (user_name,user_age,user_sex) values (?,?,?)`)
	stmt, err := db.Prepare(`INSERT test.usr_info (uid,email,tel,psw) values (?,?,?,?)`)
	checkErr(err)
	//	res, err := stmt.Exec("tom", "tomtian@we.com", 15927120431, "123dedcgs")
	res, err := stmt.Exec(usr, email, tel, psw)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}

func InsertOne(tableName,column,value string) {
	//	stmt, err := db.Prepare(`INSERT user_info (user_name,user_age,user_sex) values (?,?,?)`)
	script := "INSERT test." + tableName+"("+column+")values (?)"
	//stmt, err := db.Prepare(`INSERT test.usr_info (uid,email,tel,psw) values (?,?,?,?)`)
	stmt, err := db.Prepare(script)

	checkErr(err)
	//	res, err := stmt.Exec("tom", "tomtian@we.com", 15927120431, "123dedcgs")
	res, err := stmt.Exec(value)
	checkErr(err)
	_, err = res.LastInsertId()
	checkErr(err)
	fmt.Println("已经插入"+value)
}
func InsertTwo(tableName,column1,value1,column2,value2 string) {
	//	stmt, err := db.Prepare(`INSERT user_info (user_name,user_age,user_sex) values (?,?,?)`)
	script := "INSERT test." + tableName+"("+column1+", "+column2+")values (?,?)"
	//stmt, err := db.Prepare(`INSERT test.usr_info (uid,email,tel,psw) values (?,?,?,?)`)
	stmt, err := db.Prepare(script)

	checkErr(err)
	//	res, err := stmt.Exec("tom", "tomtian@we.com", 15927120431, "123dedcgs")
	res, err := stmt.Exec(value1,value2)
	checkErr(err)
	_, err = res.LastInsertId()
	checkErr(err)
	fmt.Println("已经插入"+value1+","+value2)
}
func InsertThree(tableName,column1,value1,column2,value2,column3,value3 string) {
	//	stmt, err := db.Prepare(`INSERT user_info (user_name,user_age,user_sex) values (?,?,?)`)
	script := "INSERT test." + tableName+"("+column1+", "+column2+", "+column3+")values (?,?,?)"
	fmt.Println(script)
	//stmt, err := db.Prepare(`INSERT test.usr_info (uid,email,tel,psw) values (?,?,?,?)`)
	stmt, err := db.Prepare(script)
	checkErr(err)
	//	res, err := stmt.Exec("tom", "tomtian@we.com", 15927120431, "123dedcgs")
	res, err := stmt.Exec(value1,value2,value3)
	checkErr(err)
	_, err = res.LastInsertId()
	checkErr(err)
	fmt.Println("已经插入"+value1+","+value2+","+value3)
}

func InsertFour(tableName,column1,value1,column2,value2,column3,value3 ,column4,value4 string) {
	//	stmt, err := db.Prepare(`INSERT user_info (user_name,user_age,user_sex) values (?,?,?)`)
	script := "INSERT test." + tableName+"("+column1+", "+column2+", "+column3+", "+column4+")values (?,?,?,?)"
	fmt.Println(script)
	//stmt, err := db.Prepare(`INSERT test.usr_info (uid,email,tel,psw) values (?,?,?,?)`)
	stmt, err := db.Prepare(script)
	checkErr(err)
	//	res, err := stmt.Exec("tom", "tomtian@we.com", 15927120431, "123dedcgs")
	res, err := stmt.Exec(value1,value2,value3,value4)
	checkErr(err)
	_, err = res.LastInsertId()
	checkErr(err)
	fmt.Println("已经插入"+value1+","+value2+","+value3)
}
//查询demo
func Query(script string) (value string, isnull int) {
	rows, err := db.Query(script)
	checkErr(err)

	//字典类型
	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
				value = string(col.([]byte))

			}
		}
		if len(record) >= 1 {
			isnull = 1
		} else {
			isnull = 0
		}

	}
	return
}

//更新数据
func UpdateOne(tableName,column,values,pryKey,id string){
	script := "UPDATE "+tableName+" SET "+column+ " =? "+"WHERE "+ pryKey+" = ?"
	//stmt, err := db.Prepare(`UPDATE usr_info SET name=?,tel=?,email=? WHERE uid=?`)
	stmt, err := db.Prepare(script)
	checkErr(err)
	//values要改的值，ID 确认这一行
	res, err := stmt.Exec(values, id )
	checkErr(err)
	_, err = res.RowsAffected()
	checkErr(err)
}
func Update(name, tel, email, uid string) {

	stmt, err := db.Prepare(`UPDATE usr_info SET name=?,tel=?,email=? WHERE uid=?`)
	checkErr(err)
	res, err := stmt.Exec(name, tel, email, uid)
	checkErr(err)
	_, err = res.RowsAffected()
	checkErr(err)
	//	fmt.Println(num)
}
func UpdatePsw(psw,uid string)  {
	stmt, err := db.Prepare(`UPDATE usr_info SET psw=? WHERE uid=?`)
	checkErr(err)
	res, err := stmt.Exec(psw, uid)
	checkErr(err)
	_, err = res.RowsAffected()
	checkErr(err)

}

//删除数据
func Remove(tableName,key,values string) {
	//stmt, err := db.Prepare(`DELETE FROM user WHERE user_id=?`)
	script := "DELETE FROM "+tableName+" WHERE "+key+" =?"
	stmt, err := db.Prepare(script)
	checkErr(err)
	res, err := stmt.Exec(values)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(num)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func DbQuary() (value string) {
	//	db, err := sql.Open("mysql", "root:owenshen123@tcp(localhost:3306)/test?charset=utf8")
	//	if err != nil {
	//		panic(err.Error())
	//	}
	//	defer db.Close()

	rows, err := db.Query("SELECT current_Date")
	if err != nil {
		panic(err.Error())
	}
	// Get column names获取列名
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	// Make a slice for the values存储列信息
	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		// Now do something with the data.
		// Here we just print each column as a string.

		for _, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}

			//			fmt.Println(columns[i], ": ", value)
			//			fmt.Println(value)

		}
	}
	return
}
