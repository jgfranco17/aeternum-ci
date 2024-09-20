package githubclient

import (
	"context"
	"encoding/base64"
	"fmt"
	"testing"

	mock_githubclient "api/clients/githubclient/mock"

	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v56/github"
	"github.com/stretchr/testify/assert"
)

func TestGetFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock
	mockGithubClient := mock_githubclient.NewMockgithubClient(ctrl)

	fileURL := "https://github.com/api/v3/repos/some-user/my-project/git/blobs/90c519f0118369a331035cd20c559a0e477384cb"
	owner := "some-user"
	repo := "my-project"
	sha := "90c519f0118369a331035cd20c559a0e477384cb" // pragma: allowlist secret

	// Set up expectations for the mock client
	mockGithubClient.EXPECT().
		GetBlobRaw(gomock.Any(), owner, repo, sha).
		Return([]byte("this is the file content"), nil, nil)
	// Define the expected content
	expectedContent := "this is the file content"

	ctx := context.Background()
	content, sha2, _ := getFile(ctx, mockGithubClient, fileURL)
	// Assert that the returned content matches the expected content
	assert.Equal(t, expectedContent, content)
	assert.Equal(t, sha, sha2)
}

func TestGetFileLatest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock
	mockGithubClient := mock_githubclient.NewMockgithubClient(ctrl)

	repoUrl := "https://github.com/some-user/my-project"
	owner := "some-user"
	repo := "my-project"
	path := "/.woven/part_definitions.json"

	// Set up expectations for the mock client

	fileContent := github.RepositoryContent{
		Type:            new(string),
		Target:          new(string),
		Encoding:        new(string),
		Size:            new(int),
		Name:            new(string),
		Path:            new(string),
		Content:         new(string),
		SHA:             new(string),
		URL:             new(string),
		GitURL:          new(string),
		HTMLURL:         new(string),
		DownloadURL:     new(string),
		SubmoduleGitURL: new(string),
	}

	expectedContent := "this is test content"
	contentDecoded := base64.StdEncoding.EncodeToString([]byte(expectedContent))

	*fileContent.Type = "file"
	*fileContent.Encoding = "base64"
	*fileContent.Size = len(contentDecoded)
	*fileContent.Name = "part_definition.json"
	*fileContent.Content = string(contentDecoded)

	mockGithubClient.EXPECT().
		GetContents(gomock.Any(), owner, repo, path, gomock.Any()).
		Return(&fileContent, nil, nil, nil)

	ctx := context.Background()
	content, _ := getFileLatest(ctx, mockGithubClient, repoUrl, path)
	// Assert that the returned content matches the expected content
	assert.Equal(t, expectedContent, content)
}

func TestGetDefaultBranchNameSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock
	mockGithubClient := mock_githubclient.NewMockgithubClient(ctrl)

	repoUrl := "https://github.com/some-user/my-project"
	owner := "some-user"
	repo := "my-project"
	branchName := "main"

	repositoryContent := github.Repository{
		DefaultBranch: &branchName,
	}

	mockGithubClient.EXPECT().
		Get(gomock.Any(), owner, repo).
		Return(&repositoryContent, nil, nil)

	ctx := context.Background()
	actualBranchName, _ := getDefaultBranchName(ctx, mockGithubClient, repoUrl)
	// Assert that the returned content matches the expected content
	assert.Equal(t, actualBranchName, branchName)
}

