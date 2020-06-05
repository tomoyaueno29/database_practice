package main

import "database/sql"

type Post struct{
	Id int
	Content string
	Author string
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=tomoyaueno dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}