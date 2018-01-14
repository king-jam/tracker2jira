package taskservice

import (
	"errors"
	"sync"
	"time"
)

// Task interface provides the behavior for the runner to periodically execute
// a synchronization job. Structs that implement this interface can be executed
// by this scheduling service.
type Task interface {
	Run() error
	ID() string
	SetRunning() error
	SetFailed() error
}

// defaultDelay is the default time in seconds between runs
const defaultDelay = 30

// TaskRunner holds references to all actively scheduled tasks.
// This is primarily a structure to hold a map back to the schedule task threads
// for cleanup & monitoring.
type TaskRunner struct {
	tasks map[string]taskRef
	wg    sync.WaitGroup
	mu    sync.Mutex
}

// taskRef is a helper struct to hold the Task definition and the stop control to
// cancel tasks
type taskRef struct {
	task   Task
	stopCh chan bool
}

// NewTaskRunner returns a new empty TaskRunner
func NewTaskRunner() (*TaskRunner, error) {
	return &TaskRunner{}, nil
}

// RunTask schedules a new Task to be periodically called
func (t *TaskRunner) RunTask(task Task) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if _, ok := t.tasks[task.ID()]; ok {
		return errors.New("a task already exists with the ID provided")
	}
	if t.tasks == nil {
		t.tasks = make(map[string]taskRef)
	}
	t.tasks[task.ID()] = t.schedule(task, defaultDelay*time.Second)
	task.SetRunning()
	return nil
}

// CancelTask stops a running task and removes its reference from the TaskRunner
// struct
func (t *TaskRunner) CancelTask(task Task) error {
	ref, ok := t.tasks[task.ID()]
	if !ok {
		return errors.New("Could not cancel tasks")
	}
	close(ref.stopCh)
	delete(t.tasks, task.ID())
	return nil
}

// schedule is the composer function that sets up the task thread and creates
// the necessary control channels to cancel the task later.
func (t *TaskRunner) schedule(task Task, delay time.Duration) taskRef {
	stop := make(chan bool)

	go func() {
		for {
			if err := task.Run(); err != nil {
				task.SetFailed()
				t.CancelTask(task)
			}
			select {
			case <-time.After(delay):
			case <-stop:
				return
			}
		}
	}()

	return taskRef{
		task:   task,
		stopCh: stop,
	}
}
