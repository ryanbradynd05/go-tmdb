package tmdb

import (
	"fmt"
)

// Job struct
type Job struct {
	Jobs []struct {
		Department string
		JobList    []string `json:"job_list"`
	}
}

// GetJobList gets a list of valid jobs
// https://developers.themoviedb.org/3/configuration/get-jobs
func (tmdb *TMDb) GetJobList() (*Job, error) {
	var jobList Job
	uri := fmt.Sprintf("%s/job/list?api_key=%s", baseURL, tmdb.apiKey)
	result, err := getTmdb(uri, &jobList)
	return result.(*Job), err
}
