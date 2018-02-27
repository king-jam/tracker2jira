package synchronizer

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/king-jam/go-pivotaltracker/v5/pivotal"
	"github.com/trivago/tgo/tcontainer"
	jira "gopkg.in/andygrunwald/go-jira.v1"
)

// StoryActivityHandler is a fill in for types we do not want to handle
type StoryActivityHandler struct {
}

// Synchronize is the general handler function for synchronizing story activities
func (sa StoryActivityHandler) Synchronize(activity *pivotal.Activity, pt *pivotal.Client, j *jira.Client) error {
	switch kind := activity.Kind; kind {
	case "story_create_activity":
		if len(activity.Changes) > 1 {
			return fmt.Errorf("story handler: invalid create activities")
		}
		projectIDStrings := strings.Split(activity.GUID, "_")
		projectID, err := strconv.Atoi(projectIDStrings[0])
		if err != nil {
			return fmt.Errorf("story handler failed: %v", err)
		}
		storyDetails, _, err := pt.Stories.Get(projectID, activity.Changes[0].ID)
		if err != nil {
			return err
		}
		// TODO: this customfield is magic, we need to fix this from configuration
		jql := fmt.Sprintf("cf[10400]~%d", storyDetails.ID)
		issues, _, err := j.Issue.Search(jql, nil)
		if err != nil {
			return err
		}
		if len(issues) >= 1 {
			return fmt.Errorf("story already exists, cannot recreate story")
		}
		issueType := jira.IssueType{
			Name: "Story",
		}
		project := jira.Project{
			Key: "KATDEV",
		}
		tmap := tcontainer.NewMarshalMap()
		tmap["customfield_10401"] = storyDetails.URL
		tmap["customfield_10400"] = strconv.Itoa(storyDetails.ID)
		fields := jira.IssueFields{
			Type:        issueType,
			Summary:     storyDetails.Name,
			Project:     project,
			Description: storyDetails.Description,
			Unknowns:    tmap,
		}
		jiraIssue := jira.Issue{
			Fields: &fields,
		}
		_, _, err = j.Issue.Create(&jiraIssue)
		if err != nil {
			return err
		}
		// 1. Get story details from the activity
		// 2. Search JIRA for an existing story
		// 3. If no story exists, create a new story.
		// 4. Update external_id field of PT story.
		// 5. Create comment on story in Jira with PT story URL.
		return nil
	case "story_delete_activity":
		// TODO: this customfield is magic, we need to fix this from configuration
		jql := fmt.Sprintf("cf[10400]~%d", activity.Changes[0].ID)
		issues, _, err := j.Issue.Search(jql, nil)
		if err != nil {
			return err
		}
		if len(issues) > 1 {
			return fmt.Errorf("cannot delete story: not unique")
		} else if len(issues) == 1 {
			_, err = j.Issue.Delete(issues[0].ID)
			if err != nil {
				return err
			}
		}
		// 1. Get story details from the activity
		// 2. Search JIRA for an existing story
		// 3. Close the story with a comment about deletion.
		// will need to embed epic handling functionality for certain changes
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
		if len(activity.Changes) > 1 {
			return fmt.Errorf("story handler: invalid create activities")
		}
		projectIDStrings := strings.Split(activity.GUID, "_")
		projectID, err := strconv.Atoi(projectIDStrings[0])
		if err != nil {
			return fmt.Errorf("epic handler failed: %v", err)
		}
		storyDetails, _, err := pt.Stories.Get(projectID, activity.Changes[0].ID)
		if err != nil {
			return err
		}
		// TODO: this customfield is magic, we need to fix this from configuration
		jql := fmt.Sprintf("cf[10400]~%d", storyDetails.ID)
		issues, _, err := j.Issue.Search(jql, nil)
		if err != nil {
			return err
		}
		if len(issues) != 1 {
			return fmt.Errorf("story already exists, cannot recreate story")
		}
		issueType := jira.IssueType{
			Name: "Story",
		}
		project := jira.Project{
			Key: "KATDEV",
		}
		tmap := tcontainer.NewMarshalMap()
		tmap["customfield_10401"] = storyDetails.URL
		tmap["customfield_10400"] = strconv.Itoa(storyDetails.ID)
		fields := jira.IssueFields{
			Type:        issueType,
			Summary:     storyDetails.Name,
			Project:     project,
			Description: storyDetails.Description,
			Unknowns:    tmap,
		}
		jiraIssue := jira.Issue{
			Key:    issues[0].Key,
			ID:     issues[0].ID,
			Fields: &fields,
		}
		_, _, err = j.Issue.Update(&jiraIssue)
		if err != nil {
			return err
		}
		// 1. Get story details from the activity
		// 2. Get the JIRA story ID from the external_id field.
		// 3. Do a check on the kind of update activity.
		// 4. Do a story update in Jira based on the kind of activity.
		// 5. Create comment on story in Jira with PT story URL.
		// will need to embed epic handling functionality for certain changes
		return nil
	default:
		return fmt.Errorf("unsupported kind of Story: this shouldn't happen")
	}
}
