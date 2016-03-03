package io

import (
	"io"
	"os"
)

// OpenForAppendOrNew Open a file and seek to end otherwise create new file
func OpenForAppendOrNew(fileName string) (io.WriteCloser, bool, error) {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		if os.IsNotExist(err) {
			f, err := os.Create(fileName)
			return f, true, err
		}
		return nil, false, err
	}
	return f, false, err
}
