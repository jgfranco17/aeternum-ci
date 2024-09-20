// Package github client provides functionality for managing GitHub artifacts and their access rights.
package githubclient

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"api/logger"

	"github.com/google/go-github/v56/github"
)

const (
	// github.TreeEntry.Mode valid values, refer to https://docs.github.com/en/rest/git/trees?apiVersion=2022-11-28#create-a-tree
	modeFile         string = "100644"
	modeExecutable   string = "100755"
	modeSubdirectory string = "040000"
	modeSubmodule    string = "160000"
	modeSymlinkPath  string = "120000"
	// github.TreeEntry.Type valid values, refer to https://docs.github.com/en/rest/git/trees?apiVersion=2022-11-28#create-a-tree
	typeBlob   string = "blob"
	typeTree   string = "tree"
	typeCommit string = "commit"
)

// Function Description: parse the provided file URL and return the required info
// [IN]: fileURL; the file URL to be parsed
// example for the fileURL: // "https://github.com/api/v3/repos/owner/repository-name/git/blobs/90c519f0118369a331035cd20c559a0e477384cb"
// [RETURN]: string; repo owner which is "some-user" in the above example
// [RETURN]: string; repo name which is "repository-name" in the above example
// [RETURN]: string; repo file SHA which is "90c519f0118369a331035cd20c559a0e477384cb" in the above example // pragma: allowlist secret
// [RETURN]: error; for error propagation
func parseFileURL(fileURL string) (string, string, string, error) {

	// The url should look like https://github.com/api/v3/repos/owner/repository-name/git/blobs/90c519f0118369a331035cd20c559a0e477384cb
	// 11 parts are expected:    0    1          2              3   4   5         6          7     8    9      10
	sections := strings.Split(fileURL, "/")
	if len(sections) < 11 {
		return "", "", "", fmt.Errorf("invalid Github file URL; missing fields")
	}
	repoOwner := sections[6]
	repo := sections[7]
	SHA := sections[10]

	return repoOwner, repo, SHA, nil
}

// Function Description: parse the provided repo URL and return the required info
// [IN]: repoUrl; the file URL to be parsed
// [RETURN]: string; repo owner which is "some-user" in the above example
// [RETURN]: string; repo name which is "repository-name" in the above example
// [RETURN]: string; branch name "if exist" which is "testBranch2" in the above example2
// [RETURN]: error; for error propagation
func parseRepoURL(repoURL string) (string, string, string, error) {

	urlRegex := regexp.MustCompile(`^https:\/\/([^\/]+)\/([^\/]+)\/([^\/]+)(\/tree\/(.+))?$`)
	matches := urlRegex.FindStringSubmatch(repoURL)

	if matches == nil {
		return "", "", "", fmt.Errorf("invalid url format")
	}
	if matches[5] == "" {
		// if no branch info, return main as the branch name
		matches[5] = "main"
	}

	return matches[2], matches[3], matches[5], nil
}

// wrapper interface for the used github package functions
type githubClient interface {
	ListCommits(ctx context.Context, owner string, repo string, opts *github.CommitsListOptions) ([]*github.RepositoryCommit, *github.Response, error)
	GetTree(ctx context.Context, owner string, repo string, sha string, recursive bool) (*github.Tree, *github.Response, error)
	GetBlobRaw(ctx context.Context, owner string, repo string, sha string) ([]byte, *github.Response, error)
	GetContents(ctx context.Context, owner string, repo string, path string, opts *github.RepositoryContentGetOptions) (fileContent *github.RepositoryContent, directoryContent []*github.RepositoryContent, resp *github.Response, err error)
	ListBranches(ctx context.Context, owner string, repo string, opts *github.BranchListOptions) ([]*github.Branch, *github.Response, error)
	Get(ctx context.Context, owner, repo string) (*github.Repository, *github.Response, error)
	CreateRef(ctx context.Context, owner string, repo string, ref *github.Reference) (*github.Reference, *github.Response, error)
	GetCommit(ctx context.Context, owner string, repo string, sha string) (*github.Commit, *github.Response, error)
	GetRef(ctx context.Context, owner string, repo string, ref string) (*github.Reference, *github.Response, error)
	CreateTree(ctx context.Context, owner string, repo string, baseTree string, entries []*github.TreeEntry) (*github.Tree, *github.Response, error)
	UpdateRef(ctx context.Context, owner string, repo string, ref *github.Reference, force bool) (*github.Reference, *github.Response, error)
	CreateCommit(ctx context.Context, owner string, repo string, commit *github.Commit, opts *github.CreateCommitOptions) (*github.Commit, *github.Response, error)
}

