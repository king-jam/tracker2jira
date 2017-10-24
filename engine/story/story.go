package story

// /projects/{project_id}/activity long polling using 'since_version'

// project_version =0
// for interval loop calls api ^^
//
// iterating through array of responses
// translate each array index object to corresponding jira api client object
// post object ^^
// if success //  bump/store/set project_version to api.get.project_version
// if failure retry exponential backoff ( plann error handling around jira going out to lunch)

// projects object needs to be creaded or found in client

// key into KIND

// interval for loop to call api
// get current project data ()
// call api with ?since_version=<value from project>
// counter = project version
// for range to read responses and handle posting to jira
