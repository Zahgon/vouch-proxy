/*

Copyright 2020 The Vouch Proxy Authors.
Use of this source code is governed by The MIT License (MIT) that
can be found in the LICENSE file. Software distributed under The
MIT License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied.

*/

package handlers

import (
	"fmt"
	"net/http"

	"github.com/vouch/vouch-proxy/pkg/cfg"
)

var errUnauthRedirURL = fmt.Errorf("/logout The requested url is not present in `%s.post_logout_redirect_uris`", cfg.Branding.LCName)

// LogoutHandler /logout
// Destroys Vouch session
// If oauth.end_session_endpoint present in conf, also redirects to destroy session at oauth provider
// If "url" param present in request, also redirects to that (after destroying one or both sessions)
func LogoutHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// Make sure that redirectURL, if given, is allowed by config

// If provider logout URL is configured, redirect to it (and pass redirectURL along)
// If provider logout URL is not configured, redirect directly to redirectURL

// Optional in spec, required by some providers (Okta, for example)
