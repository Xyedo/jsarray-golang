package arrayJS

import (
	"fmt"
	"reflect"
)

// import "fmt"

// func main() {
// 	testSlice := []int{1, 2, 3, 4}

// }
func New(length int) JSArray {
	return make(JSArray, length)
}

func IsArray(src any) bool {
	return reflect.Slice == reflect.TypeOf(src).Kind()
}

type JSArray []any

func (s JSArray) Length() int {
	return len(s)
}
func (s JSArray) At(i int) any {
	for i < 0 {
		i = s.Length() + i
	}
	return s[i]
}
func (s JSArray) CopyWithin(target int, index ...int) JSArray {
	for target < 0 {
		target = s.Length() + target
	}
	if target >= s.Length() {
		return s
	}
	start := 0
	end := s.Length()
	switch len(index) {
	case 1:
		start = index[0]
	case 2:
		start = index[0]
		end = index[1]
	}
	for start < 0 {
		start = s.Length() + start
	}

	for end < 0 {
		end = s.Length() + end
	}
	if start == end {
		return s
	}

	fmt.Println(target, start, end)
	scopy := New(s.Length())
	copy(scopy, s)
	for i, t := start, target; i < end; i, t = i+1, t+1 {
		if t >= s.Length() {
			break
		}

		s[t] = scopy[i]

	}
	return s
}

func (s JSArray) Every(callback func(element any, index int, array JSArray) bool) bool {
	len := s.Length()
	for i := 0; i < len; i++ {
		val := s[i]
		if !callback(val, i, s) {
			return false
		}
	}
	return true

}
func (s JSArray) FindIndex(callback func(element any, index int, array JSArray) bool) int {
	len := s.Length()
	for i := 0; i < len; i++ {
		val := s[i]
		if callback(val, i, s) {
			return i
		}
	}
	return -1
}

func (s JSArray) Find(callback func(element any, index int, array JSArray) bool) any {
	index := s.FindIndex(callback)
	if index == -1 {
		return nil
	}
	return s[index]
}
func (s JSArray) Concat(value ...any) JSArray {
	len := s.Length()
	res := New(len)
	copy(res, s)
	res = res.Push(value...)
	return res
}

// func (s JSArray) Flat(depth int) JSArray {
// 	if depth < 1 || !IsArray(s) {
// 		return s
// 	}
// 	return s.Reduce(func(previousValue, curentValue any, currentIndex int, array JSArray) any {
// 		res := New(10)
// 		res.Push(previousValue)

// 		return res.Concat(curentValue)
// 	})

// }
func (s JSArray) Reduce(callback func(previousValue, curentValue any, currentIndex int, array JSArray) any, initValue any) any {
	len := s.Length()
	acc := initValue

	startIndex := 0

	if initValue == nil {
		acc = s[0]
		startIndex = 1
	}
	for i := startIndex; i < len; i++ {
		val := s[i]
		acc = callback(acc, val, i, s)
	}
	return acc
}
func (s JSArray) Fill(val any, params ...int) JSArray {
	start := 0
	end := s.Length()
	switch len(params) {
	case 1:
		start = params[0]
	case 2:
		start = params[0]
		end = params[1]
	}
	for start < 0 {
		start = s.Length() + start
	}
	for end < 0 {
		end = s.Length() + end
	}
	for i := start; i < end; i++ {
		s[i] = val
	}
	return s
}
func (s JSArray) Reverse() JSArray {
	res := New(s.Length())
	lastIndex := s.Length() - 1
	for i := lastIndex; i >= 0; i-- {
		val := s[i]
		res[lastIndex-i] = val
	}
	return res
}
func (s JSArray) Slice(params ...int) JSArray {
	start := 0
	end := s.Length()
	switch len(params) {
	case 1:
		start = params[0]
	case 2:
		start = params[0]
		end = params[1]
	}
	res := New(end - start)
	for i := start; i < end; i++ {
		if i < s.Length() {
			res.Push(s[i])
		}
	}
	return res
}
func (s JSArray) Some(callback func(element any, index int, array JSArray) bool) bool {
	len := s.Length()
	for i := 0; i < len; i++ {
		if callback(s[i], i, s) {
			return true
		}
	}
	return false
}
func (s JSArray) Filter(callback func(element any, index int, array JSArray) bool) JSArray {
	res := JSArray{}
	len := s.Length()
	for i := 0; i < len; i++ {
		val := s[i]
		if callback(val, i, s) {
			res = res.Push(val)
		}
	}
	return res
}
func (s JSArray) Pop() (JSArray, any) {
	return s[:s.Length()-1], s[s.Length()-1]
}

func (s JSArray) Push(elem ...any) JSArray {
	return append(s, elem...)
}

func (s JSArray) Shift() (JSArray, any) {
	return s[1:], s[0]
}

func (s JSArray) Unshift(element ...any) JSArray {
	return append(element, s...)
}

func (slice JSArray) Splice(pointIndex int, removeCount int, element ...any) (JSArray, JSArray) {

	for pointIndex < 0 {
		pointIndex = slice.Length() - 1 + pointIndex
	}
	if pointIndex > slice.Length()-1 {
		pointIndex = slice.Length() - 1
	}
	if removeCount > pointIndex+slice.Length() {
		removeCount = slice.Length()
	}
	if pointIndex == slice.Length()-1 {
		if removeCount == 0 {
			return slice.Push(element...), []any{}
		}
		return slice[:slice.Length()-removeCount].Push(element...), slice[slice.Length()-removeCount:]

	}
	if pointIndex == 0 {
		if removeCount == 0 {
			return slice.Unshift(element...), []any{}
		}
		return slice.Unshift(element...), slice[:removeCount]
	}
	if removeCount == 0 {
		end := JSArray(element).Push(slice[pointIndex:]...)
		return slice[:pointIndex].Push(end...), []any{}
	}

	deletedEnd := pointIndex + removeCount
	if deletedEnd > slice.Length() {
		deletedEnd = slice.Length()
	}
	deleted := New(removeCount)
	copy(deleted, slice[pointIndex:deletedEnd])
	undeleted := slice[:pointIndex].Push(slice[pointIndex+removeCount:]...)
	end := JSArray(element).Push(undeleted[pointIndex:]...)
	return undeleted[:pointIndex].Push(end...), deleted

}

func (s JSArray) Includes(searchedVal any) bool {
	cb := func(element any, index int, array JSArray) bool {
		return element == searchedVal
	}
	return s.Some(cb)
}
func (s JSArray) IndexOf(searchedVal any) int {
	return s.FindIndex(func(element any, index int, array JSArray) bool {
		return element == searchedVal
	})
}
func (s JSArray) Map(callback func(element any, index int, array JSArray) any) JSArray {
	len := s.Length()
	res := New(len)

	for i := 0; i < len; i++ {
		val := s[i]
		res[i] = callback(val, i, s)
	}
	return res

}
