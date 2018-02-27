package synchronizer

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/king-jam/go-pivotaltracker/v5/pivotal"
	"github.com/king-jam/tracker2jira/backend"
	"github.com/king-jam/tracker2jira/rest/models"
	log "github.com/sirupsen/logrus"
	jira "gopkg.in/andygrunwald/go-jira.v1"
)

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
	srcProject, err := s.db.GetProjectByID(dbTask.Source)
	if err != nil {
		return err
	}
	// Get the source Project Username/Password Details
	sourceUser, err := s.db.GetUserByID(srcProject.AdminUserID)
	if err != nil {
		return err
	}
	// create a new pivotal tracker client
	ptclient := pivotal.NewClient(sourceUser.ExternalCredentials.Token.String())
	trackerProjectID, err := strconv.Atoi(srcProject.ExternalID)
	if err != nil {
		return err
	}
	// this is our default sort order
	sortOrder := "asc"
	// if this is the first run, initialize a default version state
	if dbTask.LastSynchronizedVersion == 0 {
		tSortOrder := "desc"
		var createTime time.Time
		createTime, err = time.Parse(time.RFC3339, dbTask.CreatedAt.String())
		if err != nil {
			return err
		}
		var tCursor *pivotal.ActivityCursor
		tCursor, err = ptclient.Activity.Iterate(trackerProjectID, &tSortOrder, &createTime, nil, nil)
		if err != nil {
			return fmt.Errorf("task failed: unable to get Tracker cursor")
		}
		var tActivity *pivotal.Activity
		tActivity, err = tCursor.Next()
		if err != nil {
			return fmt.Errorf("task failed: unable to read activities")
		}
		task, err := s.db.GetTaskByID(dbTask.TaskID.String())
		if err != nil {
			return err
		}
		task.LastSynchronizedVersion = int64(tActivity.ProjectVersion)
		_, err = s.db.PutTask(task)
		if err != nil {
			return fmt.Errorf("task failed: failed to put the project")
		}
	}
	// create the activity iterator based on the synchronization version
	currVersion := int(dbTask.LastSynchronizedVersion)
	c, err := ptclient.Activity.Iterate(trackerProjectID, &sortOrder, nil, nil, &currVersion)
	if err != nil {
		return fmt.Errorf("task failed: unable to get Tracker cursor")
	}
	// get the destination project details
	dstProject, err := s.db.GetProjectByID(dbTask.Destination)
	if err != nil {
		return err
	}
	// Get the destination Project Username/Password Details
	dstUser, err := s.db.GetUserByID(dstProject.AdminUserID)
	if err != nil {
		return err
	}
	// create a new empty JIRA client
	j, err := jira.NewClient(nil, dstProject.ProjectURL)
	if err != nil {
		return err
	}
	// setup the authentication for jira
	j.Authentication.SetBasicAuth(dstUser.ExternalCredentials.Username, dstUser.ExternalCredentials.Password.String())
	authenticated, err := j.Authentication.AcquireSessionCookie(dstUser.ExternalCredentials.Username, dstUser.ExternalCredentials.Password.String())
	if err != nil {
		return err
	}
	if !authenticated {
		return fmt.Errorf("failure to authenticate to Jira")
	}
	for {
		activity, err := c.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("task failed: unable to read activities")
		}
		log.Debugf("\n%+v\n", activity.Kind)
		handler, exist := activityHandlers[activity.Kind]
		if !exist {
			return fmt.Errorf("update failed: no valid handler for activity type")
		}
		log.Debugf("\n%+v\v", activity)
		err = handler.Synchronize(activity, ptclient, j)
		if err != nil {
			return fmt.Errorf("update activity failed in handler function: %v", err)
		}
		// update the dbTask version
		// get a new version just in case
		task, err := s.db.GetTaskByID(dbTask.TaskID.String())
		if err != nil {
			return err
		}
		task.LastSynchronizedVersion = int64(activity.ProjectVersion)
		_, err = s.db.PutTask(task)
		if err != nil {
			return err
		}
	}
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
