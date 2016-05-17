// This file was generated by counterfeiter
package gpfilterfakes

import (
	"sync"

	"github.com/pivotal-cf-experimental/go-pivnet"
	"github.com/pivotal-cf-experimental/pivnet-resource/gp_filter"
)

type FakeFilter struct {
	ReleasesByReleaseTypeStub        func(releases []pivnet.Release, releaseType string) ([]pivnet.Release, error)
	releasesByReleaseTypeMutex       sync.RWMutex
	releasesByReleaseTypeArgsForCall []struct {
		releases    []pivnet.Release
		releaseType string
	}
	releasesByReleaseTypeReturns struct {
		result1 []pivnet.Release
		result2 error
	}
	ReleasesByVersionStub        func(releases []pivnet.Release, version string) ([]pivnet.Release, error)
	releasesByVersionMutex       sync.RWMutex
	releasesByVersionArgsForCall []struct {
		releases []pivnet.Release
		version  string
	}
	releasesByVersionReturns struct {
		result1 []pivnet.Release
		result2 error
	}
	DownloadLinksByGlobStub        func(downloadLinks map[string]string, glob []string) (map[string]string, error)
	downloadLinksByGlobMutex       sync.RWMutex
	downloadLinksByGlobArgsForCall []struct {
		downloadLinks map[string]string
		glob          []string
	}
	downloadLinksByGlobReturns struct {
		result1 map[string]string
		result2 error
	}
	DownloadLinksStub        func(p []pivnet.ProductFile) map[string]string
	downloadLinksMutex       sync.RWMutex
	downloadLinksArgsForCall []struct {
		p []pivnet.ProductFile
	}
	downloadLinksReturns struct {
		result1 map[string]string
	}
	invocations map[string][][]interface{}
}

func (fake *FakeFilter) ReleasesByReleaseType(releases []pivnet.Release, releaseType string) ([]pivnet.Release, error) {
	var releasesCopy []pivnet.Release
	if releases != nil {
		releasesCopy = make([]pivnet.Release, len(releases))
		copy(releasesCopy, releases)
	}
	fake.releasesByReleaseTypeMutex.Lock()
	fake.releasesByReleaseTypeArgsForCall = append(fake.releasesByReleaseTypeArgsForCall, struct {
		releases    []pivnet.Release
		releaseType string
	}{releasesCopy, releaseType})
	fake.guard("ReleasesByReleaseType")
	fake.invocations["ReleasesByReleaseType"] = append(fake.invocations["ReleasesByReleaseType"], []interface{}{releasesCopy, releaseType})
	fake.releasesByReleaseTypeMutex.Unlock()
	if fake.ReleasesByReleaseTypeStub != nil {
		return fake.ReleasesByReleaseTypeStub(releases, releaseType)
	} else {
		return fake.releasesByReleaseTypeReturns.result1, fake.releasesByReleaseTypeReturns.result2
	}
}

func (fake *FakeFilter) ReleasesByReleaseTypeCallCount() int {
	fake.releasesByReleaseTypeMutex.RLock()
	defer fake.releasesByReleaseTypeMutex.RUnlock()
	return len(fake.releasesByReleaseTypeArgsForCall)
}

func (fake *FakeFilter) ReleasesByReleaseTypeArgsForCall(i int) ([]pivnet.Release, string) {
	fake.releasesByReleaseTypeMutex.RLock()
	defer fake.releasesByReleaseTypeMutex.RUnlock()
	return fake.releasesByReleaseTypeArgsForCall[i].releases, fake.releasesByReleaseTypeArgsForCall[i].releaseType
}

