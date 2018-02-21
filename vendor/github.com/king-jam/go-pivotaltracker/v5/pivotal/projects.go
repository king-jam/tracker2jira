package pivotal

import (
	"fmt"
	"net/http"
	"time"
)

// Day is a dat
type Day string

const (
	// DayMonday is ..
	DayMonday Day = "Monday"
	// DayTuesday is ..
	DayTuesday Day = "Tuesday"
	// DayWednesday is ..
	DayWednesday Day = "Wednesday"
	// DayThursday is ..
	DayThursday Day = "Thursday"
	// DayFriday is ..
	DayFriday Day = "Friday"
	// DaySaturday is ..
	DaySaturday Day = "Saturday"
	// DaySunday is ..
	DaySunday Day = "Sunday"
)

const (
	// ProjectTypePublic is ..
	ProjectTypePublic = "public"
	// ProjectTypePrivate is ..
	ProjectTypePrivate = "private"
	// ProjectTypeDemo is ..
	ProjectTypeDemo = "demo"
)

// AccountingType is ..
type AccountingType string

const (
	// AccountingTypeUnbillable is ..
	AccountingTypeUnbillable AccountingType = "unbillable"
	// AccountingTypeBillable is ..
	AccountingTypeBillable AccountingType = "billable"
	// AccountingTypeOverhead is ..
	AccountingTypeOverhead AccountingType = "overhead"
)

// Project is a project
type Project struct {
	ID                           int            `json:"id,omitempty"`
	Name                         string         `json:"name,omitempty"`
	Version                      int            `json:"version,omitempty"`
	IterationLength              int            `json:"iteration_length,omitempty"`
	WeekStartDay                 Day            `json:"week_start_day,omitempty"`
	PointScale                   string         `json:"point_scale,omitempty"`
	PointScaleIsCustom           bool           `json:"point_scale_is_custom,omitempty"`
	BugsAndChoresAreEstimatable  bool           `json:"bugs_and_chores_are_estimatable,omitempty"`
	AutomaticPlanning            bool           `json:"automatic_planning,omitempty"`
	EnableTasks                  bool           `json:"enable_tasks,omitempty"`
	StartDate                    *Date          `json:"start_date,omitempty"`
	TimeZone                     *TimeZone      `json:"time_zone,omitempty"`
	VelocityAveragedOver         int            `json:"velocity_averaged_over,omitempty"`
	ShownIterationsStartTime     *time.Time     `json:"shown_iterations_start_time,omitempty"`
	StartTime                    *time.Time     `json:"start_time,omitempty"`
	NumberOfDoneIterationsToShow int            `json:"number_of_done_iterations_to_show,omitempty"`
	HasGoogleDomain              bool           `json:"has_google_domain,omitempty"`
	Description                  string         `json:"description,omitempty"`
	ProfileContent               string         `json:"profile_content,omitempty"`
	EnableIncomingEmails         bool           `json:"enable_incoming_emails,omitempty"`
	InitialVelocity              int            `json:"initial_velocity,omitempty"`
	ProjectType                  string         `json:"project_type,omitempty"`
	Public                       bool           `json:"public,omitempty"`
	AtomEnabled                  bool           `json:"atom_enabled,omitempty"`
	CurrentIterationNumber       int            `json:"current_iteration_number,omitempty"`
	CurrentVelocity              int            `json:"current_velocity,omitempty"`
	CurrentVolatility            float64        `json:"current_volatility,omitempty"`
	AccountID                    int            `json:"account_id,omitempty"`
	AccountingType               AccountingType `json:"accounting_type,omitempty"`
	Featured                     bool           `json:"featured,omitempty"`
	StoryIds                     []int          `json:"story_ids,omitempty"`
	EpicIds                      []int          `json:"epic_ids,omitempty"`
	MembershipIds                []int          `json:"membership_ids,omitempty"`
	LabelIds                     []int          `json:"label_ids,omitempty"`
	IntegrationIds               []int          `json:"integration_ids,omitempty"`
	IterationOverrideNumbers     []int          `json:"iteration_override_numbers,omitempty"`
	CreatedAt                    *time.Time     `json:"created_at,omitempty"`
	UpdatedAt                    *time.Time     `json:"updated_at,omitempty"`
	Kind                         string         `json:"kind,omitempty"`
}

// ProjectService is ..
type ProjectService struct {
	client *Client
}

func newProjectService(client *Client) *ProjectService {
	return &ProjectService{client}
}

// List returns all active projects for the current user.
func (service *ProjectService) List() ([]*Project, *http.Response, error) {
	req, err := service.client.NewRequest("GET", "projects", nil)
	if err != nil {
		return nil, nil, err
	}

	var projects []*Project
	resp, err := service.client.Do(req, &projects)
	if err != nil {
		return nil, resp, err
	}

	return projects, resp, err
}

// Get is ..
func (service *ProjectService) Get(projectID int) (*Project, *http.Response, error) {
	u := fmt.Sprintf("projects/%v", projectID)
	req, err := service.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var project Project
	resp, err := service.client.Do(req, &project)
	if err != nil {
		return nil, resp, err
	}

	return &project, resp, err
}
