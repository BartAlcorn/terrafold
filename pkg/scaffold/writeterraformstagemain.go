package scaffold

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"github.com/bartalcorn/terrafold/pkg/efs"
	"github.com/ttacon/chalk"
)

// WriteTerraformStageMain creates the main.tf Terraform file per stage. It is stage specific.
func WriteTerraformStageMain(fldr Folder) {
	if _, err := os.Stat(fldr.LPath + "/iac/" + fldr.Lambda.Stage + "/main.tf"); os.IsNotExist(err) || fldr.Lambda.Overwrite {

		f, err := os.Create(fldr.LPath + "/iac/" + fldr.Lambda.Stage + "/main.tf")
		if err != nil {
			fmt.Println("ERROR ", err)
		}
		defer f.Close()

		var buff bytes.Buffer

		// get embedded file
		ft, err := efs.GetFile("iac-main-header.tmpl")
		if err != nil {
			fmt.Println("ERROR, Embedded FS: ", err)
		}

		tmp, err := template.New("iacBaseMain").Parse(string(ft))
		if err != nil {
			panic(err)
		}
		err = tmp.Execute(&buff, fldr.Lambda)
		if err != nil {
			panic(err)
		}

		for _, t := range fldr.Lambda.Triggers {
			fldr.Lambda.Trigger = t
			ft, err := efs.GetFile("iac-main-lambda-sns.tmpl")
			if err != nil {
				fmt.Println("ERROR, Embedded FS: ", err)
			}
			tmp, err := template.New("iacBaseMain").Parse(string(ft))
			if err != nil {
				panic(err)
			}
			err = tmp.Execute(&buff, fldr.Lambda)
			if err != nil {
				panic(err)
			}
		}

		// write static header
		_, err = f.WriteString(buff.String())
		if err != nil {
			fmt.Println("ERROR ", err)
		}
		// Commit to file
		err = f.Sync()
		if err != nil {
			fmt.Println("ERROR ", err)
		}

		fmt.Println(chalk.Green.Color("writing  " + fldr.LPath + "/iac/" + fldr.Lambda.Stage + "/main.tf"))
	} else {
		fmt.Println(chalk.Yellow.Color("skipping " + fldr.LPath + "/iac/" + fldr.Lambda.Stage + "/main.tf"))
	}

}
