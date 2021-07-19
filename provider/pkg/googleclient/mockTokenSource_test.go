// Copyright 2021, Pulumi Corporation.  All rights reserved.

package googleclient

import (
	"github.com/stretchr/testify/mock"
	"golang.org/x/oauth2"
)

// MockTokenSource is a mock type for the oauth2.TokenSource type
type MockTokenSource struct {
	mock.Mock
}

// Token provides a mock function with given fields:
func (_m *MockTokenSource) Token() (*oauth2.Token, error) {
	ret := _m.Called()

	var r0 *oauth2.Token
	if rf, ok := ret.Get(0).(func() *oauth2.Token); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oauth2.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
