package arrayJS

import (
	"errors"
	"fmt"
	"reflect"
)

// import "fmt"

// func main() {
// 	testSlice := []int{1, 2, 3, 4}

// }
func New[T comparable](values ...T) JSArray[T] {

	jsarry := make(JSArray[T], len(values))
	copy(jsarry, values)
	return jsarry
}

func IsArray(src any) bool {
	return reflect.Slice == reflect.TypeOf(src).Kind()
}

type JSArray[T comparable] []T

func (s JSArray[T]) Length() int {
	return len(s)
}
func (s JSArray[T]) At(i int) any {
	for i < 0 {
		i = s.Length() + i
	}
	return s[i]
}
func (s JSArray[T]) CopyWithin(target int, index ...int) JSArray[T] {
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
	scopy := make(JSArray[T], s.Length())
	copy(scopy, s)
	for i, t := start, target; i < end; i, t = i+1, t+1 {
		if t >= s.Length() {
			break
		}

		s[t] = scopy[i]

	}
	return s
}

func (s JSArray[T]) Every(cb func(element T, index int, array JSArray[T]) bool) bool {

	len := s.Length()
	for i := 0; i < len; i++ {
		val := s[i]
		if !cb(val, i, s) {
			return false
		}
	}
	return true

}
func (s JSArray[T]) FindIndex(cb func(element T, index int, array JSArray[T]) bool) int {

	len := s.Length()
	for i := 0; i < len; i++ {
		val := s[i]
		if cb(val, i, s) {
			return i
		}
	}
	return -1
}

var errNotFound = errors.New("notFound")

func (s JSArray[T]) Find(cb func(element T, index int, array JSArray[T]) bool) (T, error) {
	index := s.FindIndex(cb)
	var defVal T
	if index == -1 {
		return defVal, errNotFound
	}
	return s[index], nil
}
func (s JSArray[T]) Concat(value ...T) JSArray[T] {
	len := s.Length()
	res := make(JSArray[T], len)
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
func (s JSArray[T]) Reduce(callback func(previousValue, curentValue T, currentIndex int, array JSArray[T]) T, initValue ...T) T {
	arrlength := s.Length()
	var acc T

	startIndex := 0

	if len(initValue) == 0 {
		acc = s[0]
		startIndex = 1
	} else {
		acc = initValue[0]
	}
	for i := startIndex; i < arrlength; i++ {
		val := s[i]
		acc = callback(acc, val, i, s)
	}
	return acc
}
func (s JSArray[T]) Fill(val T, params ...int) JSArray[T] {
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
func (s JSArray[T]) Reverse() JSArray[T] {
	res := make(JSArray[T], s.Length())
	lastIndex := s.Length() - 1
	for i := lastIndex; i >= 0; i-- {
		val := s[i]
		res[lastIndex-i] = val
	}
	copy(s, res)
	return res
}
func (s JSArray[T]) Slice(params ...int) JSArray[T] {
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
	res := JSArray[T]{}
	if start >= s.Length() {
		return res
	}

	for i := start; i < end; i++ {
		if i < s.Length() {
			res = res.Push(s[i])
		}
	}
	return res
}
func (s JSArray[T]) Some(cb func(element T, index int, array JSArray[T]) bool) bool {

	len := s.Length()
	for i := 0; i < len; i++ {
		if cb(s[i], i, s) {
			return true
		}
	}
	return false
}
func (s JSArray[T]) Filter(cb func(element T, index int, array JSArray[T]) bool) JSArray[T] {

	res := JSArray[T]{}
	len := s.Length()
	for i := 0; i < len; i++ {
		val := s[i]
		if cb(val, i, s) {
			res = res.Push(val)
		}
	}
	return res
}
func (s JSArray[T]) Pop() (jsArr JSArray[T], poppedVal T) {
	return s[:s.Length()-1], s[s.Length()-1]
}

func (s JSArray[T]) Push(elem ...T) JSArray[T] {
	return append(s, elem...)
}

func (s JSArray[T]) Shift() (jsArr JSArray[T], shiftedVal T) {
	return s[1:], s[0]
}

func (s JSArray[T]) Unshift(element ...T) JSArray[T] {
	return append(element, s...)
}

func (slice JSArray[T]) Splice(pointIndex int, removeCount int, element ...T) (jsArr JSArray[T], splicedArr JSArray[T]) {

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
			return slice.Push(element...), splicedArr
		}
		return slice[:slice.Length()-removeCount].Push(element...), slice[slice.Length()-removeCount:]

	}
	if pointIndex == 0 {
		if removeCount == 0 {
			return slice.Unshift(element...), splicedArr
		}
		return slice.Unshift(element...), slice[:removeCount]
	}
	if removeCount == 0 {
		end := JSArray[T](element).Push(slice[pointIndex:]...)
		return slice[:pointIndex].Push(end...), splicedArr
	}

	deletedEnd := pointIndex + removeCount
	if deletedEnd > slice.Length() {
		deletedEnd = slice.Length()
	}
	deleted := make(JSArray[T], removeCount)
	copy(deleted, slice[pointIndex:deletedEnd])
	undeleted := slice[:pointIndex].Push(slice[pointIndex+removeCount:]...)
	end := JSArray[T](element).Push(undeleted[pointIndex:]...)
	return undeleted[:pointIndex].Push(end...), deleted

}

func (s JSArray[T]) Includes(searchedVal T) bool {
	cb := func(element T, index int, array JSArray[T]) bool {
		return element == searchedVal
	}
	return s.Some(cb)
}
func (s JSArray[T]) IndexOf(searchedVal any) int {
	return s.FindIndex(func(element T, index int, array JSArray[T]) bool {
		return element == searchedVal
	})
}
func (s JSArray[T]) Map(callback func(element T, index int, array JSArray[T]) T) JSArray[T] {
	len := s.Length()
	res := make(JSArray[T], len)

	for i := 0; i < len; i++ {
		val := s[i]
		res[i] = callback(val, i, s)
	}
	return res
}
