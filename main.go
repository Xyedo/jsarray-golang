package main

import (
	"fmt"

	"xyedo.dev/learnGo/arrayJS"
)

func main() {
	arrayJS.New(10)
	myFish := arrayJS.JSArray{"angel", "clown", "mandarin", "sturgeon"}
	ex1, _ := myFish.Splice(2, 0, "drum")
	fmt.Println("ex1:", ex1)
	forEx2 := arrayJS.JSArray{"angel", "clown", "mandarin", "sturgeon"}
	ex2, _ := forEx2.Splice(2, 0, "drum", "guitar")
	fmt.Println("ex2:", ex2)
	forEx3 := arrayJS.JSArray{"angel", "clown", "drum", "mandarin", "sturgeon"}
	ex3, removed3 := forEx3.Splice(3, 1)
	fmt.Println("ex3:", ex3, removed3)
	forEx4 := arrayJS.JSArray{"angel", "clown", "drum", "sturgeon"}
	ex4, removed4 := forEx4.Splice(2, 1, "trumpet")
	fmt.Println("ex4:", ex4, removed4)
	forEx5 := arrayJS.JSArray{"angel", "clown", "drum", "sturgeon"}
	ex5, removed5 := forEx5.Splice(0, 2, "parrot", "anemone", "blue")
	fmt.Println("ex5:", ex5, removed5)
	forEx6 := arrayJS.JSArray{"parrot", "anemone", "blue", "trumpet", "sturgeon"}
	ex6, removed6 := forEx6.Splice(2, 2)
	fmt.Println("ex6:", ex6, removed6)

	forEx7 := arrayJS.JSArray{"angel", "clown", "mandarin", "sturgeon"}
	ex7, removed7 := forEx7.Splice(2, 2)
	fmt.Println("ex7:", ex7, removed7)
	ex8 := arrayJS.JSArray{1, 2, 3, 4, 5}
	ex8, poppedVal := ex8.Pop()
	fmt.Println(ex8, poppedVal)
	ex8 = ex8.Push(1, 3, 6)
	fmt.Println(ex8)
	ex8, shiftedVal := ex8.Shift()
	fmt.Println(ex8, shiftedVal)
	ex8 = ex8.Unshift(10, 11)
	fmt.Println(ex8)
}
