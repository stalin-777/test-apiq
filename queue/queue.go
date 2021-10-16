package queue

import (
	"time"

	apiq "github.com/stalin-777/test-apiq"
)

type Queue struct {
	//List of tasks, in all statuses (in the queue / in progress / completed)
	taskList []*apiq.Task
	//The queue of tasks to be processed
	tasksForProcessing []*apiq.Task
	//ID of the last created task. Used to increment the id of the following tasks
	idLastCreatedTask int
	//The size of the task queue for processing
	len int
	//Time to life
	TTL time.Duration
}

//New - Create a new task queue
func New(ttl time.Duration) *Queue {

	return &Queue{
		taskList: []*apiq.Task{},
		TTL:      ttl,
	}
}

//Push - Adds a task to the list of tasks and to the queue of tasks for processing
func (s *Queue) Push(t *apiq.Task) {

	s.idLastCreatedTask++
	t.ID = s.idLastCreatedTask
	s.taskList = append(s.taskList, t)
	s.tasksForProcessing = append(s.tasksForProcessing, t)
	s.len++
}

//Pop - Retrieves a task from the queue
func (s *Queue) Pop() *apiq.Task {

	task := s.tasksForProcessing[0]
	s.tasksForProcessing = s.tasksForProcessing[1:]
	s.len--

	return task
}

//Removes the first task from the list
func (s *Queue) Kick() {

	s.taskList = s.taskList[1:]
}

//GetList - Returns a list of tasks in all statuses (in the queue / in progress / completed)
func (s *Queue) GetList() []*apiq.Task {

	return s.taskList
}

//Len - Returns the size of the queue
func (s *Queue) Len() int {

	return len(s.tasksForProcessing)
}
