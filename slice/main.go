package main

import "fmt"

/*	Slice 주소에 대해 굉장히 주의해야 한다.

	# 동적배열이란 실제 정적배열을 포인트하고 있는 것이다.
	# 길이가 증가하게되면 새로운 정적배열을 생성하고 기존의 정적배열을 복사해 포인트를 변경한다.
	# Capacity를 미리 여유있게 지정을 해두면 기존의 정적배열의 주소값에 연결해서 값을 저장하므로 새로 복사되는 배열이 없다.

	동적배열에서 실제 사용하는 길이 이외에 확보한 공간이 존재한다.
	배열안의 공간을 확장할 때에는 2, 4, 8 순으로 2배씩 증가하고
	공간에서 실제 사용하는 공간을 length로 가진다.
	2배씩 증가하는 부분을 Capacity라고 한다.

	배열안에 항목을 추가할 때에는 append라는 내장함수를 활용한다.
*/

func main() {
	slice()
	sliceCapacity()
	sliceCopy()
	sliceRef()
	sliceRemove() // slice는 복사가 아니기 때문에 주의해야 한다!

}

func sliceRemove() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < 5; i++ {
		var back int
		a, back = removeBack(a)
		fmt.Printf("%d, ", back)
	}
	fmt.Println()
	fmt.Println(a)

	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < 5; i++ {
		var front int
		b, front = removeFront(b)
		fmt.Printf("%d, ", front)
	}
	fmt.Println()
	fmt.Println(b)
}
func removeBack(a []int) ([]int, int) {
	return a[:len(a)-1], a[len(a)-1]
}

func removeFront(a []int) ([]int, int) {
	return a[1:], a[0]
}

/*
	이름에서 알 수 있듯이 잘라내서 사용한다.
	잘라낸 부분은 새로운 값이 아닌 주소값에서 잘라서 같은 곳을 바라보는 상태이다!
	같은 메모리 영역을 가리키고 있다.
*/
func sliceRef() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	b := a[4:8]
	c := a[4:]
	d := a[:4]

	c[0] = 2
	d[2] = 4

	fmt.Printf("%p %d\n", a, a)
	fmt.Printf("%p %d\n", b, b)
	fmt.Printf("%p %d\n", c, c)
	fmt.Printf("%p %d\n", d, d)
}

// 공간을 다르게 쓰고싶은게 확실한 경우
func sliceCopy() {
	a := []int{1, 2}

	b := make([]int, len(a))

	for i := 0; i < len(a); i++ {
		b[i] = a[i]

	}

	b = append(a, 3)

	fmt.Printf("%p %p\n", a, b)

	b[0] = 4
	b[1] = 5

	fmt.Println(a)
	fmt.Println(b)
}

func sliceCapacity() {
	// 이 부분에서 capacity가 있기 때문에 a와 b가 동일한 주소값으로 연결되어 있다.
	// 만약 capacity가 없다면 새로운 주소로 값이 복사되어 a와 b가 달라진다.
	a := make([]int, 2, 4)
	a[0] = 1
	a[1] = 2

	b := append(a, 3)

	fmt.Printf("%p %p\n", a, b)

	fmt.Println(a)
	fmt.Println(b)

	b[0] = 4
	b[1] = 5

	fmt.Println(a)
	fmt.Println(b)
}

func slice() {
	var a []int
	b := make([]int, 0, 8) // 배열을 생성할 때 초기 값을 아무것도 주지 않고, 공간을 8만큼 확보해둔다.
	a = append(a, 1)
	b = append(b, 1)

	fmt.Printf("len(a) = %d\n", len(a))
	fmt.Printf("cap(a) = %d\n", cap(a))

	fmt.Printf("len(b) = %d\n", len(b))
	fmt.Printf("cap(b) = %d\n", cap(b))

}
