package main

import "time"

type Options struct {

	//Username - username to lookup for
	Username string `short:"u" long:"username" description:"user to look for"`

	// OutputFile - file where to write results of scanning
	OutputFile string `short:"o" long:"output" description:"output file for results"`

	// ConfigFile - file with service information
	ConfigFile string `long:"config" description:"config file with services in it"`

	// Timeout - timeout for each request
	Timeout time.Duration `short:"t" long:"timeout" description:"timeout for each request"`

	// PrintFound - whether to print only platform with found user or not
	PrintFound bool `long:"all" description:"print only found results"`

	// LogLevel - level of logs to write with
	LogLevel string `short:"v" long:"verbosity" description:"verbosity level"`

	// LogOutput - file where to place logs
	LogOutput string `long:"log_output" description:"source where to put logs"`
}
