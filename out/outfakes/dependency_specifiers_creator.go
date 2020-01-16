// Code generated by counterfeiter. DO NOT EDIT.
package outfakes

import (
	"sync"

	pivnet "github.com/pivotal-cf/go-pivnet/v4"
)

type DependencySpecifiersCreator struct {
	CreateDependencySpecifiersStub        func(release pivnet.Release) error
	createDependencySpecifiersMutex       sync.RWMutex
	createDependencySpecifiersArgsForCall []struct {
		release pivnet.Release
	}
	createDependencySpecifiersReturns struct {
		result1 error
	}
	createDependencySpecifiersReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *DependencySpecifiersCreator) CreateDependencySpecifiers(release pivnet.Release) error {
	fake.createDependencySpecifiersMutex.Lock()
	ret, specificReturn := fake.createDependencySpecifiersReturnsOnCall[len(fake.createDependencySpecifiersArgsForCall)]
	fake.createDependencySpecifiersArgsForCall = append(fake.createDependencySpecifiersArgsForCall, struct {
		release pivnet.Release
	}{release})
	fake.recordInvocation("CreateDependencySpecifiers", []interface{}{release})
	fake.createDependencySpecifiersMutex.Unlock()
	if fake.CreateDependencySpecifiersStub != nil {
		return fake.CreateDependencySpecifiersStub(release)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createDependencySpecifiersReturns.result1
}

func (fake *DependencySpecifiersCreator) CreateDependencySpecifiersCallCount() int {
	fake.createDependencySpecifiersMutex.RLock()
	defer fake.createDependencySpecifiersMutex.RUnlock()
	return len(fake.createDependencySpecifiersArgsForCall)
}

func (fake *DependencySpecifiersCreator) CreateDependencySpecifiersArgsForCall(i int) pivnet.Release {
	fake.createDependencySpecifiersMutex.RLock()
	defer fake.createDependencySpecifiersMutex.RUnlock()
	return fake.createDependencySpecifiersArgsForCall[i].release
}

func (fake *DependencySpecifiersCreator) CreateDependencySpecifiersReturns(result1 error) {
	fake.CreateDependencySpecifiersStub = nil
	fake.createDependencySpecifiersReturns = struct {
		result1 error
	}{result1}
}

func (fake *DependencySpecifiersCreator) CreateDependencySpecifiersReturnsOnCall(i int, result1 error) {
	fake.CreateDependencySpecifiersStub = nil
	if fake.createDependencySpecifiersReturnsOnCall == nil {
		fake.createDependencySpecifiersReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createDependencySpecifiersReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *DependencySpecifiersCreator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createDependencySpecifiersMutex.RLock()
	defer fake.createDependencySpecifiersMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *DependencySpecifiersCreator) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
