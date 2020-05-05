package main

import "fmt"

type bread struct {
	val string
}

type jam struct {
}

func (b *bread) putJam(jam *jam) {
	b.val += jam.getVal()
}

func (b *bread) String() string {
	return b.val
}

func (j *jam) getVal() string {
	return " + jam"
}

func main() {
	bread := &bread{val: "bread"}
	jam := &jam{}

	bread.putJam(jam)

	fmt.Println(bread)
}
