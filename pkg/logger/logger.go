package logger

import (
	"os"
	"time"

	"github.com/isjyi/os/global"
	"github.com/isjyi/os/tools/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

/**
 * 获取日志
 */
func Setup() {

	if config.OSConfig.Logger.EnabledBUS || config.OSConfig.Application.Mode == "dev" {
		global.Logger = zap.New(newCore("/bus/"), zap.AddCaller())
	}

	if config.OSConfig.Logger.EnabledREQ || config.OSConfig.Application.Mode == "dev" {
		global.ReqLogger = zap.New(newCore("/access/"), zap.AddCaller())
	}

}

/**
* zapcore构造
* filePath 日志文件路径
* level 日志级别
* maxSize 每个日志文件保存的最大尺寸 单位：M
* maxBackups 日志文件最多保存多少个备份
* maxAge 文件最多保存多少天
* compress 是否压缩
 */
func newCore(addr string) zapcore.Core {
	//公用编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "ts",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "file",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder, // 将级别转换成大写
		EncodeTime:    zapcore.ISO8601TimeEncoder,  // ISO8601 UTC 时间格式
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
		EncodeCaller: zapcore.FullCallerEncoder, // 全路径编码器
		EncodeName:   zapcore.FullNameEncoder,
	}
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	errLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zap.ErrorLevel
	})

	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zap.ErrorLevel && lvl >= zap.DebugLevel
	})

	return zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(getWriter(addr+"info.log")), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(getWriter(addr+"error.log")), errLevel),
		//日志都会在console中展示
		zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), zapcore.Level(config.OSConfig.Logger.Level)),
	)
}

func getWriter(addr string) (hook *lumberjack.Logger) {
	hook = &lumberjack.Logger{
		Filename:   config.OSConfig.Logger.Path + addr, // 日志文件路径
		MaxSize:    config.OSConfig.Logger.MaxSize,     // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: config.OSConfig.Logger.MaxBackups,  // 日志文件最多保存多少个备份
		MaxAge:     config.OSConfig.Logger.MaxAge,      // 文件最多保存多少天
		LocalTime:  true,                               //是否使用本地时间 默认UTC
		Compress:   config.OSConfig.Logger.Compress,    // 是否压缩
	}
	return
}
