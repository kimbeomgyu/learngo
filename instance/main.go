package main

import "fmt"

// Student is student
type Student struct {
	name  string
	age   int
	grade int
}

/*
	instance는 컴퓨터 입장에서 보면 포인터타입의 student이다.
	사람의 입장에서 봤을때는 어떠한 대상으로써 추상적인 이미지이다.
*/
func main() {
	a := Student{"aaa", 20, 10}

	b := &a
	// 주소를 변수에 지정하고 값을 변경할 경우 기존의 값도 변경된다.
	b.age = 30

	fmt.Println(a)
	fmt.Println(*b)

	a.setName("bbb")
	a.setAge(25)

	fmt.Println(a)
	fmt.Println(*b)

}

func (s *Student) setName(name string) {
	s.name = name
}

func (s *Student) setAge(age int) {
	s.age = age
}
