package scaffold

import (
	"fmt"
	"os"
)

// MakePath creates the named path, if it does not already exist.
func MakePath(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			fmt.Println("ERROR Make Path", err)
		}
		return err
	}
	return nil
}
