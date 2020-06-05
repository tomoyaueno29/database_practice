package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

		rand.Seed(time.Now().UnixNano())
			dise := rand.Intn(6)
			fmt.Println(dise)
	switch dise + 1 {

	}
			if dise==6{
				fmt.Println("大吉")
			}else if dise==5 || dise==4{
				fmt.Println("中吉")
			}else if dise==3 || dise==2{
				fmt.Println("吉")
			}else{
				fmt.Println("凶")
			}
}
