/*
 * SPDX-FileCopyrightText: 2023 Siemens AG
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * Author: Michael Adler <michael.adler@siemens.com>
 */

// Code generated by mockery v2.40.2. DO NOT EDIT.

package persistence

import (
	context "context"

	model "github.com/siemens/wfx/generated/model"
	mock "github.com/stretchr/testify/mock"
)

// MockStorage is an autogenerated mock type for the Storage type
type MockStorage struct {
	mock.Mock
}

type MockStorage_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStorage) EXPECT() *MockStorage_Expecter {
	return &MockStorage_Expecter{mock: &_m.Mock}
}

// CheckHealth provides a mock function with given fields: ctx
func (_m *MockStorage) CheckHealth(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for CheckHealth")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStorage_CheckHealth_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckHealth'
type MockStorage_CheckHealth_Call struct {
	*mock.Call
}

// CheckHealth is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockStorage_Expecter) CheckHealth(ctx interface{}) *MockStorage_CheckHealth_Call {
	return &MockStorage_CheckHealth_Call{Call: _e.mock.On("CheckHealth", ctx)}
}

func (_c *MockStorage_CheckHealth_Call) Run(run func(ctx context.Context)) *MockStorage_CheckHealth_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockStorage_CheckHealth_Call) Return(_a0 error) *MockStorage_CheckHealth_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStorage_CheckHealth_Call) RunAndReturn(run func(context.Context) error) *MockStorage_CheckHealth_Call {
	_c.Call.Return(run)
	return _c
}

// CreateJob provides a mock function with given fields: ctx, job
func (_m *MockStorage) CreateJob(ctx context.Context, job *model.Job) (*model.Job, error) {
	ret := _m.Called(ctx, job)

	if len(ret) == 0 {
		panic("no return value specified for CreateJob")
	}

	var r0 *model.Job
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Job) (*model.Job, error)); ok {
		return rf(ctx, job)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.Job) *model.Job); ok {
		r0 = rf(ctx, job)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Job)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.Job) error); ok {
		r1 = rf(ctx, job)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorage_CreateJob_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateJob'
type MockStorage_CreateJob_Call struct {
	*mock.Call
}

// CreateJob is a helper method to define mock.On call
//   - ctx context.Context
//   - job *model.Job
func (_e *MockStorage_Expecter) CreateJob(ctx interface{}, job interface{}) *MockStorage_CreateJob_Call {
	return &MockStorage_CreateJob_Call{Call: _e.mock.On("CreateJob", ctx, job)}
}

func (_c *MockStorage_CreateJob_Call) Run(run func(ctx context.Context, job *model.Job)) *MockStorage_CreateJob_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.Job))
	})
	return _c
}

func (_c *MockStorage_CreateJob_Call) Return(_a0 *model.Job, _a1 error) *MockStorage_CreateJob_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorage_CreateJob_Call) RunAndReturn(run func(context.Context, *model.Job) (*model.Job, error)) *MockStorage_CreateJob_Call {
	_c.Call.Return(run)
	return _c
}

// CreateWorkflow provides a mock function with given fields: ctx, workflow
func (_m *MockStorage) CreateWorkflow(ctx context.Context, workflow *model.Workflow) (*model.Workflow, error) {
	ret := _m.Called(ctx, workflow)

	if len(ret) == 0 {
		panic("no return value specified for CreateWorkflow")
	}

	var r0 *model.Workflow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Workflow) (*model.Workflow, error)); ok {
		return rf(ctx, workflow)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.Workflow) *model.Workflow); ok {
		r0 = rf(ctx, workflow)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Workflow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.Workflow) error); ok {
		r1 = rf(ctx, workflow)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorage_CreateWorkflow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateWorkflow'
type MockStorage_CreateWorkflow_Call struct {
	*mock.Call
}

// CreateWorkflow is a helper method to define mock.On call
//   - ctx context.Context
//   - workflow *model.Workflow
func (_e *MockStorage_Expecter) CreateWorkflow(ctx interface{}, workflow interface{}) *MockStorage_CreateWorkflow_Call {
	return &MockStorage_CreateWorkflow_Call{Call: _e.mock.On("CreateWorkflow", ctx, workflow)}
}

func (_c *MockStorage_CreateWorkflow_Call) Run(run func(ctx context.Context, workflow *model.Workflow)) *MockStorage_CreateWorkflow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.Workflow))
	})
	return _c
}

