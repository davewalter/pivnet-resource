// This file was generated by counterfeiter
package releasefakes

import (
	"sync"

	go_pivnet "github.com/pivotal-cf/go-pivnet/v4"
)

type FinalizerClient struct {
	GetReleaseStub        func(productSlug string, releaseVersion string) (go_pivnet.Release, error)
	getReleaseMutex       sync.RWMutex
	getReleaseArgsForCall []struct {
		productSlug    string
		releaseVersion string
	}
	getReleaseReturns struct {
		result1 go_pivnet.Release
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FinalizerClient) GetRelease(productSlug string, releaseVersion string) (go_pivnet.Release, error) {
	fake.getReleaseMutex.Lock()
	fake.getReleaseArgsForCall = append(fake.getReleaseArgsForCall, struct {
		productSlug    string
		releaseVersion string
	}{productSlug, releaseVersion})
	fake.recordInvocation("GetRelease", []interface{}{productSlug, releaseVersion})
	fake.getReleaseMutex.Unlock()
	if fake.GetReleaseStub != nil {
		return fake.GetReleaseStub(productSlug, releaseVersion)
	} else {
		return fake.getReleaseReturns.result1, fake.getReleaseReturns.result2
	}
}

func (fake *FinalizerClient) GetReleaseCallCount() int {
	fake.getReleaseMutex.RLock()
	defer fake.getReleaseMutex.RUnlock()
	return len(fake.getReleaseArgsForCall)
}

func (fake *FinalizerClient) GetReleaseArgsForCall(i int) (string, string) {
	fake.getReleaseMutex.RLock()
	defer fake.getReleaseMutex.RUnlock()
	return fake.getReleaseArgsForCall[i].productSlug, fake.getReleaseArgsForCall[i].releaseVersion
}

func (fake *FinalizerClient) GetReleaseReturns(result1 go_pivnet.Release, result2 error) {
	fake.GetReleaseStub = nil
	fake.getReleaseReturns = struct {
		result1 go_pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FinalizerClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getReleaseMutex.RLock()
	defer fake.getReleaseMutex.RUnlock()
	return fake.invocations
}

func (fake *FinalizerClient) recordInvocation(key string, args []interface{}) {
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
