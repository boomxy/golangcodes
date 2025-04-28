package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // SQLite3 驱动，注意要使用 _ 前缀导入
	"log"
	"os"
)

func main() {
	// 打开数据库，如果不存在则创建
	// 注意：这里 指定的 sqlite3 驱动需要安装
	// 可以使用 go get github.com/mattn/go-sqlite3 命令安装
	// 也可以使用 go mod tidy 命令自动安装
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// 创建表，插入数据
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS test (id INTEGER PRIMARY KEY, name TEXT);
	INSERT INTO test (name) VALUES ('Alice');
	INSERT INTO test (name) VALUES ('Bob');
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
	// 查询数据
	rows, err := db.Query("SELECT id, name FROM test")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var id int
	var name string
	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d: %s\n", id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}
