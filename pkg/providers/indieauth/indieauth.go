/*

Copyright 2020 The Vouch Proxy Authors.
Use of this source code is governed by The MIT License (MIT) that
can be found in the LICENSE file. Software distributed under The
MIT License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied.

*/

package indieauth

import (
	"net/http"

	"golang.org/x/oauth2"

	"github.com/vouch/vouch-proxy/pkg/structs"
	"go.uber.org/zap"
)

// Provider provider specific functions
type Provider struct{}

var log *zap.SugaredLogger

// Configure see main.go configure()
func (Provider) Configure() { _ = "STUB: not implemented"; return }

// GetUserInfo provider specific call to get userinfomation
func (Provider) GetUserInfo(r *http.Request, user *structs.User, customClaims *structs.CustomClaims, ptokens *structs.PTokens, opts ...oauth2.AuthCodeOption) (rerr error) {
	_ = "STUB: not implemented"
	// indieauth sends the "me" setting in json back to the callback, so just pluck it from the callback
	return nil
}

// v.Set("code", code)

// v.Set("redirect_uri", cfg.GenOAuth.RedirectURL)

// v.Set("client_id", cfg.GenOAuth.ClientID)

// v := url.Values{}
// userinfo, err := client.PostForm(cfg.GenOAuth.UserInfoURL, v)

// http.Error(w, err.Error(), http.StatusBadRequest)
