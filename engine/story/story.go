package story

import (
	"encoding/json"
	"log"
	"time"

	"github.com/king-jam/go-pivotaltracker/v5/pivotal"
	"github.com/king-jam/tracker2jira/backend"
)

// /projects/{project_id}/activity long polling using 'since_version'

// interval for loop to call api
// get current project data ()
// call api with ?since_version=<value from project.version ^^>
// for range to read responses and handle posting to jira
// iterating through array of responses
// translate each array index object to corresponding jira api client object
// post object ^^
// if success //  bump/store/set project_version to api.get.project_version ( write back to backend to update project version PUT PROJECT from backend)
// if failure retry exponential backoff ( plann error handling around jira going out to lunch)

type ExternalCredentials struct {
	Jira    Jira
	Tracker Tracker
}

type Jira struct {
	Password string
	Username string
}

type Tracker struct {
	Token string
}

var pollTime = 60 * time.Second

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func pollAPI(t time.Time) {
	log.Printf("Polling: %v", t)
	b, _ := backend.GetDB()
	project, err := b.GetProjectByID("SomeString")
	if err != nil {
		// return err
		log.Println("Cant get project")
	}

	user, err := b.GetUserByID("someUserID")

	// credMap := make(map[string]map[string]string)
	// for k,v := range credMap {
	//     if v == "tracker" {
	//         tCred :=
	//         for k,v := range
	//     }
	// }

	data, err := json.Marshal(user.ExternalCredentials)
	externalCreds := ExternalCredentials{}
	json.Unmarshal(data, &externalCreds)

	client := pivotal.NewClient(externalCreds.Tracker.Token)
	client.Do(req, v)

}

func poll() {
	doEvery(pollTime, pollAPI)
}
