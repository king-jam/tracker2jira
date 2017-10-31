package taskservice

// Scheduler ...
type Scheduler interface {
	Start() error
	Stop() error
}

// TaskService ...
type TaskService struct {
	scheduler Scheduler
	runner    Runner

	started bool
}

// NewTaskService ...
func NewTaskService() (*TaskService, error) {
	runner, err := NewTaskRunner()
	if err != nil {
		return &TaskService{}, err
	}
	scheduler, err := NewTaskScheduler(runner)
	if err != nil {
		return &TaskService{}, err
	}
	t := &TaskService{
		scheduler: scheduler,
		runner:    runner,
	}
	if err := t.start(); err != nil {
		return &TaskService{}, err
	}
	t.started = true
	return t, nil
}

// Start ...
func (t *TaskService) start() error {
	return t.scheduler.Start()
}

func (t *TaskService) stop() error {
	return t.scheduler.Stop()
}
