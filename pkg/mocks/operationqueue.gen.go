// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/trustbloc/sidetree-core-go/pkg/api/batch"
)

type OperationQueue struct {
	AddStub        func(data *batch.OperationInfo) (uint, error)
	addMutex       sync.RWMutex
	addArgsForCall []struct {
		data *batch.OperationInfo
	}
	addReturns struct {
		result1 uint
		result2 error
	}
	addReturnsOnCall map[int]struct {
		result1 uint
		result2 error
	}
	RemoveStub        func(num uint) (uint, uint, error)
	removeMutex       sync.RWMutex
	removeArgsForCall []struct {
		num uint
	}
	removeReturns struct {
		result1 uint
		result2 uint
		result3 error
	}
	removeReturnsOnCall map[int]struct {
		result1 uint
		result2 uint
		result3 error
	}
	PeekStub        func(num uint) ([]*batch.OperationInfo, error)
	peekMutex       sync.RWMutex
	peekArgsForCall []struct {
		num uint
	}
	peekReturns struct {
		result1 []*batch.OperationInfo
		result2 error
	}
	peekReturnsOnCall map[int]struct {
		result1 []*batch.OperationInfo
		result2 error
	}
	LenStub        func() uint
	lenMutex       sync.RWMutex
	lenArgsForCall []struct{}
	lenReturns     struct {
		result1 uint
	}
	lenReturnsOnCall map[int]struct {
		result1 uint
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *OperationQueue) Add(data *batch.OperationInfo) (uint, error) {
	fake.addMutex.Lock()
	ret, specificReturn := fake.addReturnsOnCall[len(fake.addArgsForCall)]
	fake.addArgsForCall = append(fake.addArgsForCall, struct {
		data *batch.OperationInfo
	}{data})
	fake.recordInvocation("Add", []interface{}{data})
	fake.addMutex.Unlock()
	if fake.AddStub != nil {
		return fake.AddStub(data)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.addReturns.result1, fake.addReturns.result2
}

func (fake *OperationQueue) AddCallCount() int {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return len(fake.addArgsForCall)
}

func (fake *OperationQueue) AddArgsForCall(i int) *batch.OperationInfo {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return fake.addArgsForCall[i].data
}

func (fake *OperationQueue) AddReturns(result1 uint, result2 error) {
	fake.AddStub = nil
	fake.addReturns = struct {
		result1 uint
		result2 error
	}{result1, result2}
}

func (fake *OperationQueue) AddReturnsOnCall(i int, result1 uint, result2 error) {
	fake.AddStub = nil
	if fake.addReturnsOnCall == nil {
		fake.addReturnsOnCall = make(map[int]struct {
			result1 uint
			result2 error
		})
	}
	fake.addReturnsOnCall[i] = struct {
		result1 uint
		result2 error
	}{result1, result2}
}

func (fake *OperationQueue) Remove(num uint) (uint, uint, error) {
	fake.removeMutex.Lock()
	ret, specificReturn := fake.removeReturnsOnCall[len(fake.removeArgsForCall)]
	fake.removeArgsForCall = append(fake.removeArgsForCall, struct {
		num uint
	}{num})
	fake.recordInvocation("Remove", []interface{}{num})
	fake.removeMutex.Unlock()
	if fake.RemoveStub != nil {
		return fake.RemoveStub(num)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.removeReturns.result1, fake.removeReturns.result2, fake.removeReturns.result3
}

func (fake *OperationQueue) RemoveCallCount() int {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return len(fake.removeArgsForCall)
}

func (fake *OperationQueue) RemoveArgsForCall(i int) uint {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return fake.removeArgsForCall[i].num
}

func (fake *OperationQueue) RemoveReturns(result1 uint, result2 uint, result3 error) {
	fake.RemoveStub = nil
	fake.removeReturns = struct {
		result1 uint
		result2 uint
		result3 error
	}{result1, result2, result3}
}

func (fake *OperationQueue) RemoveReturnsOnCall(i int, result1 uint, result2 uint, result3 error) {
	fake.RemoveStub = nil
	if fake.removeReturnsOnCall == nil {
		fake.removeReturnsOnCall = make(map[int]struct {
			result1 uint
			result2 uint
			result3 error
		})
	}
	fake.removeReturnsOnCall[i] = struct {
		result1 uint
		result2 uint
		result3 error
	}{result1, result2, result3}
}

func (fake *OperationQueue) Peek(num uint) ([]*batch.OperationInfo, error) {
	fake.peekMutex.Lock()
	ret, specificReturn := fake.peekReturnsOnCall[len(fake.peekArgsForCall)]
	fake.peekArgsForCall = append(fake.peekArgsForCall, struct {
		num uint
	}{num})
	fake.recordInvocation("Peek", []interface{}{num})
	fake.peekMutex.Unlock()
	if fake.PeekStub != nil {
		return fake.PeekStub(num)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.peekReturns.result1, fake.peekReturns.result2
}

func (fake *OperationQueue) PeekCallCount() int {
	fake.peekMutex.RLock()
	defer fake.peekMutex.RUnlock()
	return len(fake.peekArgsForCall)
}

func (fake *OperationQueue) PeekArgsForCall(i int) uint {
	fake.peekMutex.RLock()
	defer fake.peekMutex.RUnlock()
	return fake.peekArgsForCall[i].num
}

func (fake *OperationQueue) PeekReturns(result1 []*batch.OperationInfo, result2 error) {
	fake.PeekStub = nil
	fake.peekReturns = struct {
		result1 []*batch.OperationInfo
		result2 error
	}{result1, result2}
}

func (fake *OperationQueue) PeekReturnsOnCall(i int, result1 []*batch.OperationInfo, result2 error) {
	fake.PeekStub = nil
	if fake.peekReturnsOnCall == nil {
		fake.peekReturnsOnCall = make(map[int]struct {
			result1 []*batch.OperationInfo
			result2 error
		})
	}
	fake.peekReturnsOnCall[i] = struct {
		result1 []*batch.OperationInfo
		result2 error
	}{result1, result2}
}

func (fake *OperationQueue) Len() uint {
	fake.lenMutex.Lock()
	ret, specificReturn := fake.lenReturnsOnCall[len(fake.lenArgsForCall)]
	fake.lenArgsForCall = append(fake.lenArgsForCall, struct{}{})
	fake.recordInvocation("Len", []interface{}{})
	fake.lenMutex.Unlock()
	if fake.LenStub != nil {
		return fake.LenStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.lenReturns.result1
}

func (fake *OperationQueue) LenCallCount() int {
	fake.lenMutex.RLock()
	defer fake.lenMutex.RUnlock()
	return len(fake.lenArgsForCall)
}

func (fake *OperationQueue) LenReturns(result1 uint) {
	fake.LenStub = nil
	fake.lenReturns = struct {
		result1 uint
	}{result1}
}

func (fake *OperationQueue) LenReturnsOnCall(i int, result1 uint) {
	fake.LenStub = nil
	if fake.lenReturnsOnCall == nil {
		fake.lenReturnsOnCall = make(map[int]struct {
			result1 uint
		})
	}
	fake.lenReturnsOnCall[i] = struct {
		result1 uint
	}{result1}
}

func (fake *OperationQueue) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	fake.peekMutex.RLock()
	defer fake.peekMutex.RUnlock()
	fake.lenMutex.RLock()
	defer fake.lenMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *OperationQueue) recordInvocation(key string, args []interface{}) {
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
