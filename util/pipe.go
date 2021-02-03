package util

import (
	"os"
)

// IsPipe check if a os.File is a pipe
// https://stackoverflow.com/a/26567513/45786
func IsPipe(f *os.File) bool {
	fi, err := f.Stat()
	if err != nil {
		return false
	}
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		return true
	}
	return false
}
