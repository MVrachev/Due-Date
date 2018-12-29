package components

// ListElementsByPriority changes how the user will see the elements.
// From the moment this function is called
// the user will see the elements sorted by priority
func (elementsTable ElemsTable) ListElementsByPriority() {
	elementsTable.sortByPriority = true
	elementsTable.sortElements()
}

// ListElementsBydate changes how the user will see the elements.
// From the moment this function is called
// the user will see the elements sorted by date
func (elementsTable ElemsTable) ListElementsBydate() {
	elementsTable.sortByPriority = false
	elementsTable.sortElements()
}

// AddElemIntoTheTable adds a new Element to the existing table.
// It finds where the element should be inserted and keeps the elements
// sorted by date/priority
func (elementsTable *ElemsTable) AddElemIntoTheTable(newElem Elem) {
	insertAfterIndex := elementsTable.searchIndex(newElem)

	resultSlice := make([]Elem, insertAfterIndex)
	copy(resultSlice, elementsTable.elements[:insertAfterIndex])
	resultSlice = append(resultSlice, newElem)
	resultSlice = append(resultSlice, elementsTable.elements[insertAfterIndex:]...)
	elementsTable.elements = resultSlice
}
