package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go sayHi(ch)

	for i := range ch {
		fmt.Println(i)
	}

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
}

//fmt.Println("1")
//fmt.Println("1")
//fmt.Println("1")
//fmt.Println("1")
//fmt.Println("1")
//fmt.Println(<-ch)

//fmt.Println("1")
//fmt.Println("2")
//fmt.Println("3")
//fmt.Println("4")
//fmt.Println("5")

//time.Sleep(2 * time.Second)

func say(word string) {
	fmt.Println(word)
}

func sayHi(exit chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		exit <- i
	}

	close(exit)

}
