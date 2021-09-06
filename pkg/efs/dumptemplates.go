package efs

import (
	"fmt"
	"io/fs"
	"os"
)

// DumpTemplates writes the emebbed templates to a folder.
func DumpTemplates() {
	path := "./terrafoldTemplates/"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			fmt.Println("ERROR ", err)
		}
	}

	files, err := embeddedFiles.ReadDir("templates")
	if err != nil {
		fmt.Println("ERROR ReadingDir ", err)
	}

	for _, tmpl := range files {
		fmt.Println(tmpl.Name())

		f, err := os.Create(path + tmpl.Name())
		if err != nil {
			fmt.Println("ERROR ", err)
		}
		defer f.Close()

		ft, err := fs.ReadFile(fsys, tmpl.Name())
		if err != nil {
			fmt.Println("ERROR, Embedded FS: ", err)
		}

		_, err = f.Write(ft)
		if err != nil {
			fmt.Println("ERROR, Writing templates: ", err)
		}
	}

}
