// Code generated by counterfeiter. DO NOT EDIT.
package outfakes

import (
	"sync"

	pivnet "github.com/pivotal-cf/go-pivnet/v4"
)

type ReleaseDependenciesAdder struct {
	AddReleaseDependenciesStub        func(release pivnet.Release) error
	addReleaseDependenciesMutex       sync.RWMutex
	addReleaseDependenciesArgsForCall []struct {
		release pivnet.Release
	}
	addReleaseDependenciesReturns struct {
		result1 error
	}
	addReleaseDependenciesReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ReleaseDependenciesAdder) AddReleaseDependencies(release pivnet.Release) error {
	fake.addReleaseDependenciesMutex.Lock()
	ret, specificReturn := fake.addReleaseDependenciesReturnsOnCall[len(fake.addReleaseDependenciesArgsForCall)]
	fake.addReleaseDependenciesArgsForCall = append(fake.addReleaseDependenciesArgsForCall, struct {
		release pivnet.Release
	}{release})
	fake.recordInvocation("AddReleaseDependencies", []interface{}{release})
	fake.addReleaseDependenciesMutex.Unlock()
	if fake.AddReleaseDependenciesStub != nil {
		return fake.AddReleaseDependenciesStub(release)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.addReleaseDependenciesReturns.result1
}

func (fake *ReleaseDependenciesAdder) AddReleaseDependenciesCallCount() int {
	fake.addReleaseDependenciesMutex.RLock()
	defer fake.addReleaseDependenciesMutex.RUnlock()
	return len(fake.addReleaseDependenciesArgsForCall)
}

func (fake *ReleaseDependenciesAdder) AddReleaseDependenciesArgsForCall(i int) pivnet.Release {
	fake.addReleaseDependenciesMutex.RLock()
	defer fake.addReleaseDependenciesMutex.RUnlock()
	return fake.addReleaseDependenciesArgsForCall[i].release
}

func (fake *ReleaseDependenciesAdder) AddReleaseDependenciesReturns(result1 error) {
	fake.AddReleaseDependenciesStub = nil
	fake.addReleaseDependenciesReturns = struct {
		result1 error
	}{result1}
}

func (fake *ReleaseDependenciesAdder) AddReleaseDependenciesReturnsOnCall(i int, result1 error) {
	fake.AddReleaseDependenciesStub = nil
	if fake.addReleaseDependenciesReturnsOnCall == nil {
		fake.addReleaseDependenciesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addReleaseDependenciesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ReleaseDependenciesAdder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addReleaseDependenciesMutex.RLock()
	defer fake.addReleaseDependenciesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ReleaseDependenciesAdder) recordInvocation(key string, args []interface{}) {
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
