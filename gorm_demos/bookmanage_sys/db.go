package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/test"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(20) // 最大连接数
	db.SetMaxIdleConns(10) // 最大空闲连接数
	return
}

func insertBook(title string, price float64) (err error) {
	sqlStr := "INSERT INTO book(title, price, publisher_id) VALUES (?, ?, ?)"
	_, err = db.Exec(sqlStr, title, price, 1)
	if err != nil {
		fmt.Println("insert book failed, err:", err)
	}
	return
}

func queryAllBook() (bookList []*Book, err error) {
	sqlStr := "SELECT id, title, price FROM book;"
	err = db.Select(&bookList, sqlStr)
	if err != nil {
		fmt.Println("query all book failed, err:", err)
	}
	return
}

// queryOneBook 值拷贝，开销较大
func queryOneBook(id int64) (book Book, err error) {
	sqlStr := "SELECT id, title, price FROM book WHERE id=?"
	err = db.Get(&book, sqlStr, id)
	if err != nil {
		fmt.Printf("query book %d failed, err: %s\n", id, err.Error())
	}
	return
}

// queryOneBookV2 返回指针
func queryOneBookV2(id int64) (book *Book, err error) {
	// book = new(Book)
	book = &Book{}
	sqlStr := "select id, title, price from book where id=?"
	err = db.Get(book, sqlStr, id)
	if err != nil {
		fmt.Printf("query book %d failed, err: %s\n", id, err.Error())
	}
	return
}

func queryBookInfo(id int64) (book Book, err error) {
	// sqlStr := "SELECT id, title, price, publisher_id FROM book WHERE id=?"
	sqlStr := `
		SELECT b.id, title, price, p.province, p.city
		FROM book b
		JOIN publisher p
		ON b.publisher_id=p.id
		WHERE book.id=?`

	// 跨表查询
	// 方式一：直接使用 sql
	// 方式二：一张一张查，然后用代码组合
	err = db.Get(&book, sqlStr, id)
	if err != nil {
		fmt.Printf("query book %d failed, err: %s\n", id, err.Error())
	}
	return
}

func editBook(bookId int64, title string, price float64) (err error) {
	sqlStr := "UPDATE book SET title=?, price=? WHERE id=?"
	_, err = db.Exec(sqlStr, title, price, bookId)
	if err != nil {
		fmt.Println("edit book failed, err:", err)
	}
	return
}

func deleteBook(id int64) (err error) {
	sqlStr := "DELETE FROM book WHERE id=?"
	_, err = db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("delete book failed, err:", err)
	}
	return
}
