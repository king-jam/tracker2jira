package synchronizer

import (
	"github.com/king-jam/go-pivotaltracker/v5/pivotal"
	jira "gopkg.in/andygrunwald/go-jira.v1"
)

// DefaultActivityHandler is a fill in for types we do not want to handle
type DefaultActivityHandler struct {
}

// Synchronize is the null handler function for synchronizing activities
func (d DefaultActivityHandler) Synchronize(activity *pivotal.Activity, pt *pivotal.Client, j *jira.Client) error {
	// Default activity is to do nothing
	return nil
}
