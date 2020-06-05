package main

import "fmt"

type Post struct {
	Id int
	Content string
	Author string
}

var PostById map[int]*Post
var PostsByAuthor map[string][]*Post

func store(post Post){
	PostById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func main() {
	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := Post{Id:1, Content:"Hello!", Author:"Sau Sheong"}
	post2 := Post{Id:2, Content:"World!", Author:"Pedro"}

	store(post1)
	store(post2)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostsByAuthor["Sau Sheong"] {
		fmt.Println(post)
	}
	for _, post := range PostsByAuthor["Pedro"] {
		fmt.Println(post)
	}
}