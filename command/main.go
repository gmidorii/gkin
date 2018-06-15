package main

import (
	"flag"
	"log"

	"github.com/midorigreen/gkin"
)

func main() {
	gkfile := *flag.String("y", ".gkin.yml", "gkin job file")
	flag.Parse()

	gk, err := gkin.Parse(gkfile)
	if err != nil {
		log.Fatalln(err)
	}
	if err := gkin.Run(gkin.Argument{Gkin: gk}); err != nil {
		log.Fatalln(err)
	}
}
