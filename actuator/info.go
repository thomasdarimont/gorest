package actuator

import (
	"net/http"
	"encoding/json"
	"os"
)

//provided by govvv at build-time
var GitCommit, BuildDate, Version string

type Info struct {
	CommitSha string `json: commit_sha`
	BuildDate string `json: build_date`
	Version string  `json: version`
	Hostname string `json: hostname`
}

func InfoEndpoint(w http.ResponseWriter, _ *http.Request) {

	hostname, _ := os.Hostname()
	json.NewEncoder(w).Encode(&Info{CommitSha: GitCommit, BuildDate: BuildDate, Version: Version, Hostname: hostname})
}
