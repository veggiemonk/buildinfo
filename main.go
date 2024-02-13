package main

import (
	"debug/buildinfo"
	"encoding/json"
	"flag"
	"fmt"
	"log"
)

func main() {
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
