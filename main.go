package main

import (
	"fmt"

	"github.com/Xyedo/jsarray-golang/arrayJS"
)

func main() {
	test := arrayJS.New(10)
	isarr := arrayJS.IsArray(test)
	fmt.Println(isarr)
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

	ex9 := arrayJS.JSArray{1, 2, 3, 4, 5}
	fmt.Println(ex9.Length())
	fmt.Println(ex9.CopyWithin(-2))
	fmt.Println(arrayJS.JSArray{1, 2, 3, 4, 5}.CopyWithin(0, 3))
	fmt.Println(arrayJS.JSArray{1, 2, 3, 4, 5}.CopyWithin(0, 3, 4))
	fmt.Println(arrayJS.JSArray{1, 2, 3, 4, 5}.CopyWithin(-2, -3, -1))
	fmt.Println(arrayJS.JSArray{1, 2, 3, 4, 5}.CopyWithin(0, 2))
	fmt.Println(arrayJS.JSArray{1, 2, 3, 4, 5}.CopyWithin(0, 3, 4))
	isBigEnough := func(element any, index int, array arrayJS.JSArray) bool {
		el, ok := element.(int)
		if !ok {
			return false
		}
		return el >= 10
	}
	fmt.Println(arrayJS.JSArray{12, 5, 8, 130, 44}.Every(isBigEnough))
	fmt.Println(arrayJS.JSArray{12, 54, 18, 130, 44}.Every(isBigEnough))
	// isSubset = func(arr1, arr2 arrayJS.JSArray) bool {
	// 	arr2.Every(func(element any, index int, array arrayJS.JSArray) bool {

	// 	})
	// }
	fmt.Println(arrayJS.JSArray{12, 5, 8, 130, 44}.Every(func(element any, index int, array arrayJS.JSArray) bool {
		el, ok := element.(int)
		if !ok {
			return false
		}
		return el >= 10
	}))
	fmt.Println(arrayJS.JSArray{12, 54, 18, 130, 44}.Every(func(element any, index int, array arrayJS.JSArray) bool {
		el, ok := element.(int)
		if !ok {
			return false
		}
		return el >= 10
	}))
	sum := arrayJS.JSArray{0, 1, 2, 3}.Reduce(func(previousValue, curentValue any, currentIndex int, array arrayJS.JSArray) any {
		prevVal, ok := previousValue.(int)
		if !ok {
			return nil
		}
		currVal, ok := curentValue.(int)
		if !ok {
			return nil
		}
		return prevVal + currVal
	}, 0)
	fmt.Println(sum)
	fmt.Println(arrayJS.JSArray{2, 5, 8, 1, 4}.Some(isBigEnough))
	fmt.Println(arrayJS.JSArray{12, 5, 8, 1, 4}.Some(isBigEnough))
	array1 := arrayJS.JSArray{"a", "b", "c"}
	array2 := arrayJS.JSArray{1, 2, 3}
	fmt.Println(array1.Concat(array2...))

	num1 := arrayJS.JSArray{1, 2, 3}
	num2 := arrayJS.JSArray{4, 5, 6}
	num3 := arrayJS.JSArray{7, 8, 9}

	fmt.Println(num1.Concat(num2...).Concat(num3...))

}
