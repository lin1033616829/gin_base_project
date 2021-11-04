package initialize

import (
	"fmt"
	apostates "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"mmfile/global"
	"mmfile/utils"
	"os"
	"path"
	"time"
)

var level zapcore.Level

func InitZap() (logger *zap.Logger) {
	fmt.Println("初始化log组件...")

	if ok, _ := utils.PathExists(global.ServerConf.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.ServerConf.Zap.Director)
		_ = os.Mkdir(global.ServerConf.Zap.Director, os.ModePerm)
	}

	switch global.ServerConf.Zap.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger = zap.New(getEncoderCore(), zap.AddStacktrace(level))
	} else {
		logger = zap.New(getEncoderCore())
	}
	if global.ServerConf.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

// getEncoderConfig 获取zapCore.EncoderConfig
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.ServerConf.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case global.ServerConf.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case global.ServerConf.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case global.ServerConf.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case global.ServerConf.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder 获取zapCore.Encoder
func getEncoder() zapcore.Encoder {
	if global.ServerConf.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapCore.Core
func getEncoderCore() (core zapcore.Core) {
	writer, err := getWriteSyncer() // 使用file-rotateLogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}
	return zapcore.NewCore(getEncoder(), writer, level)
}

// CustomTimeEncoder 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.ServerConf.Zap.Prefix + global.TimeFmt))
}


func getWriteSyncer() (zapcore.WriteSyncer, error) {
	fileWriter, err := apostates.New(
		path.Join(global.ServerConf.Zap.Director, "%Y-%m-%d.log"),
		apostates.WithLinkName(global.ServerConf.Zap.LinkName),
		apostates.WithMaxAge(7*24*time.Hour),
		apostates.WithRotationTime(24*time.Hour),
	)

	if global.ServerConf.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}

	return zapcore.AddSync(fileWriter), err
}
