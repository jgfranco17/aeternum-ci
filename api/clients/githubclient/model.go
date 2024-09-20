package githubclient

import "time"

type OriginInfo struct {
	Origin  string `json:"origin"`
	Version string `json:"version"`
}

// required git info for each test case
type TestCaseInfo struct {
	Name         string    // file name with its full path
	LastEditDate time.Time // last edit date/time for this file relative to the specified commit
	URL          string    // the file url including its SHA
}

// required git info for each branch
type GithubBranchesInfo struct {
	Name      string // branch name
	Uri       string // branch uri, example: https://github.tmc-stargate.com/arene-vertex/vertest/tree/testBranch2
	CommitSha string // the branch SHA
}

// feature file
type FeatureFile struct {
	FilePath *string
	Contents *string
}
