package log

import (
	"io"
	"os"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

func SetOutput(writes ...io.Writer) *Logger {
	return std.SetOutput(writes...)
}

func AddOutput(writes ...io.Writer) *Logger {
	return std.AddOutput(writes...)
}

func GetOutputWriter(filename string) io.Writer {
	if filepath.Dir(filename) != filename && !isDir(filepath.Dir(filename)) {
		err := os.MkdirAll(filepath.Dir(filename), os.ModePerm)
		if err != nil {
			Errorf("err:%v", err)
		}
	}

	hook, err := rotatelogs.New(filename)
	if err != nil {
		std.Panicf("err:%v", err)
	}
	return hook
}

func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func GetOutputWriterHourly(filename string) Writer {
	if filepath.Dir(filename) != filename && !isDir(filepath.Dir(filename)) {
		err := os.MkdirAll(filepath.Dir(filename), os.ModePerm)
		if err != nil {
			Errorf("err:%v", err)
		}
	}

	hook, err := rotatelogs.
		New(filename+"%Y%m%d%H.log",
			rotatelogs.WithLinkName(filename+".log"),
			rotatelogs.WithRotationTime(time.Hour),
			rotatelogs.WithRotationCount(24),
		)
	if err != nil {
		std.Panicf("err:%v", err)
	}

	return hook
}
