package main

import "fmt"

func main() {

	//for i:=0; i<=100; i++{
	//	if i%2==0{
	//		fmt.Println(i, "even")
	//	}else if i%3==0{
	//		fmt.Println(i, "odd")
	//	}
	//}

	for i:=1; i<=100; i++{
		switch {
		case i%2==0:
			fmt.Println(i, "even")
		default:
			fmt.Println(i, "odds")
		}
	}
}
