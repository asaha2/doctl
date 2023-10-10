// Code generated by MockGen. DO NOT EDIT.
// Source: serverless.go
//
// Generated by this command:
//
//	mockgen -source serverless.go -package=mocks ServerlessService
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	exec "os/exec"
	reflect "reflect"

	whisk "github.com/apache/openwhisk-client-go/whisk"
	do "github.com/digitalocean/doctl/do"
	gomock "go.uber.org/mock/gomock"
)

// MockServerlessService is a mock of ServerlessService interface.
type MockServerlessService struct {
	ctrl     *gomock.Controller
	recorder *MockServerlessServiceMockRecorder
}

// MockServerlessServiceMockRecorder is the mock recorder for MockServerlessService.
type MockServerlessServiceMockRecorder struct {
	mock *MockServerlessService
}

// NewMockServerlessService creates a new mock instance.
func NewMockServerlessService(ctrl *gomock.Controller) *MockServerlessService {
	mock := &MockServerlessService{ctrl: ctrl}
	mock.recorder = &MockServerlessServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServerlessService) EXPECT() *MockServerlessServiceMockRecorder {
	return m.recorder
}

// CheckServerlessStatus mocks base method.
func (m *MockServerlessService) CheckServerlessStatus() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckServerlessStatus")
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckServerlessStatus indicates an expected call of CheckServerlessStatus.
func (mr *MockServerlessServiceMockRecorder) CheckServerlessStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckServerlessStatus", reflect.TypeOf((*MockServerlessService)(nil).CheckServerlessStatus))
}

// CleanNamespace mocks base method.
func (m *MockServerlessService) CleanNamespace() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanNamespace")
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanNamespace indicates an expected call of CleanNamespace.
func (mr *MockServerlessServiceMockRecorder) CleanNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanNamespace", reflect.TypeOf((*MockServerlessService)(nil).CleanNamespace))
}

