/*

Copyright 2020 The Vouch Proxy Authors.
Use of this source code is governed by The MIT License (MIT) that
can be found in the LICENSE file. Software distributed under The
MIT License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied.

*/

package cfg

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logging struct {
	Logger          *zap.SugaredLogger
	FastLogger      *zap.Logger
	AtomicLogLevel  zap.AtomicLevel
	DefaultLogLevel zapcore.Level
}

var (
	logger *zap.Logger
	log    *zap.SugaredLogger

	// Logging is the public interface to logging
	Logging = &logging{
		AtomicLogLevel:  zap.NewAtomicLevel(),
		DefaultLogLevel: zap.InfoLevel,
	}
)

const cmdLineLoggingDefault = -2

func init() {
	Logging.AtomicLogLevel = zap.NewAtomicLevel()
	// zap needs to start at zapcore.DebugLevel so that it can then be decreased to a lesser level
	Logging.AtomicLogLevel.SetLevel(zapcore.DebugLevel)
	encoderCfg := zap.NewProductionEncoderConfig()
	logger = zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		Logging.AtomicLogLevel,
	))

	defer logger.Sync() // flushes buffer, if any
	log = logger.Sugar()
	Logging.FastLogger = logger
	Logging.Logger = log
	// 	Logging.FastLogger = zap.L()
	// 	Logging.Logger = zap.S()
	// 	log = Logging.Logger
	// log.Info("logger set")

}

func (logging) setLogLevel(lvl zapcore.Level) {
	_ = "STUB: not implemented"
	// https://github.com/uber-go/zap/blob/master/zapcore/level.go#L59
	return
}

func (logging) setLogLevelString(str string) { _ = "STUB: not implemented"; return }

func (logging) setDevelopmentLogger() {
	_ = "STUB: not implemented"
	// then configure the logger for development output
	return
}

// func(zapcore.Core) zapcore.Core {

// zap.ReplaceGlobals(clone)

// Logging.FastLogger = log.Desugar()
// Logging.Logger = log

var configured = false

func (logging) configure() {
	_ = "STUB: not implemented"
	// logging
	return
}

// then we weren't configured via command line, check the config file

// then we weren't configured via the config file, set the default

// log.Errorf("Logging.configure() Logging.LogLevel %s Cfg.LogLevel %s", Logging.LogLeveLogging.String(), Cfg.LogLevel)

// if we're supposed to run tests, run tests and exit

func (logging) configureFromCmdline() { _ = "STUB: not implemented"; return }

// defaults to Logging.DefaultLogLevel which is zap.InfoLevel

// if we're supposed to run tests, run tests and exit

// in support of `./do.sh test_logging`
func (logging) cmdlineTestLogs() { _ = "STUB: not implemented"; return }

// Logging.Logger.Panic("panic")
