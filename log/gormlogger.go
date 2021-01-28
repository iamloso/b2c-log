package log

import (
	"fmt"
	"go.uber.org/zap"
	"strings"
	"time"
)

// GormLogger struct
type GormLogger struct{}

var layout = "2006-01-02T15:04:05.000Z0700"

// Print - Log Formatter
func (*GormLogger) Print(v ...interface{}) {
	// defer修饰，会在包裹defer的函数执行完毕之后执行defer的内容
	defer func() {
		//recover函数将会捕获到当前的panic
		if err := recover(); err != nil {
			Bg().Info(
				"sql",
				zap.Any("src", v[1]),
				zap.Any("duration", v[2]),
				zap.Any("sql", v[3]),
				zap.Any("values", v[4]),
				zap.Any("rows_returned", v[5]),
			)
			return
		}
	}()
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
