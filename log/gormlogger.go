package log

import (
	"fmt"
	"time"
)

// GormLogger struct
type GormLogger struct{}

var layout = "2006-01-02T15:04:05.000Z0700"

// Print - Log Formatter
func (*GormLogger) Print(v ...interface{}) {

	switch v[0] {
	case "sql":
		str := time.Now().Format(layout) + "    " + "SQL" + "    " + fmt.Sprintf("%v    %v    %v    %v\n", v[1], v[2], v[3], v[5])
		_, _ = Hook().Write([]byte(str))
		fmt.Print(str)
	case "log":
		str := time.Now().Format(layout) + "    " + "GORM_ERROR" + "    " + fmt.Sprintf("%v\n", v[2])
		_, _ = Hook().Write([]byte(str))
		fmt.Print(str)
	}
}
