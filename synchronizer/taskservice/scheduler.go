package taskservice

import "github.com/king-jam/tracker2jira/backend"

// TaskScheduler is the top-level scheduler service class
// This class wraps the embedded polling service to start and create new synchronizer
// jobs. This should probably be moved to something more "eventy" in the future
// but this works for now.
type TaskScheduler struct {
	poller Poller
}

// NewTaskScheduler composes all the underlying services to create the task
// scheduler and synchronization service.
func NewTaskScheduler(db backend.Database) (*TaskScheduler, error) {
	source, err := NewTaskSource(db)
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
