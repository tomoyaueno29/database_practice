package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

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

func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil{
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	if err != nil {
		return
	}
	return
}

func getPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) update() (err error){
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

func (post  *Post) delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}


func main() {
	post := Post{Content: "Hello!", Author: "Sau"}
	fmt.Println(post)
	post.Create()
	fmt.Println(post)
	readPost, _ := getPost(1)
	fmt.Println(readPost)
	readPost.delete()
}