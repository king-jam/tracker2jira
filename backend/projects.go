package backend

import (
	log "github.com/sirupsen/logrus"

	"github.com/king-jam/tracker2jira/rest/models"
)

const projectsPath = "projects"

// GetProjects is ...
func (b *Backend) GetProjects() ([]*models.Project, error) {
	key := b.GetProjectsBase()
	values, err := b.store.List(key)
	if len(values) == 0 {
		return []*models.Project{}, nil
	}
	if err != nil {
		return []*models.Project{}, err
	}
	projects := []*models.Project{}
	for _, v := range values {
		project := &models.Project{}
		err = project.UnmarshalBinary(v.Value)
		if err != nil {
			return []*models.Project{}, err
		}
		projects = append(projects, project)
	}
	return projects, nil
}

// GetProjectByID ...
func (b *Backend) GetProjectByID(projectid string) (*models.Project, error) {
	key := b.GetProjectsBase() + projectid
	pair, err := b.store.Get(key)
	if err != nil {
		log.Printf("no version")
	}
	project := &models.Project{}
	err = project.UnmarshalBinary(pair.Value)
	if err != nil {
		return project, err
	}
	return project, nil
}

// PutProject ...
func (b *Backend) PutProject(project *models.Project) (*models.Project, error) {
	key := b.GetProjectsBase() + project.ProjectID.String()
	value, err := project.MarshalBinary()
	if err != nil {
		return nil, err
	}
	err = b.store.Put(key, value, nil)
	if err != nil {
		return nil, err
	}
	return project, nil
}

// DeleteProject ...
func (b *Backend) DeleteProject(projectid string) error {
	key := b.GetProjectsBase() + projectid
	err := b.store.Delete(key)
	if err != nil {
		return err
	}
	return nil
}

// GetProjectsBase returns the user base path
func (b *Backend) GetProjectsBase() string {
	return b.instanceID + "/" + projectsPath + "/"
}