func TestCommitMultipleFilesToBranch(t *testing.T) {
	t.Skip("Skipping testing in CI environment")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock
	mockGithubClient := mock_githubclient.NewMockgithubClient(ctrl)

	owner := "some-user"
	repo := "my-project"
	branchRef := "refs/heads/branch-name"
	latestCommitSHA := "0108e3c4f3100134a42fa333d103464498669ea5" // pragma: allowlist secret
	baseTreeSha := "0108e3c4f3100134a42fa444d103464498669ea5"     // pragma: allowlist secret
	newSha := "0108e3c4f3100134a42fa532d103464498669ea5"          // pragma: allowlist secret
	updatedFilesTree := []*github.TreeEntry{
		{
			Path:    github.String("contents/fileC.txt"),
			Mode:    github.String("100644"),
			Type:    github.String("blob"),
			Content: github.String("this is a test file#1"),
		},
		{
			Path:    github.String("contents/fileD.txt"),
			Mode:    github.String("100644"),
			Type:    github.String("blob"),
			Content: github.String("this is a test file#2"),
		},
		{
			Path: github.String("contents/fileA.txt"),
			Mode: github.String("100644"),
			Type: github.String("blob"),
		},
	}

	commitMsg := "this is a test commit"

	ref := github.Reference{
		Object: &github.GitObject{
			SHA: github.String(latestCommitSHA),
		},
	}
	mockGithubClient.EXPECT().
		GetRef(gomock.Any(), owner, repo, branchRef).
		Return(&ref, nil, nil)

	commit := github.Commit{
		SHA: github.String(baseTreeSha),
		Tree: &github.Tree{
			SHA: github.String(baseTreeSha),
			Entries: []*github.TreeEntry{
				{
					Path:    github.String("contents/fileA.txt"),
					Mode:    github.String("100644"),
					Type:    github.String("blob"),
					Content: github.String("this is existing file #1"),
				},
				{
					Path:    github.String("contents/fileB.txt"),
					Mode:    github.String("100644"),
					Type:    github.String("blob"),
					Content: github.String("this is existing file #2"),
				},
			},
			Truncated: new(bool),
		},
	}
	mockGithubClient.EXPECT().
		GetCommit(gomock.Any(), owner, repo, latestCommitSHA).
		Return(&commit, nil, nil)

	tree := github.Tree{
		SHA: github.String(baseTreeSha),
		Entries: []*github.TreeEntry{
			{
				Path:    github.String("contents/fileB.txt"),
				Mode:    github.String("100644"),
				Type:    github.String("blob"),
				Content: github.String("this is existing file #2"),
			},
			{
				Path:    github.String("contents/fileC.txt"),
				Mode:    github.String("100644"),
				Type:    github.String("blob"),
				Content: github.String("this is a test file#1"),
			},
			{
				Path:    github.String("contents/fileD.txt"),
				Mode:    github.String("100644"),
				Type:    github.String("blob"),
				Content: github.String("this is a test file#2"),
			},
		},
	}
	mockGithubClient.EXPECT().
		CreateTree(gomock.Any(), owner, repo, baseTreeSha, gomock.InAnyOrder(updatedFilesTree)).
		Return(&tree, nil, nil)

	newCommit := github.Commit{
		SHA:     github.String(newSha),
		Message: github.String(commitMsg),
		Tree:    &tree,
		HTMLURL: github.String("commit URL"),
	}
	mockGithubClient.EXPECT().
		CreateCommit(gomock.Any(), owner, repo,
			&github.Commit{
				Parents: []*github.Commit{{SHA: github.String(latestCommitSHA)}},
				Tree:    &tree,
				Message: github.String(commitMsg),
			},
			&github.CreateCommitOptions{
				Signer: nil,
			}).
		Return(&newCommit, nil, nil)

	commitRef := &github.Reference{
		Ref: github.String("refs/heads/branch-name"),
		Object: &github.GitObject{
			SHA: github.String(*newCommit.SHA),
		},
	}
	mockGithubClient.EXPECT().
		UpdateRef(gomock.Any(), owner, repo, commitRef, false).
		Return(nil, nil, nil)

	branchName := "branch-name"
	repoUrl := "https://github.com/some-user/my-project"
	fileChanges := make(map[string]string)
	fileChanges["contents/fileC.txt"] = "this is a test file#1"
	fileChanges["contents/fileD.txt"] = "this is a test file#2"
	filesToDelete := []string{"contents/fileA.txt"}

	ctx := context.Background()
	branchInfo, err := commitMultipleFilesToBranch(ctx, mockGithubClient, repoUrl, branchName, commitMsg, fileChanges, filesToDelete)

	// Assert that no errors occurred
	expectedSHA := newSha
	expectedURL := "commit URL"

	assert.Nil(t, err)
	assert.Equal(t, branchInfo.CommitSha, expectedSHA)
	assert.Equal(t, branchInfo.Uri, expectedURL)

}

