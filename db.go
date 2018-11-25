package main

import (
	"database/sql"
	"fmt"
)
type Book struct{
	Name string
	Year string
	Length string
}
const (
	DB_USER = "sysudmlbylkzhn"
	DB_PASSWORD = "606e1d22ba3a82107f0fe77589301c6f886fb7fe724a4c0f3cd1902a4f22fff4"
	DB_NAME = "d7om04nn29ehre"
	DB_HOST = "ec2-54-83-8-246.compute-1.amazonaws.com"
)
func dbConnect() error {
	var err error
	db, err = sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable",
	DB_USER, DB_PASSWORD, DB_NAME, DB_HOST))
	if err != nil {
		return err
	}
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS books (book_name text,book_year text,book_length text)"); err != nil {
		return err
	}
	return nil
}
func dbAddBook(name, year, length string) error {
	sqlstmt := "INSERT INTO books VALUES ($1, $2, $3)"
	_, err := db.Exec(sqlstmt, name, year, length)
	if err != nil {
		return err
	}
	return nil
}
func dbGetBooks() ([]Book, error) {
	var books []Book
	stmt, err := db.Prepare("SELECT book_name, book_year, book_length FROM books")
	if err != nil {
		return books, err
	}
	res, err := stmt.Query()
	if err != nil {
		return books, err
	}
	var tempBook Book
	for res.Next() {
		err = res.Scan(&tempBook.Name, &tempBook.Year, &tempBook.Length)
		if err != nil {
			return books, err
		}
		books = append(books, tempBook)
	}
	return books, err
}
