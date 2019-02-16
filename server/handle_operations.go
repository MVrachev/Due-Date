package server

import (
	"fmt"
	"strconv"
	"time"

	"github.com/end-date/components"
	"github.com/gorilla/websocket"
)

// ------------------------------------- Add -------------------------------------

func convertToInt(str string) int {
	res, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return res
}

// AddTask adds a task element
func (s *Server) AddTask(conn *websocket.Conn, owner string) {
	var info components.Information
	if err := conn.ReadJSON(&info); err != nil {
		panic(err)
	}

	timeElem := time.Date(convertToInt(info.Year), time.Month(convertToInt(info.Month)), convertToInt(info.Day), 0, 0, 0, 0, time.UTC)
	newTask := components.NewTask(owner, timeElem, convertToInt(info.Priority), info.Description)
	s.db.Debug().Create(&newTask)
}

// ------------------------------------- Lists -------------------------------------

func print(tasks []components.Task) {
	for _, task := range tasks {
		fmt.Println(task)
	}
}

// ListTasksByDueDate lists all tasks sorted by DueDate attribute
func (s *Server) ListTasksByDueDate(userName string) {
	var tasks []components.Task
	s.db.Where("name = ?", userName).Order("date").Find(&tasks)
	print(tasks)
}

// ListTasksByPriority lists all tasks sorted by priority
func (s *Server) ListTasksByPriority(userName string) {
	var tasks []components.Task
	s.db.Where("name = ?", userName).Order("priority").Find(&tasks)
	print(tasks)
}

// ------------------------------------- Finish -------------------------------------

// FinishTask updates the given task  with status "Done"
func (s *Server) FinishTask(task components.Task) {
	t := components.Task{}
	s.db.Where(&task).First(&t)
	t.Status = "Done"
	s.db.Save(t)
}

// ------------------------------------- Updates -------------------------------------

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

// Delete deletes a given task
func (s *Server) Delete(task components.Task) {
	s.db.Where(&task).Delete(&task)
}
