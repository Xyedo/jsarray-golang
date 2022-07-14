# jsarray-golang

My Journey Learning Golang by Creating Port of Array.Prototype method from Javascript

since Javascript array is a dynamic array with references type, we use slice type in golang
slice is a dynamic array with a references types

## Available method
#### Creating JS Array [] -> JSArray{}
```go
import arrayJS "github.com/Xyedo/jsarray-golang/arrayJS"
func main() {
  arr :=arrayJS.JSArray{1,2,3,4} //creating JSarray using array literal
}
```
#### JS Array() -> New(length) JSArray
```go
import arrayJS "github.com/Xyedo/jsarray-golang/arrayJS"
func main() {
  arr :=arrayJS.New(10) //will create JSarray with 10 capacity and 0 length
}
```
#### JS Array.isArray()  -> IsArray(src any)
```go
import arrayJS "github.com/Xyedo/jsarray-golang/arrayJS"
func main() {
  arr :=arrayJS.New(10) //will create JSarray with 10 capacity and 0 length
  arrayJS.IsArray(arr) //will return true
}
```

#### JS Array.prototype.length -> JSArray.Length()
```go
import arrayJS "github.com/Xyedo/jsarray-golang/arrayJS"
func main() {
  arr :=arrayJS.New(10) //will create JSarray with 10 capacity and 0 length
  arr.Length() //will return 0
}
```
#### JS Array.prototype.at() -> JSArray.At(i int) any
```go

func main() {
  arrayJS.JSArray{1,2,3,4}.At(-2)
  //3
}
```
#### JS Array.prototype.concat() -> JSArray.Concat(value ...any) JSArray
```go
import arrayJS "github.com/Xyedo/jsarray-golang/arrayJS"
func main() {
  array1 := arrayJS.JSArray{"a", "b", "c"}
	array2 := arrayJS.JSArray{1, 2, 3}
	fmt.Println(array1.Concat(array2...))
  //[a,b,c,1,2,3]

	num1 := arrayJS.JSArray{1, 2, 3}
	num2 := arrayJS.JSArray{4, 5, 6}
	num3 := arrayJS.JSArray{7, 8, 9}

	fmt.Println(num1.Concat(num2...).Concat(num3...))
  //[1,2,3,4,5,6,7,8,9]
 }
```
#### JS Array.prototype.copyWithin -> JSArray.CopyWithin(target int, index ...int) JSArray
mutates the original array
```go
import arrayJS "github.com/Xyedo/jsarray-golang/arrayJS"
func main() {
  arrayJS.JSArray{1, 2, 3, 4, 5}.CopyWithin(-2)
  //{1,2,3,1,2}
  arrayJS.JSArray{1, 2, 3, 4, 5}.CopyWithin(0,3)
  //{4, 5, 3, 4, 5}
  arrayJS.JSArray{1, 2, 3, 4, 5}.CopyWithin(0, 3, 4)
  //{4, 2, 3, 4, 5}
  arrayJS.JSArray{1, 2, 3, 4, 5}.CopyWithin(-2, -3, -1)
  //{1,2,3,3,4}
}
```
#### JS Array.prototype.every() -> JSArray.Every(callback func(element any, index int, array JSArray) bool) bool
```go
import arrayJS "github.com/Xyedo/jsarray-golang/arrayJS"
func isBigEnough(element any, index int, array arrayJS.JSArray) bool {
  el, ok := element.(int) // you have to typecast it to your array type
  if !ok {
		return false
	}
	return el >= 10
}
func main() {
  arrayJS.JSArray{12, 5, 8, 130, 44}.Every(isBighEnough) //false
  arrayJS.JSArray{12, 54, 18, 130, 44}.Every(isBighEnough) //true
  
}

```
#### JS Array.prototype.fill() -> JSArray.Fill(val any, params ...int) JSArray
params is used for startIndex, and EndIndex
mutate the original array
Fill(val, startIndex =0, endIndex = array.length)
```go
import arrayJS "github.com/Xyedo/jsarray-golang/arrayJS"

func main() {
  arrayJS.JSArray{1,2,3,4}.Fill(0,2,4)
  //[1,2,0,0]
  arrayJS.JSArray{1,2,3,4}.Fill(5,1)
  //[1,5,5,5]
  arrayJS.JSArray{1,2,3,4}.Fill(6)
  //[6,6,6]
  arrayJS.JSArray{1,2,3}.Fill(4, -3, -2)
  //[4,2,3]
  
}
```
#### JS Array.prototype.filter() -> JSArray.Filter(callback func(element any, index int, array JSArray) bool) JSArray

