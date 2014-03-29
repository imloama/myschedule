package router

import (
	//"encoding/json"
	"fmt"
	"github.com/itwarcraft/myschedule/model"
	"html/template"
	"log"
	"net/http"
	"strconv"
	//"os"
)

//接收用户请求，判断cookie中是否有数据，如果有数据，则获取用户名
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	if err != nil {
		//log.Fatal("没有找到cookie")
		fmt.Println("没有找到cookie")
		toLogin(w)
	} else {
		//找到cookie
		username := cookie.Value
		fmt.Printf("cookie中的用户名是:%s\n", username)
		//查询数据，打开Index界面
		/*
			user, err2 := model.GetUserByName(username)
			checkError(err2)
			todos, err3 := model.GetTodayTodos(user.Id)
			checkError(err3)

			message := model.Message{Code: 0, Msg: "获取数据成功", Data: todos}
			toIndex(w, message)
		*/
		message, _ := getTodosByUserName(username)
		toIndex(w, message)
	}

}

func getTodosByUserName(name string) (*model.Message, error) {
	user, err2 := model.GetUserByName(name)
	checkError(err2)
	todos, err3 := model.GetTodayTodos(user.Id)
	checkError(err3)

	m := make(map[string]interface{})

	m["todos"] = todos
	m["user"] = user

	message := model.Message{Code: 0, Msg: "获取数据成功", Data: m}
	return &message, nil
}

func toLogin(w http.ResponseWriter) {
	t, err := template.ParseFiles("./template/login.html", "./template/header.html", "./template/jslib.html", "./template/footer.html")
	checkError(err)
	err = t.Execute(w, nil)
	checkError(err)
}

func toIndex(w http.ResponseWriter, data interface{}) {
	t, err := template.ParseFiles("./template/index.html", "./template/header.html", "./template/jslib.html", "./template/footer.html")
	checkError(err)
	err = t.Execute(w, data)
	checkError(err)
}

/**
 * 用户登陆
 */
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var message model.Message
	r.ParseForm()
	user.Name = r.Form.Get("username")     //r.FormValue("username")
	user.Password = r.Form.Get("password") //r.FormValue("password")
	err := user.Login()

	fmt.Printf("用户登陆，ID:%d,用户名:%s,密码：%s\n", user.Id, user.Name, user.Password)

	if err != nil {
		log.Fatal("用户登陆失败")
		message.Code = 0
		message.Msg = "登陆失败"
		toLogin(w)
	} else {
		msg, _ := getTodosByUserName(user.Name)
		cookie := http.Cookie{Name: "username", Value: user.Name, MaxAge: 1000 * 60 * 60 * 24 * 7}
		http.SetCookie(w, &cookie)
		//idstr := fmt.Sprintf("%d", user.Id)
		idstr := strconv.Itoa(user.Id)
		cookie2 := http.Cookie{Name: "userid", Value: idstr, MaxAge: 1000 * 60 * 60 * 24 * 7}
		http.SetCookie(w, &cookie2)
		toIndex(w, msg)
	}
	//b, _ := json.Marshal(message)

	//w.Header().Set("Content-Type", "application/json;charset=utf-8")
	//w.Write(b)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		//os.Exit(1)
	}
}
