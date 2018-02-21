package synchronizer

import (
	"fmt"

	"github.com/king-jam/go-pivotaltracker/v5/pivotal"
	jira "gopkg.in/andygrunwald/go-jira.v1"
)

var handlers map[string]ActivityHandler

func init() {
	handlers = map[string]ActivityHandler{
		"blocker_create_activity":                         DefaultActivityHandler{}, // don't care
		"blocker_delete_activity":                         DefaultActivityHandler{}, // don't care
		"blocker_update_activity":                         DefaultActivityHandler{}, // don't care
		"blocking_create_activity":                        DefaultActivityHandler{}, // don't care
		"blocking_delete_activity":                        DefaultActivityHandler{}, // don't care
		"branch_create_activity":                          DefaultActivityHandler{}, // don't care
		"branch_delete_activity":                          DefaultActivityHandler{}, // don't care
		"comment_create_activity":                         DefaultActivityHandler{}, // don't care
		"comment_delete_activity":                         DefaultActivityHandler{}, // don't care
		"comment_update_activity":                         DefaultActivityHandler{}, // don't care
		"epic_create_activity":                            EpicActivityHandler{},    // care
		"epic_delete_activity":                            EpicActivityHandler{},    // care
		"epic_move_activity":                              EpicActivityHandler{},    // care
		"epic_update_activity":                            EpicActivityHandler{},    // care
		"follower_create_activity":                        DefaultActivityHandler{}, // don't care
		"follower_delete_activity":                        DefaultActivityHandler{}, // don't care
		"iteration_update_activity":                       DefaultActivityHandler{}, // don't care
		"label_create_activity":                           DefaultActivityHandler{}, // don't care
		"label_delete_activity":                           DefaultActivityHandler{}, // don't care
		"label_update_activity":                           DefaultActivityHandler{}, // don't care
		"model_import_activity":                           DefaultActivityHandler{}, // don't care
		"project_membership_create_activity":              DefaultActivityHandler{}, // don't care
		"project_membership_delete_activity":              DefaultActivityHandler{}, // don't care
		"project_membership_update_activity":              DefaultActivityHandler{}, // don't care
		"project_update_activity":                         DefaultActivityHandler{}, // don't care
		"pull_request_create_activity":                    DefaultActivityHandler{}, // don't care
		"pull_request_delete_activity":                    DefaultActivityHandler{}, // don't care
		"story_create_activity":                           StoryActivityHandler{},   // care
		"story_delete_activity":                           StoryActivityHandler{},   // care
		"story_move_activity":                             StoryActivityHandler{},   // care
		"story_move_from_project_activity":                StoryActivityHandler{},   // care
		"story_move_into_project_activity":                StoryActivityHandler{},   // care
		"story_move_into_project_and_prioritize_activity": StoryActivityHandler{},   // care
		"story_update_activity":                           StoryActivityHandler{},   // care
		"task_create_activity":                            DefaultActivityHandler{}, // don't care
		"task_delete_activity":                            DefaultActivityHandler{}, // don't care
		"task_update_activity":                            DefaultActivityHandler{}, // don't care
	}
}

// ActivityHandler wraps the functionality to have a class factory for handling
// activity events from Pivotal Tracker
type ActivityHandler interface {
	Synchronize(*pivotal.Activity, *pivotal.Client, *jira.Client) error
}

// DefaultActivityHandler is a fill in for types we do not want to handle
type DefaultActivityHandler struct {
}

// Synchronize is the null handler function for synchronizing activities
func (d DefaultActivityHandler) Synchronize(activity *pivotal.Activity, pt *pivotal.Client, j *jira.Client) error {
	// Default activity is to do nothing
	return nil
}

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

// StoryActivityHandler is a fill in for types we do not want to handle
type StoryActivityHandler struct {
}

// Synchronize is the general handler function for synchronizing story activities
func (sa StoryActivityHandler) Synchronize(activity *pivotal.Activity, pt *pivotal.Client, j *jira.Client) error {
	switch kind := activity.Kind; kind {
	case "story_create_activity":
		// 1. Get story details from the activity
		// 2. Search JIRA for an existing story
		// 3. If no story exists, create a new story.
		// 4. Update external_id field of PT story.
		// 5. Create comment on story in Jira with PT story URL.
		return nil
	case "story_delete_activity":
		// 1. Get story details from the activity
		// 2. Search JIRA for an existing story
		// 3. Close the story with a comment about deletion.
		return nil
	case "story_move_activity":
		// don't care for now
		return nil
	case "story_move_from_project_activity":
		// don't care for now
		return nil
	case "story_move_into_project_activity":
		// don't care for now
		return nil
	case "story_move_into_project_and_prioritize_activity":
		// don't care for now
		return nil
	case "story_update_activity":
		// 1. Get story details from the activity
		// 2. Get the JIRA story ID from the external_id field.
		// 3. Do a check on the kind of update activity.
		// 4. Do a story update in Jira based on the kind of activity.
		// 5. Create comment on story in Jira with PT story URL.
		return nil
	default:
		return fmt.Errorf("unsupported kind of Story: this shouldn't happen")
	}
}
