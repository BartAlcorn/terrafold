package scaffold

import (
	"fmt"
	"os"

	"github.com/bartalcorn/terrafold/pkg/templates"
	"github.com/ttacon/chalk"
)

// Make is a generic template runner.
func Make(path string, filename string, templateName string, source interface{}, overwrite bool) {
	err := MakePath(path)
	if err != nil {
		fmt.Println("ERROR ", err)
	}

	if _, err := os.Stat(path + "/" + filename); os.IsNotExist(err) || overwrite {
		templates.Make(path+"/"+filename, templateName, source)
		fmt.Println(chalk.Green.Color("writing  " + path + "/" + filename))
	} else {
		fmt.Println(chalk.Yellow.Color("skipping " + path + "/" + filename))
	}
}
