package main

import (
	"flag"
	"os"

	"github.com/ecator/gomeeting/fun"
	"github.com/ecator/gomeeting/server"
)

var (
	flagAddr     string
	flagPort     uint
	flagHelp     bool
	flagConfig   string
	flagFrontend string
	flagLog      string
)

func init() {
	flag.StringVar(&flagAddr, "a", "localhost", "The listen address")
	flag.UintVar(&flagPort, "p", 7728, "The listen port")
	flag.BoolVar(&flagHelp, "h", false, "Show usage")
	flag.StringVar(&flagConfig, "c", "config.yml", "The config file")
	flag.StringVar(&flagFrontend, "f", "assets", "The assets folder includes html/css/js")
	flag.StringVar(&flagLog, "l", "server.log", "The log file")
}

func main() {
	flag.Parse()
	if flagHelp {
		flag.Usage()
		os.Exit(0)
	}

	flagFrontend = fun.ToAbs(flagFrontend)
	// frontend folder must exsit
	if _, err := os.Stat(flagFrontend); err != nil {
		panic(err)
	}
	flagLog = fun.ToAbs(flagLog)
	server.StartServer(flagAddr, flagPort, flagFrontend, flagConfig, flagLog)
}
