package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

type Post struct{
	Id int
	Content string
	Author string
}

func main() {
	csvFile, err := os.Create("posts.csv")
	if err !=  nil{
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []Post{
		Post{Id:1, Content:"Hello!", Author:"Sau Sheong"},
		Post{Id:2, Content:"World!", Author:"Pedro"},
	}

	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts{
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}

		err := writer.Write(line)
		if 
	}
}