package model

import (
	"fmt"
	"github.com/itwarcraft/myschedule/util"
	"log"
	"time"
)

//create table todos(Id integer primary key,Title text,Content text,
//Isfinish boolean,Level integer,Starttime text,Finishtime text,UserId integer)
type Todo struct {
	Id         int
	Title      string
	Content    string
	Isfinish   int //0为false 1为true
	Level      int
	Starttime  string
	Finishtime string
	UserId     int
}

func (todo *Todo) Add() error {
	sql := fmt.Sprintf("insert into todos(title,isfinish,level,starttime,userid)values('%s',0,%d,'%s',%d);", todo.Title, todo.Level, todo.Starttime, todo.UserId)
	db := util.NewDb("./myschedule.db")
	fmt.Println(sql)
	err := db.Exec(sql)
	if err != nil {
		fmt.Println("保存失败！")
		return err
	} else {
		sql = fmt.Sprintf("select Id,Title,Level,Starttime from todos where userid=%d  order by id limit 1", todo.UserId)
		fmt.Println("[SQL]" + sql)
		db = util.NewDb("./myschedule.db")
		rows, _ := db.Query(sql)
		for rows.Next() {
			rows.Scan(&todo.Id, &todo.Title, &todo.Level, &todo.Starttime)
		}
		return nil
	}

}

//根据用户ID，查找所有该用户的今日工作
func GetTodayTodos(userid int) ([]Todo, error) {
	sql := fmt.Sprintf("select Id,Title,Level,Starttime from todos where UserId=%d and isfinish=0 order by id ", userid)
	db := util.NewDb("./myschedule.db")
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal("查询数据库操作!!!")
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	todos := []Todo{}
	//t := Todo{Id: 1}
	//append(todos, t)
	for rows.Next() {
		var todo Todo
		rows.Scan(&todo.Id, &todo.Title, &todo.Level, &todo.Starttime)
		fmt.Println(todo.Starttime)
		todo.UserId = userid
		fmt.Printf("todo[id:%d,title:%s,level:%d,Starttime:%s]\n", todo.Id, todo.Title, todo.Level, todo.Starttime)
		todos = append(todos, todo)
	}
	return todos, nil
}

func TodoFinish(todoid int) (Message, error) {
	finishtime := util.Format(time.Now())
	sql := fmt.Sprintf("update todos set isfinish=1,finishtime='%s' where id=%d", finishtime, todoid)
	fmt.Println("执行的SQL语句:" + sql)
	db := util.NewDb("./myschedule.db")
	err := db.Exec(sql)
	var msg Message
	if err != nil {
		fmt.Println("执行sql语句失败！")
		fmt.Println(err)
		msg.Code = 0
		msg.Msg = "执行失败！"
	} else {
		msg.Code = 1
		msg.Msg = "执行成功！"
	}
	return msg, err

}

func GetFinishTodos(userid int) ([]Todo, error) {

	tnow := time.Now()
	seven := tnow.AddDate(0, 0, -7)
	tnowStr := util.Format(tnow)
	sevenStr := util.Format(seven)
	fmt.Printf("start:%s,end:%s\n", tnowStr, sevenStr)

	sql := fmt.Sprintf("select Id,Title,Level,Starttime,Finishtime from todos where UserId=%d and isfinish=1 and finishtime >= '%s'  order by id desc ", userid, sevenStr)
	db := util.NewDb("./myschedule.db")
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal("查询数据库操作!!!")
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	todos := []Todo{}
	//t := Todo{Id: 1}
	//append(todos, t)
	for rows.Next() {
		var todo Todo
		rows.Scan(&todo.Id, &todo.Title, &todo.Level, &todo.Starttime, &todo.Finishtime)
		fmt.Println(todo.Starttime)
		todo.UserId = userid
		fmt.Printf("todo[id:%d,title:%s,level:%d,Starttime:%s,Finishtime:%s]\n", todo.Id, todo.Title, todo.Level, todo.Starttime, todo.Finishtime)
		todos = append(todos, todo)
	}
	return todos, nil
}
