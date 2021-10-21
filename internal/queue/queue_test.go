package queue

import (
	"testing"

	apiq "github.com/stalin-777/test-apiq"
)

func TestQueue_Len(t *testing.T) {
	type fields struct {
		tasksForProcessing []*apiq.Task
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "1",
			fields: fields{
				tasksForProcessing: []*apiq.Task{{}},
			},
			want: 1,
		},
		{
			name: "2",
			fields: fields{
				tasksForProcessing: []*apiq.Task{{}, {}},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Queue{
				tasksForProcessing: tt.fields.tasksForProcessing,
			}
			if got := s.Len(); got != tt.want {
				t.Errorf("Queue.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}
