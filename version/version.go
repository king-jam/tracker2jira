package version

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/king-jam/tracker2jira/rest/models"
	"github.com/king-jam/tracker2jira/rest/server/operations/general"
)

var (
	// BuildDate contains a string with the build date.
	BuildDate = "unknown"
	// CommitHash contains a string with the git commit hash.
	CommitHash = "unknown"
	// ReleaseVersion contains a string with the compiled release version.
	ReleaseVersion = "dev"
)

// Handler ...
func Handler(params general.VersionParams) middleware.Responder {
	return &general.VersionOK{
		Payload: &models.Version{
			BuildDate:      BuildDate,
			CommitHash:     CommitHash,
			ReleaseVersion: ReleaseVersion,
		},
	}
}
