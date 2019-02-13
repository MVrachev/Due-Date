package server

import (
	"fmt"
	"time"

	"github.com/end-date/components"
)

// AddTask adds a task element
func (s *Server) AddTask(newTask components.Task) {
	s.db.Create(&newTask)
}

func print(tasks []components.Task) {
	for _, task := range tasks {
		fmt.Println(task)
	}
}

// ListTasksByDueDate lists all tasks sorted by DueDate attribute
func (s *Server) ListTasksByDueDate() {
	var tasks []components.Task
	s.db.Order("date").Find(&tasks)
	print(tasks)
}

// ListTasksByPriority lists all tasks sorted by priority
func (s *Server) ListTasksByPriority() {
	var tasks []components.Task
	s.db.Order("priority").Find(&tasks)
	print(tasks)
}

// FinishTask updates the given task  with status "Done"
func (s *Server) FinishTask(task components.Task) {
	t := components.Task{}
	s.db.Where(&task).First(&t)
	t.Status = "Done"
	s.db.Save(t)
}

// UpdateDueDate updates due date of the given task
func (s *Server) UpdateDueDate(task components.Task, newDueDate time.Time) {
	t := components.Task{}
	s.db.Where(&t).First(&t)
	t.DueDate = newDueDate
	s.db.Save(t)
}

// UpdatePriority updates the priority a given task
func (s *Server) UpdatePriority(task components.Task, newPriority int) {
	t := components.Task{}
	s.db.Where(&task).First(&t)
	t.Priority = newPriority
	s.db.Save(t)
}

// UpdateDescription updates the description of a given task
func (s *Server) UpdateDescription(task components.Task, newDescription string) {
	t := components.Task{}
	s.db.Where(&task).First(&t)
	t.Description = newDescription
	s.db.Save(t)
}
