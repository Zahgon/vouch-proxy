/*

Copyright 2020 The Vouch Proxy Authors.
Use of this source code is governed by The MIT License (MIT) that
can be found in the LICENSE file. Software distributed under The
MIT License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied.

*/

package common

import (
	"net/http"

	"go.uber.org/zap"
	"golang.org/x/oauth2"

	"github.com/vouch/vouch-proxy/pkg/structs"
)

var log *zap.SugaredLogger

// Configure see main.go configure()
func Configure() { _ = "STUB: not implemented"; return }

// PrepareTokensAndClient setup the client, usually for a UserInfo request
func PrepareTokensAndClient(r *http.Request, ptokens *structs.PTokens, setProviderToken bool, opts ...oauth2.AuthCodeOption) (*http.Client, *oauth2.Token, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Certain providers (eg. gitea) don't provide an id_token
// and it's not necessary for the authentication phase

// MapClaims populate CustomClaims from userInfo for each configure claims header
func MapClaims(claims []byte, customClaims *structs.CustomClaims) error {
	_ = "STUB: not implemented"
	return nil
}
