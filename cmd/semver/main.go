package main

import (
	"fmt"

	"github.com/bartalcorn/terrafold/pkg/deploy"
)

func main() {

	err := deploy.BumpPackage("./cmd/semver/package.json")
	if err != nil {
		fmt.Println("ERROR reading file", err)
	}

}
