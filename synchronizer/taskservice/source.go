package taskservice

import (
	"github.com/king-jam/tracker2jira/backend"
	"github.com/king-jam/tracker2jira/rest/models"
	"github.com/king-jam/tracker2jira/synchronizer/synchronizer"
)

// TaskSource implements the Source interface. This is a DB/backend specific
// implementation that provides implementation of the synchronization jobs to the
// task scheduler & runner services.
type TaskSource struct {
	db backend.Database
}

// NewTaskSource composes a TaskSource object that implements the Source interface
func NewTaskSource(db backend.Database) (*TaskSource, error) {
	return &TaskSource{
		db: db,
	}, nil
}

// GetAllTasks provides Synchronizers that implement the Task interface. This is
// primarily used to do initial scheduling on system startup.
func (s *TaskSource) GetAllTasks() ([]Task, error) {
	tasks := []Task{}
	dbTasks, err := s.db.GetTasks()
	if err != nil {
		return tasks, err
	}
	for _, dbTask := range dbTasks {
		task := synchronizer.NewSynchronizer(s.db, dbTask)
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// GetPendingTasks provides Synchronizers that implement the Task interface. This
// is used to get all tasks in a pending state that are awaiting scheduling.
func (s *TaskSource) GetPendingTasks() ([]Task, error) {
	tasks := []Task{}
	dbTasks, err := s.db.GetTasks()
	if err != nil {
		return tasks, err
	}
	for _, dbTask := range dbTasks {
		if dbTask.Status == models.TaskStatusPending {
			task := synchronizer.NewSynchronizer(s.db, dbTask)
			tasks = append(tasks, task)
		}
	}
	return tasks, nil
}

// GetCancelledTasks provides Synchronizers that implement the Task interface. This
// is used to get all tasks in a pending state that are awaiting scheduling.
func (s *TaskSource) GetCancelledTasks() ([]string, error) {
	taskIDs := []string{}
	dbTasks, err := s.db.GetTasks()
	if err != nil {
		return taskIDs, err
	}
	for _, dbTask := range dbTasks {
		if dbTask.Status == models.TaskStatusCancel {
			taskIDs = append(taskIDs, dbTask.TaskID.String())
		}
	}
	return taskIDs, nil
}
