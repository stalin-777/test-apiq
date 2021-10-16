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
func DoTask(t *apiq.Task) {

	t.StartedAt = time.Now().Format(time.RFC3339)
	t.State = apiq.StateInProgress
	t.CurrentVal = t.StartValue

	for i := 1; i <= t.NumElements; i++ {

		t.CurrentVal += t.Delta
		t.CurrentIter++
		<-time.NewTimer(time.Second * time.Duration(t.Interval)).C
	}

	t.State = apiq.StateCompleted
	t.CompletedAt = time.Now().Format(time.RFC3339)
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