// Cmd mocks base method.
func (m *MockServerlessService) Cmd(arg0 string, arg1 []string) (*exec.Cmd, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cmd", arg0, arg1)
	ret0, _ := ret[0].(*exec.Cmd)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Cmd indicates an expected call of Cmd.
func (mr *MockServerlessServiceMockRecorder) Cmd(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cmd", reflect.TypeOf((*MockServerlessService)(nil).Cmd), arg0, arg1)
}

// CreateNamespace mocks base method.
func (m *MockServerlessService) CreateNamespace(arg0 context.Context, arg1, arg2 string) (do.ServerlessCredentials, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNamespace", arg0, arg1, arg2)
	ret0, _ := ret[0].(do.ServerlessCredentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNamespace indicates an expected call of CreateNamespace.
func (mr *MockServerlessServiceMockRecorder) CreateNamespace(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNamespace", reflect.TypeOf((*MockServerlessService)(nil).CreateNamespace), arg0, arg1, arg2)
}

// CredentialsPath mocks base method.
func (m *MockServerlessService) CredentialsPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CredentialsPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// CredentialsPath indicates an expected call of CredentialsPath.
func (mr *MockServerlessServiceMockRecorder) CredentialsPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CredentialsPath", reflect.TypeOf((*MockServerlessService)(nil).CredentialsPath))
}

// DeleteFunction mocks base method.
func (m *MockServerlessService) DeleteFunction(arg0 string, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFunction", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFunction indicates an expected call of DeleteFunction.
func (mr *MockServerlessServiceMockRecorder) DeleteFunction(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFunction", reflect.TypeOf((*MockServerlessService)(nil).DeleteFunction), arg0, arg1)
}

// DeleteNamespace mocks base method.
func (m *MockServerlessService) DeleteNamespace(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNamespace", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNamespace indicates an expected call of DeleteNamespace.
func (mr *MockServerlessServiceMockRecorder) DeleteNamespace(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNamespace", reflect.TypeOf((*MockServerlessService)(nil).DeleteNamespace), arg0, arg1)
}

// DeletePackage mocks base method.
func (m *MockServerlessService) DeletePackage(arg0 string, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePackage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePackage indicates an expected call of DeletePackage.
func (mr *MockServerlessServiceMockRecorder) DeletePackage(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePackage", reflect.TypeOf((*MockServerlessService)(nil).DeletePackage), arg0, arg1)
}

// DeleteTrigger mocks base method.
func (m *MockServerlessService) DeleteTrigger(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTrigger", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTrigger indicates an expected call of DeleteTrigger.
func (mr *MockServerlessServiceMockRecorder) DeleteTrigger(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTrigger", reflect.TypeOf((*MockServerlessService)(nil).DeleteTrigger), arg0, arg1)
}

// Exec mocks base method.
func (m *MockServerlessService) Exec(arg0 *exec.Cmd) (do.ServerlessOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exec", arg0)
	ret0, _ := ret[0].(do.ServerlessOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exec indicates an expected call of Exec.
func (mr *MockServerlessServiceMockRecorder) Exec(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockServerlessService)(nil).Exec), arg0)
}

// GetActivation mocks base method.
func (m *MockServerlessService) GetActivation(arg0 string) (whisk.Activation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActivation", arg0)
	ret0, _ := ret[0].(whisk.Activation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActivation indicates an expected call of GetActivation.
func (mr *MockServerlessServiceMockRecorder) GetActivation(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActivation", reflect.TypeOf((*MockServerlessService)(nil).GetActivation), arg0)
}

// GetActivationCount mocks base method.
func (m *MockServerlessService) GetActivationCount(arg0 whisk.ActivationCountOptions) (whisk.ActivationCount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActivationCount", arg0)
	ret0, _ := ret[0].(whisk.ActivationCount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActivationCount indicates an expected call of GetActivationCount.
func (mr *MockServerlessServiceMockRecorder) GetActivationCount(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActivationCount", reflect.TypeOf((*MockServerlessService)(nil).GetActivationCount), arg0)
}

// GetActivationLogs mocks base method.
func (m *MockServerlessService) GetActivationLogs(arg0 string) (whisk.Activation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActivationLogs", arg0)
	ret0, _ := ret[0].(whisk.Activation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActivationLogs indicates an expected call of GetActivationLogs.
func (mr *MockServerlessServiceMockRecorder) GetActivationLogs(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActivationLogs", reflect.TypeOf((*MockServerlessService)(nil).GetActivationLogs), arg0)
}

// GetActivationResult mocks base method.
func (m *MockServerlessService) GetActivationResult(arg0 string) (whisk.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActivationResult", arg0)
	ret0, _ := ret[0].(whisk.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActivationResult indicates an expected call of GetActivationResult.
func (mr *MockServerlessServiceMockRecorder) GetActivationResult(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActivationResult", reflect.TypeOf((*MockServerlessService)(nil).GetActivationResult), arg0)
}

// GetConnectedAPIHost mocks base method.
func (m *MockServerlessService) GetConnectedAPIHost() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConnectedAPIHost")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConnectedAPIHost indicates an expected call of GetConnectedAPIHost.
func (mr *MockServerlessServiceMockRecorder) GetConnectedAPIHost() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConnectedAPIHost", reflect.TypeOf((*MockServerlessService)(nil).GetConnectedAPIHost))
}

// GetFunction mocks base method.
func (m *MockServerlessService) GetFunction(arg0 string, arg1 bool) (whisk.Action, []do.FunctionParameter, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFunction", arg0, arg1)
	ret0, _ := ret[0].(whisk.Action)
	ret1, _ := ret[1].([]do.FunctionParameter)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetFunction indicates an expected call of GetFunction.
func (mr *MockServerlessServiceMockRecorder) GetFunction(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFunction", reflect.TypeOf((*MockServerlessService)(nil).GetFunction), arg0, arg1)
}

// GetHostInfo mocks base method.
func (m *MockServerlessService) GetHostInfo(arg0 string) (do.ServerlessHostInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostInfo", arg0)
	ret0, _ := ret[0].(do.ServerlessHostInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHostInfo indicates an expected call of GetHostInfo.
func (mr *MockServerlessServiceMockRecorder) GetHostInfo(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostInfo", reflect.TypeOf((*MockServerlessService)(nil).GetHostInfo), arg0)
}

// GetNamespace mocks base method.
func (m *MockServerlessService) GetNamespace(arg0 context.Context, arg1 string) (do.ServerlessCredentials, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespace", arg0, arg1)
	ret0, _ := ret[0].(do.ServerlessCredentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNamespace indicates an expected call of GetNamespace.
func (mr *MockServerlessServiceMockRecorder) GetNamespace(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespace", reflect.TypeOf((*MockServerlessService)(nil).GetNamespace), arg0, arg1)
}

// GetNamespaceFromCluster mocks base method.
func (m *MockServerlessService) GetNamespaceFromCluster(arg0, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespaceFromCluster", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNamespaceFromCluster indicates an expected call of GetNamespaceFromCluster.
func (mr *MockServerlessServiceMockRecorder) GetNamespaceFromCluster(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespaceFromCluster", reflect.TypeOf((*MockServerlessService)(nil).GetNamespaceFromCluster), arg0, arg1)
}

// GetServerlessNamespace mocks base method.
func (m *MockServerlessService) GetServerlessNamespace(arg0 context.Context) (do.ServerlessCredentials, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServerlessNamespace", arg0)
	ret0, _ := ret[0].(do.ServerlessCredentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServerlessNamespace indicates an expected call of GetServerlessNamespace.
func (mr *MockServerlessServiceMockRecorder) GetServerlessNamespace(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServerlessNamespace", reflect.TypeOf((*MockServerlessService)(nil).GetServerlessNamespace), arg0)
}

// GetTrigger mocks base method.
func (m *MockServerlessService) GetTrigger(arg0 context.Context, arg1 string) (do.ServerlessTrigger, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrigger", arg0, arg1)
	ret0, _ := ret[0].(do.ServerlessTrigger)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTrigger indicates an expected call of GetTrigger.
func (mr *MockServerlessServiceMockRecorder) GetTrigger(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrigger", reflect.TypeOf((*MockServerlessService)(nil).GetTrigger), arg0, arg1)
}

// InstallServerless mocks base method.
func (m *MockServerlessService) InstallServerless(arg0 string, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstallServerless", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InstallServerless indicates an expected call of InstallServerless.
func (mr *MockServerlessServiceMockRecorder) InstallServerless(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstallServerless", reflect.TypeOf((*MockServerlessService)(nil).InstallServerless), arg0, arg1)
}

// InvokeFunction mocks base method.
func (m *MockServerlessService) InvokeFunction(arg0 string, arg1 any, arg2, arg3 bool) (any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvokeFunction", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InvokeFunction indicates an expected call of InvokeFunction.
func (mr *MockServerlessServiceMockRecorder) InvokeFunction(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeFunction", reflect.TypeOf((*MockServerlessService)(nil).InvokeFunction), arg0, arg1, arg2, arg3)
}

// InvokeFunctionViaWeb mocks base method.
func (m *MockServerlessService) InvokeFunctionViaWeb(arg0 string, arg1 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InvokeFunctionViaWeb", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InvokeFunctionViaWeb indicates an expected call of InvokeFunctionViaWeb.
func (mr *MockServerlessServiceMockRecorder) InvokeFunctionViaWeb(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvokeFunctionViaWeb", reflect.TypeOf((*MockServerlessService)(nil).InvokeFunctionViaWeb), arg0, arg1)
}

// ListActivations mocks base method.
func (m *MockServerlessService) ListActivations(arg0 whisk.ActivationListOptions) ([]whisk.Activation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListActivations", arg0)
	ret0, _ := ret[0].([]whisk.Activation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListActivations indicates an expected call of ListActivations.
func (mr *MockServerlessServiceMockRecorder) ListActivations(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListActivations", reflect.TypeOf((*MockServerlessService)(nil).ListActivations), arg0)
}

// ListFunctions mocks base method.
func (m *MockServerlessService) ListFunctions(arg0 string, arg1, arg2 int) ([]whisk.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFunctions", arg0, arg1, arg2)
	ret0, _ := ret[0].([]whisk.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFunctions indicates an expected call of ListFunctions.
func (mr *MockServerlessServiceMockRecorder) ListFunctions(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFunctions", reflect.TypeOf((*MockServerlessService)(nil).ListFunctions), arg0, arg1, arg2)
}

// ListNamespaces mocks base method.
func (m *MockServerlessService) ListNamespaces(arg0 context.Context) (do.NamespaceListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNamespaces", arg0)
	ret0, _ := ret[0].(do.NamespaceListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNamespaces indicates an expected call of ListNamespaces.
func (mr *MockServerlessServiceMockRecorder) ListNamespaces(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNamespaces", reflect.TypeOf((*MockServerlessService)(nil).ListNamespaces), arg0)
}

// ListPackages mocks base method.
func (m *MockServerlessService) ListPackages() ([]whisk.Package, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPackages")
	ret0, _ := ret[0].([]whisk.Package)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPackages indicates an expected call of ListPackages.
func (mr *MockServerlessServiceMockRecorder) ListPackages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPackages", reflect.TypeOf((*MockServerlessService)(nil).ListPackages))
}

// ListTriggers mocks base method.
func (m *MockServerlessService) ListTriggers(arg0 context.Context, arg1 string) ([]do.ServerlessTrigger, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTriggers", arg0, arg1)
	ret0, _ := ret[0].([]do.ServerlessTrigger)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTriggers indicates an expected call of ListTriggers.
func (mr *MockServerlessServiceMockRecorder) ListTriggers(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTriggers", reflect.TypeOf((*MockServerlessService)(nil).ListTriggers), arg0, arg1)
}

// ReadCredentials mocks base method.
func (m *MockServerlessService) ReadCredentials() (do.ServerlessCredentials, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadCredentials")
	ret0, _ := ret[0].(do.ServerlessCredentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadCredentials indicates an expected call of ReadCredentials.
func (mr *MockServerlessServiceMockRecorder) ReadCredentials() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadCredentials", reflect.TypeOf((*MockServerlessService)(nil).ReadCredentials))
}

// ReadProject mocks base method.
func (m *MockServerlessService) ReadProject(arg0 *do.ServerlessProject, arg1 []string) (do.ServerlessOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadProject", arg0, arg1)
	ret0, _ := ret[0].(do.ServerlessOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadProject indicates an expected call of ReadProject.
func (mr *MockServerlessServiceMockRecorder) ReadProject(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadProject", reflect.TypeOf((*MockServerlessService)(nil).ReadProject), arg0, arg1)
}

// SetEffectiveCredentials mocks base method.
func (m *MockServerlessService) SetEffectiveCredentials(auth, apihost string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetEffectiveCredentials", auth, apihost)
}

// SetEffectiveCredentials indicates an expected call of SetEffectiveCredentials.
func (mr *MockServerlessServiceMockRecorder) SetEffectiveCredentials(auth, apihost any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEffectiveCredentials", reflect.TypeOf((*MockServerlessService)(nil).SetEffectiveCredentials), auth, apihost)
}

// Stream mocks base method.
func (m *MockServerlessService) Stream(arg0 *exec.Cmd) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stream", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stream indicates an expected call of Stream.
func (mr *MockServerlessServiceMockRecorder) Stream(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stream", reflect.TypeOf((*MockServerlessService)(nil).Stream), arg0)
}

// UpdateTrigger mocks base method.
func (m *MockServerlessService) UpdateTrigger(arg0 context.Context, arg1 string, arg2 *do.UpdateTriggerRequest) (do.ServerlessTrigger, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTrigger", arg0, arg1, arg2)
	ret0, _ := ret[0].(do.ServerlessTrigger)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTrigger indicates an expected call of UpdateTrigger.
func (mr *MockServerlessServiceMockRecorder) UpdateTrigger(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTrigger", reflect.TypeOf((*MockServerlessService)(nil).UpdateTrigger), arg0, arg1, arg2)
}

// WriteCredentials mocks base method.
func (m *MockServerlessService) WriteCredentials(arg0 do.ServerlessCredentials) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteCredentials", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteCredentials indicates an expected call of WriteCredentials.
func (mr *MockServerlessServiceMockRecorder) WriteCredentials(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteCredentials", reflect.TypeOf((*MockServerlessService)(nil).WriteCredentials), arg0)
}

// WriteProject mocks base method.
func (m *MockServerlessService) WriteProject(arg0 do.ServerlessProject) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteProject", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteProject indicates an expected call of WriteProject.
func (mr *MockServerlessServiceMockRecorder) WriteProject(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteProject", reflect.TypeOf((*MockServerlessService)(nil).WriteProject), arg0)
}
