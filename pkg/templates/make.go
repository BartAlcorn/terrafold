package templates

import (
	"fmt"
	"html/template"
	"os"

	"github.com/bartalcorn/terrafold/pkg/efs"
)

// Make merges data from the named source and the template, and writes the resulting file.
func Make(fileName string, templateName string, source interface{}) {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("ERROR ", err)
	}
	defer f.Close()

	ft, err := efs.GetFile(templateName)
	if err != nil {
		fmt.Println("ERROR, Embedded FS: ", err)
	}

	tmp, err := template.New("newtemplate").Parse(string(ft))
	if err != nil {
		fmt.Println("ERROR, Parse Template ", err)
		panic(err)
	}
	err = tmp.Execute(f, source)
	if err != nil {
		fmt.Println("ERROR, Execute Template ", err)
		panic(err)
	}
}