func (_c *MockStorage_CreateWorkflow_Call) Return(_a0 *model.Workflow, _a1 error) *MockStorage_CreateWorkflow_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorage_CreateWorkflow_Call) RunAndReturn(run func(context.Context, *model.Workflow) (*model.Workflow, error)) *MockStorage_CreateWorkflow_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteJob provides a mock function with given fields: ctx, jobID
func (_m *MockStorage) DeleteJob(ctx context.Context, jobID string) error {
	ret := _m.Called(ctx, jobID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteJob")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, jobID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStorage_DeleteJob_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteJob'
type MockStorage_DeleteJob_Call struct {
	*mock.Call
}

// DeleteJob is a helper method to define mock.On call
//   - ctx context.Context
//   - jobID string
func (_e *MockStorage_Expecter) DeleteJob(ctx interface{}, jobID interface{}) *MockStorage_DeleteJob_Call {
	return &MockStorage_DeleteJob_Call{Call: _e.mock.On("DeleteJob", ctx, jobID)}
}

func (_c *MockStorage_DeleteJob_Call) Run(run func(ctx context.Context, jobID string)) *MockStorage_DeleteJob_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockStorage_DeleteJob_Call) Return(_a0 error) *MockStorage_DeleteJob_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStorage_DeleteJob_Call) RunAndReturn(run func(context.Context, string) error) *MockStorage_DeleteJob_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteWorkflow provides a mock function with given fields: ctx, name
func (_m *MockStorage) DeleteWorkflow(ctx context.Context, name string) error {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for DeleteWorkflow")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStorage_DeleteWorkflow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteWorkflow'
type MockStorage_DeleteWorkflow_Call struct {
	*mock.Call
}

// DeleteWorkflow is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockStorage_Expecter) DeleteWorkflow(ctx interface{}, name interface{}) *MockStorage_DeleteWorkflow_Call {
	return &MockStorage_DeleteWorkflow_Call{Call: _e.mock.On("DeleteWorkflow", ctx, name)}
}

func (_c *MockStorage_DeleteWorkflow_Call) Run(run func(ctx context.Context, name string)) *MockStorage_DeleteWorkflow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockStorage_DeleteWorkflow_Call) Return(_a0 error) *MockStorage_DeleteWorkflow_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStorage_DeleteWorkflow_Call) RunAndReturn(run func(context.Context, string) error) *MockStorage_DeleteWorkflow_Call {
	_c.Call.Return(run)
	return _c
}

// GetJob provides a mock function with given fields: ctx, jobID, fetchParams
func (_m *MockStorage) GetJob(ctx context.Context, jobID string, fetchParams FetchParams) (*model.Job, error) {
	ret := _m.Called(ctx, jobID, fetchParams)

	if len(ret) == 0 {
		panic("no return value specified for GetJob")
	}

	var r0 *model.Job
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, FetchParams) (*model.Job, error)); ok {
		return rf(ctx, jobID, fetchParams)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, FetchParams) *model.Job); ok {
		r0 = rf(ctx, jobID, fetchParams)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Job)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, FetchParams) error); ok {
		r1 = rf(ctx, jobID, fetchParams)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorage_GetJob_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetJob'
type MockStorage_GetJob_Call struct {
	*mock.Call
}

// GetJob is a helper method to define mock.On call
//   - ctx context.Context
//   - jobID string
//   - fetchParams FetchParams
func (_e *MockStorage_Expecter) GetJob(ctx interface{}, jobID interface{}, fetchParams interface{}) *MockStorage_GetJob_Call {
	return &MockStorage_GetJob_Call{Call: _e.mock.On("GetJob", ctx, jobID, fetchParams)}
}

func (_c *MockStorage_GetJob_Call) Run(run func(ctx context.Context, jobID string, fetchParams FetchParams)) *MockStorage_GetJob_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(FetchParams))
	})
	return _c
}

func (_c *MockStorage_GetJob_Call) Return(_a0 *model.Job, _a1 error) *MockStorage_GetJob_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorage_GetJob_Call) RunAndReturn(run func(context.Context, string, FetchParams) (*model.Job, error)) *MockStorage_GetJob_Call {
	_c.Call.Return(run)
	return _c
}

// GetWorkflow provides a mock function with given fields: ctx, name
func (_m *MockStorage) GetWorkflow(ctx context.Context, name string) (*model.Workflow, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetWorkflow")
	}

	var r0 *model.Workflow
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.Workflow, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Workflow); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Workflow)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorage_GetWorkflow_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetWorkflow'
type MockStorage_GetWorkflow_Call struct {
	*mock.Call
}

