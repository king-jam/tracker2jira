// nolint
package pivotal

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// PageLimit is ...Number of items to fetch at once when getting paginated response.
const PageLimit = 10

const (
	// StoryTypeFeature ...
	StoryTypeFeature = "feature"
	// StoryTypeBug ...
	StoryTypeBug = "bug"
	// StoryTypeChore ...
	StoryTypeChore = "chore"
	// StoryTypeRelease ...
	StoryTypeRelease = "release"
)

const (
	// StoryStateUnscheduled ...
	StoryStateUnscheduled = "unscheduled"
	// StoryStatePlanned ...
	StoryStatePlanned = "planned"
	// StoryStateUnstarted ..
	StoryStateUnstarted = "unstarted"
	// StoryStateStarted ...
	StoryStateStarted = "started"
	// StoryStateFinished ...
	StoryStateFinished = "finished"
	// StoryStateDelivered ...
	StoryStateDelivered = "delivered"
	// StoryStateAccepted ...
	StoryStateAccepted = "accepted"
	// StoryStateRejected ...
	StoryStateRejected = "rejected"
)

// Story is ..
type Story struct {
	ID            int        `json:"id,omitempty"`
	ProjectID     int        `json:"project_id,omitempty"`
	Name          string     `json:"name,omitempty"`
	Description   string     `json:"description,omitempty"`
	Type          string     `json:"story_type,omitempty"`
	State         string     `json:"current_state,omitempty"`
	Estimate      *float64   `json:"estimate,omitempty"`
	AcceptedAt    *time.Time `json:"accepted_at,omitempty"`
	Deadline      *time.Time `json:"deadline,omitempty"`
	RequestedByID int        `json:"requested_by_id,omitempty"`
	OwnerIds      []int      `json:"owner_ids,omitempty"`
	LabelIds      []int      `json:"label_ids,omitempty"`
	Labels        []*Label   `json:"labels,omitempty"`
	TaskIds       []int      `json:"task_ids,omitempty"`
	Tasks         []int      `json:"tasks,omitempty"`
	FollowerIds   []int      `json:"follower_ids,omitempty"`
	CommentIds    []int      `json:"comment_ids,omitempty"`
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
	IntegrationID int        `json:"integration_id,omitempty"`
	ExternalID    string     `json:"external_id,omitempty"`
	URL           string     `json:"url,omitempty"`
}

// StoryRequest is ..
type StoryRequest struct {
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Type        string    `json:"story_type,omitempty"`
	State       string    `json:"current_state,omitempty"`
	Estimate    *float64  `json:"estimate,omitempty"`
	OwnerIds    *[]int    `json:"owner_ids,omitempty"`
	LabelIds    *[]int    `json:"label_ids,omitempty"`
	Labels      *[]*Label `json:"labels,omitempty"`
	TaskIds     *[]int    `json:"task_ids,omitempty"`
	Tasks       *[]int    `json:"tasks,omitempty"`
	FollowerIds *[]int    `json:"follower_ids,omitempty"`
	CommentIds  *[]int    `json:"comment_ids,omitempty"`
}

// Label is
type Label struct {
	ID        int        `json:"id,omitempty"`
	ProjectID int        `json:"project_id,omitempty"`
	Name      string     `json:"name,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Kind      string     `json:"kind,omitempty"`
}

// Task is a task
type Task struct {
	ID          int        `json:"id,omitempty"`
	StoryID     int        `json:"story_id,omitempty"`
	Description string     `json:"description,omitempty"`
	Position    int        `json:"position,omitempty"`
	Complete    bool       `json:"complete,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

// Person is
type Person struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Initials string `json:"initials,omitempty"`
	Username string `json:"username,omitempty"`
	Kind     string `json:"kind,omitempty"`
}

// Comment is
type Comment struct {
	ID                  int        `json:"id,omitempty"`
	StoryID             int        `json:"story_id,omitempty"`
	EpicID              int        `json:"epic_id,omitempty"`
	PersonID            int        `json:"person_id,omitempty"`
	Text                string     `json:"text,omitempty"`
	FileAttachmentIds   []int      `json:"file_attachment_ids,omitempty"`
	GoogleAttachmentIds []int      `json:"google_attachment_ids,omitempty"`
	CommitType          string     `json:"commit_type,omitempty"`
	CommitIdentifier    string     `json:"commit_identifier,omitempty"`
	CreatedAt           *time.Time `json:"created_at,omitempty"`
	UpdatedAt           *time.Time `json:"updated_at,omitempty"`
}

