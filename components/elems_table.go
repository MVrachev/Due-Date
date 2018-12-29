package components

import (
	"fmt"
	"sort"
)

// ElemsTable is a type managing the elements in the table
// and their representation
type ElemsTable struct {
	elements       []Elem
	sortByPriority bool
}

// NewElemsTable Creates a new Element table object
func NewElemsTable(elements []Elem) *ElemsTable {
	return &ElemsTable{
		elements: elements,
	}
}

// Implements the String interface for the Elem object
func (e *ElemsTable) String() string {
	var str string
	for index, elem := range e.elements {
		str += fmt.Sprintf("%v: %v\n", index, elem)
	}
	return str
}

// less reports whether the right element should sort before the left element.
func (e *ElemsTable) less(rightElem, leftElem Elem) bool {
	if !e.sortByPriority {
		return leftElem.date.After(rightElem.date) || leftElem.date.Equal(rightElem.date)
	}
	return rightElem.priority <= leftElem.priority
}

// Sorts the elements by date/priority
func (e *ElemsTable) sortElements() {
	sort.Sort(e)
}

/*
 Gives an index where the new element should be in the new slice
* Returns 0 if the element should be inserted in the begining
*/
func (e *ElemsTable) searchIndex(elem Elem) int {
	elements := e.elements
	len := len(elements)

	if len == 0 || e.less(elem, elements[0]) {
		// This means that the element should be placed in front of all elements
		return 0
	}
	for i := 0; i < len-1; i++ {
		if e.less(elements[i], elem) && e.less(elem, elements[i+1]) {
			return i + 1
		}
	}
	// If function did not returned a result until here it means that the element
	// is larger than all others and it must be placed in the end
	return len
}

// ------------------------------ Sort interface implementation ------------------------------

// For reference: https://golang.org/pkg/sort/#Interface

// Len is the number of elements in the collection.
func (e *ElemsTable) Len() int {
	return len(e.elements)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (e *ElemsTable) Less(i, j int) bool {
	return e.less(e.elements[i], e.elements[j])
}

// Swap swaps the elements with indexes i and j.
func (e *ElemsTable) Swap(i, j int) {
	temp := e.elements[i]
	e.elements[i] = e.elements[j]
	e.elements[j] = temp
}

// ------------------------------------------------------------------------------------------
