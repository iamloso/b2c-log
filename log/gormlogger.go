package log

import (
	"fmt"
	"strings"
	"time"
)

// GormLogger struct
type GormLogger struct{}

var layout = "2006-01-02T15:04:05.000Z0700"

// Print - Log Formatter
func (*GormLogger) Print(v ...interface{}) {

	switch v[0] {
	case "sql":
		content := fmt.Sprintf("%v", v[3])
		params := strings.Split(ArrayToString(fmt.Sprintf("%v", v[4])), ",")
		for i := 0; i < len(params); i++ {
			content = strings.Replace(content, "?", params[i], 1)
		}
		str := time.Now().Format(layout) + "    " + "SQL" + "    " + fmt.Sprintf("%v    %v    %v   %v\n", v[1], v[2], content, v[5])
		_, _ = Hook().Write([]byte(str))
		fmt.Print(str)
	case "log":
		str := time.Now().Format(layout) + "    " + "GORM_ERROR" + "    " + fmt.Sprintf("%v\n", v[2])
		_, _ = Hook().Write([]byte(str))
		fmt.Print(str)
	}
}
func ArrayToString(array interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}