// Publicly exposed struct
type GithubService struct {
	client githubClient
}

// Function Description: get the contents of the provided test case URL
// [IN]: ctx; context
// [IN]: fileURL; the file URL for the required file, and it shall contain its SHA
// example for the fileURL: // "https://github.com/api/v3/repos/owner/repository-name/git/blobs/90c519f0118369a331035cd20c559a0e477384cb"
// [RETURN]: string; the file contents as a string
// [RETURN]: error; for error propagation
func (s *GithubService) GetFile(ctx context.Context, fileURL string) (string, string, error) {
	return getFile(ctx, s.client, fileURL)
}

// Function Description: get the contents of the provided file
// [IN]: ctx; context
// [IN]: repoUrl; the target repo URL
// example for the repoUrl: // "https://github.com/owner/repository-name"
// [IN]: filePath; the filePath inside the repo including its name
// [RETURN]: string; the file contents as a string
// [RETURN]: error; for error propagation
func (s *GithubService) GetFileLatest(ctx context.Context, repoURL, branchName, filePath string) (string, error) {
	return getFileLatest(ctx, s.client, repoURL+"/tree/"+branchName, filePath)
}

// Function Description: get the contents of the provided test case URL
// [IN]: ctx; context
// [IN]: repoUrl; the file URL to be parsed
// example for the repoUrl: // "https://github.com/owner/repository-name"
// [RETURN]: []GithubBranchesInfo; list of the branch info
// [RETURN]: error; for error propagation
func (s *GithubService) GetListOfBranches(ctx context.Context, repoURL string) ([]GithubBranchesInfo, error) {
	return getListOfBranches(ctx, s.client, repoURL)
}

// Function Description: get the default branch name for a specified repository
// [IN]: ctx; context
// [IN]: repoUrl; the file URL to be parsed
// [RETURN]: string; the name of the default branch
// [RETURN]: error; for error propagation
func (s *GithubService) GetDefaultBranchName(ctx context.Context, repoURL string) (string, error) {
	return getDefaultBranchName(ctx, s.client, repoURL)
}

// Function Description: create a branch on the specified repo with the specified branch name
// [IN]: ctx; context
// [IN]: repoUrl; the target repo URL
// example for the repoUrl: // "https://github.com/owner/repository-name"
// [IN]: branchName; the branch name to be created
// [RETURN]: branchSHA; the branch SHA of the created branch
// [RETURN]: branchURL; the URL of the created branch
// [RETURN]: error; for error propagation
func (s *GithubService) CreateBranch(ctx context.Context, repoUrl string, branchName string) (*GithubBranchesInfo, error) {
	return createBranch(ctx, s.client, repoUrl, branchName)
}

// Function Description: commit files changes on the specified repo/branch
// [IN]: ctx; context
// [IN]: repoUrl: the target repo
// [IN]: branchName; the target branch
// [IN]: commitMessage; message to be used with the github commit
// [IN]: fileChanges; context; map of the changed files, with the files paths as keys and the contents as values
// [RETURN]: commitSHA; the commit SHA of the created commit
// [RETURN]: commitURL; the URL of the created commit
// [RETURN]: error; for error propagation
func (s *GithubService) CommitMultipleFilesToBranch(ctx context.Context, repoUrl, branchName, commitMessage string, fileChanges map[string]string, filesToDelete []string) (*GithubBranchesInfo, error) {
	return commitMultipleFilesToBranch(ctx, s.client, repoUrl, branchName, commitMessage, fileChanges, filesToDelete)
}

