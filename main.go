// Copyright 2019 Yoshi Yamaguchi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"math/rand"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const duration = 10 * time.Millisecond

func main() {
	logger, err := initLogger()
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}

	ticker := time.NewTicker(duration)
	for {
		select {
		case <-ticker.C:
			sendLog(logger)
		}
	}
}

func sendLog(l *zap.SugaredLogger) {
	i := rand.Int31()
	switch {
	case i%3571 == 0:
		l.Errorf("the value was devided by 3571: %d", i)
	case i%1123 == 0:
		l.Warnf("the value was devided by 1123: %d", i)
	case i%9 == 0 || i%11 == 0 || i%13 == 0:
		l.Infof("the value was devided by 9, 11 or 13: %d", i)
	default:
		l.Debugf("the value is %d", i)
	}
}

func initLogger() (*zap.SugaredLogger, error) {
	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "message",
			LevelKey:       "severity",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			TimeKey:        "timestamp",
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			CallerKey:      "caller",
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	var err error
	var l *zap.Logger
	l, err = cfg.Build()
	if err != nil {
		return nil, err
	}
	return l.Sugar(), nil
}
