package main

import (
	"fmt"
	"time"
)

/*
	select를 채널의 switch문이라고 이해하면 된다.
	time package 안에 Tick과 After 를 사용하면 주기적으로 체크할때 편하게 사용 할 수 있다.

*/
func push(c chan int) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		c <- i
		i++

	}
}

func main() {
	c := make(chan int)

	go push(c)

	timerChan := time.After(10 * time.Second)
	tickTimerChan := time.Tick(2 * time.Second)
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-timerChan:
			fmt.Println("timeout")
			return
		case <-tickTimerChan:
			fmt.Println("Tick")
		}
	}
}
