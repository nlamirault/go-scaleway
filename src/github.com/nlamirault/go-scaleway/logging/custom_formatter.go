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
	"bytes"
	"fmt"
	"sort"
	//"strings"
	"time"

	log "github.com/Sirupsen/logrus"
)

const (
	nocolor = 0
	red     = 31
	green   = 32
	yellow  = 33
	blue    = 34
)

var (
	baseTimestamp time.Time
	isTerminal    bool
)

func init() {
	baseTimestamp = time.Now()
	isTerminal = log.IsTerminal()
}

func miniTS() int {
	return int(time.Since(baseTimestamp) / time.Second)
}

// CustomFormatter is a Logrus TextFormatter whith different colors for each
// log level
type CustomFormatter struct {
	ForceColors   bool
	DisableColors bool
}

// Format print the entry
func (f *CustomFormatter) Format(entry *log.Entry) ([]byte, error) {

	var keys []string
	for k := range entry.Data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	b := &bytes.Buffer{}

	prefixFieldClashes(entry)

	isColored := (f.ForceColors || isTerminal) && !f.DisableColors

	if isColored {
		printColored(b, entry, keys)
	} else {
		f.appendKeyValue(b, "time", entry.Time.Format(time.RFC3339))
		f.appendKeyValue(b, "level", entry.Level.String())
		f.appendKeyValue(b, "msg", entry.Message)
		for _, key := range keys {
			f.appendKeyValue(b, key, entry.Data[key])
		}
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}
func prefixFieldClashes(entry *log.Entry) {
	_, ok := entry.Data["time"]
	if ok {
		entry.Data["fields.time"] = entry.Data["time"]
	}

	_, ok = entry.Data["msg"]
	if ok {
		entry.Data["fields.msg"] = entry.Data["msg"]
	}

	_, ok = entry.Data["level"]
	if ok {
		entry.Data["fields.level"] = entry.Data["level"]
	}
}

func printColored(b *bytes.Buffer, entry *log.Entry, keys []string) {
	var levelColor int
	switch entry.Level {
	case log.DebugLevel:
		levelColor = nocolor
	case log.InfoLevel:
		levelColor = green
	case log.WarnLevel:
		levelColor = yellow
	case log.ErrorLevel, log.FatalLevel, log.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	//levelText := strings.ToUpper(entry.Level.String())[0:4]

	fmt.Fprintf(b, "\x1b[%dm %-44s \x1b[0m",
		levelColor, entry.Message)
	for _, k := range keys {
		v := entry.Data[k]
		fmt.Fprintf(b, "\x1b[%dm%s\x1b[0m=%v", levelColor, k, v)
	}
}
func (f *CustomFormatter) appendKeyValue(b *bytes.Buffer, key, value interface{}) {
	switch value.(type) {
	case string, error:
		fmt.Fprintf(b, "%v=%q ", key, value)
	default:
		fmt.Fprintf(b, "%v=%v ", key, value)
	}
}