// GetWorkflow is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockStorage_Expecter) GetWorkflow(ctx interface{}, name interface{}) *MockStorage_GetWorkflow_Call {
	return &MockStorage_GetWorkflow_Call{Call: _e.mock.On("GetWorkflow", ctx, name)}
}

func (_c *MockStorage_GetWorkflow_Call) Run(run func(ctx context.Context, name string)) *MockStorage_GetWorkflow_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockStorage_GetWorkflow_Call) Return(_a0 *model.Workflow, _a1 error) *MockStorage_GetWorkflow_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorage_GetWorkflow_Call) RunAndReturn(run func(context.Context, string) (*model.Workflow, error)) *MockStorage_GetWorkflow_Call {
	_c.Call.Return(run)
	return _c
}

// Initialize provides a mock function with given fields: ctx, options
func (_m *MockStorage) Initialize(ctx context.Context, options string) error {
	ret := _m.Called(ctx, options)

	if len(ret) == 0 {
		panic("no return value specified for Initialize")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStorage_Initialize_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Initialize'
type MockStorage_Initialize_Call struct {
	*mock.Call
}

// Initialize is a helper method to define mock.On call
//   - ctx context.Context
//   - options string
func (_e *MockStorage_Expecter) Initialize(ctx interface{}, options interface{}) *MockStorage_Initialize_Call {
	return &MockStorage_Initialize_Call{Call: _e.mock.On("Initialize", ctx, options)}
}

func (_c *MockStorage_Initialize_Call) Run(run func(ctx context.Context, options string)) *MockStorage_Initialize_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockStorage_Initialize_Call) Return(_a0 error) *MockStorage_Initialize_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockStorage_Initialize_Call) RunAndReturn(run func(context.Context, string) error) *MockStorage_Initialize_Call {
	_c.Call.Return(run)
	return _c
}

// QueryJobs provides a mock function with given fields: ctx, filterParams, sortParams, paginationParams
func (_m *MockStorage) QueryJobs(ctx context.Context, filterParams FilterParams, sortParams SortParams, paginationParams PaginationParams) (*model.PaginatedJobList, error) {
	ret := _m.Called(ctx, filterParams, sortParams, paginationParams)

	if len(ret) == 0 {
		panic("no return value specified for QueryJobs")
	}

	var r0 *model.PaginatedJobList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, FilterParams, SortParams, PaginationParams) (*model.PaginatedJobList, error)); ok {
		return rf(ctx, filterParams, sortParams, paginationParams)
	}
	if rf, ok := ret.Get(0).(func(context.Context, FilterParams, SortParams, PaginationParams) *model.PaginatedJobList); ok {
		r0 = rf(ctx, filterParams, sortParams, paginationParams)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PaginatedJobList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, FilterParams, SortParams, PaginationParams) error); ok {
		r1 = rf(ctx, filterParams, sortParams, paginationParams)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorage_QueryJobs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryJobs'
type MockStorage_QueryJobs_Call struct {
	*mock.Call
}

// QueryJobs is a helper method to define mock.On call
//   - ctx context.Context
//   - filterParams FilterParams
//   - sortParams SortParams
//   - paginationParams PaginationParams
func (_e *MockStorage_Expecter) QueryJobs(ctx interface{}, filterParams interface{}, sortParams interface{}, paginationParams interface{}) *MockStorage_QueryJobs_Call {
	return &MockStorage_QueryJobs_Call{Call: _e.mock.On("QueryJobs", ctx, filterParams, sortParams, paginationParams)}
}

func (_c *MockStorage_QueryJobs_Call) Run(run func(ctx context.Context, filterParams FilterParams, sortParams SortParams, paginationParams PaginationParams)) *MockStorage_QueryJobs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(FilterParams), args[2].(SortParams), args[3].(PaginationParams))
	})
	return _c
}

func (_c *MockStorage_QueryJobs_Call) Return(_a0 *model.PaginatedJobList, _a1 error) *MockStorage_QueryJobs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorage_QueryJobs_Call) RunAndReturn(run func(context.Context, FilterParams, SortParams, PaginationParams) (*model.PaginatedJobList, error)) *MockStorage_QueryJobs_Call {
	_c.Call.Return(run)
	return _c
}

