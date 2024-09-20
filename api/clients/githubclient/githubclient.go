package githubclient

import (
	"context"
	"fmt"

	"api/logger"

	"github.com/google/go-github/v56/github"
)

type GithubClient struct {
	client *github.Client
}

func (Obj GithubClient) ListCommits(ctx context.Context, owner string, repo string, opts *github.CommitsListOptions) ([]*github.RepositoryCommit, *github.Response, error) {
	return Obj.client.Repositories.ListCommits(ctx, owner, repo, opts)
}

func (Obj GithubClient) GetTree(ctx context.Context, owner string, repo string, sha string, recursive bool) (*github.Tree, *github.Response, error) {
	return Obj.client.Git.GetTree(ctx, owner, repo, sha, recursive)
}

func (Obj GithubClient) GetBlobRaw(ctx context.Context, owner string, repo string, sha string) ([]byte, *github.Response, error) {
	return Obj.client.Git.GetBlobRaw(ctx, owner, repo, sha)
}

func (Obj GithubClient) GetContents(ctx context.Context, owner string, repo string, path string, opts *github.RepositoryContentGetOptions) (fileContent *github.RepositoryContent, directoryContent []*github.RepositoryContent, resp *github.Response, err error) {
	return Obj.client.Repositories.GetContents(ctx, owner, repo, path, opts)
}

func (Obj GithubClient) ListBranches(ctx context.Context, owner string, repo string, opts *github.BranchListOptions) ([]*github.Branch, *github.Response, error) {
	return Obj.client.Repositories.ListBranches(ctx, owner, repo, opts)
}

func (Obj GithubClient) Get(ctx context.Context, owner, repo string) (*github.Repository, *github.Response, error) {
	return Obj.client.Repositories.Get(ctx, owner, repo)
}

func (Obj GithubClient) CreateRef(ctx context.Context, owner string, repo string, ref *github.Reference) (*github.Reference, *github.Response, error) {
	return Obj.client.Git.CreateRef(ctx, owner, repo, ref)
}

func (Obj GithubClient) GetCommit(ctx context.Context, owner string, repo string, sha string) (*github.Commit, *github.Response, error) {
	return Obj.client.Git.GetCommit(ctx, owner, repo, sha)
}

func (Obj GithubClient) GetRef(ctx context.Context, owner string, repo string, ref string) (*github.Reference, *github.Response, error) {
	return Obj.client.Git.GetRef(ctx, owner, repo, ref)
}

func (Obj GithubClient) CreateTree(ctx context.Context, owner string, repo string, baseTree string, entries []*github.TreeEntry) (*github.Tree, *github.Response, error) {
	return Obj.client.Git.CreateTree(ctx, owner, repo, baseTree, entries)
}

func (Obj GithubClient) UpdateRef(ctx context.Context, owner string, repo string, ref *github.Reference, force bool) (*github.Reference, *github.Response, error) {
	return Obj.client.Git.UpdateRef(ctx, owner, repo, ref, force)
}

func (Obj GithubClient) CreateCommit(ctx context.Context, owner string, repo string, commit *github.Commit, opts *github.CreateCommitOptions) (*github.Commit, *github.Response, error) {
	return Obj.client.Git.CreateCommit(ctx, owner, repo, commit, opts)
}

// Function Description: create an authenticated client to the provided Github base URL
// [IN]: ctx; context
// [IN]: baseGithubURL; the base URL for the github repo; ex: "https://github.tmc-stargate.com/"
// [IN]: token; security access token with for the provided github repo
// [RETURN]: *github.Client; an authenticated configured github client
// [RETURN]: error; for error propagation
func createClient(ctx context.Context, baseGithubURL string, token string) (*GithubClient, error) {
	log := logger.FromContext(ctx)
	log.Debug("Creating github client")

	client, err := github.NewClient(nil).WithAuthToken(token).WithEnterpriseURLs(baseGithubURL, baseGithubURL)
	if err != nil {
		return nil, fmt.Errorf("unable to create authenticated github client: %w", err)
	}

	//check the github APIs rate limits for the authenticated user
	rateLimits, _, err := client.RateLimits(ctx)
	if err != nil {
		// no limits configured
	} else if rateLimits.Core.Limit != 0 {
		limitPercent := float32(rateLimits.Core.Remaining) / float32(rateLimits.Core.Limit)
		if limitPercent > 0.7 {
			log.Warn("Github Rate limit reached critical level:", rateLimits)
		}
	}

	return &GithubClient{client: client}, nil
}
