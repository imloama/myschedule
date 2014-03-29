package main

import (
	//"crypto/md5"
	//"encoding/hex"
	"fmt"
	"github.com/itwarcraft/myschedule/model"
	"github.com/itwarcraft/myschedule/util"
	"time"
)

func main() {
	//h := md5.New()
	//h.Write([]byte("admin"))
	//fmt.Printf("admin,(md5):%s\n", hex.EncodeToString(h.Sum(nil)))

	//time test
	//mtime()

	//todo
	getTodo()
}

func mtime() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	t1, _ := time.Parse("2006-01-02 15:04:05", "2014-03-28 00:00:00")
	fmt.Println(t1)
	fmt.Println(t1.Format("2006-01-02 15:04:05"))
}

func getTodo() {
	sql := "select Id,Title,Level,Starttime from todos"
	db := util.NewDb("./myschedule.db")
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var todo model.Todo
		rows.Scan(&todo.Id, &todo.Title, &todo.Level, &todo.Starttime)
		fmt.Println(todo.Starttime)
		fmt.Printf("todo[id:%d,title:%s,level:%d,Starttime:%s]\n", todo.Id, todo.Title, todo.Level, todo.Starttime)
	}

}
