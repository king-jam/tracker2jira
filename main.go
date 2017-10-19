package main

import (
	"fmt"
	"os"

	"github.com/king-jam/tracker2jira/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v\n", err)
		os.Exit(1)
	}
}
