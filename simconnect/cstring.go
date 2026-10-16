package simconnect

import (
	"strings"

	"golang.org/x/sys/windows"
)

// CString converts a Go string to a null-terminated ANSI string (SimConnect uses LPCSTR).
func CString(s string) (*byte, error) {
	b, err := windows.BytePtrFromString(s)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// TrimCString converts a null-terminated C char buffer to a Go string.
func TrimCString(b []byte) string {
	if i := strings.IndexByte(string(b), 0); i >= 0 {
		return string(b[:i])
	}
	return string(b)
}
