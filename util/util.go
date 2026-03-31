package util

import (
	"errors"
	"os"
)

func Exists(path string) bool {
    _, err := os.Stat(path)
    return err == nil || !errors.Is(err, os.ErrNotExist)
}