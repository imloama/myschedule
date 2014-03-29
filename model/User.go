package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/itwarcraft/myschedule/util"
	"log"
)

/**
 * 登陆用户
 */
type User struct {
	Id       int
	Name     string
	Password string
	Email    string
}

func (user *User) Save() error {
	h := md5.New()
	h.Write([]byte(user.Password))

	sql := fmt.Sprintf("insert into sys_user(name,password,email)values('%s','%s','%s');", user.Name, hex.EncodeToString(h.Sum(nil)), user.Email)
	db := util.NewDb("./myschedule.db")
	err := db.Exec(sql)
	return err
}

func (user *User) Login() error {
	h := md5.New()
	h.Write([]byte(user.Password))
	sql := fmt.Sprintf("select Id,Name,Password,Email from sys_user where name='%s' and password='%s';", user.Name, hex.EncodeToString(h.Sum(nil)))
	db := util.NewDb("./myschedule.db")
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal("user login error!!!")
		log.Fatal(err)
		return err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&user.Id, &user.Name, &user.Password, &user.Email)
	}
	return nil
}

func GetUserByName(name string) (*User, error) {
	sql := fmt.Sprintf("select Id,Name,Password,Email from sys_user where name='%s';", name)
	fmt.Printf("查询sql语句:%s\n", sql)
	db := util.NewDb("./myschedule.db")
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal("user login error!!!")
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	var user User
	for rows.Next() {
		rows.Scan(&user.Id, &user.Name, &user.Password, &user.Email)
	}
	fmt.Printf("用户[id:%d,name:%s,password:%s,email:%s]\n", user.Id, user.Name, user.Password, user.Email)
	return &user, nil
}
