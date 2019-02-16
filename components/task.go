package components

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// Task is object describing what an element in the table is
// Owner is the owner of the element
// DueDate is the date when the task should be finished
type Task struct {
	gorm.Model
	Owner       string
	DueDate     time.Time
	Priority    int
	Description string
	Status      string
}

// NewTask creates a new task
func NewTask(owner string, dueDate time.Time, priority int, description string) Task {
	return Task{
		Owner:       owner,
		DueDate:     dueDate,
		Priority:    priority,
		Description: description,
		Status:      "In progress",
	}
}

// Implements the String interface for the Elem object
func (t Task) String() string {
	layout := "02 January 2006 15:04:05"

	return fmt.Sprintf("ID: %d,  %v; Due date: %v; 	priority: %v; description: %v; status: %v;",
		t.ID, fmt.Sprintf(t.DueDate.Format(layout)), t.Priority, t.Description, t.Status)
}
