package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// argsProg := os.Args[1.]
	// fmt.Println(argsProg)

	bumppkg := flag.String("bumppackage", "/package.json", "path to package.json")
	flag.Parse()

	if *bumppkg == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Println("bumppkg:", *bumppkg)

}
