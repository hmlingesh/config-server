// Code generated by counterfeiter. DO NOT EDIT.
package typesfakes

import (
	"sync"

	"github.com/cloudfoundry/config-server/types"
)

type FakeValueGenerator struct {
	GenerateStub        func(interface{}) (interface{}, error)
	generateMutex       sync.RWMutex
	generateArgsForCall []struct {
		arg1 interface{}
	}
	generateReturns struct {
		result1 interface{}
		result2 error
	}
	generateReturnsOnCall map[int]struct {
		result1 interface{}
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeValueGenerator) Generate(arg1 interface{}) (interface{}, error) {
	fake.generateMutex.Lock()
	ret, specificReturn := fake.generateReturnsOnCall[len(fake.generateArgsForCall)]
	fake.generateArgsForCall = append(fake.generateArgsForCall, struct {
		arg1 interface{}
	}{arg1})
	fake.recordInvocation("Generate", []interface{}{arg1})
	fake.generateMutex.Unlock()
	if fake.GenerateStub != nil {
		return fake.GenerateStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.generateReturns.result1, fake.generateReturns.result2
}

func (fake *FakeValueGenerator) GenerateCallCount() int {
	fake.generateMutex.RLock()
	defer fake.generateMutex.RUnlock()
	return len(fake.generateArgsForCall)
}

func (fake *FakeValueGenerator) GenerateArgsForCall(i int) interface{} {
	fake.generateMutex.RLock()
	defer fake.generateMutex.RUnlock()
	return fake.generateArgsForCall[i].arg1
}

func (fake *FakeValueGenerator) GenerateReturns(result1 interface{}, result2 error) {
	fake.GenerateStub = nil
	fake.generateReturns = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeValueGenerator) GenerateReturnsOnCall(i int, result1 interface{}, result2 error) {
	fake.GenerateStub = nil
	if fake.generateReturnsOnCall == nil {
		fake.generateReturnsOnCall = make(map[int]struct {
			result1 interface{}
			result2 error
		})
	}
	fake.generateReturnsOnCall[i] = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeValueGenerator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.generateMutex.RLock()
	defer fake.generateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeValueGenerator) recordInvocation(key string, args []interface{}) {
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

var _ types.ValueGenerator = new(FakeValueGenerator)
