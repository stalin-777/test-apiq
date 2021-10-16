package apiq

type Task struct {
	ID          int     `json:"id"`                    // Номер в очереди
	State       string  `json:"state"`                 // Статус: В процессе/В очереди/Завершена
	NumElements int     `json:"n"`                     // количество элементов
	Delta       float64 `json:"d"`                     // дельта между элементами последовательности
	StartValue  float64 `json:"n1"`                    // Стартовое значение
	Interval    float64 `json:"I"`                     // интервал в секундах между итерациями
	TTL         float64 `json:"TTL"`                   // время хранения результата в секундах
	CurrentIter int     `json:"currentIter"`           // Текущая итерация
	CurrentVal  float64 `json:"currentVal"`            // Текущее значение
	CreatedAt   string  `json:"createdt"`              // Время постановки задачи
	StartedAt   string  `json:"startedAt,omitempty"`   // Время старта задачи
	CompletedAt string  `json:"completedAt,omitempty"` // Время окончания задачи
}

const (
	//Task states
	StateInProgress = "in progress"
	StateInQueue    = "in queue"
	StateCompleted  = "completed"
)

type TaskService interface {
	FindTasks() ([]*Task, error)
	CreateTask(t *Task) error
}
