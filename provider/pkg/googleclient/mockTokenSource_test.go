// Copyright 2016-2022, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
