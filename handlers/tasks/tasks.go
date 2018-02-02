package tasks

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/king-jam/tracker2jira/backend"
	"github.com/king-jam/tracker2jira/rest/models"
	"github.com/king-jam/tracker2jira/rest/server/operations/tasks"
	uuid "github.com/satori/go.uuid"
)

const defaultTaskState = models.TaskStatusPending

// GetTask ...
func GetTask(db backend.Database, params tasks.GetTaskByIDParams) middleware.Responder {
	value, err := db.GetTaskByID(params.TaskID)
	if err != nil {
		return &tasks.GetTaskByIDNotFound{}
	}
	return &tasks.GetTaskByIDOK{
		Payload: value,
	}
}

// GetTasks ...
func GetTasks(db backend.Database, params tasks.GetTasksParams) middleware.Responder {
	values, err := db.GetTasks()
	if err != nil {
		return &tasks.GetTasksBadRequest{}
	}
	return &tasks.GetTasksOK{
		Payload: values,
	}
}

// PostTask ...
func PostTask(db backend.Database, params tasks.PostTaskParams) middleware.Responder {
	uuid := uuid.NewV4()
	params.Body.TaskID = strfmt.UUID4(uuid.String())
	params.Body.Status = defaultTaskState // set the status to default until scheduled
	value, err := db.PutTask(params.Body)
	if err != nil {
		return &tasks.PostTaskBadRequest{}
	}
	return &tasks.PostTaskAccepted{
		Payload: value,
	}
}

// DeleteTask ...
func DeleteTask(db backend.Database, params tasks.DeleteTaskByIDParams) middleware.Responder {
	err := db.DeleteTask(params.TaskID)
	if err != nil {
		return &tasks.DeleteTaskByIDNotFound{}
	}
	return &tasks.DeleteTaskByIDNoContent{}
}
