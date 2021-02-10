package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/apex/log"
	"github.com/tufin/oasdiff/diff"
	"github.com/tufin/oasdiff/load"
)

var manualSwagger, autoSwagger, prefix, filter string

func init() {
	flag.StringVar(&manualSwagger, "base", "", "base swagger file")
	flag.StringVar(&autoSwagger, "revision", "", "revised swagger file")
	flag.StringVar(&prefix, "prefix", "", "path prefix that exists in base swagger but not in revision swagger")
	flag.StringVar(&filter, "filter", "", "regex to filter results")
}

func main() {
	flag.Parse()

	base, err := load.Load(manualSwagger)
	if err != nil {
		return
	}

	revision, err := load.Load(autoSwagger)
	if err != nil {
		return
	}

	bytes, err := json.MarshalIndent(diff.GetDiffResponse(base, revision, prefix, filter), "", " ")
	if err != nil {
		log.Errorf("failed to marshal result with '%v'", err)
		return
	}

	fmt.Printf("%s\n", bytes)
}