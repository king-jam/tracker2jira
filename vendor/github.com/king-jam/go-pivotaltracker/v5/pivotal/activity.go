package pivotal

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Changes is ...
type Changes struct {
	Kind           string      `json:"kind,omitempty"`
	GUID           string      `json:"id,omitempty"`
	Name           string      `json:"name,omitempty"`
	ChangeType     string      `json:"change_type,omitempty"`
	StoryType      string      `json:"story_type,omitempty"`
	OriginalValues interface{} `json:"original_values,omitempty"`
	NewValues      interface{} `json:"new_values,omitempty"`
	URL            string      `json:"url,omitempty"`
}

// Activity is ...
type Activity struct {
	Kind               string        `json:"kind,omitempty"`
	GUID               string        `json:"guid,omitempty"`
	ProjectVersion     int           `json:"project_version,omitempty"`
	Message            string        `json:"message,omitempty"`
	Highlight          string        `json:"highlight,omitempty"`
	Changes            []Changes     `json:"changes,omitempty"`
	PrimaryResources   []interface{} `json:"primary_resources,omitempty"`
	SecondaryResources []interface{} `json:"secondary_resources,omitempty"`
	Project            Project       `json:"project,omitempty"`
	PerformedBy        Person        `json:"performed_by,omitempty"`
	OccurredAt         time.Time     `json:"occurred_at,omitempty"`
}

// ActivityService is ...
type ActivityService struct {
	client *Client
}

func newActivitiesService(client *Client) *ActivityService {
	return &ActivityService{client}
}

// List returns all activities matching the filter in case the filter is specified.
//
// List actually sends 2 HTTP requests - one to get the total number of activities,
// another to retrieve the activities using the right pagination setup. The reason
// for this is that the filter might require to fetch all the activities at once
// to get the right results. The response is default sorted in DESCENDING order so
// leverage the sortAsc variable to control sort order.
func (service *ActivityService) List(projectID int, sortOrder *string, limit *int, offset *int, occurredBefore *time.Time, occurredAfter *time.Time, sinceVersion *int) ([]*Activity, error) {
	reqFunc := newActivitiesRequestFunc(service.client, projectID, sortOrder, limit, offset, occurredBefore, occurredAfter, sinceVersion)
	cursor, err := newCursor(service.client, reqFunc, 0)
	if err != nil {
		return nil, err
	}

	var activities []*Activity
	if err := cursor.all(&activities); err != nil {
		return nil, err
	}
	return activities, nil
}

// newActivitiesRequestFunc takes in pointers to a bunch of types, there reason for this is so we can pass in nil values and create a query string accordingly
// this could be wrapped a different way to accomplish a similar goal but the nil value is the desired behavior
func newActivitiesRequestFunc(client *Client, projectID int, sortOrder *string, limit *int, offset *int, occurredBefore *time.Time, occurredAfter *time.Time, sinceVersion *int) func() *http.Request {
	return func() *http.Request {
		activityPath := fmt.Sprintf("projects/%v/activity", projectID)
		queryParams := url.Values{}
		if sortOrder != nil {
			queryParams.Add("sort_order", url.QueryEscape(*sortOrder))
		}
		if limit != nil {
			queryParams.Add("limit", url.QueryEscape(strconv.Itoa(*limit)))
		}
		if offset != nil {
			queryParams.Add("offset", url.QueryEscape(strconv.Itoa(*offset)))
		}
		if occurredBefore != nil {
			queryParams.Add("occurred_before", url.QueryEscape(occurredBefore.String()))
		}
		if occurredAfter != nil {
			queryParams.Add("occurred_after", url.QueryEscape(occurredAfter.String()))
		}
		if sinceVersion != nil {
			queryParams.Add("since_version", url.QueryEscape(strconv.Itoa(*sinceVersion)))
		}
		if len(queryParams) > 0 {
			activityPath += "?"
			activityPath += queryParams.Encode()
		}
		req, _ := client.NewRequest("GET", activityPath, nil)
		return req
	}
}

// ActivityCursor is...
type ActivityCursor struct {
	*cursor
	buff []*Activity
}

// Next returns the next story.
//
// In case there are no more stories, io.EOF is returned as an error.
func (c *ActivityCursor) Next() (s *Activity, err error) {
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

// Iterate returns a cursor that can be used to iterate over the activities specified
// by the filter. More stories are fetched on demand as needed.
func (service *ActivityService) Iterate(projectID int, sortOrder *string, limit *int, offset *int, occurredBefore *time.Time, occurredAfter *time.Time, sinceVersion *int) (c *ActivityCursor, err error) {
	reqFunc := newActivitiesRequestFunc(service.client, projectID, sortOrder, limit, offset, occurredBefore, occurredAfter, sinceVersion)
	cursor, err := newCursor(service.client, reqFunc, PageLimit)
	if err != nil {
		return nil, err
	}
	return &ActivityCursor{cursor, make([]*Activity, 0)}, nil
}

func (service *ActivityService) validateSortOrder(order string) error {
	validValues := []string{"asc", "desc"}
	for _, value := range validValues {
		if value == order {
			return nil
		}
	}
	return fmt.Errorf("%s is not a valid sort_order", order)
}
