package taskservice

import (
	"log"
	"sync"
	"time"

	"github.com/king-jam/tracker2jira/backend"
	"github.com/king-jam/tracker2jira/rest/models"
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

// TaskScheduler ...
type TaskScheduler struct {
	runner *TaskRunner

	db        *backend.Backend
	pollDelay time.Duration

	started bool
	stopCh  chan bool
	wg      sync.WaitGroup
	once    sync.Once
}

// NewTaskScheduler ...
func NewTaskScheduler(runner *TaskRunner) (*TaskScheduler, error) {
	backend, err := backend.GetDB()
	if err != nil {
		return &TaskScheduler{}, err
	}
	return &TaskScheduler{
		started:   false,
		runner:    runner,
		db:        backend,
		stopCh:    nil,
		pollDelay: 30 * time.Second,
		wg:        sync.WaitGroup{},
	}, nil
}

// Start ...
func (t *TaskScheduler) Start() error {
	if t.started {
		return nil
	}
	t.wg.Add(1)
	t.once.Do(t.scheduleAllTasks)
	t.stopCh = t.startDBPoller()
	t.started = true
	return nil
}

// Stop ...
func (t *TaskScheduler) Stop() error {
	if t.started {
		close(t.stopCh)
		t.wg.Done()
	}
	return nil
}

func (t *TaskScheduler) startDBPoller() chan bool {
	stop := make(chan bool)

	go func() {
		for {
			t.taskScheduler()
			select {
			case <-time.After(t.pollDelay):
			case <-stop:
				return
			}
		}
	}()

	return stop
}

func (t *TaskScheduler) taskScheduler() {
	t.schedulePendingTasks()
}

func (t *TaskScheduler) schedulePendingTasks() {
	log.Println("SCHEDULING PENDING TASKS")
	tasks, err := t.db.GetTasks()
	if err != nil {
		log.Println("DB Access Error for Tasks")
	}
	for _, v := range tasks {
		if v.Status == "pending" {
			t.scheduleTask(v)
		}
	}
}

func (t *TaskScheduler) scheduleAllTasks() {
	log.Println("SCHEDULING ALL TASKS")
	tasks, err := t.db.GetTasks()
	if err != nil {
		log.Println("DB Access Error for Tasks")
	}
	for _, v := range tasks {
		t.scheduleTask(v)
	}
}

func (t *TaskScheduler) scheduleTask(task *models.Task) {
	task.Status = "scheduled"
	t.db.PutTask(task)
	log.Printf("Scheduling: %+v", task)
	t.runner.ScheduleTask(task)
}
