//go:build windows
// +build windows

package utils

import (
	"log"

	"golang.org/x/sys/windows"
)

func SetHidden(path string) string {
	filenameW, err := windows.UTF16PtrFromString(path)
	if err != nil {
		log.Fatal(err)
	}

	err = windows.SetFileAttributes(filenameW, windows.FILE_ATTRIBUTE_HIDDEN)
	if err != nil {
		log.Fatal(err)
	}

	return path
}
