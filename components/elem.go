package components

import (
	"fmt"
	"time"
)

// Elem is object describing what an element in the table is
type Elem struct {
	date        time.Time
	priority    int
	description string
	status      int // Can be enums ??
}

// NewElem Creates a new Element
func NewElem(date time.Time, priority int, description string, status int) Elem {
	return Elem{
		date:        date,
		priority:    priority,
		description: description,
		status:      status,
	}
}

// Implements the String interface for the Elem object
func (e Elem) String() string {
	layout := "02 January	2006 15:04:05"
	return fmt.Sprintf("Date: %v; 	priority: %v; description: %v; status: %v;",
		fmt.Sprintf(e.date.Format(layout)), e.priority, e.description, e.status)
}

// Equal Checks if the values of the right element are equal of those of the left element
func Equal(right, left Elem) bool {
	if right.date.Equal(left.date) && right.description == left.description &&
		right.priority == left.priority && right.status == left.status {
		return true
	}
	return false
}
