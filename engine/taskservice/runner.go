package taskservice

import (
	"errors"
	"sync"
	"time"
)

// Task ...
type Task interface {
	Run() error
	ID() string
	SetRunning() error
	SetFailed() error
}

const defaultDelay = 30

// TaskRunner ...
type TaskRunner struct {
	tasks map[string]taskRef
	wg    sync.WaitGroup
	mu    sync.Mutex
}

type taskRef struct {
	task   Task
	stopCh chan bool
}

// NewTaskRunner ...
func NewTaskRunner() (*TaskRunner, error) {
	return &TaskRunner{}, nil
}

// RunTask ...
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

// CancelTask ...
func (t *TaskRunner) CancelTask(task Task) error {
	ref, ok := t.tasks[task.ID()]
	if !ok {
		return errors.New("Could not cancel tasks")
	}
	close(ref.stopCh)
	delete(t.tasks, task.ID())
	return nil
}

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
