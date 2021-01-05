// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//interface 会自动转换给zap.Field
// Logger is a simplified abstraction of the zap.Logger
type Logger interface {
	Info(msg string, params ...interface{})
	Error(msg string, params ...interface{})
	Debug(msg string, params ...interface{})
	Fatal(msg string, params ...interface{})
	With(fields ...zapcore.Field) Logger
}

// logger delegates all calls to the underlying zap.Logger
type logger struct {
	logger *zap.Logger
}

// Info logs an info msg with fields
func (l logger) Info(msg string, params ...interface{}) {
	l.logger.Info(msg, ToFields(params)...)
}

// Error logs an error msg with fields
func (l logger) Error(msg string, params ...interface{}) {
	l.logger.Error(msg, ToFields(params)...)
}

// Debug logs an error msg with fields
func (l logger) Debug(msg string, params ...interface{}) {
	l.logger.Debug(msg, ToFields(params)...)
}

// Fatal logs a fatal error msg with fields
func (l logger) Fatal(msg string, params ...interface{}) {
	l.logger.Fatal(msg, ToFields(params)...)
}

// With creates a child logger, and optionally adds some context fields to that logger.
func (l logger) With(fields ...zapcore.Field) Logger {
	return logger{logger: l.logger.With(fields...)}
}

//convert interface 2 field
func ToFields(params ...interface{}) (fields []zapcore.Field) {
	for _, param := range params {
		for _, par := range param.([]interface{}) {
			if e, ok := par.(zapcore.Field); ok {
				fields = append(fields, e)
			} else {
				fields = append(fields, zap.Reflect(fmt.Sprintf("%T", par), par))
			}
		}
	}
	return
}
