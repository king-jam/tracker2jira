package cli

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	// BuildDate contains a string with the build date.
	BuildDate = "unknown"
	// CommitHash contains a string with the git commit hash.
	CommitHash = "unknown"
	// ReleaseVersion contains a string with the compiled release version.
	ReleaseVersion = "dev"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Long:  `Display version and build information about tracker2jira.`,
	Run: func(cli *cobra.Command, args []string) {
		fmt.Printf("tracker2jira %s\n", CommitHash)
		fmt.Printf("  Build date: %s\n", BuildDate)
		fmt.Printf("  Built with: %s\n", runtime.Version())
	},
}
