package main

import (
	"fmt"

	"github.com/myProject/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	err := dictionary.Delete("base")

	if err != nil {
		fmt.Println(err)
	}

	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)

}
