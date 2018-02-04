package synchronizer

import (
	"fmt"
	"io"

	"github.com/king-jam/tracker2jira/backend"
	"github.com/king-jam/tracker2jira/rest/models"
)

// ExternalCredentials ...
type ExternalCredentials struct {
	Jira    Jira
	Tracker Tracker
}

// Jira ...
type Jira struct {
	Password string
	Username string
}

// Tracker ..
type Tracker struct {
	Token string
}

// Synchronizer ...
type Synchronizer struct {
	db     backend.Database
	taskID string
}

// NewSynchronizer ...
func NewSynchronizer(db backend.Database, task *models.Task) *Synchronizer {
	s := new(Synchronizer)
	s.db = db
	s.taskID = task.TaskID.String()
	return s
}

// Run ...
func (s *Synchronizer) Run() error {
	// get the current state of the task from the DB
	dbTask, err := s.db.GetTaskByID(s.taskID)
	if err != nil {
		return err
	}
	// short circuit if it is a failing task
	// this is a bad method to short circuit, we should tell the runner to skip
	if dbTask.Status == models.TaskStatusFailed {
		return fmt.Errorf("task skipped: already in a failing state already")
	}
	// this is where the errors can really start
	pt, err := NewTrackerClient(s.db, dbTask.Master)
	if err != nil {
		return err
	}

	c, err := pt.UpdateCursor()
	if err != nil {
		return fmt.Errorf("task failed: unable to get Tracker cursor")
	}
	for {
		activity, err := c.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("task failed: unable to read activities")
		}
		fmt.Printf("\n\n DO ACTION: %+v \n\n", activity)
		project, err := s.db.GetProjectByID(dbTask.Master)
		if err != nil {
			return fmt.Errorf("failed to get the project")
		}
		project.ProjectVersion = int64(activity.ProjectVersion)
		_, err = s.db.PutProject(project)
		if err != nil {
			return fmt.Errorf("failed to put the project")
		}
	}
	//masterUser.ExternalCredentials.

	// slaveProject, err := s.db.GetProjectByID(dbTask.SLAVE)
	// if err != nil {
	// 	return err
	// }
	// //slaveUser, err := s.db.GetUserByID(slaveProject.AdminUserID)
	// s.db.GetUserByID(slaveProject.AdminUserID)

	// // /projects/{project_id}/activity long polling using 'since_version'
	//
	// // interval for loop to call api
	// // get current project data ()
	// // call api with ?since_version=<value from project.version ^^>
	// // for range to read responses and handle posting to jira
	// // iterating through array of responses
	// // translate each array index object to corresponding jira api client object
	// // post object ^^
	// // if success //  bump/store/set project_version to api.get.project_version ( write back to backend to update project version PUT PROJECT from backend)
	// // if failure retry exponential backoff ( plann error handling around jira going out to lunch)
	return nil
}

// ID ...
func (s *Synchronizer) ID() string {
	return s.taskID
}

// SetRunning ...
func (s *Synchronizer) SetRunning() error {
	task, err := s.db.GetTaskByID(s.taskID)
	if err != nil {
		return err
	}
	task.Status = models.TaskStatusRunning
	_, err = s.db.PutTask(task)
	return err
}

// SetFailed ...
func (s *Synchronizer) SetFailed() error {
	task, err := s.db.GetTaskByID(s.taskID)
	if err != nil {
		return err
	}
	task.Status = models.TaskStatusFailed
	_, err = s.db.PutTask(task)
	return err
}

// SetStopped ...
func (s *Synchronizer) SetStopped() error {
	task, err := s.db.GetTaskByID(s.taskID)
	if err != nil {
		return err
	}
	task.Status = models.TaskStatusStopped
	_, err = s.db.PutTask(task)
	return err
}
