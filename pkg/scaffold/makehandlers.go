package scaffold

import (
	"fmt"
)

// MakeHandlers creates each of the named handlers from templates
func MakeHandlers(fldr Folder) {
	fmt.Println("Handlers...")

	for _, t := range fldr.Lambda.Triggers {
		hpath := fldr.LPath + "/handler/" + t
		fmt.Printf("Handler for %v...\n", t)
		Make(hpath, "Dockerfile", "dockerfiletmpl.tmpl", fldr.Lambda, fldr.Lambda.Overwrite)
		Make(hpath, "Makefile", "makefile.tmpl", fldr.Lambda, fldr.Lambda.Overwrite)
		Make(hpath, "main.go", "main-"+t+".tmpl", fldr.Lambda, fldr.Lambda.Overwrite)
		Make(hpath, "package.json", "package.tmpl", fldr.Lambda, fldr.Lambda.Overwrite)

	}
	fmt.Println()

}
