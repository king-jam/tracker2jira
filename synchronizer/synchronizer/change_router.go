package synchronizer

import (
	"github.com/king-jam/go-pivotaltracker/v5/pivotal"
	jira "gopkg.in/andygrunwald/go-jira.v1"
)

var changeHandlers map[string]ChangeHandler

func init() {
	changeHandlers = map[string]ChangeHandler{
		// "blocker_create_activity":  DefaultActivityHandler{}, // don't care
		// "blocker_delete_activity":  DefaultActivityHandler{}, // don't care
		// "blocker_update_activity":  DefaultActivityHandler{}, // don't care
		// "blocking_create_activity": DefaultActivityHandler{}, // don't care
		// "blocking_delete_activity": DefaultActivityHandler{}, // don't care
		// "branch_create_activity":   DefaultActivityHandler{}, // don't care
		// "branch_delete_activity":   DefaultActivityHandler{}, // don't care
		// "comment_create_activity":  DefaultActivityHandler{}, // don't care
	}
}

// ChangeHandler wraps the functionality to have a class factory for handling
// change events from Pivotal Tracker
type ChangeHandler interface {
	HandleChanges(*pivotal.Change, *pivotal.Client, *jira.Client) error
}
