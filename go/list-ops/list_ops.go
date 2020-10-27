// Package listops implements basic list operations
package listops

// IntList represents an integer slice
type IntList []int

// Append adds all items in the second list to the end of the first list
func (first IntList) Append(second IntList) (combinedList IntList) {
	combinedList = make(IntList, first.Length()+second.Length())
	for i, v := range first {
		combinedList[i] = v
	}
	for i, v := range second {
		index := i + len(first)
		combinedList[index] = v
	}
	return
}

// Concat combines all items in all lists into one flattened list
func (first IntList) Concat(lists []IntList) (combinedList IntList) {
	if first.Length() == 0 {
		return IntList{}
	}
	combinedList = combinedList.Append(first)
	for _, list := range lists {
		combinedList = combinedList.Append(list)
	}
	return
}

type predFunc func(n int) bool

// Filter returns a list of all items in the list for which the predicate is true
func (first IntList) Filter(f predFunc) (filteredList IntList) {
	if first.Length() == 0 {
		return IntList{}
	}
	for _, v := range first {
		if f(v) {
			filteredList = filteredList.Append(IntList{v})
		}
	}

	return
}

// Length returns the total number of items within the list
func (first IntList) Length() (size int) {
	for range first {
		size++
	}
	return
}

type unaryFunc func(x int) int

// Map returns a list of results after applying the function to all items in the list
func (first IntList) Map(u unaryFunc) (mapList IntList) {
	if first.Length() == 0 {
		return IntList{}
	}
	for _, v := range first {
		mapList = mapList.Append(IntList{u(v)})
	}
	return
}

// Reverse returns a list with all the original items, but in reverse order
func (first IntList) Reverse() (reversedList IntList) {
	length := first.Length()
	reversedList = make(IntList, length)
	for i, v := range first {
		reversedList[length-i-1] = v
	}
	return
}

type binFunc func(x, y int) int

// Foldl folds/reduces each item into the accumulator, starting from the left
func (first IntList) Foldl(fn binFunc, initial int) int {
	for _, v := range first {
		if initial == 0 {
			break
		}
		initial = fn(v, initial)
	}
	return initial
}

// Foldr folds/reduces each item into the accumulator, starting from the right
func (first IntList) Foldr(fn binFunc, initial int) int {
	index := first.Length() - 1
	for range first {
		if initial == 0 {
			break
		}
		initial = fn(first[index], initial)
		index--
	}
	return initial
}
