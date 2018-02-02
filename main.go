// Code generated by go-swagger; DO NOT EDIT.

package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	loads "github.com/go-openapi/loads"
	flag "github.com/spf13/pflag"

	"github.com/king-jam/tracker2jira/backend"
	"github.com/king-jam/tracker2jira/engine/taskservice"
	"github.com/king-jam/tracker2jira/rest/server"
	"github.com/king-jam/tracker2jira/rest/server/operations"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

// This file was generated by the swagger tool.
// Make sure not to overwrite this file after you generated it because all your edits would be lost!

func main() {
	// this initializes the global database object to be injected into all underlying components
	db, err := backend.InitializeDB(".")
	if err != nil {
		log.Fatalf("DB Init Failed: %+v\n", err)
	}

	// this creates the task scheduler service which monitors the DB for new
	// synchronization tasks and handles starting and monitoring tasks
	taskservice.NewTaskScheduler(db)

	swaggerSpec, err := loads.Analyzed(server.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	var httpServer *server.Server
	//var httpServer *server.Server // make sure init is called

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, "Usage:\n")
		fmt.Fprint(os.Stderr, "  t2j-server [OPTIONS]\n\n")

		title := "Tracker 2 JIRA"
		fmt.Fprint(os.Stderr, title+"\n\n")
		desc := "Pivotal Tracker to JIRA synchronization service."
		if desc != "" {
			fmt.Fprintf(os.Stderr, desc+"\n\n")
		}
		fmt.Fprintln(os.Stderr, flag.CommandLine.FlagUsages())
	}
	// parse the CLI flags
	flag.Parse()

	api := operations.NewT2jAPI(swaggerSpec)
	// get server with flag values filled out
	httpServer = server.NewServer(api)
	defer httpServer.Shutdown()

	httpServer.ConfigureAPIWithDependencies(db)

	if err := httpServer.Serve(); err != nil {
		log.Fatalln(err)
	}
}
