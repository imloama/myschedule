/**
 * 启动服务器
 */
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/itwarcraft/myschedule/router"
	"log"
	"net/http"
	"os"
)

type App struct {
	//端口
	Port string
}

var app App

//在init函数中初始化
func init() {
	f, err := os.Open("config.json")
	if err != nil {
		app.Port = "7890"
		fmt.Println("读取配置文件config.json错误！采用默认7890端口")
	} else {
		defer f.Close()
		var bs bytes.Buffer
		buf := make([]byte, 1024)
		for {
			n, _ := f.Read(buf)
			if 0 == n {
				break
				bs.Write(buf[:n])
			}
		}
		er := json.Unmarshal(bs.Bytes(), &app)
		if er != nil {
			app.Port = "7890"
		}
	}

}

func main() {

	m := mux.NewRouter()
	m.HandleFunc("/", router.IndexHandler).Methods("GET")
	m.HandleFunc("/login", router.LoginHandler).Methods("POST")
	m.HandleFunc("/login", router.LoginHandler).Methods("GET")
	m.HandleFunc("/list", router.TodoListHandler)
	m.HandleFunc("/todo/add", router.TodoAddHandler).Methods("POST")
	m.HandleFunc("/todo/list/{userid:[0-9]+}", router.TodoUserListHandler).Methods("GET")
	m.HandleFunc("/todo/finish/{todoid:[0-9]+}", router.TodoFinishHandler)
	m.HandleFunc("/todo/finishlist/{userid:[0-9]+}", router.TodoUserFinishListHandler).Methods("GET")

	http.Handle("/public/", http.FileServer(http.Dir(".")))
	http.Handle("/", m)
	url := fmt.Sprintf(":%s", app.Port)
	//http.ListenAndServe(url, nil)
	fmt.Printf("服务器启动成功,地址：%s\n", url)
	log.Fatal("ListenAndServe: ", http.ListenAndServe(url, nil))

}
