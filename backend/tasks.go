package backend

import (
	"log"

	"github.com/king-jam/tracker2jira/rest/models"
	uuid "github.com/satori/go.uuid"
)

const tasksPath = "tasks"
const defaultTaskState = "pending"

// GetTasks is ...
func (b *Backend) GetTasks() ([]*models.Task, error) {
	key := b.GetTaskBase()
	values, err := b.store.List(key)
	if len(values) == 0 {
		return []*models.Task{}, nil
	}
	if err != nil {
		return []*models.Task{}, err
	}
	tasks := []*models.Task{}
	for _, v := range values {
		task := &models.Task{}
		err = task.UnmarshalBinary(v.Value)
		if err != nil {
			return []*models.Task{}, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// GetTaskByID ...
func (b *Backend) GetTaskByID(taskid string) (*models.Task, error) {
	key := b.GetTaskBase() + taskid
	pair, err := b.store.Get(key)
	if err != nil {
		log.Printf("no version")
	}
	task := &models.Task{}
	err = task.UnmarshalBinary(pair.Value)
	if err != nil {
		return task, err
	}
	return task, nil
}

// PutTask ...
func (b *Backend) PutTask(task *models.Task) (*models.Task, error) {
	uuid := uuid.NewV4()
	key := b.GetTaskBase() + uuid.String()
	task.TaskID = uuid.String()
	task.Status = defaultTaskState // set the status to default until scheduled
	value, err := task.MarshalBinary()
	if err != nil {
		return nil, err
	}
	err = b.store.Put(key, value, nil)
	if err != nil {
		return nil, err
	}
	return task, nil
}

// DeleteTask ...
func (b *Backend) DeleteTask(taskid string) error {
	key := b.GetTaskBase() + taskid
	err := b.store.Delete(key)
	if err != nil {
		return err
	}
	return nil
}

// GetTaskBase returns the user base path
func (b *Backend) GetTaskBase() string {
	return b.instanceID + "/" + tasksPath + "/"
}
