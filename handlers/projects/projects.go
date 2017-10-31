package projects

import (
	"github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/king-jam/tracker2jira/backend"
	"github.com/king-jam/tracker2jira/rest/server/operations/projects"
	uuid "github.com/satori/go.uuid"
)

const defaultProjectVersion = 0

// GetProject ...
func GetProject(db *backend.Backend, params projects.GetProjectByIDParams) middleware.Responder {
	value, err := db.GetProjectByID(params.ProjectID)
	if err != nil {
		return &projects.GetProjectByIDNotFound{}
	}
	return &projects.GetProjectByIDOK{
		Payload: value,
	}
}

// GetProjects ...
func GetProjects(db *backend.Backend, params projects.GetProjectsParams) middleware.Responder {
	values, err := db.GetProjects()
	if err != nil {
		return &projects.GetProjectsBadRequest{}
	}
	return &projects.GetProjectsOK{
		Payload: values,
	}
}

// PostProject ...// init to 0 for post create
func PostProject(db *backend.Backend, params projects.PostProjectParams) middleware.Responder {
	uuid := uuid.NewV4()
	params.Body.ProjectID = strfmt.UUID4(uuid.String())
	params.Body.ProjectVersion = defaultProjectVersion
	value, err := db.PutProject(params.Body)
	if err != nil {
		return &projects.PostProjectBadRequest{}
	}
	return &projects.PostProjectCreated{
		Payload: value,
	}
}

// DeleteProject ...
func DeleteProject(db *backend.Backend, params projects.DeleteProjectByIDParams) middleware.Responder {
	err := db.DeleteProject(params.ProjectID)
	if err != nil {
		return &projects.DeleteProjectByIDNotFound{}
	}
	return &projects.DeleteProjectByIDNoContent{}
}
