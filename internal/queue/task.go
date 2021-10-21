package queue

import (
	"context"
	"time"

	apiq "github.com/stalin-777/test-apiq"
)

type TaskService struct {
	Queue *Queue
}

//FindTasks - Returns a list of all tasks
func (s *TaskService) FindTasks() ([]*apiq.Task, error) {

	return s.Queue.GetList(), nil
}

//CreateTask - Creates a task and queues it.
//The task is removed from the list after TTL expires
func (s *TaskService) CreateTask(t *apiq.Task) error {

	prepareToPush(t, s)

	s.Queue.Push(t)

	go s.waitForTTLEnd(t)

	return nil
}

//DoTask - Calculates arithmetic progression
func DoTask(s *apiq.Task) {

	s.StartedAt = time.Now().Format(time.RFC3339)
	s.State = apiq.StateInProgress
	s.CurrentVal = s.StartValue

	for i := 1; i <= s.NumElements; i++ {

		s.CurrentVal += s.Delta
		s.CurrentIter++
		<-time.NewTimer(time.Second * time.Duration(s.Interval)).C
	}

	s.State = apiq.StateCompleted
	s.CompletedAt = time.Now().Format(time.RFC3339)
}

func prepareToPush(t *apiq.Task, ts *TaskService) {

	t.State = apiq.StateInQueue
	t.CurrentIter = 0
	t.CurrentVal = 0.0
	t.CreatedAt = time.Now().Format(time.RFC3339)
	t.StartedAt = ""
	t.CompletedAt = ""
	t.TTL = float64(ts.Queue.TTL) / float64(time.Second)
}

func (s *TaskService) waitForTTLEnd(t *apiq.Task) {

	ctx, cancel := context.WithTimeout(context.Background(), s.Queue.TTL)
	defer cancel()

	<-ctx.Done()
	s.Queue.Kick()
}
