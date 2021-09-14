package utils

import (
	apostates "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
	"mmrights/global"
	"os"
	"path"
	"time"
)

// PathExists 文件目录是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func GetWriteSyncer() (zapcore.WriteSyncer, error) {
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
