package components

import (
	"fmt"
	"time"
)

type Status bool

const (
	IN_PROGRESS Status = false
	DONE        Status = true
)

//status      int // Can be enums ??

// Elem is object describing what an element in the table is
type Elem struct {
	date        time.Time
	priority    int
	description string
	status      Status
}

// NewElem Creates a new Element
func NewElem(date time.Time, priority int, description string) Elem {
	return Elem{
		date:        date,
		priority:    priority,
		description: description,
		status:      IN_PROGRESS,
	}
}

func (status Status) String() string {
	if status == IN_PROGRESS {
		return "IN PROGRESS"
	} else {
		return "DONE"
	}
}

// Implements the String interface for the Elem object
func (e Elem) String() string {
	layout := "02 January	2006 15:04:05"

	return fmt.Sprintf("Date: %v; 	priority: %v; description: %v; status: %v;",
		fmt.Sprintf(e.date.Format(layout)), e.priority, e.description, e.status.String())
}

// Equal Checks if the values of the right element are equal of those of the left element
func Equal(right, left Elem) bool {
	if right.date.Equal(left.date) && right.description == left.description &&
		right.priority == left.priority && right.status == left.status {
		return true
	}
	return false
}

// Equal Checks if the values of the right element are equal of those of the left element
func Equal(right, left Elem) bool {
	if right.date.Equal(left.date) && right.description == left.description &&
		right.priority == left.priority && right.status == left.status {
		return true
	}
	return false
}
