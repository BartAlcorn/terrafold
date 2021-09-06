package efs

import (
	"fmt"
	"io/fs"
)

// GetFile retrieves the named frile from the embedded file system.
func GetFile(name string) ([]byte, error) {
	ft, err := fs.ReadFile(fsys, name)
	if err != nil {
		err = fmt.Errorf("ERROR, Embedded FS: %v: %v", name, err)
	}
	return ft, err
}
