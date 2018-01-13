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

// Poller ...
type Poller interface {
	Start() error
	Stop() error
}

// Runner ...
type Runner interface {
	RunTask(task Task) error
	CancelTask(task Task) error
}

// Source ...
type Source interface {
	GetAllTasks() ([]Task, error)
	GetPendingTasks() ([]Task, error)
}

// TaskPoller ...
type TaskPoller struct {
	runner    Runner
	source    Source
	pollDelay time.Duration
	started   bool
	stopCh    chan bool
	wg        sync.WaitGroup
	once      sync.Once
}

// NewTaskPoller ...
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

// Start ...
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

// Stop ...
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
}

func (t *TaskPoller) runPendingTasks() {
	log.Debugf("SCHEDULING PENDING TASKS")
	tasks, err := t.source.GetPendingTasks()
	if err != nil {
		log.Warnf("Error getting pending Tasks")
	}
	t.runTask(tasks)
}

func (t *TaskPoller) runAllTasks() {
	log.Debugf("SCHEDULING ALL TASKS")
	tasks, err := t.source.GetAllTasks()
	if err != nil {
		log.Warnf("DB Access Error for Tasks")
	}
	t.runTask(tasks)
}

func (t *TaskPoller) runTask(tasks []Task) {
	for _, task := range tasks {
		log.Debugf("Scheduling: %+v", task)
		t.runner.RunTask(task)
	}
}
