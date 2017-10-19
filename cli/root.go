// Package cli handles the command-line interface for tracker2jira.
package cli

import (
	"github.com/spf13/cobra"
)

// RootCmd is the root for all hello commands.
var RootCmd = &cobra.Command{
	Use:           "tracker2jira",
	Short:         "start me",
	Long:          `stuff and things.`,
	SilenceErrors: true,
}
