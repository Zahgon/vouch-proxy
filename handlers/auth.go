/*

Copyright 2020 The Vouch Proxy Authors.
Use of this source code is governed by The MIT License (MIT) that
can be found in the LICENSE file. Software distributed under The
MIT License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied.

*/

package handlers

import (
	"net/http"

	"github.com/vouch/vouch-proxy/pkg/structs"

	"golang.org/x/oauth2"
)

// CallbackHandler /auth
// - redirects to /auth/{state}/ with the state coming from the query parameter
func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"

	// did the IdP return an error?
	return
}

// has to have a trailing / in its path, because the path of the session cookie is set to /auth/{state}/.
// see note in login.go and https://github.com/vouch/vouch-proxy/issues/373

// AuthStateHandler /auth/{state}/
// - validate info from oauth provider (Google, GitHub, OIDC, etc)
// - issue jwt in the form of a cookie
func AuthStateHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// Handle the exchange code to initiate a transport.

// is the nonce "state" valid?

// is code challenge enabled?

// verify / authz the user

// SUCCESS!! they are authorized

// issue the jwt

// get the originally requested URL so we can send them on their way

// clear out the session value

// otherwise serve an error

// verifyUser validates that the domains match for the user
func verifyUser(u interface{}) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// AllowAllUsers

// WhiteList

// TeamWhiteList

// Domains

// nothing configured, allow everyone through

func getUserInfo(r *http.Request, user *structs.User, customClaims *structs.CustomClaims, ptokens *structs.PTokens, opts ...oauth2.AuthCodeOption) error {
	_ = "STUB: not implemented"
	return nil
}
