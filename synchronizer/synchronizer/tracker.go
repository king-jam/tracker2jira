package synchronizer

import (
	"strconv"

	"github.com/king-jam/go-pivotaltracker/v5/pivotal"
	"github.com/king-jam/tracker2jira/backend"
)

// PivotalTracker ...
type PivotalTracker struct {
	client    *pivotal.Client
	projectID int
	version   int
}

// NewTrackerClient ...
func NewTrackerClient(db backend.Database, projectID string) (*PivotalTracker, error) {
	// get the master project which should always be Pivotal Tracker in our case
	tracker := new(PivotalTracker)
	masterProject, err := db.GetProjectByID(projectID)
	if err != nil {
		return tracker, err
	}
	// Get the master Project Username/Password Details
	masterUser, err := db.GetUserByID(masterProject.AdminUserID)
	if err != nil {
		return tracker, err
	}
	tracker.client = pivotal.NewClient(masterUser.ExternalCredentials.Token)
	id, err := strconv.Atoi(masterProject.ExternalID)
	if err != nil {
		return tracker, err
	}
	tracker.projectID = id
	tracker.version = int(masterProject.ProjectVersion)
	return tracker, nil
}

// Client ...
func (p *PivotalTracker) Client() *pivotal.Client {
	return p.client
}

// UpdateCursor ...
func (p *PivotalTracker) UpdateCursor() (c *pivotal.ActivityCursor, err error) {
	return p.client.Activity.Iterate(p.projectID, p.version, true)
}
