/*

Copyright 2020 The Vouch Proxy Authors.
Use of this source code is governed by The MIT License (MIT) that
can be found in the LICENSE file. Software distributed under The
MIT License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied.

*/

package handlers

import (
	"errors"
	"net/http"

	"github.com/vouch/vouch-proxy/pkg/jwtmanager"
)

var (
	errNoJWT  = errors.New("no jwt found in request")
	errNoUser = errors.New("no User found in jwt")
)

// ValidateRequestHandler /validate
func ValidateRequestHandler(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// fastlog.Debugf("response headers %+v", w.Header())
// fastlog.Debug("response header",
// 	zap.String(cfg.Cfg.Headers.User, w.Header().Get(cfg.Cfg.Headers.User)))

// good to go!!

func generateCustomClaimsHeaders(w http.ResponseWriter, claims *jwtmanager.VouchClaims) {
	_ = "STUB: not implemented"
	return
}

// Run through all the claims found

// Run through the claims we are looking for

// Check for matching claim

// convert to string

// if val, ok := v.(string); ok {

func send401or200PublicAccess(w http.ResponseWriter, r *http.Request, e error) {
	_ = "STUB: not implemented"
	return
}
