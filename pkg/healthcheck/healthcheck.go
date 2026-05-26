/*

Copyright 2020 The Vouch Proxy Authors.
Use of this source code is governed by The MIT License (MIT) that
can be found in the LICENSE file. Software distributed under The
MIT License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied.

*/

package healthcheck

import (
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func configure() {
	_ = "STUB: not implemented"
	// cfg.ConfigureLogger()
	return
}

// CheckAndExitIfIsHealthCheck healthcheck is a command line flag `-healthcheck`
func CheckAndExitIfIsHealthCheck() { _ = "STUB: not implemented"; return }

func healthcheck() { _ = "STUB: not implemented"; return }

// #nosec - turn off gosec checking which flags `http.Get(url)`
