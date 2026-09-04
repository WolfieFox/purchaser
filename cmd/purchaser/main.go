package main

import (
	"flag"
	"log"
	"strings"

	"github.com/mattkimber/purchaser/internal/processor"
)

type Flags struct {
	Scales          string
	IgnoreTemplates string
}

var flags Flags

func init() {
	flag.StringVar(&flags.Scales, "scale", "1,2", "comma-separated list of scales to render sprites at")
	flag.StringVar(&flags.Scales, "s", "1,2", "shorthand for -scale")
	flag.StringVar(&flags.IgnoreTemplates, "ignore", "na,tender", "comma-separated list of templates that don't need purchase sprites")
	flag.StringVar(&flags.IgnoreTemplates, "i", "na,tender", "comma-separated list of templates that don't need purchase sprites")
}

func main() {
	flag.Parse()

	files := flag.Args()
	for _, file := range files {
		err := processor.Process(file, strings.Split(flags.Scales, ","), strings.Split(flags.IgnoreTemplates, ","))
		if err != nil {
			log.Panicf("could not open file: %v", err)
		}
	}
}
