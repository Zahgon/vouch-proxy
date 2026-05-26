/*

Copyright 2020 The Vouch Proxy Authors.
Use of this source code is governed by The MIT License (MIT) that
can be found in the LICENSE file. Software distributed under The
MIT License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied.

*/

package timelog

import (
	"net/http"

	"go.uber.org/zap"
)

var (
	req        = int64(0)
	avgLatency = int64(0)
	log        *zap.SugaredLogger
)

// Configure see main.go configure()
func Configure() { _ = "STUB: not implemented"; return }

// TimeLog records how long it takes to process the http request and produce the response (latency)
func TimeLog(nextHandler http.Handler) func(http.ResponseWriter, *http.Request) {
	_ = "STUB: not implemented"
	return nil
}

// log.Debugf("Request received : %v", r)

// make the call

// Stop timer

// log.Debugf("Request handled successfully: %v", v.GetStatusCode())
