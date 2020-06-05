package main

import (
	"database/sql"
	"errors"
)

type Post struct {
	Id int
	Content string
	Author string
	Comments []Comment
}

type Comment struct {
	Id int
	Content string
	Author string
	Post *Post
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres",
		"user=tomoyaueno dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func (comment *Comment) Create() (err error){
	if comment.Post == nil {
		err = errors.New("No posts")
		return
	}
	err = Db.QueryRow(`insert into comments (content, author, post_id)
	values ($1, $2, $3) returning id`, comment.Content, comment.Author,
	comment.Post.Id).Scan(&comment.Id)
	return
}

func GetPost(id int) (post Post, err error){
	post = Post{}
	post.Comments = []Comment{}
	err = Db.QueryRow(`select id, content, author from posts 
	where id = $1`, id).Scan(&post.Id, &post.Content, &post.Author)

	rows, err := Db.Query("select id, content, author from comments")
	if err != nil {
		return
	}
	for rows.Next(){
		comment := Comment{Post: &post}
		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil {
			return
		}
		post.Comments = append(post.Comments, comment)
	}
	rows.Close()
	return
}