package main

import (
	"fmt"
	"time"
)

/*
	thread란 하나의 작업을 수행하는 일이라고 생각하면된다.
	cpu가 thread를 여러개를 작업할 때에 context switching 작업이 일어나는데 이 작업의 비용이 크다.
	그래서 thread가 너무 많아지면 작업이 수행되는 과정보다 더 큰 비용을 소모하게 되어 버린다.
	go에서는 thread를 직접 조작하지 않고 thread 하나를 go thread라는 개념으로 분리해 나누어서 사용하게 되고,
	context switching이 최소한으로 발생하게 만들어준다.
*/
func main() {
	go fun1() // go thread
	for i := 0; i < 20; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("main: ", i)
	}

	fmt.Scanln()
}

func fun1() {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("fun1: ", i)
	}
}
