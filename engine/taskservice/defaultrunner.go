package taskservice

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/king-jam/tracker2jira/backend"
	"github.com/king-jam/tracker2jira/rest/models"
)

// Task ...
type Task interface {
	RunTask() error
	CancelTask() error
}

// TaskRunner ...
type TaskRunner struct {
	tasks map[string]chan bool

	db *backend.Backend

	wg sync.WaitGroup
	mu sync.Mutex
}

// NewTaskRunner ...
func NewTaskRunner() (*TaskRunner, error) {
	backend, err := backend.GetDB()
	if err != nil {
		return &TaskRunner{}, err
	}
	return &TaskRunner{
		db: backend,
	}, nil
}

// ScheduleTask ...
func (t *TaskRunner) ScheduleTask(task *models.Task) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if _, ok := t.tasks[task.TaskID.String()]; ok {
		return errors.New("a task already exists with the ID provided")
	}
	if t.tasks == nil {
		t.tasks = make(map[string]chan bool)
	}
	t.tasks[task.TaskID.String()] = t.schedule(30 * time.Second)
	task.Status = "running"
	t.db.PutTask(task)
	//t.schedule(task, 60*time.Second)
	return nil
}

// CancelTask ...
func (t *TaskRunner) CancelTask(task models.Task) error {
	_, ok := t.tasks[task.TaskID.String()]
	if !ok {
		return nil
	}
	return nil
}

// CancelAllTask ...
func (t *TaskRunner) CancelAllTask() error {
	return nil
}

func (t *TaskRunner) schedule(delay time.Duration) chan bool {
	stop := make(chan bool)

	go func() {
		for {
			log.Println("DOING STATE UPDATE")
			select {
			case <-time.After(delay):
			case <-stop:
				return
			}
		}
	}()

	return stop
}
