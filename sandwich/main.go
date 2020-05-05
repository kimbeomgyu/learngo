package main

import "fmt"

type bread struct {
	val string
}

type strawberryJam struct {
	opened bool
}

type spoonOfStrawberry struct {
}

type sandwitch struct {
	val string
}

func getBreads(num int) []*bread {
	breads := make([]*bread, num)
	for i := 0; i < num; i++ {
		breads[i] = &bread{val: "bread"}
	}
	return breads
}

func openStrawberryJam(jam *strawberryJam) {
	jam.opened = true
}

func getOneSpoon(_ *strawberryJam) *spoonOfStrawberry {
	return &spoonOfStrawberry{}
}

func putJamOnBread(bread *bread, jam *spoonOfStrawberry) {
	bread.val += " + Strawberry Jam"
}

func makeSandwitch(breads []*bread) *sandwitch {
	sandwitch := &sandwitch{}
	for i := 0; i < len(breads); i++ {
		sandwitch.val += breads[i].val + " + "
	}
	return sandwitch
}

func main() {
	// 1. 빵 두개를 꺼낸다.
	breads := getBreads(2)

	jam := &strawberryJam{}

	// 2. 딸기잼 뚜껑을 연다.
	openStrawberryJam(jam)

	// 3. 딸기잼을 한스푼 뜬다.
	spoon := getOneSpoon(jam)

	// 4. 딸기잼을 빵에 바른다.
	putJamOnBread(breads[0], spoon)

	// 5. 빵을 덮는다.
	sandwitch := makeSandwitch(breads)

	// 6. 완성
	fmt.Println(sandwitch.val)
}
