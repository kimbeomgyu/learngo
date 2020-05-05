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

// Plane is Plane
type Plane struct {
	val string
}

func makeTire(carChan chan Car, planeChan chan Plane, outCarChan chan Car, outPlaneChan chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += " Tire_C, "
			outCarChan <- car
		case plane := <-planeChan:
			plane.val += " Tire_P,"
			outPlaneChan <- plane
		}
	}
}

func makeEngine(carChan chan Car, planeChan chan Plane, outCarChan chan Car, outPlaneChan chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += "Engine_C, "
			outCarChan <- car
		case plane := <-planeChan:
			plane.val += "Engine_P, "
			outPlaneChan <- plane
		}
	}
}

func startCarWork(chan1 chan Car) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Car{val: "Car" + strconv.Itoa(i)}
		i++
	}
}

func startPlaneWork(chan1 chan Plane) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Plane{val: "Plane" + strconv.Itoa(i)}
		i++
	}
}

func main() {
	carChan1 := make(chan Car)
	carChan2 := make(chan Car)
	carChan3 := make(chan Car)

	planeChan1 := make(chan Plane)
	planeChan2 := make(chan Plane)
	planeChan3 := make(chan Plane)

	go startCarWork(carChan1)
	go startPlaneWork(planeChan1)
	go makeTire(carChan1, planeChan1, carChan2, planeChan2)
	go makeEngine(carChan2, planeChan2, carChan3, planeChan3)

	for {
		select {
		case result := <-carChan3:
			fmt.Println(result.val)
		case result := <-planeChan3:
			fmt.Println(result.val)
		}
	}

}
