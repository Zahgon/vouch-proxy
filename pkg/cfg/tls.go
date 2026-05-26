/*

Copyright 2020 The Vouch Proxy Authors.
Use of this source code is governed by The MIT License (MIT) that
can be found in the LICENSE file. Software distributed under The
MIT License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied.

*/

package cfg

import (
	"crypto/tls"
)

// TLSConfig config returns a *tls.Config with the specified profile (modern, intermediate, old, default) configuration.
func TLSConfig(profile string) *tls.Config {
	_ = "STUB: not implemented"

	// Source: https://ssl-config.mozilla.org/#server=go&version=1.14&config=modern&hsts=false&guideline=5.6
	return nil
}