type Blocker struct {
	Id          int        `json:"id,omitempty"`
	StoryId     int        `json:"story_id,omitempty"`
	PersonId    int        `json:"person_id,omitempty"`
	Description string     `json:"description,omitempty"`
	Resolved    bool       `json:"resolved,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type BlockerRequest struct {
	Description string `json:"description,omitempty"`
	Resolved    *bool  `json:"resolved,omitempty"`
}

// StoryService ...
type StoryService struct {
	client *Client
}

func newStoryService(client *Client) *StoryService {
	return &StoryService{client}
}

// List returns all stories matching the filter in case the filter is specified.
//
// List actually sends 2 HTTP requests - one to get the total number of stories,
// another to retrieve the stories using the right pagination setup. The reason
// for this is that the filter might require to fetch all the stories at once
// to get the right results. Since the response as generated by Pivotal Tracker
// is not always sorted when using a filter, this approach is required to get
// the right data. Not sure whether this is a bug or a feature.
func (service *StoryService) List(projectID int, filter string) ([]*Story, error) {
	reqFunc := newStoriesRequestFunc(service.client, projectID, filter)
	cursor, err := newCursor(service.client, reqFunc, 0)
	if err != nil {
		return nil, err
	}

	var stories []*Story
	if err := cursor.all(&stories); err != nil {
		return nil, err
	}
	return stories, nil
}

func newStoriesRequestFunc(client *Client, projectID int, filter string) func() *http.Request {
	return func() *http.Request {
		u := fmt.Sprintf("projects/%v/stories", projectID)
		if filter != "" {
			u += "?filter=" + url.QueryEscape(filter)
		}
		req, _ := client.NewRequest("GET", u, nil)
		return req
	}
}

// StoryCursor is
type StoryCursor struct {
	*cursor
	buff []*Story
}

// Next returns the next story.
//
// In case there are no more stories, io.EOF is returned as an error.
func (c *StoryCursor) Next() (s *Story, err error) {
	if len(c.buff) == 0 {
		_, err = c.next(&c.buff)
		if err != nil {
			return nil, err
		}
	}

	if len(c.buff) == 0 {
		err = io.EOF
	} else {
		s, c.buff = c.buff[0], c.buff[1:]
	}
	return s, err
}

// Iterate returns a cursor that can be used to iterate over the stories specified
// by the filter. More stories are fetched on demand as needed.
func (service *StoryService) Iterate(projectID int, filter string) (c *StoryCursor, err error) {
	reqFunc := newStoriesRequestFunc(service.client, projectID, filter)
	cursor, err := newCursor(service.client, reqFunc, PageLimit)
	if err != nil {
		return nil, err
	}
	return &StoryCursor{cursor, make([]*Story, 0)}, nil
}

// Create is ..
func (service *StoryService) Create(projectID int, story *StoryRequest) (*Story, *http.Response, error) {
	if projectID == 0 {
		return nil, nil, &ErrFieldNotSet{"project_id"}
	}

	if story.Name == "" {
		return nil, nil, &ErrFieldNotSet{"name"}
	}

	u := fmt.Sprintf("projects/%v/stories", projectID)
	req, err := service.client.NewRequest("POST", u, story)
	if err != nil {
		return nil, nil, err
	}

	var newStory Story

	resp, err := service.client.Do(req, &newStory)
	if err != nil {
		return nil, resp, err
	}

	return &newStory, resp, nil
}

// Get is ..
func (service *StoryService) Get(projectID, storyID int) (*Story, *http.Response, error) {
	u := fmt.Sprintf("projects/%v/stories/%v", projectID, storyID)
	req, err := service.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var story Story
	resp, err := service.client.Do(req, &story)
	if err != nil {
		return nil, resp, err
	}

	return &story, resp, err
}

// Update is ..
func (service *StoryService) Update(projectID, storyID int, story *StoryRequest) (*Story, *http.Response, error) {
	u := fmt.Sprintf("projects/%v/stories/%v", projectID, storyID)
	req, err := service.client.NewRequest("PUT", u, story)
	if err != nil {
		return nil, nil, err
	}

	var bodyStory Story
	resp, err := service.client.Do(req, &bodyStory)
	if err != nil {
		return nil, resp, err
	}

	return &bodyStory, resp, err

}

// ListTasks is ..
func (service *StoryService) ListTasks(projectID, storyID int) ([]*Task, *http.Response, error) {
	u := fmt.Sprintf("projects/%v/stories/%v/tasks", projectID, storyID)
	req, err := service.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var tasks []*Task
	resp, err := service.client.Do(req, &tasks)
	if err != nil {
		return nil, resp, err
	}

	return tasks, resp, err
}

// AddTask IS ..
func (service *StoryService) AddTask(projectID, storyID int, task *Task) (*http.Response, error) {
	if task.Description == "" {
		return nil, &ErrFieldNotSet{"description"}
	}

	u := fmt.Sprintf("projects/%v/stories/%v/tasks", projectID, storyID)
	req, err := service.client.NewRequest("POST", u, task)
	if err != nil {
		return nil, err
	}

	return service.client.Do(req, nil)
}

// ListOwners IS ..
func (service *StoryService) ListOwners(projectID, storyID int) ([]*Person, *http.Response, error) {
	u := fmt.Sprintf("projects/%d/stories/%d/owners", projectID, storyID)
	req, err := service.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var owners []*Person
	resp, err := service.client.Do(req, &owners)
	if err != nil {
		return nil, resp, err
	}

	return owners, resp, err
}

// AddComment is ..
func (service *StoryService) AddComment(
	projectID int,
	storyID int,
	comment *Comment,
) (*Comment, *http.Response, error) {

	u := fmt.Sprintf("projects/%v/stories/%v/comments", projectID, storyID)
	req, err := service.client.NewRequest("POST", u, comment)
	if err != nil {
		return nil, nil, err
	}

	var newComment Comment
	resp, err := service.client.Do(req, &newComment)
	if err != nil {
		return nil, resp, err
	}

	return &newComment, resp, err
}

// ListComments returns the list of Comments in a Story.
func (service *StoryService) ListComments(
	projectID int,
	storyID int,
) ([]*Comment, *http.Response, error) {

	u := fmt.Sprintf("projects/%v/stories/%v/comments", projectID, storyID)
	req, err := service.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var comments []*Comment
	resp, err := service.client.Do(req, &comments)
	if err != nil {
		return nil, resp, err
	}

	return comments, resp, nil
}

// ListBlockers returns the list of Blockers in a Story.
func (service *StoryService) ListBlockers(
	projectId int,
	storyId int,
) ([]*Blocker, *http.Response, error) {

	u := fmt.Sprintf("projects/%v/stories/%v/blockers", projectId, storyId)
	req, err := service.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var blockers []*Blocker
	resp, err := service.client.Do(req, &blockers)
	if err != nil {
		return nil, resp, err
	}

	return blockers, resp, nil
}

// AddBlocker ...
func (service *StoryService) AddBlocker(projectId int, storyId int, description string) (*Blocker, *http.Response, error) {
	u := fmt.Sprintf("projects/%v/stories/%v/blockers", projectId, storyId)
	req, err := service.client.NewRequest("POST", u, BlockerRequest{
		Description: description,
	})
	if err != nil {
		return nil, nil, err
	}

	var blocker Blocker
	resp, err := service.client.Do(req, &blocker)
	if err != nil {
		return nil, resp, err
	}

	return &blocker, resp, nil
}

// UpdateBlocker ...
func (service *StoryService) UpdateBlocker(projectId, stroyId, blockerId int, blocker *BlockerRequest) (*Blocker, *http.Response, error) {
	u := fmt.Sprintf("projects/%v/stories/%v/blockers/%v", projectId, stroyId, blockerId)
	req, err := service.client.NewRequest("PUT", u, blocker)
	if err != nil {
		return nil, nil, err
	}

	var blockerResp Blocker
	resp, err := service.client.Do(req, &blockerResp)
	if err != nil {
		return nil, resp, err
	}

	return &blockerResp, resp, nil
}
