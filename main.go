package main

import (
	"log"
	"os"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"
	"github.com/king-jam/tracker2jira/backend"
	"github.com/king-jam/tracker2jira/rest/server"
	"github.com/king-jam/tracker2jira/rest/server/operations"
)

func main() {
	// startup services/workers
	if backend.ConfigureDB() != nil {
		log.Fatal("DB Init Failed")
	}

	// startup persistence
	StartAPI()
}

// StartAPI does things
func StartAPI() {
	swaggerSpec, err := loads.Analyzed(server.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	var server *server.Server // make sure init is called

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
	server = server.NewServer(api)

	defer server.Shutdown()

	server.ConfigureAPI()
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
