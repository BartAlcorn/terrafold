package deploy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tidwall/pretty"
)

// BumpPackage searches the named json file,
// locates the version field and increments sember the patch number.
func BumpPatch(filename string) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return err
	}

	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	byteValue, _ := ioutil.ReadAll(f)

	var p Package
	err = json.Unmarshal(byteValue, &p)
	if err != nil {
		fmt.Println("ERROR marshaling json", err)
	}

	p.Version, err = IncrementSemVarPatch(p.Version)
	if err != nil {
		fmt.Println("ERROR Incrementing SemVer", err)
	}

	fmt.Println(p.Version)

	pjs, err := json.Marshal(p)
	if err != nil {
		fmt.Println("ERROR Marshalling Package", err)
	}

	ppjs := pretty.Pretty(pjs)

	err = ioutil.WriteFile(filename, ppjs, os.ModePerm)
	if err != nil {
		fmt.Println("ERROR writing Package", err)
	}
	return err
}
