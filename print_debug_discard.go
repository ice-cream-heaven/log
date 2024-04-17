//go:build debug && discard

package log

import (
	"io"
)

func init() {
	SetOutput(io.Discard)
}
