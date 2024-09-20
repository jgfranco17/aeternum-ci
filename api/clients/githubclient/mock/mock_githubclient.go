// Code generated by MockGen. DO NOT EDIT.
// Source: .\githubservice.go

// Package mock_githubclient is a generated GoMock package.
package mock_githubclient

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/v56/github"
)

// MockgithubClient is a mock of githubClient interface.
type MockgithubClient struct {
	ctrl     *gomock.Controller
	recorder *MockgithubClientMockRecorder
}

// MockgithubClientMockRecorder is the mock recorder for MockgithubClient.
type MockgithubClientMockRecorder struct {
	mock *MockgithubClient
}

// NewMockgithubClient creates a new mock instance.
func NewMockgithubClient(ctrl *gomock.Controller) *MockgithubClient {
	mock := &MockgithubClient{ctrl: ctrl}
	mock.recorder = &MockgithubClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockgithubClient) EXPECT() *MockgithubClientMockRecorder {
	return m.recorder
}

// CreateCommit mocks base method.
func (m *MockgithubClient) CreateCommit(ctx context.Context, owner, repo string, commit *github.Commit, opts *github.CreateCommitOptions) (*github.Commit, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCommit", ctx, owner, repo, commit, opts)
	ret0, _ := ret[0].(*github.Commit)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateCommit indicates an expected call of CreateCommit.
func (mr *MockgithubClientMockRecorder) CreateCommit(ctx, owner, repo, commit, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCommit", reflect.TypeOf((*MockgithubClient)(nil).CreateCommit), ctx, owner, repo, commit, opts)
}

// CreateRef mocks base method.
func (m *MockgithubClient) CreateRef(ctx context.Context, owner, repo string, ref *github.Reference) (*github.Reference, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRef", ctx, owner, repo, ref)
	ret0, _ := ret[0].(*github.Reference)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateRef indicates an expected call of CreateRef.
func (mr *MockgithubClientMockRecorder) CreateRef(ctx, owner, repo, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRef", reflect.TypeOf((*MockgithubClient)(nil).CreateRef), ctx, owner, repo, ref)
}

// CreateTree mocks base method.
func (m *MockgithubClient) CreateTree(ctx context.Context, owner, repo, baseTree string, entries []*github.TreeEntry) (*github.Tree, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTree", ctx, owner, repo, baseTree, entries)
	ret0, _ := ret[0].(*github.Tree)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateTree indicates an expected call of CreateTree.
func (mr *MockgithubClientMockRecorder) CreateTree(ctx, owner, repo, baseTree, entries interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTree", reflect.TypeOf((*MockgithubClient)(nil).CreateTree), ctx, owner, repo, baseTree, entries)
}

// Get mocks base method.
func (m *MockgithubClient) Get(ctx context.Context, owner, repo string) (*github.Repository, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, owner, repo)
	ret0, _ := ret[0].(*github.Repository)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockgithubClientMockRecorder) Get(ctx, owner, repo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockgithubClient)(nil).Get), ctx, owner, repo)
}

// GetBlobRaw mocks base method.
func (m *MockgithubClient) GetBlobRaw(ctx context.Context, owner, repo, sha string) ([]byte, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlobRaw", ctx, owner, repo, sha)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetBlobRaw indicates an expected call of GetBlobRaw.
func (mr *MockgithubClientMockRecorder) GetBlobRaw(ctx, owner, repo, sha interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlobRaw", reflect.TypeOf((*MockgithubClient)(nil).GetBlobRaw), ctx, owner, repo, sha)
}

// GetCommit mocks base method.
func (m *MockgithubClient) GetCommit(ctx context.Context, owner, repo, sha string) (*github.Commit, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommit", ctx, owner, repo, sha)
	ret0, _ := ret[0].(*github.Commit)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCommit indicates an expected call of GetCommit.
func (mr *MockgithubClientMockRecorder) GetCommit(ctx, owner, repo, sha interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommit", reflect.TypeOf((*MockgithubClient)(nil).GetCommit), ctx, owner, repo, sha)
}

// GetContents mocks base method.
func (m *MockgithubClient) GetContents(ctx context.Context, owner, repo, path string, opts *github.RepositoryContentGetOptions) (*github.RepositoryContent, []*github.RepositoryContent, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContents", ctx, owner, repo, path, opts)
	ret0, _ := ret[0].(*github.RepositoryContent)
	ret1, _ := ret[1].([]*github.RepositoryContent)
	ret2, _ := ret[2].(*github.Response)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetContents indicates an expected call of GetContents.
func (mr *MockgithubClientMockRecorder) GetContents(ctx, owner, repo, path, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContents", reflect.TypeOf((*MockgithubClient)(nil).GetContents), ctx, owner, repo, path, opts)
}

// GetRef mocks base method.
func (m *MockgithubClient) GetRef(ctx context.Context, owner, repo, ref string) (*github.Reference, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRef", ctx, owner, repo, ref)
	ret0, _ := ret[0].(*github.Reference)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetRef indicates an expected call of GetRef.
func (mr *MockgithubClientMockRecorder) GetRef(ctx, owner, repo, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRef", reflect.TypeOf((*MockgithubClient)(nil).GetRef), ctx, owner, repo, ref)
}

// GetTree mocks base method.
func (m *MockgithubClient) GetTree(ctx context.Context, owner, repo, sha string, recursive bool) (*github.Tree, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTree", ctx, owner, repo, sha, recursive)
	ret0, _ := ret[0].(*github.Tree)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTree indicates an expected call of GetTree.
func (mr *MockgithubClientMockRecorder) GetTree(ctx, owner, repo, sha, recursive interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTree", reflect.TypeOf((*MockgithubClient)(nil).GetTree), ctx, owner, repo, sha, recursive)
}

// ListBranches mocks base method.
func (m *MockgithubClient) ListBranches(ctx context.Context, owner, repo string, opts *github.BranchListOptions) ([]*github.Branch, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBranches", ctx, owner, repo, opts)
	ret0, _ := ret[0].([]*github.Branch)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListBranches indicates an expected call of ListBranches.
func (mr *MockgithubClientMockRecorder) ListBranches(ctx, owner, repo, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBranches", reflect.TypeOf((*MockgithubClient)(nil).ListBranches), ctx, owner, repo, opts)
}

// ListCommits mocks base method.
func (m *MockgithubClient) ListCommits(ctx context.Context, owner, repo string, opts *github.CommitsListOptions) ([]*github.RepositoryCommit, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCommits", ctx, owner, repo, opts)
	ret0, _ := ret[0].([]*github.RepositoryCommit)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListCommits indicates an expected call of ListCommits.
func (mr *MockgithubClientMockRecorder) ListCommits(ctx, owner, repo, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCommits", reflect.TypeOf((*MockgithubClient)(nil).ListCommits), ctx, owner, repo, opts)
}

// UpdateRef mocks base method.
func (m *MockgithubClient) UpdateRef(ctx context.Context, owner, repo string, ref *github.Reference, force bool) (*github.Reference, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRef", ctx, owner, repo, ref, force)
	ret0, _ := ret[0].(*github.Reference)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateRef indicates an expected call of UpdateRef.
func (mr *MockgithubClientMockRecorder) UpdateRef(ctx, owner, repo, ref, force interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRef", reflect.TypeOf((*MockgithubClient)(nil).UpdateRef), ctx, owner, repo, ref, force)
}
