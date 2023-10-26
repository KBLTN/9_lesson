package main

import (
	"fmt"
	"time"
)

// HW: Без примитивов синхронизации (без пакета sync) написать калькулятор на горутинах.
// 2+2, 3+6, 7*7, 9/3
// 2+2=4, 3+6=9, 7*7=49, 9/3=3
// GitHub, Create Repository, Create Push Branch calculator, Write Code, Commit, Push Branch: Calculator, Create PullRequest calculator -> calc

func main() {
	data := make(chan int)
	exit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-data)
		}
		exit <- 0
	}()
	selectOne(data, exit)
}

func selectOne(data, exit chan int) {
	x := 0
	for {
		select {
		case data <- x:
			x += 1
		case <-exit:
			fmt.Println("exit")
			return
		default:
			fmt.Println("waiting")
			time.Sleep(5 * time.Millisecond)
		}
	}
}
