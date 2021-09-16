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
	for _, t := range fldr.Lambda.Triggers {
		fldr.Lambda.Trigger = t
		if _, err := os.Stat(fldr.LPath + "/iac/common/lambda-" + t + ".tf"); os.IsNotExist(err) || fldr.Lambda.Overwrite {

			f, err := os.Create(fldr.LPath + "/iac/common/lambda-" + t + ".tf")
			if err != nil {
				fmt.Println("ERROR ", err)
			}
			defer f.Close()

			var buff bytes.Buffer

			ft, err := efs.GetFile("lambda-" + t + ".tf")
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

			fmt.Println(chalk.Green.Color("writing  " + fldr.LPath + "/iac/common/lambda-" + t + ".tf"))
		} else {
			fmt.Println(chalk.Yellow.Color("skipping " + fldr.LPath + "/iac/common/lambda-" + t + ".tf"))
		}
	}
}
