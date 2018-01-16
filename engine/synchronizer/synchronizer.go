package synchronizer

import (
	"fmt"

	"github.com/andygrunwald/go-jira"
	"github.com/king-jam/go-pivotaltracker/v5/pivotal"
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
	db         *backend.Backend
	ptclient   *pivotal.Client
	jiraclient *jira.Client
	taskID     string
}

// NewSynchronizer ...
func NewSynchronizer(db *backend.Backend, task *models.Task) *Synchronizer {
	s := new(Synchronizer)
	s.db = db
	s.taskID = task.TaskID.String()
	// TODO: Add client creation / validation.
	return s
}

// Run ...
func (s *Synchronizer) Run() error {
	// Get the task with a current state value to determine what we should do
	dbTask, err := s.db.GetTaskByID(s.taskID)
	if err != nil {
		return err
	}
	// short circuit if it is a failing task
	// this is a bad method to short circuit, we should tell the runner to skip
	if dbTask.Status == models.TaskStatusFailed {
		return fmt.Errorf("task skipped: already in a failing state already")
	}
	// get the master project which should always be Pivotal Tracker in our case
	masterProject, err := s.db.GetProjectByID(dbTask.Master)
	if err != nil {
		return err
	}
	// Get the master Project Username/Password Details
	////masterUser, err := s.db.GetUserByID(masterProject.AdminUserID)
	_, err = s.db.GetUserByID(masterProject.AdminUserID)
	if err != nil {
		return err
	}
	//masterUser.ExternalCredentials.

	slaveProject, err := s.db.GetProjectByID(dbTask.SLAVE)
	if err != nil {
		return err
	}
	//slaveUser, err := s.db.GetUserByID(slaveProject.AdminUserID)
	s.db.GetUserByID(slaveProject.AdminUserID)

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
	fmt.Printf("HELLO %s\n", s.ID())
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

func (s *Synchronizer) getTrackerClient() {

}
