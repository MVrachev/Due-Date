package server

import (
	"errors"
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

func sendInfo(conn *websocket.Conn, tasks []components.Task) {
	info := components.InfoForTasks{
		InfoTasks: tasks,
	}
	if err := conn.WriteJSON(&info); err != nil {
		panic(err)
	}
}

// ListTasksByDueDate lists all tasks sorted by DueDate attribute
func (s *Server) ListTasksByDueDate(conn *websocket.Conn, userName string) {
	var tasks []components.Task
	s.db.Where("owner = ?", userName).Order("due_date").Find(&tasks)
	sendInfo(conn, tasks)
}

// ListTasksByPriority lists all tasks sorted by priority
func (s *Server) ListTasksByPriority(conn *websocket.Conn, userName string) {
	var tasks []components.Task
	s.db.Where("owner = ?", userName).Order("priority").Find(&tasks)
	sendInfo(conn, tasks)
}

// ------------------------------------- Finish -------------------------------------

func (s *Server) getID(conn *websocket.Conn) int {
	if err := conn.WriteMessage(websocket.TextMessage,
		[]byte("Task with which ID you want to modfiy?")); err != nil {
		panic(err)
	}

	_, bufID, err := conn.ReadMessage()
	if err != nil {
		panic(err)
	}
	ID, err := strconv.Atoi(string(bufID))
	if err != nil {
		panic(err)
	}
	return ID
}

func (s *Server) getCertainTask(conn *websocket.Conn, owner string) (*components.Task, error) {
	t := components.Task{}
	id := s.getID(conn)
	s.db.Where("owner = ? AND id = ?", owner, id).First(&t)
	if t.Owner == "" {
		if err := conn.WriteMessage(websocket.TextMessage,
			[]byte("Not right ID!")); err != nil {
			panic(err)
		}
		return nil, errors.New("Not the right ID!")
	}
	if err := conn.WriteMessage(websocket.TextMessage,
		[]byte("The right ID!")); err != nil {
		panic(err)
	}
	return &t, nil
}

// FinishTask updates the given task  with status "Done"
func (s *Server) FinishTask(conn *websocket.Conn, owner string) {
	t, err := s.getCertainTask(conn, owner)
	if err != nil {
		return
	}
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
