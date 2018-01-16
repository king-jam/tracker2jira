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
	ID                           int            `json:"id"`
	Name                         string         `json:"name"`
	Version                      int            `json:"version"`
	IterationLength              int            `json:"iteration_length"`
	WeekStartDay                 Day            `json:"week_start_day"`
	PointScale                   string         `json:"point_scale"`
	PointScaleIsCustom           bool           `json:"point_scale_is_custom"`
	BugsAndChoresAreEstimatable  bool           `json:"bugs_and_chores_are_estimatable"`
	AutomaticPlanning            bool           `json:"automatic_planning"`
	EnableTasks                  bool           `json:"enable_tasks"`
	StartDate                    *Date          `json:"start_date"`
	TimeZone                     *TimeZone      `json:"time_zone"`
	VelocityAveragedOver         int            `json:"velocity_averaged_over"`
	ShownIterationsStartTime     *time.Time     `json:"shown_iterations_start_time"`
	StartTime                    *time.Time     `json:"start_time"`
	NumberOfDoneIterationsToShow int            `json:"number_of_done_iterations_to_show"`
	HasGoogleDomain              bool           `json:"has_google_domain"`
	Description                  string         `json:"description"`
	ProfileContent               string         `json:"profile_content"`
	EnableIncomingEmails         bool           `json:"enable_incoming_emails"`
	InitialVelocity              int            `json:"initial_velocity"`
	ProjectType                  string         `json:"project_type"`
	Public                       bool           `json:"public"`
	AtomEnabled                  bool           `json:"atom_enabled"`
	CurrentIterationNumber       int            `json:"current_iteration_number"`
	CurrentVelocity              int            `json:"current_velocity"`
	CurrentVolatility            float64        `json:"current_volatility"`
	AccountID                    int            `json:"account_id"`
	AccountingType               AccountingType `json:"accounting_type"`
	Featured                     bool           `json:"featured"`
	StoryIds                     []int          `json:"story_ids"`
	EpicIds                      []int          `json:"epic_ids"`
	MembershipIds                []int          `json:"membership_ids"`
	LabelIds                     []int          `json:"label_ids"`
	IntegrationIds               []int          `json:"integration_ids"`
	IterationOverrideNumbers     []int          `json:"iteration_override_numbers"`
	CreatedAt                    *time.Time     `json:"created_at"`
	UpdatedAt                    *time.Time     `json:"updated_at"`
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
