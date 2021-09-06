package efs

import (
	"embed"
	"fmt"
	"io/fs"
)

//go:embed templates
var embeddedFiles embed.FS
var fsys fs.FS
var err error

func init() {

	// get embedded file
	fsys, err = fs.Sub(embeddedFiles, "templates")
	if err != nil {
		fmt.Println("Error initializing embedded file system: ", err)
	}

}
