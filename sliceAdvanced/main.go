package main

import "fmt"

/*
	slice도 내부를 까보면 struct로 구성되어 있다!

	대략 이런느낌

	type Slice struct {
		pointer -> 시작주소
		len -> 갯수
		cap -> 최대 갯수
	}


*/

func main() {
	s := make([]int, 3)
	s[0] = 100
	s[1] = 200
	s[2] = 300

	fmt.Println(s, len(s), cap(s))

	t := append(s, 400)

	fmt.Println("////////////////")
	fmt.Println(s, len(s), cap(s))
	fmt.Println(t, len(t), cap(t))

	u := append(t, 500)

	fmt.Println("////////////////")
	fmt.Println(s, len(s), cap(s))
	fmt.Println(t, len(t), cap(t))
	fmt.Println(u, len(u), cap(u))

	u[0] = 9999

	fmt.Println("////////////////")
	fmt.Println(s, len(s), cap(s))
	fmt.Println(t, len(t), cap(t))
	fmt.Println(u, len(u), cap(u))

}
