// This file was generated by counterfeiter
package fakes

import (
	. "github.com/cloudfoundry/cli/cf/actors"
	"sync"
)

type FakeServicePlanActor struct {
	UpdateAllPlansForServiceStub        func(string, bool) (bool, error)
	updateAllPlansForServiceMutex       sync.RWMutex
	updateAllPlansForServiceArgsForCall []struct {
		arg1 string
		arg2 bool
	}
	updateAllPlansForServiceReturns struct {
		result1 bool
		result2 error
	}
	UpdateOrgForServiceStub        func(string, string, bool) (bool, error)
	updateOrgForServiceMutex       sync.RWMutex
	updateOrgForServiceArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 bool
	}
	updateOrgForServiceReturns struct {
		result1 bool
		result2 error
	}
	UpdateSinglePlanForServiceStub        func(string, string, bool) (Access, error)
	updateSinglePlanForServiceMutex       sync.RWMutex
	updateSinglePlanForServiceArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 bool
	}
	updateSinglePlanForServiceReturns struct {
		result1 Access
		result2 error
	}
	UpdatePlanAndOrgForServiceStub        func(string, string, string, bool) (Access, error)
	updatePlanAndOrgForServiceMutex       sync.RWMutex
	updatePlanAndOrgForServiceArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 bool
	}
	updatePlanAndOrgForServiceReturns struct {
		result1 Access
		result2 error
	}
}

func (fake *FakeServicePlanActor) UpdateAllPlansForService(arg1 string, arg2 bool) (bool, error) {
	fake.updateAllPlansForServiceMutex.Lock()
	defer fake.updateAllPlansForServiceMutex.Unlock()
	fake.updateAllPlansForServiceArgsForCall = append(fake.updateAllPlansForServiceArgsForCall, struct {
		arg1 string
		arg2 bool
	}{arg1, arg2})
	if fake.UpdateAllPlansForServiceStub != nil {
		return fake.UpdateAllPlansForServiceStub(arg1, arg2)
	} else {
		return fake.updateAllPlansForServiceReturns.result1, fake.updateAllPlansForServiceReturns.result2
	}
}

func (fake *FakeServicePlanActor) UpdateAllPlansForServiceCallCount() int {
	fake.updateAllPlansForServiceMutex.RLock()
	defer fake.updateAllPlansForServiceMutex.RUnlock()
	return len(fake.updateAllPlansForServiceArgsForCall)
}

func (fake *FakeServicePlanActor) UpdateAllPlansForServiceArgsForCall(i int) (string, bool) {
	fake.updateAllPlansForServiceMutex.RLock()
	defer fake.updateAllPlansForServiceMutex.RUnlock()
	return fake.updateAllPlansForServiceArgsForCall[i].arg1, fake.updateAllPlansForServiceArgsForCall[i].arg2
}

func (fake *FakeServicePlanActor) UpdateAllPlansForServiceReturns(result1 bool, result2 error) {
	fake.updateAllPlansForServiceReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeServicePlanActor) UpdateOrgForService(arg1 string, arg2 string, arg3 bool) (bool, error) {
	fake.updateOrgForServiceMutex.Lock()
	defer fake.updateOrgForServiceMutex.Unlock()
	fake.updateOrgForServiceArgsForCall = append(fake.updateOrgForServiceArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 bool
	}{arg1, arg2, arg3})
	if fake.UpdateOrgForServiceStub != nil {
		return fake.UpdateOrgForServiceStub(arg1, arg2, arg3)
	} else {
		return fake.updateOrgForServiceReturns.result1, fake.updateOrgForServiceReturns.result2
	}
}

func (fake *FakeServicePlanActor) UpdateOrgForServiceCallCount() int {
	fake.updateOrgForServiceMutex.RLock()
	defer fake.updateOrgForServiceMutex.RUnlock()
	return len(fake.updateOrgForServiceArgsForCall)
}

func (fake *FakeServicePlanActor) UpdateOrgForServiceArgsForCall(i int) (string, string, bool) {
	fake.updateOrgForServiceMutex.RLock()
	defer fake.updateOrgForServiceMutex.RUnlock()
	return fake.updateOrgForServiceArgsForCall[i].arg1, fake.updateOrgForServiceArgsForCall[i].arg2, fake.updateOrgForServiceArgsForCall[i].arg3
}

