package main

import "fmt"

// Student is 학생이다
type Student struct {
	name  string
	class int

	sungjuk Sungjuk
}

// Sungjuk is 학생의 성적이다
type Sungjuk struct {
	name  string
	grade string
}

// ViewSungjuk is 성적을 출력한다
func (s Student) ViewSungjuk() {
	fmt.Println(s.sungjuk)
}

func main() {
	var s Student

	s.name = "Smith"
	s.class = 1

	s.sungjuk.name = "수학"
	s.sungjuk.grade = "C"

	s.ViewSungjuk()
}
