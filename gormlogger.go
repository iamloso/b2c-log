package b2c_log

import "go.uber.org/zap"

// GormLogger struct
type GormLogger struct{}

// Print - Log Formatter
func (*GormLogger) Print(v ...interface{}) {
	switch v[0] {
	case "sql":
		zap.L().Info(
			"sql",
			zap.Any("src", v[1]),
			zap.Any("duration", v[2]),
			zap.Any("sql", v[3]),
			zap.Any("values", v[4]),
			zap.Any("rows_returned", v[5]),
		)
	case "log":
		zap.L().Error("error", zap.Any("gorm", v[2]))
	}
}
