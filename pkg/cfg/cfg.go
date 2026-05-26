/*

Copyright 2020 The Vouch Proxy Authors.
Use of this source code is governed by The MIT License (MIT) that
can be found in the LICENSE file. Software distributed under The
MIT License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied.

*/

package cfg

import (
	"embed"
	"errors"
	"flag"
	"io/fs"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Config vouch jwt cookie configuration
// Note to developers!  Any new config elements
// should use `snake_case` such as `post_logout_redirect_uris`
// in certain situations you'll need to add both a `mapstructure` tag used by viper
// as well as a `envconfig` tag used by https://github.com/kelseyhightower/envconfig
// though most of the time envconfig will use the struct key's name: VOUCH_PORT VOUCH_JWT_MAXAGE
// default values should be set in .defaults.yml
type Config struct {
	LogLevel      string   `mapstructure:"logLevel"`
	Listen        string   `mapstructure:"listen"`
	Port          int      `mapstructure:"port"`
	SocketMode    int      `mapstructure:"socket_mode"`
	SocketGroup   string   `mapstructure:"socket_group"`
	DocumentRoot  string   `mapstructure:"document_root" envconfig:"document_root"`
	WriteTimeout  int      `mapstructure:"writeTimeout"`
	ReadTimeout   int      `mapstructure:"readTimeout"`
	IdleTimeout   int      `mapstructure:"idleTimeout"`
	Domains       []string `mapstructure:"domains"`
	WhiteList     []string `mapstructure:"whitelist"`
	TeamWhiteList []string `mapstructure:"teamWhitelist"`
	AllowAllUsers bool     `mapstructure:"allowAllUsers"`
	PublicAccess  bool     `mapstructure:"publicAccess"`
	TLS           struct {
		Cert    string `mapstructure:"cert"`
		Key     string `mapstructure:"key"`
		Profile string `mapstructure:"profile"`
	}
	JWT struct {
		SigningMethod  string `mapstructure:"signing_method"`
		MaxAge         int    `mapstructure:"maxAge"` // in minutes
		Issuer         string `mapstructure:"issuer"`
		Secret         string `mapstructure:"secret"`
		PrivateKeyFile string `mapstructure:"private_key_file"`
		PublicKeyFile  string `mapstructure:"public_key_file"`
		Compress       bool   `mapstructure:"compress"`
	}
	Cookie struct {
		Name     string `mapstructure:"name"`
		Domain   string `mapstructure:"domain"`
		Secure   bool   `mapstructure:"secure"`
		HTTPOnly bool   `mapstructure:"httpOnly"`
		MaxAge   int    `mapstructure:"maxage"`
		SameSite string `mapstructure:"sameSite"`
	}

	Headers struct {
		JWT           string            `mapstructure:"jwt"`
		User          string            `mapstructure:"user"`
		QueryString   string            `mapstructure:"querystring"`
		Redirect      string            `mapstructure:"redirect"`
		Success       string            `mapstructure:"success"`
		Error         string            `mapstructure:"error"`
		ClaimHeader   string            `mapstructure:"claimheader"`
		Claims        []string          `mapstructure:"claims"`
		AccessToken   string            `mapstructure:"accesstoken"`
		IDToken       string            `mapstructure:"idtoken"`
		ClaimsCleaned map[string]string // the rawClaim is mapped to the actual claims header
	}
	Session struct {
		Name   string `mapstructure:"name"`
		MaxAge int    `mapstructure:"maxage"`
		Key    string `mapstructure:"key"`
	}
	TestURL            string   `mapstructure:"test_url"`
	TestURLs           []string `mapstructure:"test_urls"`
	Testing            bool     `mapstructure:"testing"`
	LogoutRedirectURLs []string `mapstructure:"post_logout_redirect_uris" envconfig:"post_logout_redirect_uris"`
}

type branding struct {
	LCName   string // lower case vouch
	UCName   string // UPPER CASE VOUCH
	CcName   string // camelCase Vouch
	FullName string // Vouch Proxy
	URL      string // https://github.com/vouch/vouch-proxy
}

var (
	// Branding that's our name
	Branding = branding{"vouch", "VOUCH", "Vouch", "Vouch Proxy", "https://github.com/vouch/vouch-proxy"}

	// RootDir is where Vouch Proxy looks for ./config/config.yml and ./data
	RootDir string

	secretFile string

	// CmdLine command line arguments
	CmdLine = &cmdLineFlags{
		IsHealthCheck: flag.Bool("healthcheck", false, "invoke healthcheck (check process return value)"),
		port:          flag.Int("port", -1, "port"),
		configFile:    flag.String("config", "", "specify alternate config.yml file as command line arg"),
		// https://github.com/uber-go/zap/blob/master/flag.go
		logLevel: zap.LevelFlag("loglevel", cmdLineLoggingDefault, "set log level to one of: panic, error, warn, info, debug"),
		logTest:  flag.Bool("logtest", false, "print a series of log messages and exit (used for testing)"),
	}

	// Cfg the main exported config variable
	Cfg = &Config{}
	// IsHealthCheck see main.go
	IsHealthCheck = false

	errConfigNotFound = errors.New("configuration file not found")
	// TODO: audit errors and use errConfigIsBad
	// errConfigIsBad    = errors.New("configuration file is malformed")

	// Templates are loaded from the file system with a go:embed directive in main.go
	Templates fs.FS

	// Defaults are loaded from the file system with a go:embed directive in main.go
	Defaults embed.FS
)

type cmdLineFlags struct {
	IsHealthCheck *bool
	port          *int
	configFile    *string
	logLevel      *zapcore.Level
	logTest       *bool
}

const (
	// for a Base64 string we need 44 characters to get 32bytes (6 bits per char)
	minBase64Length = 44
	base64Bytes     = 32

	// ErrCtxKey set or check the http request context to see if it has errored
	// see `responses.Error401` and `jwtmanager.JWTCacheHandler` for example
	ErrCtxKey ctxKey = 0
)

// use a typed ctxKey to avoid context key collision
// https://blog.golang.org/context#TOC_3.2.
type ctxKey int

// Configure called at the very top of main()
// the order of config follows the Viper conventions...
//
// The priority of the sources is the following:
// 1. command line flags
// 2. env. variables
// 3. config file
// 4. defaults
//
// so we process these in backwards order (defaults then config file)
func Configure() { _ = "STUB: not implemented"; return }

// bail if we're testing

// then it's probably config file not found

// using envconfig
// https://github.com/kelseyhightower/envconfig
func configureFromEnv() bool { _ = "STUB: not implemented"; return false }

// did anything change?

// set logLevel before calling Log.Debugf()

// log.Debugf("preEnvConfig %+v", preEnvConfig)

// ValidateConfiguration confirm the Configuration is valid
func ValidateConfiguration() error {
	_ = "STUB: not implemented"

	// Logging.setLogLevel(zap.DebugLevel)
	return nil
}

func setRootDir() {
	_ = "STUB: not implemented"
	// set RootDir from VOUCH_ROOT env var, or to the executable's directory
	return
}

// parseConfig parse the config file
func parseConfigFile() error { _ = "STUB: not implemented"; return nil }

// Find and read the config file
// Handle errors reading the config file

// don't log the secret!
// log.Debugf("secret: %s", string(Cfg.JWT.Secret))

// consolidate config related Log.Debugf() calls so that they can be placed *after* we set the logLevel
func logConfigIfDebug() { _ = "STUB: not implemented"; return }

// log.Debugf("viper settings %+v", viper.AllSettings())

// Mask sensitive configuration items before logging

func fixConfigOptions() { _ = "STUB: not implemented"; return }

// headers defaults

// jwt defaults

// use viper and mapstructure check to see if
// https://pkg.go.dev/github.com/spf13/viper@v1.20.1?tab=doc#Unmarshal
// https://github.com/go-viper/mapstructure
func checkConfigFileWellFormed() error { _ = "STUB: not implemented"; return nil }

// UnmarshalKey populate struct from contents of cfg tree at key
func UnmarshalKey(key string, rawVal interface{}) error { _ = "STUB: not implemented"; return nil }

// Get string value for key
func Get(key string) string { _ = "STUB: not implemented"; return "" }

// basicTest just a quick sanity check to see if the config is sound
func basicTest() error {
	_ = "STUB: not implemented"
	// check oauth config
	return nil
}

// Domains is required _unless_ Cfg.AllowAllUsers is set

// issue a warning if the secret is too small

// HMAC
// RSA
// ECDSA

// check tls config

// setDefaults set default options for most items from `.defaults.yml` in the root dir
func setDefaults() {
	_ = "STUB: not implemented"

	// viper.SetConfigName(".defaults")
	return
}

// viper.AddConfigPath(RootDir)
// viper.ReadInConfig()

// keep this here for development, we're still pre configurating of LogLevel
// log.Debugf("setDefaults from .defaults.yml %+v", Cfg)

// bare minimum for healthcheck achieved

func claimToHeader(claim string) (string, error) {
	_ = "STUB: not implemented"

	// Auth0 allows "namespaceing" of claims and represents them as URLs
	return "", nil
}

// not allowed in header: "(),/:;<=>?@[\]{}"
// https://greenbytes.de/tech/webdav/rfc7230.html#rfc.section.3.2.6
// and we don't allow underscores `_` or periods `.` because nginx doesn't like them
// "Valid names are composed of English letters, digits, hyphens, and possibly underscores"
// as per http://nginx.org/en/docs/http/ngx_http_core_module.html#underscores_in_headers

// The field-name must be composed of printable ASCII characters (i.e., characters)
// that have values between 33. and 126., decimal, except colon).
// https://github.com/vouch/vouch-proxy/issues/183#issuecomment-564427548
// get the rune (char) for each claim character

// log.Debugf("claimToHeader rune %c - %d", r, r)

// log.Errorf("%s.header.claims %s will be forwarded in the Header %s", Branding.CcName, was, claim)

// fix the claims headers
// https://github.com/vouch/vouch-proxy/issues/183

func cleanClaimsHeaders() error { _ = "STUB: not implemented"; return nil }

// InitForTestPurposes is called by most *_testing.go files in Vouch Proxy
func InitForTestPurposes() { _ = "STUB: not implemented"; return }

// InitForTestPurposesWithProvider just for testing
func InitForTestPurposesWithProvider(provider string) {
	_ = "STUB: not implemented"
	// clear it out since we're called multiple times from subsequent tests
	return
}

// _, b, _, _ := runtime.Caller(0)
// basepath := filepath.Dir(b)

// Configure()
// setRootDir()

// can't use setDefaults for testing which is go:embed based so we do it the old way
// setDefaults()

// this also mimics the go:embed testing setup

// setDevelopmentLogger()

// Needed to override the provider, which is otherwise set via yml

func DecryptionKey() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

// signingMethod should already have been validated, this should not happen

func SigningKey() (interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

// We should have validated this before

// Check that we have read permission for this file
// https://stackoverflow.com/questions/60128401/how-to-check-if-a-file-is-executable-in-go
func canRead(file string) bool { _ = "STUB: not implemented"; return false }

// detect if we're in a docker environment
func isDocker() bool { _ = "STUB: not implemented"; return false }

func logSysInfo() { _ = "STUB: not implemented"; return }
