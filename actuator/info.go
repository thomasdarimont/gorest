package actuator

import (
	"net/http"
	"encoding/json"
)

//provided by govvv at build-time
var GitCommit, BuildDate, Version string

type Info struct {
	CommitSha string `json: commit_sha`
	BuildDate string `json: build_date`
	Version string  `json: version`
}

func InfoEndpoint(w http.ResponseWriter, _ *http.Request) {
	json.NewEncoder(w).Encode(&Info{CommitSha: GitCommit, BuildDate: BuildDate, Version: Version})
}
