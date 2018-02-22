package synchronizer

import (
	"fmt"

	"github.com/king-jam/go-pivotaltracker/v5/pivotal"
	jira "gopkg.in/andygrunwald/go-jira.v1"
)

// EpicActivityHandler is the struct to wrap all Epic related updates
type EpicActivityHandler struct {
}

// Synchronize is the general handler function for synchronizing epic activities
func (e EpicActivityHandler) Synchronize(activity *pivotal.Activity, pt *pivotal.Client, j *jira.Client) error {
	switch kind := activity.Kind; kind {
	case "epic_create_activity":
		// 1. Get the Epic details from the activity
		// 2. Search JIRA epics for the one matching the PT number.
		// 3. If not already exist, create the new Epic reference in JIRA
		// 4. Post Comment in JIRA Epic referencing Pivotal Tracker Epic
		// 5. Post Comment in Pivotal Tracker Epic referencing JIRA
		return nil
	case "epic_delete_activity":
		// 1. Get the Epic details from the activity
		// 2. Search JIRA epics for the one matching the PT number.
		// 3. Close the Epic reference in JIRA with a comment about why.
		return nil
	case "epic_update_activity":
		// 1. Get the Epic details from the activity
		// 2. Search JIRA epics for the one matching the PT number.
		// 3. Update the Description or the Title details.
		return nil
	case "epic_move_activity":
		// don't care, this is a tracker only thing for priority of the Epic
		return nil
	default:
		return fmt.Errorf("unsupported kind of Epic: this shouldn't happen")
	}
}
