package scaffold

import (
	"bytes"
	"fmt"
	"html/template"
	"os"

	"github.com/bartalcorn/terrafold/pkg/efs"
	"github.com/ttacon/chalk"
)

// WriteTerraformECR writes a single ECR Terraform file.
// This Terraform file will container the necessary Terraform commands
// to deploy a named ECR per trigger for the named application.
func WriteTerraformECR(fldr Folder) {
	fmt.Println("Terraform ECR...")

	if _, err := os.Stat(fldr.LPath + "/iac/base/main.tf"); os.IsNotExist(err) || fldr.Lambda.Overwrite {

		f, err := os.Create(fldr.LPath + "/iac/base/main.tf")
		if err != nil {
			fmt.Println("ERROR ", err)
		}
		defer f.Close()

		var buff bytes.Buffer

		// get embedded file
		ft, err := efs.GetFile("header.tf")
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
			ft, err := efs.GetFile("ecr.tf")
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

		fmt.Println(chalk.Green.Color("writing " + fldr.LPath + "/iac/base/main.tf"))
	} else {
		fmt.Println(chalk.Yellow.Color("skipping " + fldr.LPath + "/iac/base/main.tf"))
	}

}