func (fake *FakeServicePlanActor) UpdateOrgForServiceReturns(result1 bool, result2 error) {
	fake.updateOrgForServiceReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeServicePlanActor) UpdateSinglePlanForService(arg1 string, arg2 string, arg3 bool) (Access, error) {
	fake.updateSinglePlanForServiceMutex.Lock()
	defer fake.updateSinglePlanForServiceMutex.Unlock()
	fake.updateSinglePlanForServiceArgsForCall = append(fake.updateSinglePlanForServiceArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 bool
	}{arg1, arg2, arg3})
	if fake.UpdateSinglePlanForServiceStub != nil {
		return fake.UpdateSinglePlanForServiceStub(arg1, arg2, arg3)
	} else {
		return fake.updateSinglePlanForServiceReturns.result1, fake.updateSinglePlanForServiceReturns.result2
	}
}

func (fake *FakeServicePlanActor) UpdateSinglePlanForServiceCallCount() int {
	fake.updateSinglePlanForServiceMutex.RLock()
	defer fake.updateSinglePlanForServiceMutex.RUnlock()
	return len(fake.updateSinglePlanForServiceArgsForCall)
}

func (fake *FakeServicePlanActor) UpdateSinglePlanForServiceArgsForCall(i int) (string, string, bool) {
	fake.updateSinglePlanForServiceMutex.RLock()
	defer fake.updateSinglePlanForServiceMutex.RUnlock()
	return fake.updateSinglePlanForServiceArgsForCall[i].arg1, fake.updateSinglePlanForServiceArgsForCall[i].arg2, fake.updateSinglePlanForServiceArgsForCall[i].arg3
}

func (fake *FakeServicePlanActor) UpdateSinglePlanForServiceReturns(result1 Access, result2 error) {
	fake.updateSinglePlanForServiceReturns = struct {
		result1 Access
		result2 error
	}{result1, result2}
}

func (fake *FakeServicePlanActor) UpdatePlanAndOrgForService(arg1 string, arg2 string, arg3 string, arg4 bool) (Access, error) {
	fake.updatePlanAndOrgForServiceMutex.Lock()
	defer fake.updatePlanAndOrgForServiceMutex.Unlock()
	fake.updatePlanAndOrgForServiceArgsForCall = append(fake.updatePlanAndOrgForServiceArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 bool
	}{arg1, arg2, arg3, arg4})
	if fake.UpdatePlanAndOrgForServiceStub != nil {
		return fake.UpdatePlanAndOrgForServiceStub(arg1, arg2, arg3, arg4)
	} else {
		return fake.updatePlanAndOrgForServiceReturns.result1, fake.updatePlanAndOrgForServiceReturns.result2
	}
}

func (fake *FakeServicePlanActor) UpdatePlanAndOrgForServiceCallCount() int {
	fake.updatePlanAndOrgForServiceMutex.RLock()
	defer fake.updatePlanAndOrgForServiceMutex.RUnlock()
	return len(fake.updatePlanAndOrgForServiceArgsForCall)
}

func (fake *FakeServicePlanActor) UpdatePlanAndOrgForServiceArgsForCall(i int) (string, string, string, bool) {
	fake.updatePlanAndOrgForServiceMutex.RLock()
	defer fake.updatePlanAndOrgForServiceMutex.RUnlock()
	return fake.updatePlanAndOrgForServiceArgsForCall[i].arg1, fake.updatePlanAndOrgForServiceArgsForCall[i].arg2, fake.updatePlanAndOrgForServiceArgsForCall[i].arg3, fake.updatePlanAndOrgForServiceArgsForCall[i].arg4
}

func (fake *FakeServicePlanActor) UpdatePlanAndOrgForServiceReturns(result1 Access, result2 error) {
	fake.updatePlanAndOrgForServiceReturns = struct {
		result1 Access
		result2 error
	}{result1, result2}
}

var _ ServicePlanActor = new(FakeServicePlanActor)