// Returns a factory of GithubServices
// This one is mostly used for tests, where we can customize the client
type GithubServiceFactory func(ctx context.Context, token string, baseUrl string) (*GithubService, error)

func NewGithubServiceFactory(client githubClient) GithubServiceFactory {
	return func(ctx context.Context, token string, baseUrl string) (*GithubService, error) {
		return &GithubService{client: client}, nil
	}
}

// Returns a factory of GithubServices using the default github client
func DefaultGithubServiceFactory() GithubServiceFactory {
	return func(ctx context.Context, token string, baseUrl string) (*GithubService, error) {
		client, err := createClient(ctx, baseUrl, token)
		if err != nil {
			return nil, err
		}
		return &GithubService{client: client}, nil
	}
}

// Function Description: get the last edit date for the specified file relative to a specific file SHA
// [IN]: ctx; context
// [IN]: githubClient; an authenticated github client
// [IN]: repoOwner; the repo owner name; either user or organization
// [IN]: repo; the repo name
// [IN]: filePath; the filePath inside the repo including its name
// [IN]: fileSHA; the SHA of this specific file revision
// [RETURN]: time.Time; time object that provide the date/time information for the last edit date for the provided file
// [RETURN]: error; for error propagation
func getLastEditDateForFile(ctx context.Context, githubClient githubClient, repoOwner string, repo string, filePath string, fileSHA string) (time.Time, error) {
	log := logger.FromContext(ctx)
	log.Debug("Getting last commit date of test cases:", filePath)

	var lastUpdateDate time.Time

	// get commit list on this repo that involve changes on the specified filepath
	opts := &github.CommitsListOptions{Path: filePath}
	commitList, _, err := githubClient.ListCommits(ctx, repoOwner, repo, opts)
	if err != nil {
		return lastUpdateDate, fmt.Errorf("unable to get the commit list: %w", err)
	}

	for _, commit := range commitList {
		// get the git tree for each commit
		tree, _, err := githubClient.GetTree(ctx, repoOwner, repo, *commit.SHA, true)
		if err != nil {
			return lastUpdateDate, fmt.Errorf("unable to get the commit tree: %w", err)
		}
		// loop over the git tree
		for _, entry := range tree.Entries {
			// check if the provided filePath has the same SHA
			if entry.GetSHA() == fileSHA {
				// if yes, then get the date/time information of this commit
				lastUpdateDate = commit.Commit.Author.GetDate().Time
				// break the all outer loop
				return lastUpdateDate, nil
			}
		}
	}
	// exit the loop without finding the specified file/SHA, this shouldn't happen and means there is something wrong
	return lastUpdateDate, fmt.Errorf("unable to find the specified file/SHA")
}

// Function Description: get the contents of the provided test case URL
// [IN]: ctx; context
// [IN]: githubClient; an authenticated github client
// [IN]: fileURL; the file URL for the required file, and it shall contain its SHA
// example for the fileURL: // "https://github.com/api/v3/repos/owner/repository-name/git/blobs/90c519f0118369a331035cd20c559a0e477384cb"
// [RETURN]: string; the file contents as a string
// [RETURN]: string; the sha of the file
// [RETURN]: error; for error propagation
func getFile(ctx context.Context, githubClient githubClient, fileURL string) (string, string, error) {
	log := logger.FromContext(ctx)
	// parse the fileURL to get the required info
	repoOwner, repo, SHA, err := parseFileURL(fileURL)
	if err != nil {
		return "", "", fmt.Errorf("unable to parse the URL: %w", err)
	}
	log.Debugf("Loading file with: owner: %s, repo: %s, sha: %s", repoOwner, repo, SHA)

	// fetch the information of the provided fileURL
	fileContent, _, err := githubClient.GetBlobRaw(ctx, repoOwner, repo, SHA)
	if err != nil {
		return "", "", fmt.Errorf("unable to get the file contents: %w", err)
	}

	return string(fileContent), SHA, nil
}

