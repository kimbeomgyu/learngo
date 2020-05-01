package main

import "fmt"

// Student is 학생이다
type Student struct {
	name string
	age  int

	class string
	grade string
}

// ViewSungjuk is 성적을 출력한다
func (s *Student) ViewSungjuk() {
	fmt.Println(s.class, s.grade)
}

// InputSungjuk is 성적을 입력한다
func (s *Student) InputSungjuk(class string, grade string) {
	s.class = class
	s.grade = grade
}

// 포인터를 함수 인자로 받으면 메모리 주소만 복사
// 값을 함수 인자로 받으면 전체 속성을 복사

func main() {
	var s Student = Student{name: "Beomgyu", age: 23, class: "수학", grade: "A+"}

	s.InputSungjuk("과학", "C")
	s.ViewSungjuk()
}
