package components

// ListElementsByPriority changes how the user will see the elements.
// From the moment this function is called
// the user will see the elements sorted by priority
func (e *ElemsTable) ListElementsByPriority() {
	e.sortByPriority = true
	e.sortElements()
}

// ListElementsBydate changes how the user will see the elements.
// From the moment this function is called
// the user will see the elements sorted by date
func (e *ElemsTable) ListElementsBydate() {
	e.sortByPriority = false
	e.sortElements()
}

// AddElemIntoTheTable adds a new Element to the existing table.
// It finds where the element  should be inserted and keeps the elements
// sorted by date/priority
func (e *ElemsTable) AddElemIntoTheTable(newElem Elem) {
	insertAtIndex := e.searchInsertIndex(newElem)

	resultSlice := make([]Elem, insertAtIndex)
	copy(resultSlice, e.elements[:insertAtIndex])
	resultSlice = append(resultSlice, newElem)
	resultSlice = append(resultSlice, e.elements[insertAtIndex:]...)
	e.elements = resultSlice
}

// DeleteElemFromTheTable deletes the given element from the
// element table. Returns true when succeeding, false otherwise.
func (e *ElemsTable) DeleteElemFromTheTable(elemForDeletion Elem) bool {
	if len(e.elements) == 0 {
		return false
	}

	deleteAtIndex := e.findElemIndex(elemForDeletion)
	if deleteAtIndex == -1 {
		return false
	}
	e.elements = append(e.elements[:deleteAtIndex], e.elements[deleteAtIndex+1:]...)
	return true
}
