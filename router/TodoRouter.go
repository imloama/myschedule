//任务相应的路由功能实现类
package router

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/itwarcraft/myschedule/model"
	//"github.com/itwarcraft/myschedule/util"
	"net/http"
	"strconv"
	//"time"
)

func TodoListHandler(w http.ResponseWriter, r *http.Request) {

}

//添加新的记录，返回json,为message类型
func TodoAddHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("添加新的工作")
	r.ParseForm()
	var todo model.Todo

	todo.Title = r.FormValue("Title")     //PostFormValue("Title")
	starttime := r.FormValue("Starttime") //r.PostFormValue("Starttime")
	level := r.FormValue("Level")         //r.PostFormValue("Level")
	fmt.Printf("提交到后台的数据，title:%s,starttime:%s,level:%s", todo.Title, starttime, level)
	i, err := strconv.Atoi(level)

	if err != nil {
		i = 0
		fmt.Printf("转化level失败，level:%s\n", level)
	}
	todo.Level = i
	if len(starttime) == 10 {
		starttime = starttime + " 00:00:00"
	}

	todo.Starttime = starttime
	userid_str := r.FormValue("UserId") //r.PostFormValue("UserId")
	uid, _ := strconv.Atoi(userid_str)
	todo.UserId = uid
	err = todo.Add()
	var msg model.Message
	if err != nil {
		fmt.Println("保存失败！")
		fmt.Println(err)
		msg.Code = 0
		msg.Msg = "保存失败！"
	} else {
		msg.Code = 1
		msg.Msg = "保存成功！"
		msg.Data = todo
	}
	b, _ := json.Marshal(msg)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)

}

//根据当前用户编号，查询该用户当前未完成的工作
func TodoUserListHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	useridstr := vars["userid"]
	userid, _ := strconv.Atoi(useridstr)
	todos, err := model.GetTodayTodos(userid)
	if err != nil {
		fmt.Println("获取当前用户的当前任务失败！")
		fmt.Println(err)
	}
	b, _ := json.Marshal(todos)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)

}

//接收请求，将某个任务变成已完成
func TodoFinishHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoidStr := vars["todoid"]
	todoid, _ := strconv.Atoi(todoidStr)

	msg, _ := model.TodoFinish(todoid)
	b, _ := json.Marshal(msg)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)
}

func TodoUserFinishListHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	useridstr := vars["userid"]
	userid, _ := strconv.Atoi(useridstr)
	todos, err := model.GetFinishTodos(userid)
	if err != nil {
		fmt.Println("获取当前用户的当前任务失败！")
		fmt.Println(err)
	}
	b, _ := json.Marshal(todos)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)

}
