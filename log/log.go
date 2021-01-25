package log

import (
	"fmt"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

var (
	_globalL Factory
	_hook    io.Writer
)

func Hook() io.Writer {
	return _hook
}
func SetLogs(logLevel string, logDir string, logSoftLink string) {

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       TimeKey,
		LevelKey:      LevelKey,
		NameKey:       NameKey,
		CallerKey:     CallerKey,
		MessageKey:    MessageKey,
		StacktraceKey: StacktraceKey,
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalColorLevelEncoder, // 大写编码器
		EncodeTime:    zapcore.ISO8601TimeEncoder,       // ISO8601 UTC 时间格式
		// ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,     // 短路径编码器(相对路径+行号)
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志输出格式
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	// 设置日志级别,debug可以打印出info,debug,warn；info级别可以打印warn，info；warn只能打印warn
	// debug->info->warn->error
	var zapLevel zapcore.Level

	switch logLevel {
	case "debug":
		zapLevel = zap.DebugLevel
	case "info":
		zapLevel = zap.InfoLevel
	case "error":
		zapLevel = zap.ErrorLevel
	default:
		zapLevel = zap.InfoLevel
	}
	_hook = getWriter(logDir, logSoftLink)
	// 最后创建具体的Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(_hook)), zap.NewAtomicLevelAt(zapLevel)), // 日志级别

	)

	opts := []zap.Option{
		// 开启文件及行号
		zap.AddCaller(),
		// 开启开发模式，堆栈跟踪
		zap.Development(),
		zap.AddStacktrace(zapcore.ErrorLevel),
		// 避免zap始终将包装器(wrapper)代码报告为调用方。
		zap.AddCallerSkip(1),
	}
	// 构造日志
	logger := zap.New(core, opts...)
	zap.ReplaceGlobals(logger)
	_globalL = newFactory(zap.L())
}

func getWriter(logDir string, logSoftLink string) io.Writer {

	if ok, _ := pathExists(logDir); !ok {
		// directory not exist
		fmt.Println("create log directory")
		_ = os.Mkdir(logDir, os.ModePerm)
	}
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// filename.log是指向最新日志的链接%Y-%m-%d-%H-%M
	// 每24小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		logDir+string(os.PathSeparator)+"%Y-%m-%d-%H-%M.log",
		// generate soft link, point to latest log file
		rotatelogs.WithLinkName(logSoftLink),
		// time period of log file switching
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		panic(err)
	}
	return hook
}

// @title    PathExists
// @description   文件目录是否存在
// @auth                     （2020/04/05  20:22）
// @param     path            string
// @return    err             error

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
