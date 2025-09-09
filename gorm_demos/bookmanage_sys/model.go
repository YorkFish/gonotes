package main

type Publisher struct {
	Id       int64  `db:"id"`
	Province string `db:"province"`
	City     string `db:"city"`
}

type Book struct {
	Id    int64   `db:"id"`
	Title string  `db:"title"`
	Price float64 `db:"price"`
	Publisher
}
