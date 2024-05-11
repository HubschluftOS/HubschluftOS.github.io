package cmd

import (
	"flag"
)

type Args struct {
	flagWeb *bool
}

func Arguments() *Args {
	args := Args{}
	args.flagWeb = flag.Bool("web", false, "start a web server")
	flag.Parse()

	if *args.flagWeb {
		Server()
	}

	return &args
}
