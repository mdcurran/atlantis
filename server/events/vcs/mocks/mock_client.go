// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events/vcs (interfaces: Client)

package mocks

import (
	pegomock "github.com/petergtz/pegomock"
	models "github.com/runatlantis/atlantis/server/events/models"
	"reflect"
	"time"
)

type MockClient struct {
	fail func(message string, callerSkip ...int)
}

func NewMockClient(options ...pegomock.Option) *MockClient {
	mock := &MockClient{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockClient) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockClient) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockClient) GetModifiedFiles(repo models.Repo, pull models.PullRequest) ([]string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{repo, pull}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetModifiedFiles", params, []reflect.Type{reflect.TypeOf((*[]string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 []string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClient) CreateComment(repo models.Repo, pullNum int, comment string) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{repo, pullNum, comment}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CreateComment", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockClient) PullIsApproved(repo models.Repo, pull models.PullRequest) (bool, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{repo, pull}
	result := pegomock.GetGenericMockFrom(mock).Invoke("PullIsApproved", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 bool
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClient) PullIsMergeable(repo models.Repo, pull models.PullRequest) (bool, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{repo, pull}
	result := pegomock.GetGenericMockFrom(mock).Invoke("PullIsMergeable", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 bool
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClient) UpdateStatus(repo models.Repo, pull models.PullRequest, state models.CommitStatus, src string, description string, url string) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{repo, pull, state, src, description, url}
	result := pegomock.GetGenericMockFrom(mock).Invoke("UpdateStatus", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockClient) MergePull(pull models.PullRequest) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClient().")
	}
	params := []pegomock.Param{pull}
	result := pegomock.GetGenericMockFrom(mock).Invoke("MergePull", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockClient) UpdateLabels(pull models.PullRequest) error {
	return nil
}

func (mock *MockClient) VerifyWasCalledOnce() *VerifierMockClient {
	return &VerifierMockClient{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockClient) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierMockClient {
	return &VerifierMockClient{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockClient) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierMockClient {
	return &VerifierMockClient{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockClient) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierMockClient {
	return &VerifierMockClient{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockClient struct {
	mock                   *MockClient
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockClient) GetModifiedFiles(repo models.Repo, pull models.PullRequest) *MockClient_GetModifiedFiles_OngoingVerification {
	params := []pegomock.Param{repo, pull}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetModifiedFiles", params, verifier.timeout)
	return &MockClient_GetModifiedFiles_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_GetModifiedFiles_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_GetModifiedFiles_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest) {
	repo, pull := c.GetAllCapturedArguments()
	return repo[len(repo)-1], pull[len(pull)-1]
}

func (c *MockClient_GetModifiedFiles_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
	}
	return
}

func (verifier *VerifierMockClient) CreateComment(repo models.Repo, pullNum int, comment string) *MockClient_CreateComment_OngoingVerification {
	params := []pegomock.Param{repo, pullNum, comment}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CreateComment", params, verifier.timeout)
	return &MockClient_CreateComment_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_CreateComment_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_CreateComment_OngoingVerification) GetCapturedArguments() (models.Repo, int, string) {
	repo, pullNum, comment := c.GetAllCapturedArguments()
	return repo[len(repo)-1], pullNum[len(pullNum)-1], comment[len(comment)-1]
}

func (c *MockClient_CreateComment_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []int, _param2 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]int, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(int)
		}
		_param2 = make([]string, len(params[2]))
		for u, param := range params[2] {
			_param2[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockClient) PullIsApproved(repo models.Repo, pull models.PullRequest) *MockClient_PullIsApproved_OngoingVerification {
	params := []pegomock.Param{repo, pull}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "PullIsApproved", params, verifier.timeout)
	return &MockClient_PullIsApproved_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_PullIsApproved_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_PullIsApproved_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest) {
	repo, pull := c.GetAllCapturedArguments()
	return repo[len(repo)-1], pull[len(pull)-1]
}

func (c *MockClient_PullIsApproved_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
	}
	return
}

func (verifier *VerifierMockClient) PullIsMergeable(repo models.Repo, pull models.PullRequest) *MockClient_PullIsMergeable_OngoingVerification {
	params := []pegomock.Param{repo, pull}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "PullIsMergeable", params, verifier.timeout)
	return &MockClient_PullIsMergeable_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_PullIsMergeable_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_PullIsMergeable_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest) {
	repo, pull := c.GetAllCapturedArguments()
	return repo[len(repo)-1], pull[len(pull)-1]
}

func (c *MockClient_PullIsMergeable_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
	}
	return
}

func (verifier *VerifierMockClient) UpdateStatus(repo models.Repo, pull models.PullRequest, state models.CommitStatus, src string, description string, url string) *MockClient_UpdateStatus_OngoingVerification {
	params := []pegomock.Param{repo, pull, state, src, description, url}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "UpdateStatus", params, verifier.timeout)
	return &MockClient_UpdateStatus_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_UpdateStatus_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_UpdateStatus_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest, models.CommitStatus, string, string, string) {
	repo, pull, state, src, description, url := c.GetAllCapturedArguments()
	return repo[len(repo)-1], pull[len(pull)-1], state[len(state)-1], src[len(src)-1], description[len(description)-1], url[len(url)-1]
}

func (c *MockClient_UpdateStatus_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest, _param2 []models.CommitStatus, _param3 []string, _param4 []string, _param5 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
		_param2 = make([]models.CommitStatus, len(params[2]))
		for u, param := range params[2] {
			_param2[u] = param.(models.CommitStatus)
		}
		_param3 = make([]string, len(params[3]))
		for u, param := range params[3] {
			_param3[u] = param.(string)
		}
		_param4 = make([]string, len(params[4]))
		for u, param := range params[4] {
			_param4[u] = param.(string)
		}
		_param5 = make([]string, len(params[5]))
		for u, param := range params[5] {
			_param5[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockClient) MergePull(pull models.PullRequest) *MockClient_MergePull_OngoingVerification {
	params := []pegomock.Param{pull}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "MergePull", params, verifier.timeout)
	return &MockClient_MergePull_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockClient_MergePull_OngoingVerification struct {
	mock              *MockClient
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockClient_MergePull_OngoingVerification) GetCapturedArguments() models.PullRequest {
	pull := c.GetAllCapturedArguments()
	return pull[len(pull)-1]
}

func (c *MockClient_MergePull_OngoingVerification) GetAllCapturedArguments() (_param0 []models.PullRequest) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.PullRequest, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(models.PullRequest)
		}
	}
	return
}
