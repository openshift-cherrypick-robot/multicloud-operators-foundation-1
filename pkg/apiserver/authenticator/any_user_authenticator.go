// licensed Materials - Property of IBM
// 5737-E67
// (C) Copyright IBM Corporation 2016, 2019 All Rights Reserved
// US Government Users Restricted Rights - Use, duplication or disclosure restricted by GSA ADP Schedule Contract with IBM Corp.

/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package authenticator

import (
	"net/http"

	"k8s.io/apiserver/pkg/authentication/authenticator"
	"k8s.io/apiserver/pkg/authentication/user"
)

// AnyUserAuthenticator is an authenticator that can be used in tests to
// allow for test users to be added to the context. The user name used
// is from the Basic auth of the http request.
type AnyUserAuthenticator struct{}

var _ authenticator.Request = &AnyUserAuthenticator{}

// AuthenticateRequest gets the user name from the Basic auth. If there is
// no basic auth or it is not valid for any reason, then an empty username
// is used.
func (a *AnyUserAuthenticator) AuthenticateRequest(req *http.Request) (*authenticator.Response, bool, error) {
	username, _, _ := req.BasicAuth()
	userInfo := &user.DefaultInfo{
		Name: username,
	}

	return &authenticator.Response{
		User: userInfo,
	}, true, nil
}
