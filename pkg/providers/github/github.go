/*

Copyright 2020 The Vouch Proxy Authors.
Use of this source code is governed by The MIT License (MIT) that
can be found in the LICENSE file. Software distributed under The
MIT License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied.

*/

package github

import (
	"net/http"

	"github.com/vouch/vouch-proxy/pkg/structs"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

// Provider provider specific functions
type Provider struct {
	PrepareTokensAndClient func(r *http.Request, ptokens *structs.PTokens, setProviderToken bool, opts ...oauth2.AuthCodeOption) (*http.Client, *oauth2.Token, error)
}

var log *zap.SugaredLogger

// Configure see main.go configure()
func (Provider) Configure() { _ = "STUB: not implemented"; return }

// GetUserInfo github user info, calls github api for org and teams
func (me Provider) GetUserInfo(r *http.Request, user *structs.User, customClaims *structs.CustomClaims, ptokens *structs.PTokens, opts ...oauth2.AuthCodeOption) (rerr error) {
	_ = "STUB: not implemented"
	return nil
}

// user = &ghUser.User

// only organization given

func getOrgMembershipStateFromGitHub(client *http.Client, user *structs.User, orgID string) (isMember bool, rerr error) {
	_ = "STUB: not implemented"
	return false, nil
}

func getTeamMembershipStateFromGitHub(client *http.Client, user *structs.User, orgID string, team string) (isMember bool, rerr error) {
	_ = "STUB: not implemented"
	return false, nil
}
