package synchronizer

import (
	"strconv"

	"github.com/king-jam/tracker2jira/backend"
	jira "gopkg.in/andygrunwald/go-jira.v1"
)

// Jira ...
type Jira struct {
	client    *jira.Client
	projectID int
}

// NewJiraClient ...
func NewJiraClient(db backend.Database, projectID string) (*Jira, error) {
	// get the master project which should always be Pivotal Tracker in our case
	j := new(Jira)
	slaveProject, err := db.GetProjectByID(projectID)
	if err != nil {
		return j, err
	}
	// Get the slave Project Username/Password Details
	slaveUser, err := db.GetUserByID(slaveProject.AdminUserID)
	if err != nil {
		return j, err
	}
	j.client, err = jira.NewClient(nil, slaveProject.ProjectURL)
	if err != nil {
		return j, err
	}
	j.client.Authentication.SetBasicAuth(slaveUser.ExternalCredentials.Username, slaveUser.ExternalCredentials.Password)
	id, err := strconv.Atoi(slaveProject.ExternalID)
	if err != nil {
		return j, err
	}
	j.projectID = id
	return j, nil
}
