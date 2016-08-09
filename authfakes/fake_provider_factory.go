// This file was generated by counterfeiter
package authfakes

import (
	"sync"

	"github.com/concourse/atc/auth"
	"github.com/concourse/atc/auth/provider"
	"github.com/concourse/atc/db"
)

type FakeProviderFactory struct {
	GetProvidersStub        func(db.SavedTeam) (provider.Providers, error)
	getProvidersMutex       sync.RWMutex
	getProvidersArgsForCall []struct {
		arg1 db.SavedTeam
	}
	getProvidersReturns struct {
		result1 provider.Providers
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeProviderFactory) GetProviders(arg1 db.SavedTeam) (provider.Providers, error) {
	fake.getProvidersMutex.Lock()
	fake.getProvidersArgsForCall = append(fake.getProvidersArgsForCall, struct {
		arg1 db.SavedTeam
	}{arg1})
	fake.recordInvocation("GetProviders", []interface{}{arg1})
	fake.getProvidersMutex.Unlock()
	if fake.GetProvidersStub != nil {
		return fake.GetProvidersStub(arg1)
	} else {
		return fake.getProvidersReturns.result1, fake.getProvidersReturns.result2
	}
}

func (fake *FakeProviderFactory) GetProvidersCallCount() int {
	fake.getProvidersMutex.RLock()
	defer fake.getProvidersMutex.RUnlock()
	return len(fake.getProvidersArgsForCall)
}

func (fake *FakeProviderFactory) GetProvidersArgsForCall(i int) db.SavedTeam {
	fake.getProvidersMutex.RLock()
	defer fake.getProvidersMutex.RUnlock()
	return fake.getProvidersArgsForCall[i].arg1
}

func (fake *FakeProviderFactory) GetProvidersReturns(result1 provider.Providers, result2 error) {
	fake.GetProvidersStub = nil
	fake.getProvidersReturns = struct {
		result1 provider.Providers
		result2 error
	}{result1, result2}
}

func (fake *FakeProviderFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getProvidersMutex.RLock()
	defer fake.getProvidersMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeProviderFactory) recordInvocation(key string, args []interface{}) {
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

var _ auth.ProviderFactory = new(FakeProviderFactory)
