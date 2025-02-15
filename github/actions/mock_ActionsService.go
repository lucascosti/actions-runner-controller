// Code generated by mockery v2.16.0. DO NOT EDIT.

package actions

import (
	context "context"

	uuid "github.com/google/uuid"
	mock "github.com/stretchr/testify/mock"
)

// MockActionsService is an autogenerated mock type for the ActionsService type
type MockActionsService struct {
	mock.Mock
}

// AcquireJobs provides a mock function with given fields: ctx, runnerScaleSetId, messageQueueAccessToken, requestIds
func (_m *MockActionsService) AcquireJobs(ctx context.Context, runnerScaleSetId int, messageQueueAccessToken string, requestIds []int64) ([]int64, error) {
	ret := _m.Called(ctx, runnerScaleSetId, messageQueueAccessToken, requestIds)

	var r0 []int64
	if rf, ok := ret.Get(0).(func(context.Context, int, string, []int64) []int64); ok {
		r0 = rf(ctx, runnerScaleSetId, messageQueueAccessToken, requestIds)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, string, []int64) error); ok {
		r1 = rf(ctx, runnerScaleSetId, messageQueueAccessToken, requestIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateMessageSession provides a mock function with given fields: ctx, runnerScaleSetId, owner
func (_m *MockActionsService) CreateMessageSession(ctx context.Context, runnerScaleSetId int, owner string) (*RunnerScaleSetSession, error) {
	ret := _m.Called(ctx, runnerScaleSetId, owner)

	var r0 *RunnerScaleSetSession
	if rf, ok := ret.Get(0).(func(context.Context, int, string) *RunnerScaleSetSession); ok {
		r0 = rf(ctx, runnerScaleSetId, owner)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*RunnerScaleSetSession)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, string) error); ok {
		r1 = rf(ctx, runnerScaleSetId, owner)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRunnerScaleSet provides a mock function with given fields: ctx, runnerScaleSet
func (_m *MockActionsService) CreateRunnerScaleSet(ctx context.Context, runnerScaleSet *RunnerScaleSet) (*RunnerScaleSet, error) {
	ret := _m.Called(ctx, runnerScaleSet)

	var r0 *RunnerScaleSet
	if rf, ok := ret.Get(0).(func(context.Context, *RunnerScaleSet) *RunnerScaleSet); ok {
		r0 = rf(ctx, runnerScaleSet)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*RunnerScaleSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *RunnerScaleSet) error); ok {
		r1 = rf(ctx, runnerScaleSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteMessage provides a mock function with given fields: ctx, messageQueueUrl, messageQueueAccessToken, messageId
func (_m *MockActionsService) DeleteMessage(ctx context.Context, messageQueueUrl string, messageQueueAccessToken string, messageId int64) error {
	ret := _m.Called(ctx, messageQueueUrl, messageQueueAccessToken, messageId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64) error); ok {
		r0 = rf(ctx, messageQueueUrl, messageQueueAccessToken, messageId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteMessageSession provides a mock function with given fields: ctx, runnerScaleSetId, sessionId
func (_m *MockActionsService) DeleteMessageSession(ctx context.Context, runnerScaleSetId int, sessionId *uuid.UUID) error {
	ret := _m.Called(ctx, runnerScaleSetId, sessionId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, *uuid.UUID) error); ok {
		r0 = rf(ctx, runnerScaleSetId, sessionId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GenerateJitRunnerConfig provides a mock function with given fields: ctx, jitRunnerSetting, scaleSetId
func (_m *MockActionsService) GenerateJitRunnerConfig(ctx context.Context, jitRunnerSetting *RunnerScaleSetJitRunnerSetting, scaleSetId int) (*RunnerScaleSetJitRunnerConfig, error) {
	ret := _m.Called(ctx, jitRunnerSetting, scaleSetId)

	var r0 *RunnerScaleSetJitRunnerConfig
	if rf, ok := ret.Get(0).(func(context.Context, *RunnerScaleSetJitRunnerSetting, int) *RunnerScaleSetJitRunnerConfig); ok {
		r0 = rf(ctx, jitRunnerSetting, scaleSetId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*RunnerScaleSetJitRunnerConfig)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *RunnerScaleSetJitRunnerSetting, int) error); ok {
		r1 = rf(ctx, jitRunnerSetting, scaleSetId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAcquirableJobs provides a mock function with given fields: ctx, runnerScaleSetId
func (_m *MockActionsService) GetAcquirableJobs(ctx context.Context, runnerScaleSetId int) (*AcquirableJobList, error) {
	ret := _m.Called(ctx, runnerScaleSetId)

	var r0 *AcquirableJobList
	if rf, ok := ret.Get(0).(func(context.Context, int) *AcquirableJobList); ok {
		r0 = rf(ctx, runnerScaleSetId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*AcquirableJobList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, runnerScaleSetId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMessage provides a mock function with given fields: ctx, messageQueueUrl, messageQueueAccessToken, lastMessageId
func (_m *MockActionsService) GetMessage(ctx context.Context, messageQueueUrl string, messageQueueAccessToken string, lastMessageId int64) (*RunnerScaleSetMessage, error) {
	ret := _m.Called(ctx, messageQueueUrl, messageQueueAccessToken, lastMessageId)

	var r0 *RunnerScaleSetMessage
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64) *RunnerScaleSetMessage); ok {
		r0 = rf(ctx, messageQueueUrl, messageQueueAccessToken, lastMessageId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*RunnerScaleSetMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, int64) error); ok {
		r1 = rf(ctx, messageQueueUrl, messageQueueAccessToken, lastMessageId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRunner provides a mock function with given fields: ctx, runnerId
func (_m *MockActionsService) GetRunner(ctx context.Context, runnerId int64) (*RunnerReference, error) {
	ret := _m.Called(ctx, runnerId)

	var r0 *RunnerReference
	if rf, ok := ret.Get(0).(func(context.Context, int64) *RunnerReference); ok {
		r0 = rf(ctx, runnerId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*RunnerReference)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, runnerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRunnerByName provides a mock function with given fields: ctx, runnerName
func (_m *MockActionsService) GetRunnerByName(ctx context.Context, runnerName string) (*RunnerReference, error) {
	ret := _m.Called(ctx, runnerName)

	var r0 *RunnerReference
	if rf, ok := ret.Get(0).(func(context.Context, string) *RunnerReference); ok {
		r0 = rf(ctx, runnerName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*RunnerReference)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, runnerName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRunnerGroupByName provides a mock function with given fields: ctx, runnerGroup
func (_m *MockActionsService) GetRunnerGroupByName(ctx context.Context, runnerGroup string) (*RunnerGroup, error) {
	ret := _m.Called(ctx, runnerGroup)

	var r0 *RunnerGroup
	if rf, ok := ret.Get(0).(func(context.Context, string) *RunnerGroup); ok {
		r0 = rf(ctx, runnerGroup)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*RunnerGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, runnerGroup)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRunnerScaleSet provides a mock function with given fields: ctx, runnerScaleSetName
func (_m *MockActionsService) GetRunnerScaleSet(ctx context.Context, runnerScaleSetName string) (*RunnerScaleSet, error) {
	ret := _m.Called(ctx, runnerScaleSetName)

	var r0 *RunnerScaleSet
	if rf, ok := ret.Get(0).(func(context.Context, string) *RunnerScaleSet); ok {
		r0 = rf(ctx, runnerScaleSetName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*RunnerScaleSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, runnerScaleSetName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRunnerScaleSetById provides a mock function with given fields: ctx, runnerScaleSetId
func (_m *MockActionsService) GetRunnerScaleSetById(ctx context.Context, runnerScaleSetId int) (*RunnerScaleSet, error) {
	ret := _m.Called(ctx, runnerScaleSetId)

	var r0 *RunnerScaleSet
	if rf, ok := ret.Get(0).(func(context.Context, int) *RunnerScaleSet); ok {
		r0 = rf(ctx, runnerScaleSetId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*RunnerScaleSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, runnerScaleSetId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RefreshMessageSession provides a mock function with given fields: ctx, runnerScaleSetId, sessionId
func (_m *MockActionsService) RefreshMessageSession(ctx context.Context, runnerScaleSetId int, sessionId *uuid.UUID) (*RunnerScaleSetSession, error) {
	ret := _m.Called(ctx, runnerScaleSetId, sessionId)

	var r0 *RunnerScaleSetSession
	if rf, ok := ret.Get(0).(func(context.Context, int, *uuid.UUID) *RunnerScaleSetSession); ok {
		r0 = rf(ctx, runnerScaleSetId, sessionId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*RunnerScaleSetSession)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, *uuid.UUID) error); ok {
		r1 = rf(ctx, runnerScaleSetId, sessionId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveRunner provides a mock function with given fields: ctx, runnerId
func (_m *MockActionsService) RemoveRunner(ctx context.Context, runnerId int64) error {
	ret := _m.Called(ctx, runnerId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, runnerId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockActionsService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockActionsService creates a new instance of MockActionsService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockActionsService(t mockConstructorTestingTNewMockActionsService) *MockActionsService {
	mock := &MockActionsService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