// Function Description: get the contents of the provided test case URL
// [IN]: ctx; context
// [IN]: githubClient; an authenticated github client
// [IN]: repoUrl; the file URL to be parsed
// example for the repoUrl: // "https://github.com/owner/repository-name"
// [IN]: filePath; the filePath inside the repo including its name
// [RETURN]: string; the file contents as a string
// [RETURN]: error; for error propagation
func getFileLatest(ctx context.Context, githubClient githubClient, repoURL, filePath string) (string, error) {

	// parse the repoURL to get the required info
	repoOwner, repo, branchName, err := parseRepoURL(repoURL)
	if err != nil {
		return "", fmt.Errorf("unable to parse the URL: %w", err)
	}

	// fetch the information of the provided filePath
	fileContent, _, _, err := githubClient.GetContents(ctx, repoOwner, repo, filePath, &github.RepositoryContentGetOptions{
		Ref: branchName,
	})
	if err != nil {
		return "", fmt.Errorf("unable to get the file contents: %w", err)
	}

	content, err := fileContent.GetContent()
	if err != nil {
		return "", fmt.Errorf("unable to extract the file contents: %w", err)
	}

	return string(content), nil

}

// Function Description: get the contents of the provided test case URL
// [IN]: ctx; context
// [IN]: githubClient; an authenticated github client
// [IN]: repoUrl; the file URL to be parsed
// example for the repoUrl: // "https://github.com/owner/repository-name"
// [RETURN]: []GithubBranchesInfo; list of the branch info
// [RETURN]: error; for error propagation
func getListOfBranches(ctx context.Context, githubClient githubClient, repoURL string) ([]GithubBranchesInfo, error) {

	log := logger.FromContext(ctx)
	// parse the repoURL to get the required info
	repoOwner, repo, _, err := parseRepoURL(repoURL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse the URL: %w", err)
	}

	totalBranches := []*github.Branch{}
	requestOptions := &github.BranchListOptions{
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	}

	for {
		branches, response, err := githubClient.ListBranches(ctx, repoOwner, repo, requestOptions)

		if err != nil {
			return nil, fmt.Errorf("unable to get the branches list info: %w", err)
		}
		log.Debugf("# branches obtained in page %d is %d", requestOptions.Page, len(branches))
		totalBranches = append(totalBranches, branches...)

		if response.NextPage == 0 {
			break
		}
		requestOptions.Page = response.NextPage
	}

	log.Debugf("total branches obtained: %d", len(totalBranches))
	var branchesList []GithubBranchesInfo
	var tmp GithubBranchesInfo

	for _, branch := range totalBranches {
		tmp.Name = *branch.Name
		tmp.Uri = *branch.Commit.URL
		tmp.CommitSha = *branch.Commit.SHA
		branchesList = append(branchesList, tmp)
	}

	return branchesList, nil
}

// Function Description: get the default branch name for a specified repository
// [IN]: ctx; context
// [IN]: gitService; an authenticated github client
// [IN]: repoUrl; the file URL to be parsed
// [RETURN]: string; the name of the default branch
// [RETURN]: error; for error propagation
func getDefaultBranchName(ctx context.Context, githubClient githubClient, repoURL string) (string, error) {
	// parse the repoURL to get the required info
	repoOwner, repo, _, err := parseRepoURL(repoURL)
	if err != nil {
		return "", fmt.Errorf("unable to parse the URL: %w", err)
	}

	repositoryInfo, _, err := githubClient.Get(ctx, repoOwner, repo)
	if err != nil {
		return "", fmt.Errorf("unable to get the repository information: %w", err)
	}

	return *repositoryInfo.DefaultBranch, nil
}

// create branch on the specified repo with the provided name
func createBranch(ctx context.Context, githubClient githubClient, repoUrl string, branchName string) (*GithubBranchesInfo, error) {
	log := logger.FromContext(ctx)
	log.Debug("creating branch from:", repoUrl)

	// parse the commitHtmlUrl to get the required info
	repoOwner, repo, _, err := parseRepoURL(repoUrl)
	if err != nil {
		return nil, fmt.Errorf("unable to parse the URL: %w", err)
	}

	// Get the reference for the main.
	mainRef, _, err := githubClient.GetRef(ctx, repoOwner, repo, "refs/heads/main")
	if err != nil {
		return nil, fmt.Errorf("failed to get the repo reference: %w", err)
	}

	ref := &github.Reference{
		Ref: github.String("refs/heads/" + branchName),
		Object: &github.GitObject{
			SHA: github.String(mainRef.Object.GetSHA()),
		},
	}

	branchRef, _, err := githubClient.CreateRef(ctx, repoOwner, repo, ref)
	if err != nil {
		return nil, fmt.Errorf("failed to create new ref for the branch: %s; %w", branchName, err)
	}

	return &GithubBranchesInfo{
		CommitSha: *branchRef.Object.SHA,
		Uri:       *branchRef.Object.URL,
		Name:      branchName,
	}, nil
}

