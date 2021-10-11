package scaffold

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tidwall/pretty"
)

// ReadBaseProfile reads the terraform.json common profile.
func ReadBaseProfile() (Profile, error) {
	prfl := Profile{}

	filename := "terrafold.json"
	if !IsLambdaDir() {
		filename = "lambdas/" + filename
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return prfl, fmt.Errorf("ERROR! Profile named %v not found in lambdas folder!\n", filename)
	}

	file, _ := ioutil.ReadFile(filename)

	err := json.Unmarshal(file, &prfl)
	if err != nil {
		return prfl, fmt.Errorf("ERROR unmarshalling json: %v.\n", err)

	}

	js, err := json.Marshal(prfl)
	if err != nil {
		return prfl, err
	}

	fmt.Println("Base Profile:")
	fmt.Println(string(pretty.Color((pretty.Pretty(js)), nil)))

	return prfl, nil

}
