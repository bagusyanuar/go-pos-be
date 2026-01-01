package config

import (
	"os"
	"sync"

	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	globalLogger *zap.Logger
	onceLogger   sync.Once
)

func NewLogger(viper *viper.Viper) *zap.Logger {
	onceLogger.Do(func() {
		logFile := &lumberjack.Logger{
			Filename:   viper.GetString("LOG_FILE"),
			MaxSize:    viper.GetInt("LOG_MAX_SIZE"),
			MaxAge:     viper.GetInt("LOG_MAX_AGE"),
			MaxBackups: viper.GetInt("LOG_MAX_BACKUPS"),
			Compress:   viper.GetBool("LOG_COMPRESS"),
		}

		// encoderConfig := zapcore.EncoderConfig{
		// 	TimeKey:        "time",
		// 	LevelKey:       "level",
		// 	NameKey:        "logger",
		// 	CallerKey:      "caller",
		// 	MessageKey:     "msg",
		// 	StacktraceKey:  "stacktrace",
		// 	LineEnding:     zapcore.DefaultLineEnding,
		// 	EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		// 	EncodeTime:     zapcore.ISO8601TimeEncoder,
		// 	EncodeDuration: zapcore.SecondsDurationEncoder,
		// 	EncodeCaller:   zapcore.ShortCallerEncoder,
		// }

		levelStr := viper.GetString("LOG_LEVEL")
		var zapLevel zapcore.Level
		if err := zapLevel.UnmarshalText([]byte(levelStr)); err != nil {
			zapLevel = zapcore.InfoLevel // fallback
		}

		consoleConfig := zap.NewDevelopmentEncoderConfig()
		consoleConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		consoleConfig.EncodeTime = zapcore.ISO8601TimeEncoder

		// Config khusus untuk File (JSON, bersih tanpa warna)
		fileConfig := zap.NewProductionEncoderConfig()
		fileConfig.EncodeTime = zapcore.ISO8601TimeEncoder

		core := zapcore.NewTee(
			zapcore.NewCore(zapcore.NewConsoleEncoder(consoleConfig), zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
			zapcore.NewCore(zapcore.NewJSONEncoder(fileConfig), zapcore.AddSync(logFile), zapcore.InfoLevel),
		)

		globalLogger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	})
	return globalLogger
}
