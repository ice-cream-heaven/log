//go:build release

package log

import (
	"os"
	"path/filepath"
)

func init() {
	//SetOutput(NewAsyncWriter(GetOutputWriterHourly(filepath.Join(os.TempDir(), "ice", "log") + "/")))
	SetOutput((GetOutputWriterHourly(filepath.Join(os.TempDir(), "ice", "log") + "/")))
}

func Debug(args ...interface{}) {
}

func Debugf(format string, args ...interface{}) {
}
