package main

import "github.com/jmoiron/sqlx"

type Post struct {
	Id int
	Content string
	AuthorName string `db: author`
}

var Db *sqlx.DB

func init() {
	var err error
	Db, err = sqlx.Open("postgres", "user=tomoyaueno dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRowx("select id, content, author from posts where id=$1", id).StructScan(&post)
	if err != nil {
		return
	}
	return
}

func (post *Post) Create() (err error) {
	err = Db.QueryRow("insert into posts (content, author) values ($1, $2) returning id", post.Content, post.AuthorName).Scan(&post.Id)
	return
}

