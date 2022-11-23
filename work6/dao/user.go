package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"redrock/work6/model"
)

var db *sql.DB
var u model.User
var com model.Comments

// 连接数据库
func InitDB() (err error) {
	dsn := "user:password@tcp(127.0.0.1:3306)/mysql_demo?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	//defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("数据库连接失败")
		return err
	}
	return nil
}

// 将数据保存进数据库
func AddUser(username, password, question, answer string) {
	sqlStr := "insert into user(username,password,question,answer) values (?,?,?,?)"
	_, err := db.Exec(sqlStr, username, password, question, answer)
	if err != nil {
		fmt.Println("insert failed,err:", err)
		return
	}
}

// 查询用户
func SelectUser(username string) bool {
	sqlStr := "select username from user where username=?"
	err := db.QueryRow(sqlStr, username).Scan(&u.Username)
	if err != nil {
		return false
	}
	return true
}
func SelectPassword(username string) string {
	sqlStr := "select password from user where username=?"
	db.QueryRow(sqlStr, username).Scan(&u.Password)
	return u.Password
}
func SelectQuestion(username string) string {
	sqlStr := "select password from user where username=?"
	db.QueryRow(sqlStr, username).Scan(&u.Question)
	return u.Question
}
func SelectAnswer(username string) string {
	sqlStr := "select answer from user where username=?"
	db.QueryRow(sqlStr, username).Scan(&u.Answer)
	return u.Answer
}

// 更新用户信息
func UpdatePassword(username, password string) {
	sqlStr := "update user set password=? where username=?"
	db.Exec(sqlStr, password, username)
}
func UpdateQuestion(username, question string) {
	sqlStr := "update user set password=? where username=?"
	db.Exec(sqlStr, question, username)
}
func UpdateAnswer(username, answer string) {
	sqlStr := "update user set answer=? where username=?"
	db.Exec(sqlStr, answer, username)
}

// 保存留言
func AddComments(yourname, myname, content string) {
	sqlStr := "insert into comments(yourname,content,myname) values(?,?,?)"
	db.Exec(sqlStr, yourname, content, myname)
}

// 查询留言
func Select(name string) (yourname, content, myname string) {
	sqlStr := "select yourname,content,myname from comments where yourname=?"
	db.QueryRow(sqlStr, name).Scan(&com.Yourname, &com.Content, &com.Myname)
	return com.Yourname, com.Content, com.Myname
}
