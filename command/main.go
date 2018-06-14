package main

import (
	"flag"
	"log"

	"github.com/midorigreen/gkin"
)

func main() {
	var arg gkin.Argument
	arg.Gkin = *flag.String("y", ".gkin.yml", "gkin job file")
	flag.Parse()

	if err := gkin.Run(arg); err != nil {
		log.Fatalln(err)
	}
}
