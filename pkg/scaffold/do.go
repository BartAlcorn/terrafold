package scaffold

import (
	"fmt"
	"os"

	"github.com/ttacon/chalk"
)

// Do is a convenince function to perform all scafflolding steps at once.
func Do(filename string) {
	prfl, err := ReadBaseProfile()
	if err != nil {
		fmt.Println(chalk.Red.Color(err.Error()))
		os.Exit(3)
	}
	fldr, err := ReadProfile(filename, prfl)
	if err != nil {
		fmt.Println(chalk.Red.Color(err.Error()))
		os.Exit(3)
	}
	fldr = CreateLambdaDir(fldr)
	CreateDirs(fldr)
	MakeHandlers(fldr)
	MakeIAC(fldr)
}
