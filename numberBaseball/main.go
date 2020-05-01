package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Result is 결과를 담는 struct
type Result struct {
	strikes int
	balls   int
}

/**
1. 무작위 숫자 3개를 받는다
2. 사용자의 입력을 받는다
3. 비교한다
3-1. 3개의 입력이 정확하면 게임이 끝난다
3-2. 비교했을때 결과가 다른 경우 현재 숫자와 무작위 숫자를 비교해
	 다시 2로 이동해 입력을 받는다

*/
func main() {
	rand.Seed(time.Now().UnixNano())
	numbers := MakeNumbers()

	count := 0
	for {
		count++

		inputNumbers := InputNumbers()

		result := CompareNumbers(numbers, inputNumbers)

		PrintResult(result)

		if IsGameEnd(result) {
			break

		}
	}

	fmt.Printf("%d 번만에 맞췄습니다\n", count)
}

// MakeNumbers is 0~9 사이의 겹치지 않는 무작위 숫자 3개를 반환한다
func MakeNumbers() [3]int {
	var result [3]int

	for i := 0; i < 3; i++ {
		for {
			n := rand.Intn(10)
			duplicated := false
			for j := 0; j < i; j++ {
				if result[j] == n {
					// 겹치는 경우 다시 뽑는다.
					duplicated = true
					break
				}
			}
			if !duplicated {
				result[i] = n
				break
			}
		}
	}
	// fmt.Println(result) // 확인용
	return result
}

// InputNumbers is 키보드로부터 0~9사이의 겹치지 않는 숫자 3개를 입력을 받아 반환한다
func InputNumbers() [3]int {
	var result [3]int

	for {
		fmt.Println("겹치지 않는 0 ~ 9 사이의 숫자 3개를 입력해 주세요")
		var no int
		_, err := fmt.Scanf("%d\n", &no)
		if err != nil {
			fmt.Println("잘못 입력하였습니다")
			continue
		}

		success := true
		idx := 0
		for no > 0 {
			n := no % 10
			no = no / 10

			duplicated := false
			for j := 0; j < idx; j++ {
				if result[j] == n {
					// 겹치는 경우 다시 뽑는다.
					duplicated = true
					break
				}
			}
			if duplicated {
				fmt.Println("숫자가 겹치지 않아야 합니다")
				success = false
				break
			}
			if idx >= 3 {
				fmt.Println("3개보다 많은 숫자를 입력하셧습니다")
				success = false
				break
			}

			result[idx] = n
			idx++
		}
		if success && idx < 3 {
			fmt.Println("3개의 숫자를 입력하세요")
			success = false
		}

		if !success {
			continue
		}

		break
	}
	result[0], result[2] = result[2], result[0] // reverse
	// fmt.Println(result) // 확인용
	return result
}

// CompareNumbers is 두개의 숫자 3개를 비교해서 결과를 반환한다
func CompareNumbers(numbers, inputNumbers [3]int) Result {
	strikes := 0
	balls := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if numbers[i] == inputNumbers[j] {
				if i == j {
					strikes++
				} else {
					balls++
				}
				break
			}
		}
	}
	return Result{strikes, balls}
}

// PrintResult is 숫자의 결과를 반환한다
func PrintResult(result Result) {
	fmt.Printf("%dS, %dB \n", result.strikes, result.balls)
}

// IsGameEnd is 비교 결과가 3 스트라이크 인지 확인
func IsGameEnd(result Result) bool {
	return 3 == result.strikes
}
