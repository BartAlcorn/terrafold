package scaffold

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/tidwall/pretty"
)

// ReadProfile reads the named terrafold profile json file.
func ReadProfile(filename string, prfl Profile) (Folder, error) {
	fldr := Folder{}

	filename = filename + ".json"
	if !IsLambdaDir() {
		filename = "lambdas/" + filename
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return fldr, fmt.Errorf("ERROR! Profile named %v not found in lambdas folder!\n", filename)
	}

	file, _ := ioutil.ReadFile(filename)

	err := json.Unmarshal(file, &fldr)
	if err != nil {
		return fldr, fmt.Errorf("ERROR unmarshalling json: %v.\n", err)

	}

	name := strings.Replace(strings.ToLower(fldr.Lambda.Name), "-", "", -1)
	fldr.LPath = name
	if !IsLambdaDir() {
		fldr.LPath = "lambdas/" + name
	}

	if fldr.Lambda.Description == "" {
		fldr.Lambda.Description = "generic description of my new lambda: " + fldr.Lambda.Name
	}

	if len(fldr.Lambda.Stages) < 1 {
		fldr.Lambda.Stages = prfl.Stages
	}

	if len(fldr.Lambda.Triggers) < 1 {
		fldr.Lambda.Triggers = prfl.Triggers
	}

	js, err := json.Marshal(fldr)
	if err != nil {
		return fldr, err
	}

	fmt.Println("Lambda Profile:")
	fmt.Println(string(pretty.Color((pretty.Pretty(js)), nil)))

	return fldr, nil

}
