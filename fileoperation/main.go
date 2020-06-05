package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main(){
	bytes := []byte("Hello!\n")
	err := ioutil.WriteFile("xfile", bytes, 0644)
	if err != nil{
		panic(err)
	}
	read1, _ := ioutil.ReadFile("xfile")
	fmt.Println(string(read1))

	files, _ := ioutil.ReadDir("/fileoperation")
	for _, file2 := range files {
		fmt.Println(file2.Name())
	}

	file3, _ := os.Create("data2")
	defer file3.Close()

	aa, _ := file3.Write(bytes)

}