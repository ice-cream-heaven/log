//go:build release && !discard

package log

import (
	"os"
	"path/filepath"
)

func init() {
	SetOutput((GetOutputWriterHourly(filepath.Join(os.TempDir(), "ice", "log") + "/")))
}

func Debug(args ...interface{}) {
}

func Debugf(format string, args ...interface{}) {
}
