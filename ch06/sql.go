package main

import "database/sql"

type Post struct {
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

func Posts(limit int) (posts []Post, err error) {

	rows, err := Db.Query("select id, content, author from posts limit $1", limit)
	if err != nil {
		return
	}
	for rows.Next(){
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	
}

func main() {

}