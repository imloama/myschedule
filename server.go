/**
 * 启动服务器
 */
package main

import (
	"github.com/gorilla/mux"
	"github.com/itwarcraft/myschedule/router"
	"log"
	"net/http"
)

//在init函数中初始化
func init() {
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
	http.ListenAndServe(":7890", nil)
	log.Fatalln("server started")

}
