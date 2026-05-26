/*

Copyright 2020 The Vouch Proxy Authors.
Use of this source code is governed by The MIT License (MIT) that
can be found in the LICENSE file. Software distributed under The
MIT License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied.

*/

package cfg

import (
	"golang.org/x/oauth2"
)

var (
	// GenOAuth exported OAuth config variable
	// TODO: GenOAuth and OAuthClient should be combined
	GenOAuth = &oauthConfig{}

	// OAuthClient is the configured client which will call the provider
	// this actually carries the oauth2 client ala oauthclient.Client(oauth2.NoContext, providerToken)
	OAuthClient *oauth2.Config
	// OAuthopts authentication options
	OAuthopts []oauth2.AuthCodeOption

	// Providers static strings to test against
	Providers = &OAuthProviders{
		Google:        "google",
		GitHub:        "github",
		IndieAuth:     "indieauth",
		ADFS:          "adfs",
		Azure:         "azure",
		OIDC:          "oidc",
		HomeAssistant: "homeassistant",
		OpenStax:      "openstax",
		Nextcloud:     "nextcloud",
		Alibaba:       "alibaba",
		Discord:       "discord",
	}
)

// OAuthProviders holds the stings for
type OAuthProviders struct {
	Google        string
	GitHub        string
	IndieAuth     string
	ADFS          string
	Azure         string
	OIDC          string
	HomeAssistant string
	OpenStax      string
	Nextcloud     string
	Alibaba       string
	Discord       string
}

// oauth config items endoint for access
// `envconfig` tag is for env var support
// https://github.com/kelseyhightower/envconfig
type oauthConfig struct {
	Provider       string   `mapstructure:"provider"`
	ClientID       string   `mapstructure:"client_id" envconfig:"client_id"`
	ClientSecret   string   `mapstructure:"client_secret" envconfig:"client_secret"`
	AuthURL        string   `mapstructure:"auth_url" envconfig:"auth_url"`
	TokenURL       string   `mapstructure:"token_url" envconfig:"token_url"`
	LogoutURL      string   `mapstructure:"end_session_endpoint"  envconfig:"end_session_endpoint"`
	RedirectURL    string   `mapstructure:"callback_url"  envconfig:"callback_url"`
	RedirectURLs   []string `mapstructure:"callback_urls"  envconfig:"callback_urls"`
	RelyingPartyId string   `mapstructure:"relying_party_id"  envconfig:"relying_party_id"`
	Scopes         []string `mapstructure:"scopes"`
	// pointer-to-pointer so that the default uninitialized value is nil
	Claims              **oauthClaimsConfig `mapstructure:"claims"`
	UserInfoURL         string              `mapstructure:"user_info_url" envconfig:"user_info_url"`
	UserTeamURL         string              `mapstructure:"user_team_url" envconfig:"user_team_url"`
	UserOrgURL          string              `mapstructure:"user_org_url" envconfig:"user_org_url"`
	PreferredDomain     string              `mapstructure:"preferredDomain"`
	AzureToken          string              `mapstructure:"azure_token" envconfig:"azure_token"`
	CodeChallengeMethod string              `mapstructure:"code_challenge_method" envconfig:"code_challenge_method"`
	// DiscordUseIDs defaults to false, maintaining the more common username checking behavior
	// If set to true, match the Discord user's ID instead of their username
	DiscordUseIDs bool `mapstructure:"discord_use_ids" envconfig:"discord_use_ids"`
}

type oauthClaimsConfig struct {
	UserInfo map[string]*oauthClaimValueConfig `mapstructure:"userinfo" json:"userinfo,omitempty"`
	IDToken  map[string]*oauthClaimValueConfig `mapstructure:"id_token" json:"id_token,omitempty"`
}

type oauthClaimValueConfig struct {
	Essential bool          `mapstructure:"essential" json:"essential,omitempty"`
	Value     interface{}   `mapstructure:"value" json:"value,omitempty"`
	Values    []interface{} `mapstructure:"values" json:"values,omitempty"`
}

func configureOauth() error {
	_ = "STUB: not implemented"
	// OAuth defaults and client configuration
	return nil
}

func oauthBasicTest() error { _ = "STUB: not implemented"; return nil }

// OAuthconfig Checks

// everyone has a clientID

// everyone except IndieAuth has a clientSecret
// ADFS and OIDC providers also do not require this, but can have it optionally set.

// everyone except IndieAuth and Google has an authURL

// everyone except IndieAuth, Google and ADFS has an userInfoURL, and Azure does not actively use it

func checkScopes(scopes []string) { _ = "STUB: not implemented"; return }

// TODO: all of these methods should become `provider.SetDefaults()` or `provider.SetDefaults(*GenOAuth)`
func setProviderDefaults() { _ = "STUB: not implemented"; return }

// setDefaultsGoogle also configures the OAuthClient

// OIDC, OpenStax, Nextcloud

func setDefaultsGoogle() { _ = "STUB: not implemented"; return }

// You have to select a scope from
// https://developers.google.com/identity/protocols/googlescopes#google_sign-in

func setDefaultsADFS() { _ = "STUB: not implemented"; return }

func setDefaultsAzure() { _ = "STUB: not implemented"; return }

func setDefaultsGitHub() {
	_ = "STUB: not implemented"
	// log.Info("configuring GitHub OAuth")
	return
}

// https://github.com/vouch/vouch-proxy/issues/63
// https://developer.github.com/apps/building-oauth-apps/understanding-scopes-for-oauth-apps/

func setDefaultsDiscord() {
	_ = "STUB: not implemented"
	// log.Info("configuring GitHub OAuth")
	return
}

//Required for UserInfo URL
//https://discord.com/developers/docs/resources/user#get-current-user

func configureOAuthClient() { _ = "STUB: not implemented"; return }

func checkCallbackConfig(url string) error { _ = "STUB: not implemented"; return nil }

func arrContains(arr []string, str string) bool { _ = "STUB: not implemented"; return false }
