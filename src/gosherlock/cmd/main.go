package main

import (
	"os"

	log "github.com/e1esm/GoSherlock/pkg/logger"
	"github.com/jessevdk/go-flags"
)

func main() {
	opts := &Options{}

	_, err := flags.ParseArgs(opts, os.Args)
	if err != nil {
		panic(err)
	}

	_, err = NewCfg(opts.ConfigFile)
	if err != nil {
		panic(err)
	}

	_, err = log.New(opts.LogLevel, opts.LogOutput)
	if err != nil {
		panic(err)
	}

}
