//go:build unix
// +build unix

package utils

import (
	"fmt"
	"path/filepath"
)

func SetHidden(path string) string {
	filename := filepath.Base(path)	
	if filename[0] == '.' {
		return path
	}
	
	newFileName := fmt.Sprintf(".%v", filename)
	newPath := filepath.Join(filepath.Dir(path), newFileName)
	
	return newPath
}