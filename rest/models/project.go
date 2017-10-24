// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Project project
// swagger:model Project

type Project struct {

	// admin user ID
	AdminUserID string `json:"adminUserID,omitempty"`

	// external ID
	ExternalID string `json:"externalID,omitempty"`

	// project ID
	ProjectID string `json:"projectID,omitempty"`

	// project overrides
	ProjectOverrides interface{} `json:"projectOverrides,omitempty"`

	// project type
	ProjectType string `json:"projectType,omitempty"`

	// project URL
	ProjectURL string `json:"projectURL,omitempty"`

	// project version
	ProjectVersion int64 `json:"projectVersion,omitempty"`
}

/* polymorph Project adminUserID false */

/* polymorph Project externalID false */

/* polymorph Project projectID false */

/* polymorph Project projectOverrides false */

/* polymorph Project projectType false */

/* polymorph Project projectURL false */

/* polymorph Project projectVersion false */

// Validate validates this project
func (m *Project) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Project) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Project) UnmarshalBinary(b []byte) error {
	var res Project
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
