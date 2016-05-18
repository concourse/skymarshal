// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc/auth"
	"github.com/concourse/atc/db"
)

type FakeAuthDB struct {
	GetTeamStub        func() (db.SavedTeam, bool, error)
	getTeamMutex       sync.RWMutex
	getTeamArgsForCall []struct{}
	getTeamReturns     struct {
		result1 db.SavedTeam
		result2 bool
		result3 error
	}
}

func (fake *FakeAuthDB) GetTeam() (db.SavedTeam, bool, error) {
	fake.getTeamMutex.Lock()
	fake.getTeamArgsForCall = append(fake.getTeamArgsForCall, struct{}{})
	fake.getTeamMutex.Unlock()
	if fake.GetTeamStub != nil {
		return fake.GetTeamStub()
	} else {
		return fake.getTeamReturns.result1, fake.getTeamReturns.result2, fake.getTeamReturns.result3
	}
}

func (fake *FakeAuthDB) GetTeamCallCount() int {
	fake.getTeamMutex.RLock()
	defer fake.getTeamMutex.RUnlock()
	return len(fake.getTeamArgsForCall)
}

func (fake *FakeAuthDB) GetTeamReturns(result1 db.SavedTeam, result2 bool, result3 error) {
	fake.GetTeamStub = nil
	fake.getTeamReturns = struct {
		result1 db.SavedTeam
		result2 bool
		result3 error
	}{result1, result2, result3}
}

var _ auth.AuthDB = new(FakeAuthDB)
