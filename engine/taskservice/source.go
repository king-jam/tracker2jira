package taskservice

import (
	"github.com/king-jam/tracker2jira/backend"
	"github.com/king-jam/tracker2jira/engine/synchronizer"
	"github.com/king-jam/tracker2jira/rest/models"
)

// TaskSource ...
type TaskSource struct {
	db *backend.Backend
}

// NewTaskSource ...
func NewTaskSource() (*TaskSource, error) {
	db, err := backend.GetDB()
	if err != nil {
		return &TaskSource{}, err
	}
	return &TaskSource{
		db: db,
	}, nil
}

// GetAllTasks ...
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

// GetPendingTasks ...
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
