package synchronizer

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/king-jam/go-pivotaltracker/v5/pivotal"
	"github.com/trivago/tgo/tcontainer"
	jira "gopkg.in/andygrunwald/go-jira.v1"
)

// EpicActivityHandler is the struct to wrap all Epic related updates
type EpicActivityHandler struct {
}

// Synchronize is the general handler function for synchronizing epic activities
func (e EpicActivityHandler) Synchronize(activity *pivotal.Activity, pt *pivotal.Client, j *jira.Client) error {
	switch kind := activity.Kind; kind {
	case "epic_create_activity":
		if len(activity.Changes) > 1 {
			return fmt.Errorf("epic handler: invalid create activities")
		}
		projectIDStrings := strings.Split(activity.GUID, "_")
		projectID, err := strconv.Atoi(projectIDStrings[0])
		if err != nil {
			return fmt.Errorf("epic handler failed: %v", err)
		}
		epicDetails, _, err := pt.Epic.Get(projectID, activity.Changes[0].ID)
		if err != nil {
			return err
		}
		// TODO: this customfield is magic, we need to fix this from configuration
		jql := fmt.Sprintf("cf[10400]~%d", epicDetails.ID)
		issues, _, err := j.Issue.Search(jql, nil)
		if err != nil {
			return err
		}
		if len(issues) >= 1 {
			return fmt.Errorf("epic already exists, cannot recreate epic")
		}
		issueType := jira.IssueType{
			Name: "Epic",
		}
		project := jira.Project{
			Key: "KATDEV",
		}
		tmap := tcontainer.NewMarshalMap()
		tmap["customfield_10003"] = epicDetails.Name
		tmap["customfield_10401"] = epicDetails.URL
		tmap["customfield_10400"] = strconv.Itoa(epicDetails.ID)
		fields := jira.IssueFields{
			Type:        issueType,
			Summary:     epicDetails.Name,
			Project:     project,
			Description: epicDetails.Description,
			Unknowns:    tmap,
		}
		jiraIssue := jira.Issue{
			Fields: &fields,
		}
		_, _, err = j.Issue.Create(&jiraIssue)
		if err != nil {
			return err
		}
		return nil
		// https://docs.atlassian.com/software/jira/docs/api/REST/7.6.1/?_ga=2.262436809.497525053.1519328036-1921000638.1519328036#api/2/
		// {"fields":{"project":{"key": "TEST"},"customfield_10401": "Epic Name 01","summary": "REST EXAMPLE1","description": "Creating an Epic via REST","issuetype": {"name": "Epic"}}}
		// 1. Get the Epic details from the activity
		// 2. Search JIRA epics for the one matching the PT number.
		// 3. If not already exist, create the new Epic reference in JIRA
		// 4. Post Comment in JIRA Epic referencing Pivotal Tracker Epic
		// 5. Post Comment in Pivotal Tracker Epic referencing JIRA
	case "epic_delete_activity":
		// TODO: this customfield is magic, we need to fix this from configuration
		jql := fmt.Sprintf("cf[10400]~%d", activity.Changes[0].ID)
		issues, _, err := j.Issue.Search(jql, nil)
		if err != nil {
			return err
		}
		if len(issues) > 1 {
			return fmt.Errorf("cannot delete epic: not unique")
		} else if len(issues) == 1 {
			_, err = j.Issue.Delete(issues[0].ID)
			if err != nil {
				return err
			}
		}
		// 1. Get the Epic details from the activity
		// 2. Search JIRA epics for the one matching the PT number.
		// 3. Close the Epic reference in JIRA with a comment about why.
		return nil
	case "epic_update_activity":
		if len(activity.Changes) > 1 {
			return fmt.Errorf("epic handler: invalid create activities")
		}
		projectIDStrings := strings.Split(activity.GUID, "_")
		projectID, err := strconv.Atoi(projectIDStrings[0])
		if err != nil {
			return fmt.Errorf("epic handler failed: %v", err)
		}
		epicDetails, _, err := pt.Epic.Get(projectID, activity.Changes[0].ID)
		if err != nil {
			return err
		}
		// TODO: this customfield is magic, we need to fix this from configuration
		jql := fmt.Sprintf("cf[10400]~%d", epicDetails.ID)
		issues, _, err := j.Issue.Search(jql, nil)
		if err != nil {
			return err
		}
		if len(issues) != 1 {
			return fmt.Errorf("cannot update epic: doesn't exist or not unique")
		}
		issueType := jira.IssueType{
			Name: "Epic",
		}
		project := jira.Project{
			Key: "KATDEV",
		}
		tmap := tcontainer.NewMarshalMap()
		tmap["customfield_10003"] = epicDetails.Name
		tmap["customfield_10401"] = epicDetails.URL
		tmap["customfield_10400"] = strconv.Itoa(epicDetails.ID)
		fields := jira.IssueFields{
			Type:        issueType,
			Summary:     epicDetails.Name,
			Project:     project,
			Description: epicDetails.Description,
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
