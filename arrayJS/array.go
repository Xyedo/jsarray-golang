package arrayJS

// import "fmt"

// func main() {
// 	testSlice := []int{1, 2, 3, 4}

// }
type JSArray []any

func New(size int) JSArray {
	return make(JSArray, size)
}
func (s JSArray) Pop() (JSArray, any) {
	return s[:len(s)-1], s[len(s)-1]
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
		pointIndex = len(slice) - 1 + pointIndex
	}
	if pointIndex > len(slice)-1 {
		pointIndex = len(slice) - 1
	}
	if removeCount > pointIndex+len(slice) {
		removeCount = len(slice)
	}
	if pointIndex == len(slice)-1 {
		if removeCount == 0 {
			return slice.Push(element...), []any{}
		}
		return slice[:len(slice)-removeCount].Push(element...), slice[len(slice)-removeCount:]

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
	if deletedEnd > len(slice) {
		deletedEnd = len(slice)
	}
	deleted := make([]any, removeCount)
	copy(deleted, slice[pointIndex:deletedEnd])
	undeleted := slice[:pointIndex].Push(slice[pointIndex+removeCount:]...)
	end := JSArray(element).Push(undeleted[pointIndex:]...)
	return undeleted[:pointIndex].Push(end...), deleted

}
