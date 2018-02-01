package backend

import (
	"github.com/king-jam/tracker2jira/rest/models"
)

const (
	// this is the key prefix for storage of all project objects, the format will be
	// /<storage prefix - should be t2j>/<t2j instance ID>/projects/<project ID>/<Project Object>
	projectsPath = "projects"
)

// ProjectBackend interface encapsulates all the implementations of the project peristence
type ProjectBackend interface {
	GetProjects() ([]*models.Project, error)
	GetProjectByID(projectid string) (*models.Project, error)
	PutProject(project *models.Project) (*models.Project, error)
	DeleteProject(projectid string) error
}

// GetProjects is returns an array of all Projects
func (b *Backend) GetProjects() ([]*models.Project, error) {
	projects := []*models.Project{}
	key := b.getProjectsBase()
	values, err := b.store.List(key)
	if len(values) == 0 {
		return projects, nil
	}
	if err != nil {
		return projects, err
	}
	for _, v := range values {
		project := &models.Project{}
		err = project.UnmarshalBinary(v.Value)
		if err != nil {
			return projects, err
		}
		projects = append(projects, project)
	}
	return projects, nil
}

// GetProjectByID returns a project object by ID
func (b *Backend) GetProjectByID(projectid string) (*models.Project, error) {
	project := &models.Project{}
	key := b.getProjectsBase() + projectid
	pair, err := b.store.Get(key)
	if err != nil {
		return project, err
	}
	err = project.UnmarshalBinary(pair.Value)
	if err != nil {
		return project, err
	}
	return project, nil
}

// PutProject ...
func (b *Backend) PutProject(project *models.Project) (*models.Project, error) {
	key := b.getProjectsBase() + project.ProjectID.String()
	value, err := project.MarshalBinary()
	if err != nil {
		return project, err
	}
	err = b.store.Put(key, value, nil)
	if err != nil {
		return project, err
	}
	return project, nil
}

// DeleteProject ...
func (b *Backend) DeleteProject(projectid string) error {
	key := b.getProjectsBase() + projectid
	err := b.store.Delete(key)
	if err != nil {
		return err
	}
	return nil
}

// getProjectsBase returns the user base path
func (b *Backend) getProjectsBase() string {
	return b.instanceID + "/" + projectsPath + "/"
}
