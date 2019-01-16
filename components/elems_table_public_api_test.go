package components

import (
	"testing"
	"time"
)

func myInit() *ElemsTable {
	elements := make([]Elem, 0)

	t := time.Date(2018, 11, 24, 13, 35, 21, 0, time.UTC)
	elem := NewElem(t, 4, "First Task!")

	elements = append(elements, elem)

	secT := time.Date(2019, 05, 24, 13, 35, 21, 0, time.UTC)
	secElem := NewElem(secT, 3, "Second Task")

	elements = append(elements, secElem)

	thirdT := time.Date(2020, 12, 24, 13, 35, 21, 0, time.UTC)
	thirdElem := NewElem(thirdT, 2, "Third elem")

	elements = append(elements, thirdElem)

	return NewElemsTable(elements)
}

func TestNewElemInEmptyElemsTable(t *testing.T) {
	elementsTable := NewElemsTable(make([]Elem, 0))
	timeElem := time.Date(2018, 12, 24, 13, 35, 21, 0, time.UTC)
	elem := NewElem(timeElem, 0, "First Task!")
	elementsTable.AddElemIntoTheTable(elem)

	if len(elementsTable.elements) != 1 {
		t.Errorf("\nExpected len of the elements table to be 1 but it was: %#v.", len(elementsTable.elements))
	}
	if !timeElem.Equal(elementsTable.elements[0].date) {
		t.Errorf("\nExpected element with index 0 to be: %+v\n but it was: %+v\n.", elem, elementsTable.elements[0])
	}
}

func TestAddInBegining(t *testing.T) {
	elementsTable := myInit()
	timeElem := time.Date(2000, 12, 24, 13, 35, 21, 0, time.UTC)
	newElem := NewElem(timeElem, 0, "Zero Task!")

	elementsTable.AddElemIntoTheTable(newElem)

	if len(elementsTable.elements) != 4 {
		t.Errorf("\nExpected len of the elements table to be 1 but it was: %#v.", len(elementsTable.elements))
	}
	if !timeElem.Equal(elementsTable.elements[0].date) {
		t.Errorf("\nExpected element with index 0 to be: %+v\n but it was: %+v\n.", newElem, elementsTable.elements[0])
	}
}

func TestAddInMedium(t *testing.T) {
	elementsTable := myInit()
	timeElem := time.Date(2019, 12, 24, 13, 35, 21, 0, time.UTC)
	newElem := NewElem(timeElem, 0, "Medium Task!")

	elementsTable.AddElemIntoTheTable(newElem)
	if len(elementsTable.elements) != 4 {
		t.Errorf("\nExpected len of the elements table to be 1 but it was: %#v.", len(elementsTable.elements))
	}
	if !timeElem.Equal(elementsTable.elements[2].date) {
		t.Errorf("\nExpected element with index 2 to be: %+v\n but it was: %+v\n.", newElem, elementsTable.elements[2])
	}
}

func TestAddInEnd(t *testing.T) {
	elementsTable := myInit()
	timeElem := time.Date(2021, 12, 24, 13, 35, 21, 0, time.UTC)
	newElem := NewElem(timeElem, 0, "End Task!")

	elementsTable.AddElemIntoTheTable(newElem)

	if len(elementsTable.elements) != 4 {
		t.Errorf("\nExpected len of the elements table to be 1 but it was: %#v.", len(elementsTable.elements))
	}
	if !timeElem.Equal(elementsTable.elements[3].date) {
		t.Errorf("\nExpected element with index 3 to be: %+v\n but it was: %+v\n.", newElem, elementsTable.elements[3])
	}
}

func TestListElementsByPriority(t *testing.T) {
	elementsTable := myInit()
	elementsTable.ListElementsByPriority()
	len := len(elementsTable.elements)
	for i := 0; i < len-1; i++ {
		if elementsTable.elements[i].priority > elementsTable.elements[i+1].priority {
			t.Errorf("\nExpected element with index %#v to have lower priority than element with index %#v", i, i+1)
		}
	}
}

func TestListElementsByDate(t *testing.T) {
	elementsTable := myInit()
	elementsTable.ListElementsByPriority()

	elementsTable.ListElementsBydate()
	len := len(elementsTable.elements)
	for i := 0; i < len-1; i++ {
		if !elementsTable.less(elementsTable.elements[i], elementsTable.elements[i+1]) {
			t.Errorf("\nExpected element with index %#v to be before than element with index %#v", i, i+1)
		}
	}
}

func TestDeleteFromEmptyTable(t *testing.T) {
	elementsTable := NewElemsTable(make([]Elem, 0))
	timeElem := time.Date(2021, 12, 24, 13, 35, 21, 0, time.UTC)
	delElement := NewElem(timeElem, 0, "End Task!")
	result := elementsTable.DeleteElemFromTheTable(delElement)

	if result || len(elementsTable.elements) != 0 {
		t.Errorf("\nExpected the deletion from an empty elements table to fail!")
	}
}

func TestDeleteNonExistingElement(t *testing.T) {
	elementsTable := myInit()
	timeElem := time.Date(2111, 12, 24, 13, 35, 21, 0, time.UTC)
	delElement := NewElem(timeElem, 0, "End Task!")

	result := elementsTable.DeleteElemFromTheTable(delElement)

	if result || len(elementsTable.elements) != 3 {
		t.Errorf("\nExpected the deletion of an non existing element in the elements table to fail!")
	}
}

func TestDeleteExistingElemFromFullTable(t *testing.T) {
	elementsTable := myInit()
	delTime := time.Date(2019, 05, 24, 13, 35, 21, 0, time.UTC)
	delElement := NewElem(delTime, 3, "Second Task")

	result := elementsTable.DeleteElemFromTheTable(delElement)

	if !result || len(elementsTable.elements) != 2 {
		t.Errorf("\nExpected the deletion of an existing element in the full elements table to succeed!")
	}
}