func (fake *FakeFilter) ReleasesByReleaseTypeReturns(result1 []pivnet.Release, result2 error) {
	fake.ReleasesByReleaseTypeStub = nil
	fake.releasesByReleaseTypeReturns = struct {
		result1 []pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakeFilter) ReleasesByVersion(releases []pivnet.Release, version string) ([]pivnet.Release, error) {
	var releasesCopy []pivnet.Release
	if releases != nil {
		releasesCopy = make([]pivnet.Release, len(releases))
		copy(releasesCopy, releases)
	}
	fake.releasesByVersionMutex.Lock()
	fake.releasesByVersionArgsForCall = append(fake.releasesByVersionArgsForCall, struct {
		releases []pivnet.Release
		version  string
	}{releasesCopy, version})
	fake.guard("ReleasesByVersion")
	fake.invocations["ReleasesByVersion"] = append(fake.invocations["ReleasesByVersion"], []interface{}{releasesCopy, version})
	fake.releasesByVersionMutex.Unlock()
	if fake.ReleasesByVersionStub != nil {
		return fake.ReleasesByVersionStub(releases, version)
	} else {
		return fake.releasesByVersionReturns.result1, fake.releasesByVersionReturns.result2
	}
}

func (fake *FakeFilter) ReleasesByVersionCallCount() int {
	fake.releasesByVersionMutex.RLock()
	defer fake.releasesByVersionMutex.RUnlock()
	return len(fake.releasesByVersionArgsForCall)
}

func (fake *FakeFilter) ReleasesByVersionArgsForCall(i int) ([]pivnet.Release, string) {
	fake.releasesByVersionMutex.RLock()
	defer fake.releasesByVersionMutex.RUnlock()
	return fake.releasesByVersionArgsForCall[i].releases, fake.releasesByVersionArgsForCall[i].version
}

func (fake *FakeFilter) ReleasesByVersionReturns(result1 []pivnet.Release, result2 error) {
	fake.ReleasesByVersionStub = nil
	fake.releasesByVersionReturns = struct {
		result1 []pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakeFilter) DownloadLinksByGlob(downloadLinks map[string]string, glob []string) (map[string]string, error) {
	var globCopy []string
	if glob != nil {
		globCopy = make([]string, len(glob))
		copy(globCopy, glob)
	}
	fake.downloadLinksByGlobMutex.Lock()
	fake.downloadLinksByGlobArgsForCall = append(fake.downloadLinksByGlobArgsForCall, struct {
		downloadLinks map[string]string
		glob          []string
	}{downloadLinks, globCopy})
	fake.guard("DownloadLinksByGlob")
	fake.invocations["DownloadLinksByGlob"] = append(fake.invocations["DownloadLinksByGlob"], []interface{}{downloadLinks, globCopy})
	fake.downloadLinksByGlobMutex.Unlock()
	if fake.DownloadLinksByGlobStub != nil {
		return fake.DownloadLinksByGlobStub(downloadLinks, glob)
	} else {
		return fake.downloadLinksByGlobReturns.result1, fake.downloadLinksByGlobReturns.result2
	}
}

func (fake *FakeFilter) DownloadLinksByGlobCallCount() int {
	fake.downloadLinksByGlobMutex.RLock()
	defer fake.downloadLinksByGlobMutex.RUnlock()
	return len(fake.downloadLinksByGlobArgsForCall)
}

func (fake *FakeFilter) DownloadLinksByGlobArgsForCall(i int) (map[string]string, []string) {
	fake.downloadLinksByGlobMutex.RLock()
	defer fake.downloadLinksByGlobMutex.RUnlock()
	return fake.downloadLinksByGlobArgsForCall[i].downloadLinks, fake.downloadLinksByGlobArgsForCall[i].glob
}

func (fake *FakeFilter) DownloadLinksByGlobReturns(result1 map[string]string, result2 error) {
	fake.DownloadLinksByGlobStub = nil
	fake.downloadLinksByGlobReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeFilter) DownloadLinks(p []pivnet.ProductFile) map[string]string {
	var pCopy []pivnet.ProductFile
	if p != nil {
		pCopy = make([]pivnet.ProductFile, len(p))
		copy(pCopy, p)
	}
	fake.downloadLinksMutex.Lock()
	fake.downloadLinksArgsForCall = append(fake.downloadLinksArgsForCall, struct {
		p []pivnet.ProductFile
	}{pCopy})
	fake.guard("DownloadLinks")
	fake.invocations["DownloadLinks"] = append(fake.invocations["DownloadLinks"], []interface{}{pCopy})
	fake.downloadLinksMutex.Unlock()
	if fake.DownloadLinksStub != nil {
		return fake.DownloadLinksStub(p)
	} else {
		return fake.downloadLinksReturns.result1
	}
}

func (fake *FakeFilter) DownloadLinksCallCount() int {
	fake.downloadLinksMutex.RLock()
	defer fake.downloadLinksMutex.RUnlock()
	return len(fake.downloadLinksArgsForCall)
}

func (fake *FakeFilter) DownloadLinksArgsForCall(i int) []pivnet.ProductFile {
	fake.downloadLinksMutex.RLock()
	defer fake.downloadLinksMutex.RUnlock()
	return fake.downloadLinksArgsForCall[i].p
}

func (fake *FakeFilter) DownloadLinksReturns(result1 map[string]string) {
	fake.DownloadLinksStub = nil
	fake.downloadLinksReturns = struct {
		result1 map[string]string
	}{result1}
}

func (fake *FakeFilter) Invocations() map[string][][]interface{} {
	return fake.invocations
}

func (fake *FakeFilter) guard(key string) {
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
}

var _ pivnet_filter.Filter = new(FakeFilter)
