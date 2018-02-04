package taskservice

import (
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

// for loop polling DB for tasks put into db bu user api
// found new tasks
// change state to "starting"
// task.master go get project object
// task.slave go get project object
// master.project.userid go get user credentials
// slave/project/userid go get user creds
// connect/login to both projects
// validate task is good to go
// create go routine to run below
// start story.go loop function
// change state to "running"

// Poller interface describes the behavior of a general polling service. It
// exposes the basic control methods.
type Poller interface {
	Start() error
	Stop() error
}

// Runner interface describes the task runner behavior for controlling active
// tasks.
type Runner interface {
	RunTask(task Task) error
	CancelTask(id string) error
}

// Source interface describes task source behavior. Implementation is mainly to
// abstract the backend datastore and returns structs that implement the Task
// interface.
type Source interface {
	GetAllTasks() ([]Task, error)
	GetPendingTasks() ([]Task, error)
	GetCancelledTasks() ([]string, error)
}

// TaskPoller struct holds all the state logic and child structures to watch a
// backend datastore and schedule tasks to run periodically. This implements the
// Poller interface.
type TaskPoller struct {
	runner    Runner
	source    Source
	pollDelay time.Duration
	started   bool
	stopCh    chan bool
	wg        sync.WaitGroup
	once      sync.Once
}

// NewTaskPoller returns a TaskPoller struct which implements the Poller
// interface. It is composed of a Runner and Source interface.
func NewTaskPoller(runner Runner, source Source) (*TaskPoller, error) {
	return &TaskPoller{
		started:   false,
		runner:    runner,
		source:    source,
		stopCh:    nil,
		pollDelay: 30 * time.Second,
		wg:        sync.WaitGroup{},
	}, nil
}

// Start implements the Start functionality of the Poller interface.
func (t *TaskPoller) Start() error {
	if t.started {
		return nil
	}
	t.wg.Add(1)
	t.once.Do(t.runAllTasks)
	t.stopCh = t.startPoller()
	t.started = true
	return nil
}

// Stop implements the Stop functionality of the Poller interface.
func (t *TaskPoller) Stop() error {
	if t.started {
		close(t.stopCh)
		t.wg.Done()
	}
	return nil
}

func (t *TaskPoller) startPoller() chan bool {
	stop := make(chan bool)

	go func() {
		for {
			t.taskPoller()
			select {
			case <-time.After(t.pollDelay):
			case <-stop:
				return
			}
		}
	}()

	return stop
}

func (t *TaskPoller) taskPoller() {
	t.runPendingTasks()
	t.stopCancelledTasks()
}

func (t *TaskPoller) runPendingTasks() {
	log.Debugf("SCHEDULING PENDING TASKS")
	tasks, err := t.source.GetPendingTasks()
	if err != nil {
		log.Warnf("Error getting pending Tasks")
	}
	t.runTasks(tasks)
}

func (t *TaskPoller) stopCancelledTasks() {
	log.Debugf("STOPPING CANCELLED TASKS")
	taskIDs, err := t.source.GetCancelledTasks()
	if err != nil {
		log.Warnf("Error getting pending Tasks")
	}
	t.cancelTasks(taskIDs)
}

func (t *TaskPoller) runAllTasks() {
	log.Debugf("SCHEDULING ALL TASKS")
	tasks, err := t.source.GetAllTasks()
	if err != nil {
		log.Warnf("DB Access Error for Tasks")
	}
	t.runTasks(tasks)
}

func (t *TaskPoller) runTasks(tasks []Task) {
	for _, task := range tasks {
		log.Debugf("Scheduling: %+v", task)
		t.runner.RunTask(task)
	}
}

func (t *TaskPoller) cancelTasks(taskIDs []string) {
	for _, id := range taskIDs {
		log.Debugf("Cancelling: %+v", id)
		t.runner.CancelTask(id)
	}
}
