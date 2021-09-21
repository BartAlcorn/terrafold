package scaffold

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tidwall/pretty"
)

//NewProfile creates a new named profile json file.
func NewProfile(name string) {
	path := name
	if !IsLambdaDir() {
		path = "lambdas/" + name + ".json"
	}
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("ERROR Creating file ", err)
		return
	}
	defer f.Close()

	l := Lambda{
		Name:     name,
		Triggers: []string{"api", "event", "invoke", "sns", "sqs"},
		Stages:   []string{"dev", "qa", "uat", "prod"},
	}
	prfl := Folder{
		Lambda: l,
	}

	js, err := json.Marshal(prfl)
	if err != nil {
		fmt.Println("ERROR Marshalling ", err)
		return
	}

	_, err = f.Write(pretty.Pretty(js))
	if err != nil {
		fmt.Println("ERROR Writing file ", err)
	}
}
