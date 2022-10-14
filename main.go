package main

import (
	"fmt"

	"github.com/Xyedo/jsarray-golang/arrayJS"
)

func main() {
	test := arrayJS.New(10)
	isarr := arrayJS.IsArray(test)
	fmt.Println(isarr)
	fmt.Println(arrayJS.JSArray[int]{1, 2, 3, 4, 5}.At(-2))
	myFish := arrayJS.JSArray[string]{"angel", "clown", "mandarin", "sturgeon"}
	ex1, removed1 := myFish.Splice(2, 0, "drum")
	fmt.Println("ex1:", ex1, removed1)
	forEx2 := arrayJS.JSArray[string]{"angel", "clown", "mandarin", "sturgeon"}
	ex2, removed2 := forEx2.Splice(2, 0, "drum", "guitar")
	fmt.Println("ex2:", ex2, removed2)
	forEx3 := arrayJS.JSArray[string]{"angel", "clown", "drum", "mandarin", "sturgeon"}
	ex3, removed3 := forEx3.Splice(3, 1)
	fmt.Println("ex3:", ex3, removed3)
	forEx4 := arrayJS.JSArray[string]{"angel", "clown", "drum", "sturgeon"}
	ex4, removed4 := forEx4.Splice(2, 1, "trumpet")
	fmt.Println("ex4:", ex4, removed4)
	forEx5 := arrayJS.JSArray[string]{"angel", "clown", "drum", "sturgeon"}
	ex5, removed5 := forEx5.Splice(0, 2, "parrot", "anemone", "blue")
	fmt.Println("ex5:", ex5, removed5)
	forEx6 := arrayJS.JSArray[string]{"parrot", "anemone", "blue", "trumpet", "sturgeon"}
	ex6, removed6 := forEx6.Splice(2, 2)
	fmt.Println("ex6:", ex6, removed6)

	forEx7 := arrayJS.JSArray[string]{"angel", "clown", "mandarin", "sturgeon"}
	ex7, removed7 := forEx7.Splice(2, 2)
	fmt.Println("ex7:", ex7, removed7)
	ex8 := arrayJS.JSArray[int]{1, 2, 3, 4, 5}
	ex8, poppedVal := ex8.Pop()
	fmt.Println(ex8, poppedVal)
	ex8 = ex8.Push(1, 3, 6)
	fmt.Println(ex8)
	ex8, shiftedVal := ex8.Shift()
	fmt.Println(ex8, shiftedVal)
	ex8 = ex8.Unshift(10, 11)
	fmt.Println(ex8)

	ex9 := arrayJS.JSArray[int]{1, 2, 3, 4, 5}
	fmt.Println(ex9.Length())
	fmt.Println(ex9.CopyWithin(-2))
	fmt.Println(arrayJS.JSArray[int]{1, 2, 3, 4, 5}.CopyWithin(0, 3))
	fmt.Println(arrayJS.JSArray[int]{1, 2, 3, 4, 5}.CopyWithin(0, 3, 4))
	fmt.Println(arrayJS.JSArray[int]{1, 2, 3, 4, 5}.CopyWithin(-2, -3, -1))
	fmt.Println(arrayJS.JSArray[int]{1, 2, 3, 4, 5}.CopyWithin(0, 2))
	fmt.Println(arrayJS.JSArray[int]{1, 2, 3, 4, 5}.CopyWithin(0, 3, 4))
	isBigEnough := func(el int, index int, array arrayJS.JSArray[int]) bool {
		return el >= 10
	}
	fmt.Println(arrayJS.JSArray[int]{12, 5, 8, 130, 44}.Every(isBigEnough))
	fmt.Println(arrayJS.JSArray[int]{12, 54, 18, 130, 44}.Every(isBigEnough))
	isSubset := func(arr1, arr2 arrayJS.JSArray[int]) bool {
		return arr2.Every(func(element int, index int, array arrayJS.JSArray[int]) bool {
			return arr1.Includes(element)
		})
	}
	fmt.Println("is Subset", isSubset(arrayJS.JSArray[int]{1, 2, 3, 4, 5, 6, 7}, arrayJS.JSArray[int]{5, 6, 7}))
	fmt.Println(arrayJS.JSArray[int]{12, 5, 8, 130, 44}.Every(func(element, index int, array arrayJS.JSArray[int]) bool {
		return element >= 10
	}))
	fmt.Println(arrayJS.JSArray[int]{12, 54, 18, 130, 44}.Every(func(element int, index int, array arrayJS.JSArray[int]) bool {

		return element >= 10
	}))
	sum := arrayJS.JSArray[int]{0, 1, 2, 3}.Reduce(func(previousValue, curentValue, currentIndex int, array arrayJS.JSArray[int]) int {
		return previousValue + curentValue
	}, 0)
	fmt.Println(sum)
	fmt.Println(arrayJS.JSArray[int]{2, 5, 8, 1, 4}.Some(isBigEnough))
	fmt.Println(arrayJS.JSArray[int]{12, 5, 8, 1, 4}.Some(isBigEnough))

	num1 := arrayJS.JSArray[int]{1, 2, 3}
	num2 := arrayJS.JSArray[int]{4, 5, 6}
	num3 := arrayJS.JSArray[int]{7, 8, 9}

	fmt.Println(num1.Concat(num2...).Concat(num3...))
	fmt.Println("Fill")
	fmt.Println(arrayJS.JSArray[int]{1, 2, 3, 4}.Fill(0, 2, 4))
	fmt.Println(arrayJS.JSArray[int]{1, 2, 3, 4}.Fill(5, 1))
	fmt.Println(arrayJS.JSArray[int]{1, 2, 3, 4}.Fill(6))
	fmt.Println(arrayJS.JSArray[int]{1, 2, 3}.Fill(4, -3, -2))
	fmt.Println("Filter")
	fmt.Println(arrayJS.JSArray[int]{12, 5, 8, 130, 44}.Filter(isBigEnough))
	fmt.Println("Find")
	fmt.Println(arrayJS.JSArray[int]{12, 5, 8, 130, 44}.Find(isBigEnough))
	fmt.Println("FindIndex")
	fmt.Println(arrayJS.JSArray[int]{12, 5, 8, 130, 44}.FindIndex(isBigEnough))
	fmt.Println(arrayJS.JSArray[int]{1, 5, 8, 0, 44}.FindIndex(isBigEnough))
	fmt.Println("IndexOf")
	fmt.Println(arrayJS.JSArray[int]{12, 5, 8, 130, 44}.IndexOf(12))
	fmt.Println(arrayJS.JSArray[int]{1, 5, 8, 0, 4}.IndexOf(0))
	fmt.Println("includes")
	fmt.Println(arrayJS.JSArray[int]{12, 5, 8, 130, 44}.Includes(5))
	fmt.Println(arrayJS.JSArray[int]{1, 5, 8, 0, 4}.Includes(-1))

	fmt.Println("Map")
	fmt.Println(
		arrayJS.JSArray[int]{12, 5, 8, 130, 44}.Map(func(element, index int, array arrayJS.JSArray[int]) int {
			return element * 2
		}))
	fmt.Println("Reverse")
	arr := arrayJS.JSArray[int]{1, 5, 8, 0, 4}
	fmt.Println(arr.Reverse())
	fmt.Println(arr)
	fmt.Println("Slice")
	slice := arrayJS.JSArray[string]{"ant", "bison", "camel", "duck", "elephant"}
	fmt.Println(slice.Slice(2))
	fmt.Println(slice.Slice(2, 4))
	fmt.Println(slice.Slice(-2))
	fmt.Println(slice.Slice(2, -1))
	fmt.Println(slice.Slice())
}
