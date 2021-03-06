// This file was generated by counterfeiter
package commands

import (
	"sync"

	. "github.com/cloudfoundry/cli/cf/commands/application"
	"github.com/cloudfoundry/cli/cf/models"
)

type FakeApplicationStopper struct {
	ApplicationStopStub        func(app models.Application, orgName string, spaceName string) (updatedApp models.Application, err error)
	applicationStopMutex       sync.RWMutex
	applicationStopArgsForCall []struct {
		arg1 models.Application
		arg2 string
		arg3 string
	}
	applicationStopReturns struct {
		result1 models.Application
		result2 error
	}
}

func (fake *FakeApplicationStopper) ApplicationStop(arg1 models.Application, arg2 string, arg3 string) (updatedApp models.Application, err error) {
	fake.applicationStopMutex.Lock()
	defer fake.applicationStopMutex.Unlock()
	fake.applicationStopArgsForCall = append(fake.applicationStopArgsForCall, struct {
		arg1 models.Application
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	if fake.ApplicationStopStub != nil {
		return fake.ApplicationStopStub(arg1, arg2, arg3)
	} else {
		return fake.applicationStopReturns.result1, fake.applicationStopReturns.result2
	}
}

func (fake *FakeApplicationStopper) ApplicationStopCallCount() int {
	fake.applicationStopMutex.RLock()
	defer fake.applicationStopMutex.RUnlock()
	return len(fake.applicationStopArgsForCall)
}

func (fake *FakeApplicationStopper) ApplicationStopArgsForCall(i int) (models.Application, string, string) {
	fake.applicationStopMutex.RLock()
	defer fake.applicationStopMutex.RUnlock()
	return fake.applicationStopArgsForCall[i].arg1, fake.applicationStopArgsForCall[i].arg2, fake.applicationStopArgsForCall[i].arg3
}

func (fake *FakeApplicationStopper) ApplicationStopReturns(result1 models.Application, result2 error) {
	fake.applicationStopReturns = struct {
		result1 models.Application
		result2 error
	}{result1, result2}
}

var _ ApplicationStopper = new(FakeApplicationStopper)