```go
import arrayJS "github.com/Xyedo/jsarray-golang/arrayJS"

func isBigEnough(element any, index int, array arrayJS.JSArray) bool {
  el, ok := element.(int) // you have to typecast it to your array type
  if !ok {
    return false
  }
  return el >= 10
}

func main() {
  arrayJS.JSArray{12, 5, 8, 130, 44}.Filter(isBigEnough)
  //[12,130, 44]
}
```
#### JS Array.prototype.find() -> JSArray.Find(callback func(element any, index int, array JSArray) bool) any
```go
import arrayJS "github.com/Xyedo/jsarray-golang/arrayJS"

func isBigEnough(element any, index int, array arrayJS.JSArray) bool {
  el, ok := element.(int) // you have to typecast it to your array type
  if !ok {
    return false
  }
  return el >= 10
}

func main() {
  arrayJS.JSArray{12, 5, 8, 130, 44}.Find(isBigEnough)
  //12
  arrayJS.JSArray{1, 5, 8, 0, 4}.Find(isBigEnough)
  //nil
}
```
#### JS Array.prototype.findIndex() -> JSArray.FindIndex(callback func(element any, index int, array JSArray) bool) int
```go
import arrayJS "github.com/Xyedo/jsarray-golang/arrayJS"

func isBigEnough(element any, index int, array arrayJS.JSArray) bool {
  el, ok := element.(int) // you have to typecast it to your array type
  if !ok {
    return false
  }
  return el >= 10
}

func main() {
  arrayJS.JSArray{12, 5, 8, 130, 44}.FindIndex(isBigEnough)
  //0
  arrayJS.JSArray{1, 5, 8, 0, 4}.FindIndex(isBigEnough)
  //-1
}
```
#### JS Array.prototype.includes() -> JSArray.Includes(searchedVal any) bool 
```go
import arrayJS "github.com/Xyedo/jsarray-golang/arrayJS"
func main() {
  arrayJS.JSArray{12, 5, 8, 130, 44}.Includes(5)
  //true
  arrayJS.JSArray{1, 5, 8, 0, 4}.Includes(-1)
  //false
}
```
#### JS Array.prototype.indexOf() -> JSArray.IndexOf(searchedVal any) bool 
```go
import arrayJS "github.com/Xyedo/jsarray-golang/arrayJS"
func main() {
  arrayJS.JSArray{12, 5, 8, 130, 44}.IndexOf(12)
  //0
  arrayJS.JSArray{1, 5, 8, 0, 4}.IndexOf(0)
  //3
}
```
#### JS Array.prototype.map() -> JSArray.Map(callback func(element any, index int, array JSArray) any) JSArray
```go
import arrayJS "github.com/Xyedo/jsarray-golang/arrayJS"
func main() {
  //using anonymous function as callback
  arrayJS.JSArray{12, 5, 8, 130, 44}.Map(func(element any, index int, array arrayJS.JSArray) any {
		el, ok := element.(int)
		if !ok {
			return nil
		}
		return el * 2
	})
  //[24,10,16,260,88]
}
```
#### JS Array.prototype.pop() -> JSArray.Pop() JSArray, any
slight bit different with JS pop() method that only return popped value and mutate the array
in Go Pop() return the mutated array and popped value
```go
import arrayJS "github.com/Xyedo/jsarray-golang/arrayJS"
func main() {
  arrayJS.JSArray{12, 5, 8, 130, 44}.Pop()
  //({12,5,8,130}, 44)
```
#### JS Array.prototype.push() -> JSArray.Push(elem ...any) JSArray
in JS Push return void and mutate original array, in Go Push return mutated array
```go
import arrayJS "github.com/Xyedo/jsarray-golang/arrayJS"
func main() {
  arrayJS.JSArray{12, 5, 8, 130, 44}.Push(ArrayJS.JSArray{1,2,3,4}...).Push(55)
  //arr =({12,5,8,130,44,1,2,3,4,55})
```
#### JS Array.prototype.reduce() JSArray.Reduce(callback func(previousValue, curentValue any, currentIndex int, array JSArray) any, initValue any) any 
```go
import arrayJS "github.com/Xyedo/jsarray-golang/arrayJS"
func main() {
  //because using any, we typecast it to our array type
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
```
#### JS Array.prototype.Reverse() -> JSArray.Reverse() JSArray
mutate the original array and return reference to that array
```go
import arrayJS "github.com/Xyedo/jsarray-golang/arrayJS"
func main() {
  //because using any, we typecast it to our array type
  arr := arrayJS.JSArray{0, 1, 2, 3,4}
  mutatedArr :=arr.Reverse()
  fmt.Println(mutatedArr) //[4,3,2,1]
  fmt.Println(arr) /[4,3,2,1]
```
#### JS Array.prototype.shift() -> JSArray.Shift() JSArray, any
in js shift mutated the original array and return the shifted arr; 
in Go Shift returned array and shifted value
```go
import arrayJS "github.com/Xyedo/jsarray-golang/arrayJS"
func main() {
  ex8 := arrayJS.JSArray{1, 2, 3, 4, 5}
  ex8 = ex8.Push(1, 3, 6)
  ex8, shiftedVal := ex8.Shift()
	fmt.Println(ex8, shiftedVal)
  //[3,6], 1
}
```
#### JS Array.prototype.slice() -> JSArray.Slice(params ...int) JSArray
params[0] for start
params[1] for end;
```go
import arrayJS "github.com/Xyedo/jsarray-golang/arrayJS"
func main() {
  slice := arrayJS.JSArray{"ant", "bison", "camel", "duck", "elephant"}
  
	slice.Slice(2)
  //[camel duck elephant]
  slice.Slice(2,4)
  //[camel, duck]
  slice.Slice(-2)
  //[duck elephant]
  slice.Slice(2, -1)
  //[camel duck]
  slice.Slice()
  //[ant bison camel duck elephant]
}
```
#### JS Array.prototype.some() -> JSArray.Some(callback func(element any, index int, array JSArray) bool) bool
```go
import arrayJS "github.com/Xyedo/jsarray-golang/arrayJS"

func isBigEnough(element any, index int, array arrayJS.JSArray) bool {
  el, ok := element.(int) // you have to typecast it to your array type
  if !ok {
    return false
  }
  return el >= 10
}

func main() {
	fmt.Println(arrayJS.JSArray{2, 5, 8, 1, 4}.Some(isBigEnough)) //false
	fmt.Println(arrayJS.JSArray{12, 5, 8, 1, 4}.Some(isBigEnough)) //true
}
```
#### JS Array.prototype.splice() -> JSArray.Splice(pointIndex int, removeCount int, element ...any) (JSArray, JSArray)
```go
import arrayJS "github.com/Xyedo/jsarray-golang/arrayJS"
func main() {
myFish := arrayJS.JSArray{"angel", "clown", "mandarin", "sturgeon"}
	ex1, removed1 := myFish.Splice(2, 0, "drum")
	fmt.Println("ex1:", ex1, removed1) //[angel clown drum mandarin sturgeon] ,[]
	forEx2 := arrayJS.JSArray{"angel", "clown", "mandarin", "sturgeon"}
	ex2, removed2 := forEx2.Splice(2, 0, "drum", "guitar") 
	fmt.Println("ex2:", ex2, removed2)//[angel clown drum guitar mandarin sturgeon] ,[]
	forEx3 := arrayJS.JSArray{"angel", "clown", "drum", "mandarin", "sturgeon"}
	ex3, removed3 := forEx3.Splice(3, 1) 
	fmt.Println("ex3:", ex3, removed3) //[angel clown drum sturgeon] [mandarin]
	forEx4 := arrayJS.JSArray{"angel", "clown", "drum", "sturgeon"}
	ex4, removed4 := forEx4.Splice(2, 1, "trumpet")
	fmt.Println("ex4:", ex4, removed4) //[angel clown trumpet sturgeon] [drum]
	forEx5 := arrayJS.JSArray{"angel", "clown", "drum", "sturgeon"}
	ex5, removed5 := forEx5.Splice(0, 2, "parrot", "anemone", "blue")
	fmt.Println("ex5:", ex5, removed5) //[parrot anemone blue angel clown drum sturgeon] [angel clown]
	forEx6 := arrayJS.JSArray{"parrot", "anemone", "blue", "trumpet", "sturgeon"}
	ex6, removed6 := forEx6.Splice(2, 2) //[parrot anemone sturgeon] [blue trumpet]
}

```
#### JS Array.prototype.unshift() -> JSArray.Unshift(element ...any) JSArray
```go
import arrayJS "github.com/Xyedo/jsarray-golang/arrayJS"
func main() {
arrayJS.JSArray{1, 2, 3, 4, 5}.Unshift(10,11) //[1,2,3,4,5,10,11]
}
```