// QueryWorkflows provides a mock function with given fields: ctx, paginationParams
func (_m *MockStorage) QueryWorkflows(ctx context.Context, paginationParams PaginationParams) (*model.PaginatedWorkflowList, error) {
	ret := _m.Called(ctx, paginationParams)

	if len(ret) == 0 {
		panic("no return value specified for QueryWorkflows")
	}

	var r0 *model.PaginatedWorkflowList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, PaginationParams) (*model.PaginatedWorkflowList, error)); ok {
		return rf(ctx, paginationParams)
	}
	if rf, ok := ret.Get(0).(func(context.Context, PaginationParams) *model.PaginatedWorkflowList); ok {
		r0 = rf(ctx, paginationParams)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.PaginatedWorkflowList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, PaginationParams) error); ok {
		r1 = rf(ctx, paginationParams)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorage_QueryWorkflows_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryWorkflows'
type MockStorage_QueryWorkflows_Call struct {
	*mock.Call
}

// QueryWorkflows is a helper method to define mock.On call
//   - ctx context.Context
//   - paginationParams PaginationParams
func (_e *MockStorage_Expecter) QueryWorkflows(ctx interface{}, paginationParams interface{}) *MockStorage_QueryWorkflows_Call {
	return &MockStorage_QueryWorkflows_Call{Call: _e.mock.On("QueryWorkflows", ctx, paginationParams)}
}

func (_c *MockStorage_QueryWorkflows_Call) Run(run func(ctx context.Context, paginationParams PaginationParams)) *MockStorage_QueryWorkflows_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(PaginationParams))
	})
	return _c
}

func (_c *MockStorage_QueryWorkflows_Call) Return(_a0 *model.PaginatedWorkflowList, _a1 error) *MockStorage_QueryWorkflows_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorage_QueryWorkflows_Call) RunAndReturn(run func(context.Context, PaginationParams) (*model.PaginatedWorkflowList, error)) *MockStorage_QueryWorkflows_Call {
	_c.Call.Return(run)
	return _c
}

// Shutdown provides a mock function with given fields:
func (_m *MockStorage) Shutdown() {
	_m.Called()
}

// MockStorage_Shutdown_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Shutdown'
type MockStorage_Shutdown_Call struct {
	*mock.Call
}

// Shutdown is a helper method to define mock.On call
func (_e *MockStorage_Expecter) Shutdown() *MockStorage_Shutdown_Call {
	return &MockStorage_Shutdown_Call{Call: _e.mock.On("Shutdown")}
}

func (_c *MockStorage_Shutdown_Call) Run(run func()) *MockStorage_Shutdown_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStorage_Shutdown_Call) Return() *MockStorage_Shutdown_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockStorage_Shutdown_Call) RunAndReturn(run func()) *MockStorage_Shutdown_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateJob provides a mock function with given fields: ctx, job, request
func (_m *MockStorage) UpdateJob(ctx context.Context, job *model.Job, request JobUpdate) (*model.Job, error) {
	ret := _m.Called(ctx, job, request)

	if len(ret) == 0 {
		panic("no return value specified for UpdateJob")
	}

	var r0 *model.Job
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Job, JobUpdate) (*model.Job, error)); ok {
		return rf(ctx, job, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.Job, JobUpdate) *model.Job); ok {
		r0 = rf(ctx, job, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Job)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.Job, JobUpdate) error); ok {
		r1 = rf(ctx, job, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorage_UpdateJob_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateJob'
type MockStorage_UpdateJob_Call struct {
	*mock.Call
}

// UpdateJob is a helper method to define mock.On call
//   - ctx context.Context
//   - job *model.Job
//   - request JobUpdate
func (_e *MockStorage_Expecter) UpdateJob(ctx interface{}, job interface{}, request interface{}) *MockStorage_UpdateJob_Call {
	return &MockStorage_UpdateJob_Call{Call: _e.mock.On("UpdateJob", ctx, job, request)}
}

func (_c *MockStorage_UpdateJob_Call) Run(run func(ctx context.Context, job *model.Job, request JobUpdate)) *MockStorage_UpdateJob_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.Job), args[2].(JobUpdate))
	})
	return _c
}

func (_c *MockStorage_UpdateJob_Call) Return(_a0 *model.Job, _a1 error) *MockStorage_UpdateJob_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorage_UpdateJob_Call) RunAndReturn(run func(context.Context, *model.Job, JobUpdate) (*model.Job, error)) *MockStorage_UpdateJob_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockStorage creates a new instance of MockStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStorage {
	mock := &MockStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}