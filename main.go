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

// InputSungjuk is 성적을 입력한다
func (s Student) InputSungjuk(name string, grade string) {
	s.sungjuk.name = name
	s.sungjuk.grade = grade
}

// ViewSungjuk is 성적을 출력한다
func ViewSungjuk(s Student) {
	fmt.Println(s.sungjuk)
}

func main() {
	var s Student

	s.name = "Smith"
	s.class = 1

	s.sungjuk.name = "수학"
	s.sungjuk.grade = "C"

	// 같은 기능이다
	s.ViewSungjuk() // 수학 C
	ViewSungjuk(s)  // 수학 C

	// 과학 A 로 변경되지 않는다 입력이 되었을때
	// 입력되는 값과 함수 내부에서 사용되는 값은 완전히 다른 메모리 주소값이다.
	// 복사로 동작해서 기능이 동작하지 않는다.
	// 이걸 해결하기 위해 포인터를 이용한다.
	s.InputSungjuk("과학", "A")

	s.ViewSungjuk() // 수학 C
}
