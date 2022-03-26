package main

import (
	"fmt"
	"github.com/pkg/errors"
	//"gorm.io/driver/mysql"
	//"gorm.io/gorm"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


/**************************
需要由wrap包装抛给上层，执行sql语句报错后，如果想知道哪个操作导致了sql.ErrNoRows，可以
用wrap包装抛出。 处理错误的地方也可以打印堆栈调用信息，更快地定位问题。

**************************/


type UserInfo struct {
	Id     uint
	Name   string
	Gender string
	Age    int
}

func UserQuery(db *sql.DB, sql string) (err error) {
	var rows int
	fmt.Println("sql: ", sql)

	rows, err = db.Query(sql)
	if err != nil {
		return errors.Wrap(err, "Mysql Query fail.")
	}

	for rows.Next() {
		var user UserInfo
		err = rows.Scan(&user.Id, &user.Name, &user.Age, &user.Gender)
		if err != nil {
			switch {
			case err == sql.ErrNoRows:
				err = errors.Wrap(err, "ErrNoRows ")
			default:
				err = errors.Wrap(err, "Scan error")
			}
			return
		}
		fmt.Println(user)
	}

	fmt.Println("Query Over")
	return
}

func main() {
	db, err := sql.Open("mysql", "root:include@tcp(127.0.0.1:3306)/yanl?charset=utf8")
	if err != nil {
		err = errors.Wrap(err, "Mysql Open Fail")
		return
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		err = errors.Wrap(err, "Mysql Connect Fail")
		return
	}

	err = UserQuery(db, "select * from UserInfo where id = 1")
	if err != nil {
		fmt.Println("UserQuery error: ", err)
		fmt.Printf("%+v", err)
	}
}
