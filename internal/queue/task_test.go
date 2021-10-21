package queue

import (
	"testing"
	"time"

	apiq "github.com/stalin-777/test-apiq"
)

func TestDoTask(t *testing.T) {

	tests := []struct {
		name string
		t    *apiq.Task
	}{
		{
			"1",
			&apiq.Task{
				ID:          1,
				State:       apiq.StateInQueue,
				NumElements: 3,
				Delta:       10,
				StartValue:  0,
				Interval:    1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go DoTask(tt.t)
			<-time.NewTicker(time.Second * 1).C
			if tt.t.State != apiq.StateInProgress {
				t.Errorf("State want: in progress, got: %s", tt.t.State)
			}
			<-time.NewTicker(time.Second * 4).C

			if tt.t.State != apiq.StateCompleted {
				t.Errorf("State want: compleated, got: %s ", tt.t.State)
			}

			if tt.t.CurrentVal != 30 {
				t.Errorf("CurrentVal want: 30, got: %v", tt.t.CurrentVal)
			}

			if tt.t.CurrentIter != 3 {
				t.Errorf("CurrentIter want: 30, got: %v", tt.t.CurrentIter)
			}
		})
	}
}
