/*

Copyright 2020 The Vouch Proxy Authors.
Use of this source code is governed by The MIT License (MIT) that
can be found in the LICENSE file. Software distributed under The
MIT License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied.

*/

package capturewriter

import (
	"net/http"

	"go.uber.org/zap"
)

// we wrap ResponseWriter so that we can store the StatusCode
// and then pull it out later for logging
// https://play.golang.org/p/wPHaX9DH-Ik

// var logger *zap.SugaredLogger
var log *zap.Logger

// Configure see main.go configure()
func Configure() {
	_ = "STUB: not implemented"
	// logger = cfg.Logging.Logger
	return
}

// CaptureWriter extends http.ResponseWriter
type CaptureWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (w *CaptureWriter) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// log.Debug("CaptureWriter.Write set w.StatusCode " + strconv.Itoa(w.StatusCode))

// Header calls http.Writer.Header()
func (w *CaptureWriter) Header() http.Header { _ = "STUB: not implemented"; return *new(http.Header) }

// WriteHeader calls http.Writer.WriteHeader(code)
func (w *CaptureWriter) WriteHeader(code int) { _ = "STUB: not implemented"; return }

// GetStatusCode return w.StatusCode
func (w *CaptureWriter) GetStatusCode() int { _ = "STUB: not implemented"; return 0 }
