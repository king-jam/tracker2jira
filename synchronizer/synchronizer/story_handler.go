package synchronizer

import (
	"fmt"

	"github.com/king-jam/go-pivotaltracker/v5/pivotal"
	jira "gopkg.in/andygrunwald/go-jira.v1"
)

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