func TestCreateBranch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock
	mockGithubClient := mock_githubclient.NewMockgithubClient(ctrl)

	owner := "some-user"
	repo := "my-project"
	latestCommitSHA := "0108e3c4f3100134a42fa333d103464498669ea5" // pragma: allowlist secret
	newCommitSHA := "0108e3c4f3100134a42fa444d103464498669ea5"    // pragma: allowlist secret

	mainRef := github.Reference{
		Object: &github.GitObject{
			SHA: github.String(latestCommitSHA),
		},
	}
	mockGithubClient.EXPECT().
		GetRef(gomock.Any(), owner, repo, "refs/heads/main").
		Return(&mainRef, nil, nil)

	newRef := &github.Reference{
		Ref: github.String("refs/heads/branch-name"),
		Object: &github.GitObject{
			SHA: github.String(latestCommitSHA),
		},
	}
	branchRef := github.Reference{
		Object: &github.GitObject{
			SHA: github.String(newCommitSHA),
			URL: github.String("commit URL"),
		},
	}
	mockGithubClient.EXPECT().
		CreateRef(gomock.Any(), owner, repo, newRef).
		Return(&branchRef, nil, nil)

	repoUrl := "https://github.com/some-user/my-project"
	branchName := "branch-name"

	ctx := context.Background()
	branchInfo, err := createBranch(ctx, mockGithubClient, repoUrl, branchName)

	// Assert that no errors occurred
	expectedSHA := newCommitSHA
	expectedURL := "commit URL"

	assert.Nil(t, err)
	assert.Equal(t, branchInfo.CommitSha, expectedSHA)
	assert.Equal(t, branchInfo.Uri, expectedURL)

}

func TestParseRepoURL(t *testing.T) {

	examples := []struct {
		description string
		repoUrl     string
		repoOwner   string
		repoName    string
		branchName  string
		parseError  error
	}{
		{
			description: "valid repo URL without branch info",
			repoUrl:     "https://github.com/some-user/my-project",
			repoOwner:   "some-user",
			repoName:    "my-project",
			branchName:  "main",
			parseError:  nil,
		},
		{
			description: "valid repo URL with branch info",
			repoUrl:     "https://github.com/some-user/my-project/tree/testBranch",
			repoOwner:   "some-user",
			repoName:    "my-project",
			branchName:  "testBranch",
			parseError:  nil,
		},
		{
			description: "valid repo URL with branch info that has forward slash",
			repoUrl:     "https://github.com/some-user/my-project/tree/ticketId/testBranch",
			repoOwner:   "some-user",
			repoName:    "my-project",
			branchName:  "ticketId/testBranch",
			parseError:  nil,
		},
		{
			description: "invalid repo URL in tree section",
			repoUrl:     "https://github.com/some-user/my-project/treeabc/testBranch",
			repoOwner:   "",
			repoName:    "",
			branchName:  "",
			parseError:  fmt.Errorf("invalid url format"),
		},
	}

	for _, example := range examples {
		t.Run(example.description, func(t *testing.T) {
			repoOwner, repoName, branchName, err := parseRepoURL(example.repoUrl)
			assert.Equal(t, example.repoOwner, repoOwner)
			assert.Equal(t, example.repoName, repoName)
			assert.Equal(t, example.branchName, branchName)
			assert.Equal(t, example.parseError, err)
		})
	}

}
