package version

import (
	"runtime"

	"github.com/go-openapi/runtime/middleware"
	"github.com/king-jam/tracker2jira/meta"
	"github.com/king-jam/tracker2jira/rest/models"
	"github.com/king-jam/tracker2jira/rest/server/operations/general"
)

// GetVersion ...
func GetVersion(params general.VersionParams) middleware.Responder {
	return &general.VersionOK{
		Payload: &models.Version{
			BuildDate:      meta.BuildDate,
			CommitHash:     meta.BuildCommitHash,
			ReleaseVersion: meta.BuildReleaseVersion,
			Runtime:        runtime.Version(),
		},
	}
}
