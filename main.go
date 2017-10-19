package main

import (
	"fmt"
	"os"

	"github.com/king-jam/tracker2jira/cli"
)

func main() {
	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v\n", err)
		os.Exit(1)
	}
}
