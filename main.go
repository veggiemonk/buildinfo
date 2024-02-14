package main

import (
	"debug/buildinfo"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Usage = usage
	flag.Parse()

	i, err := buildinfo.ReadFile(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	b, err := json.Marshal(i)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", string(b))
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [binary]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\nexample:\n")
	fmt.Fprintf(os.Stderr, `buildinfo $(which pkgsite) | jq -r '.Path' | xargs -I {} go install "{}@latest"`)
	fmt.Fprintf(os.Stderr, "\nmore info: https://github.com/veggiemonk/buildinfo\n")
	// flag.PrintDefaults()
	os.Exit(2)
}
