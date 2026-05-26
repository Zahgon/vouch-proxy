/*

Copyright 2020 The Vouch Proxy Authors.
Use of this source code is governed by The MIT License (MIT) that
can be found in the LICENSE file. Software distributed under The
MIT License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied.

*/

package cookie

import (
	"net/http"

	// "github.com/vouch/vouch-proxy/pkg/structs"

	"go.uber.org/zap"
)

const maxCookieSize = 4000

var log *zap.SugaredLogger
var sameSite http.SameSite

// Configure see main.go configure()
func Configure() { _ = "STUB: not implemented"; return }

// SetCookie http
func SetCookie(w http.ResponseWriter, r *http.Request, val string) {
	_ = "STUB: not implemented"
	return
}

// convert minutes to seconds

func setCookie(w http.ResponseWriter, r *http.Request, val string, maxAge int) {
	_ = "STUB: not implemented"
	return
}

// foreach domain

// Allow overriding the cookie domain in the config file

// Cookies have a max size of 4096 bytes, but to support most browsers, we should stay below 4000 bytes
// https://tools.ietf.org/html/rfc6265#section-6.1
// http://browsercookielimits.squawky.net/

// https://www.lifewire.com/cookie-limit-per-domain-3466809

// Cookies are named 1of3, 2of3, 3of3

// Cookie get the vouch jwt cookie
func Cookie(r *http.Request) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Get the remaining parts
// search for cookie parts in order
// this is the hotpath so we're trying to only walk once

// then its uninitialized

// combinedCookieStr := combinedCookie.String()

// ClearCookie get rid of the existing cookie
func ClearCookie(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// Allow overriding the cookie domain in the config file

// search for cookie parts

// https://stackoverflow.com/questions/5285940/correct-way-to-delete-cookies-server-side
// system dependent but usually Thu, 01 Jan 1970 00:00:00 GMT

// SameSite return cfg.Cfg.Cookie.SameSite as http.Samesite
// if cfg.Cfg.Cookie.SameSite is unconfigured return http.SameSite(0)
// see https://github.com/vouch/vouch-proxy/issues/210
func SameSite() http.SameSite { _ = "STUB: not implemented"; return *new(http.SameSite) }

// splitCookie separate string into several strings of specified length
func splitCookie(longString string, maxLen int) []string { _ = "STUB: not implemented"; return nil }
