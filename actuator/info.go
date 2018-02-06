package actuator

import (
	"net/http"
	"encoding/json"
)

var GitCommit, BuildDate string

type Info struct {
	CommitSha string `json: commit_sha`
	BuildDate string `json: build_date`
}

func InfoEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(&Info{CommitSha: GitCommit, BuildDate: BuildDate})
}
