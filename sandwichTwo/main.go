package main

import "fmt"

type Jam interface {
	GetOneSpoon() SpoonOfJam
}

type SpoonOfJam interface {
	String() string
}

type Bread struct {
	val string
}

type AppleJam struct {
}

type OrangeJam struct {
}

type StrawberryJam struct {
}

type SpoonOfAppleJam struct {
}

type SpoonOfOrangeJam struct {
}

type SpoonOfStrawberryJam struct {
}

func (s *SpoonOfStrawberryJam) String() string {
	return "+ StrawberryJam"
}

func (s *SpoonOfOrangeJam) String() string {
	return "+ OrangeJam"
}

func (s *SpoonOfAppleJam) String() string {
	return "+ AppleJam"
}

func (j *StrawberryJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfStrawberryJam{}
}

func (j *OrangeJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfOrangeJam{}
}

func (j *AppleJam) GetOneSpoon() SpoonOfJam {
	return &SpoonOfAppleJam{}
}

func (b *Bread) PutJam(jam Jam) {
	spoon := jam.GetOneSpoon()
	b.val += spoon.String()
}

func (b *Bread) String() string {
	return "bread " + b.val
}

func main() {
	bread := &Bread{}
	// jam := &StrawberryJam{}
	// jam := &OrangeJam{}
	jam := &AppleJam{}

	bread.PutJam(jam)

	fmt.Println(bread)

}
