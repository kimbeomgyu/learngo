package main

import "fmt"

/*
	멀티쓰레드를 하는 이유
	"공짜 점심이 끝났다" 무어 -> 인텔의 전 CEO -> "cpu는 1년에 2배씩 좋아질 것이다!"
	인텔 집적도가 너무 높아져 나노공정의 단계에서 물리적 한계에 부딪힘 -> 열, 자기장 방해
	-> 멀티코어로 동선변경 -> cpu가 2배씩 빨라졌다 -> 가만히 있어도 프로그램이 2배씩 빨라졌다

	멀티프로세스

	channel은 go에서 제공하는 쓰레드 queue라고 생각하면 된다.
	fixed size
	push,pop <-

*/

func pop(c chan int) {
	fmt.Println("Pop func")
	v := <-c
	fmt.Println(v)
}

func main() {
	var c chan int
	c = make(chan int)

	go pop(c)
	c <- 10

	fmt.Println("end of program")
}
