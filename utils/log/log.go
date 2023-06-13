package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap/zapcore"
	"niubi-mall/global"
	"os"
)

// GetWriteSyncer --- 日志写与日志切割
func GetWriteSyncer(file string) zapcore.WriteSyncer {
	lumerJackLogger := &lumberjack.Logger{

		Filename:   file,
		MaxSize:    10,
		MaxBackups: 200,
		MaxAge:     30,
		Compress:   true,
	}

	if global.GVA_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumerJackLogger))
	}
	return zapcore.AddSync(lumerJackLogger)
}
