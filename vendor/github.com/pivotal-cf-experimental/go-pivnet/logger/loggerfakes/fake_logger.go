// This file was generated by counterfeiter
package loggerfakes

import (
	"sync"

	"github.com/pivotal-cf-experimental/pivnet-resource/vendor/github.com/pivotal-cf-experimental/go-pivnet/logger"
)

type FakeLogger struct {
	DebugStub        func(action string, data ...logger.Data)
	debugMutex       sync.RWMutex
	debugArgsForCall []struct {
		action string
		data   []logger.Data
	}
	InfoStub        func(action string, data ...logger.Data)
	infoMutex       sync.RWMutex
	infoArgsForCall []struct {
		action string
		data   []logger.Data
	}
	invocations map[string][][]interface{}
}

func (fake *FakeLogger) Debug(action string, data ...logger.Data) {
	fake.debugMutex.Lock()
	fake.debugArgsForCall = append(fake.debugArgsForCall, struct {
		action string
		data   []logger.Data
	}{action, data})
	fake.guard("Debug")
	fake.invocations["Debug"] = append(fake.invocations["Debug"], []interface{}{action, data})
	fake.debugMutex.Unlock()
	if fake.DebugStub != nil {
		fake.DebugStub(action, data...)
	}
}

func (fake *FakeLogger) DebugCallCount() int {
	fake.debugMutex.RLock()
	defer fake.debugMutex.RUnlock()
	return len(fake.debugArgsForCall)
}

func (fake *FakeLogger) DebugArgsForCall(i int) (string, []logger.Data) {
	fake.debugMutex.RLock()
	defer fake.debugMutex.RUnlock()
	return fake.debugArgsForCall[i].action, fake.debugArgsForCall[i].data
}

func (fake *FakeLogger) Info(action string, data ...logger.Data) {
	fake.infoMutex.Lock()
	fake.infoArgsForCall = append(fake.infoArgsForCall, struct {
		action string
		data   []logger.Data
	}{action, data})
	fake.guard("Info")
	fake.invocations["Info"] = append(fake.invocations["Info"], []interface{}{action, data})
	fake.infoMutex.Unlock()
	if fake.InfoStub != nil {
		fake.InfoStub(action, data...)
	}
}

func (fake *FakeLogger) InfoCallCount() int {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return len(fake.infoArgsForCall)
}

func (fake *FakeLogger) InfoArgsForCall(i int) (string, []logger.Data) {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return fake.infoArgsForCall[i].action, fake.infoArgsForCall[i].data
}

func (fake *FakeLogger) Invocations() map[string][][]interface{} {
	return fake.invocations
}

func (fake *FakeLogger) guard(key string) {
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
}

var _ logger.Logger = new(FakeLogger)
