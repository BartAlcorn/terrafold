package main

import (
	"github.com/bartalcorn/terrafold/cmd/cli/cmd"
	"github.com/spf13/viper"
)

var version = "0.0.0-default"

func main() {
	viper.Set("version", version)

	cmd.Execute()
}
