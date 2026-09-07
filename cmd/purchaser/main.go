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

	p, err := processor.GetProcessor(strings.Split(flags.IgnoreTemplates, ","), strings.Split(flags.Scales, ","))
	if err != nil {
		log.Panicf("error attempting to instantiate processor: %v", err)
	}

	for _, file := range files {
		err := p.Process(file)
		if err != nil {
			log.Panicf("could not open file: %v", err)
		}
	}
}
