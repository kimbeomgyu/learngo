package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
	channel 을 이용해서 컨베이어 벨트 형태로 프로그램을 짤 수 있다.

*/

// Car is Car
type Car struct {
	val string
}

func makeTire(carChan chan Car, outChan chan Car) {
	for {
		car := <-carChan
		car.val += "Tire, "

		outChan <- car
	}
}

func makeEngine(carChan chan Car, outChan chan Car) {
	for {
		car := <-carChan
		car.val += "Engine, "

		outChan <- car
	}
}

func startwork(chan1 chan Car) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Car{val: "Car" + strconv.Itoa(i)}
		i++
	}
}

func main() {
	chan1 := make(chan Car)
	chan2 := make(chan Car)
	chan3 := make(chan Car)

	go startwork(chan1)
	go makeTire(chan1, chan2)
	go makeEngine(chan2, chan3)

	for {

		result := <-chan3

		fmt.Println(result.val)
	}

}
