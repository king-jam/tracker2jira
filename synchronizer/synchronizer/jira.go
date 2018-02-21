package synchronizer

import (
	"strconv"

	"github.com/king-jam/tracker2jira/backend"
	jira "gopkg.in/andygrunwald/go-jira.v1"
)

// TODO: make these less coupled to the DB backend

// Jira ...
type Jira struct {
	client    *jira.Client
	projectID int
}

// NewJiraClient ...
func NewJiraClient(db backend.Database, projectID string) (*Jira, error) {
	// get the master project which should always be Pivotal Tracker in our case
	j := new(Jira)
	dstProject, err := db.GetProjectByID(projectID)
	if err != nil {
		return j, err
	}
	// Get the slave Project Username/Password Details
	dstUser, err := db.GetUserByID(dstProject.AdminUserID)
	if err != nil {
		return j, err
	}
	j.client, err = jira.NewClient(nil, dstProject.ProjectURL)
	if err != nil {
		return j, err
	}
	j.client.Authentication.SetBasicAuth(dstUser.ExternalCredentials.Username, dstUser.ExternalCredentials.Password.String())
	id, err := strconv.Atoi(dstProject.ExternalID)
	if err != nil {
		return j, err
	}
	j.projectID = id
	return j, nil
}
