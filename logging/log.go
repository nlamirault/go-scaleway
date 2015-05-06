// Copyright (C) 2015  Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logging

import (
	//"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
)

// InitLogging define log level
func InitLogging(level log.Level) {
	//log.SetLevel(log.InfoLevel)
	log.SetOutput(os.Stderr)
	log.SetLevel(level)
}

// Debug print message using the Debug level color
func Debug(args ...interface{}) {
	log.Debug("[Scaleway] ", args)
}

// Info print message using the INFO level color
func Info(args ...interface{}) {
	log.Info("[Scaleway] ", args)
}

// Warn print message using the WARN level color
func Warn(args ...interface{}) {
	log.Warn("[Scaleway] ", args)
}

// Error print message using the ERROR level color
func Error(args ...interface{}) {
	log.Error("[Scaleway] ", args)
}
