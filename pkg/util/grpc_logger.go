/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/common/logging"
)

// ZapLogger is
type ZapLogger struct {
	logger logging.Logger
}

// NewZapLogger 创建封装了zap的对象，该对象是对LoggerV2接口的实现
func NewZapLogger(logger *logging.Logger) *ZapLogger {
	return &ZapLogger{
		logger: *logger,
	}
}

// Info returns
func (zl *ZapLogger) Info(args ...interface{}) {
	zl.logger.Info(args);
}

// Infoln returns
func (zl *ZapLogger) Infoln(args ...interface{}) {
	zl.logger.Info(args...)
}

// Infof returns
func (zl *ZapLogger) Infof(format string, args ...interface{}) {
	zl.logger.Infof(format, args...)
}

// Warning returns
func (zl *ZapLogger) Warning(args ...interface{}) {
	zl.logger.Warn(args...)
}

// Warningln returns
func (zl *ZapLogger) Warningln(args ...interface{}) {
	zl.logger.Warn(args...)
}

// Warningf returns
func (zl *ZapLogger) Warningf(format string, args ...interface{}) {
	zl.logger.Warnf(format, args...)
}

// Error returns
func (zl *ZapLogger) Error(args ...interface{}) {
	zl.logger.Error(args...)
}

// Errorln returns
func (zl *ZapLogger) Errorln(args ...interface{}) {
	zl.logger.Error(args...)
}

// Errorf returns
func (zl *ZapLogger) Errorf(format string, args ...interface{}) {
	zl.logger.Errorf(format, args...)
}

// Fatal returns
func (zl *ZapLogger) Fatal(args ...interface{}) {
	zl.logger.Fatal(args...)
}

// Fatalln returns
func (zl *ZapLogger) Fatalln(args ...interface{}) {
	zl.logger.Fatal(args...)
}

// Fatalf logs to fatal level
func (zl *ZapLogger) Fatalf(format string, args ...interface{}) {
	zl.logger.Fatalf(format, args...)
}

// V reports whether verbosity level l is at least the requested verbose level.
func (zl *ZapLogger) V(v int) bool {
	return false
}