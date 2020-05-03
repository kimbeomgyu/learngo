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
	var a []int
	b := make([]int, 0, 8) // 배열을 생성할 때 초기 값을 아무것도 주지 않고, 공간을 8만큼 확보해둔다.
	a = append(a, 1)
	b = append(b, 1)

	fmt.Printf("len(a) = %d\n", len(a))
	fmt.Printf("cap(a) = %d\n", cap(a))

	fmt.Printf("len(b) = %d\n", len(b))
	fmt.Printf("cap(b) = %d\n", cap(b))

	// 이 부분에서 capacity가 있기 때문에 c와 d가 동일한 주소값으로 연결되어 있다.
	// 만약 capacity가 없다면 새로운 주소로 값이 복사되어 c와 d가 달라진다.
	c := make([]int, 2, 4)
	c[0] = 1
	c[1] = 2

	d := append(c, 3)

	fmt.Printf("%p %p\n", c, d)

	fmt.Println(c)
	fmt.Println(d)

	d[0] = 4
	d[1] = 5

	fmt.Println(c)
	fmt.Println(d)

	// 공간을 다르게 쓰고싶은게 확실한 경우
	e := []int{1, 2}

	f := make([]int, len(e))

	for i := 0; i < len(e); i++ {
		f[i] = e[i]

	}

	f = append(e, 3)

	fmt.Printf("%p %p\n", e, f)

	f[0] = 4
	f[1] = 5

	fmt.Println(e)
	fmt.Println(f)
}
