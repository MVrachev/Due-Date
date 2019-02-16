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
	fmt.Println("will add")
	var info components.Information
	if err := conn.ReadJSON(&info); err != nil {
		panic(err)
	}

	fmt.Println("Read JSON")

	timeElem := time.Date(convertToInt(info.Year), time.Month(convertToInt(info.Month)), convertToInt(info.Day), 0, 0, 0, 0, time.UTC)
	newTask := components.NewTask(owner, timeElem, convertToInt(info.Priority), info.Description)
	fmt.Println(newTask)

	s.mutex.Lock()
	s.db.Create(&newTask)
	s.mutex.Unlock()
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

	s.mutex.Lock()
	s.db.Save(t)
	s.mutex.Unlock()
}

// ------------------------------------- Updates -------------------------------------

// UpdateDueDate updates due date of the given task
func (s *Server) UpdateDueDate(task components.Task, newDueDate time.Time) {
	t := components.Task{}
	s.db.Where(&t).First(&t)
	t.DueDate = newDueDate

	s.mutex.Lock()
	s.db.Save(t)
	s.mutex.Unlock()
}

// UpdatePriority updates the priority a given task
func (s *Server) UpdatePriority(task components.Task, newPriority int) {
	t := components.Task{}
	s.db.Where(&task).First(&t)
	t.Priority = newPriority

	s.mutex.Lock()
	s.db.Save(t)
	s.mutex.Unlock()

}

// UpdateDescription updates the description of a given task
func (s *Server) UpdateDescription(task components.Task, newDescription string) {
	t := components.Task{}
	s.db.Where(&task).First(&t)
	t.Description = newDescription

	s.mutex.Lock()
	s.db.Save(t)
	s.mutex.Unlock()
}

// Delete deletes a given task
func (s *Server) Delete(task components.Task) {
	s.mutex.Lock()
	s.db.Where(&task).Delete(&task)
	s.mutex.Unlock()
}