// commit a list of changes to a specific branch
func commitMultipleFilesToBranch(ctx context.Context, githubClient githubClient, repoUrl, branchName, commitMessage string, fileChanges map[string]string, filesToDelete []string) (*GithubBranchesInfo, error) {
	log := logger.FromContext(ctx)
	log.Debugf("commit change on repo %s - branch %s: ", repoUrl, branchName)

	// parse the branchUrl to get the required info
	repoOwner, repo, _, err := parseRepoURL(repoUrl)
	if err != nil {
		return nil, fmt.Errorf("unable to parse the URL: %w", err)
	}

	// Get the reference for the branch.
	branchRef, _, err := githubClient.GetRef(ctx, repoOwner, repo, "refs/heads/"+branchName)
	if err != nil {
		return nil, fmt.Errorf("failed to get the repo reference: %w", err)
	}

	// Get the latest commit for the branch.
	latestCommitSHA := branchRef.Object.GetSHA()
	latestCommit, _, err := githubClient.GetCommit(ctx, repoOwner, repo, latestCommitSHA)
	if err != nil {
		return nil, fmt.Errorf("failed to get the latest commit for the branch: %w", err)
	}

	// Create a tree with the updated files contents.
	updatedFilesTree := createTreeEntries(fileChanges, filesToDelete)

	tree, _, err := githubClient.CreateTree(ctx, repoOwner, repo, *latestCommit.Tree.SHA, updatedFilesTree)
	if err != nil {
		return nil, fmt.Errorf("failed to create a tree with the updated files contents: %w", err)
	}

	// Create a new commit based on the updated tree.
	commit, _, err := githubClient.CreateCommit(ctx, repoOwner, repo, &github.Commit{
		Parents: []*github.Commit{{SHA: &latestCommitSHA}},
		Tree:    tree,
		Message: github.String(commitMessage),
	}, &github.CreateCommitOptions{
		Signer: nil,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create a new commit based on the updated tree: %w", err)
	}

	// Update the branch reference to point to the new commit.
	ref := &github.Reference{
		Ref: github.String("refs/heads/" + branchName),
		Object: &github.GitObject{
			SHA: github.String(*commit.SHA),
		},
	}

	// Update the branch reference to point to the new commit
	_, _, err = githubClient.UpdateRef(ctx, repoOwner, repo, ref, false)
	if err != nil {
		return nil, fmt.Errorf("failed to update the branch reference to point to the new commit: %w", err)
	}
	return &GithubBranchesInfo{
		CommitSha: *commit.SHA,
		Uri:       *commit.HTMLURL, // FIXME: we need to use this uri for calling again the STF Backend should it be URL?
		Name:      branchName,
	}, nil
}

// createTreeEntries creates an array of tree entries for file changes.
func createTreeEntries(fileChanges map[string]string, filesToDelete []string) []*github.TreeEntry {
	entries := make([]*github.TreeEntry, 0)

	for path, content := range fileChanges {
		entries = append(entries, &github.TreeEntry{
			Path:    github.String(path),
			Mode:    github.String(modeFile),
			Type:    github.String(typeBlob),
			Content: github.String(content),
		})
	}

	for _, filePath := range filesToDelete {
		entries = append(entries, &github.TreeEntry{
			Path: github.String(filePath),
			Mode: github.String("100644"), // 100644 for file (blob)
			Type: github.String("blob"),
			// empty SHA & Content means delete the file
		})
	}

	return entries
}
