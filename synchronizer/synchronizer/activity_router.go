package synchronizer

import (
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
