package backend

import (
	"github.com/king-jam/tracker2jira/rest/models"
)

const (
	// this is the key prefix for storage of all task objects, the format will be
	// /<storage prefix - should be t2j>/<t2j instance ID>/tasks/<task ID>/<Task Object>
	tasksPath = "tasks"
)

// TaskBackend interface encapsulates all the implementations of the task peristence
type TaskBackend interface {
	GetTasks() ([]*models.Task, error)
	GetTaskByID(taskid string) (*models.Task, error)
	PutTask(task *models.Task) (*models.Task, error)
	DeleteTask(taskid string) error
}

// GetTasks is returns an array of all tasks, this may need an accessor method
// at some point (probably never) based on scale
func (b *Backend) GetTasks() ([]*models.Task, error) {
	tasks := []*models.Task{}
	key := b.getTaskBase()
	values, err := b.store.List(key)
	if len(values) == 0 {
		return tasks, nil
	}
	if err != nil {
		return tasks, err
	}
	for _, v := range values {
		task := &models.Task{}
		err = task.UnmarshalBinary(v.Value)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// GetTaskByID returns a Task by ID
func (b *Backend) GetTaskByID(taskid string) (*models.Task, error) {
	task := &models.Task{}
	key := b.getTaskBase() + taskid
	pair, err := b.store.Get(key)
	if err != nil {
		return task, err
	}
	err = task.UnmarshalBinary(pair.Value)
	if err != nil {
		return task, err
	}
	return task, nil
}

// PutTask will persist a full formed task model
func (b *Backend) PutTask(task *models.Task) (*models.Task, error) {
	key := b.getTaskBase() + task.TaskID.String()
	value, err := task.MarshalBinary()
	if err != nil {
		return task, err
	}
	err = b.store.Put(key, value, nil)
	if err != nil {
		return task, err
	}
	return task, nil
}

// DeleteTask removes a task from the database by ID
func (b *Backend) DeleteTask(taskid string) error {
	key := b.getTaskBase() + taskid
	err := b.store.Delete(key)
	if err != nil {
		return err
	}
	return nil
}

// getTaskBase returns the task base path
func (b *Backend) getTaskBase() string {
	return b.instanceID + "/" + tasksPath + "/"
}
