package b2c_log

const (
	// logFormat
	FormatJson    = "json"
	FormatConsole = "console"

	// EncoderConfig
	TimeKey       = "Time"
	LevelKey      = "Level"
	NameKey       = "Logger"
	CallerKey     = "Caller"
	MessageKey    = "Msg"
	StacktraceKey = "Stacktrace"

	// 日志归档配置项
	// 每个日志文件保存的最大尺寸 单位：M
	MaxSize = 1
	// 文件最多保存多少天
	MaxBackups = 1
	// 日志文件最多保存多少个备份
	MaxAge = 1
)
