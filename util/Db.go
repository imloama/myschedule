package util

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Db struct {
	conn *sql.DB
}

func NewDb(database string) *Db {
	var db Db
	conn, _ := sql.Open("sqlite3", database)
	db.conn = conn

	//defer db.conn.Close()
	return &db
}

/**
 * 执行sql语句
 */
func (db *Db) Exec(sql string) error {
	_, err := db.conn.Exec(sql)
	defer db.conn.Close()
	return err
}

func (db *Db) Query(sql string) (*sql.Rows, error) {
	rows, err := db.conn.Query(sql)
	defer db.conn.Close()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return rows, nil
}
