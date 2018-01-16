package pivotal

import (
	"fmt"
	"time"
)

// Date is a date
type Date time.Time

// UnMarshallJSON unmarshall ALL THE THINGS
func (date *Date) UnMarshallJSON(content []byte) error {
	s := string(content)

	parsingError := func() error {
		return fmt.Errorf(
			"pivotal.Date.UnMarshallJSON: invalid date string: %s", content)
	}

	// Check whether the leading and trailing " is there.
	if len(s) < 2 || s[0] != '"' || s[len(s)-1] != '"' {
		return parsingError()
	}

	// Strip the leading and trailing "
	s = s[:len(s)-1][1:]

	// Parse the rest.
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return parsingError()
	}

	*date = Date(t)
	return nil
}

// MarshallJSON MARSHALL ALL THE THINGS
func (date Date) MarshallJSON() ([]byte, error) {
	return []byte((time.Time)(date).Format("2006-01-02")), nil
}
