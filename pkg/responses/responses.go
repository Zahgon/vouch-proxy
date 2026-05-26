/*

Copyright 2020 The Vouch Proxy Authors.
Use of this source code is governed by The MIT License (MIT) that
can be found in the LICENSE file. Software distributed under The
MIT License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied.

*/

package responses

import (
	"html/template"
	"net/http"

	"go.uber.org/zap"
)

// Index variables passed to index.tmpl
type Index struct {
	Msg          string
	TestURLs     []string
	Testing      bool
	DocumentRoot string
}

var (
	indexTemplate *template.Template
	log           *zap.SugaredLogger
	// fastlog       *zap.Logger

	// errorTemplate *template.Template
	// errNotAuthorized = errors.New("not authorized")
)

// Configure see main.go configure()
func Configure() { _ = "STUB: not implemented"; return }

// fastlog = cfg.Logging.FastLogger

// RenderIndex render the response as an HTML page, mostly used in testing
func RenderIndex(w http.ResponseWriter, msg string) { _ = "STUB: not implemented"; return }

// renderError html error page
// something terse for the end user
func renderError(w http.ResponseWriter, msg string, status int) { _ = "STUB: not implemented"; return }

// OK200 returns "200 OK"
func OK200(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// Redirect302 redirect to the specified rURL
func Redirect302(w http.ResponseWriter, r *http.Request, rURL string) {
	_ = "STUB: not implemented"
	return
}

// Error400 Bad Request
func Error400(w http.ResponseWriter, r *http.Request, e error) { _ = "STUB: not implemented"; return }

// Error401 Unauthorized, the standard error returned when failing /validate
// this is captured by nginx, which converts the 401 into 302 to the login page
func Error401(w http.ResponseWriter, r *http.Request, e error) { _ = "STUB: not implemented"; return }

// renderError(w, "401 Unauthorized")

// Error401HTTP
func Error401HTTP(w http.ResponseWriter, r *http.Request, e error) {
	_ = "STUB: not implemented"
	return
}

// Error403 Forbidden
// if there's an error during /auth or if they don't pass validation in /auth
func Error403(w http.ResponseWriter, r *http.Request, e error) { _ = "STUB: not implemented"; return }

// Error500 Internal Error
// something is not right, hopefully this never happens
func Error500(w http.ResponseWriter, r *http.Request, e error) { _ = "STUB: not implemented"; return }

// cancelClearSetError convenience method to keep it DRY
func cancelClearSetError(w http.ResponseWriter, r *http.Request, e error) {
	_ = "STUB: not implemented"
	return
}

// cfg.ErrCtx is tested by `jwtmanager.JWTCacheHandler`
func addErrandCancelRequest(r *http.Request) { _ = "STUB: not implemented"; return }

// we're done
