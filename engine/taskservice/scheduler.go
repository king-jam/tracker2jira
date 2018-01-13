package taskservice

// TaskScheduler ...
type TaskScheduler struct {
	poller Poller
}

// NewTaskScheduler ...
func NewTaskScheduler() (*TaskScheduler, error) {
	source, err := NewTaskSource()
	if err != nil {
		return &TaskScheduler{}, err
	}
	runner, err := NewTaskRunner()
	if err != nil {
		return &TaskScheduler{}, err
	}
	poller, err := NewTaskPoller(runner, source)
	if err != nil {
		return &TaskScheduler{}, err
	}
	t := &TaskScheduler{
		poller: poller,
	}
	if err := t.start(); err != nil {
		return &TaskScheduler{}, err
	}
	return t, nil
}

// Start ...
func (t *TaskScheduler) start() error {
	return t.poller.Start()
}
