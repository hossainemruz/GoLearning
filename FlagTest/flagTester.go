package main

import (
	"flag"
	"fmt"
)

func main()  {
	var version *bool
	version = flag.Bool("version",false,"1.0.1")
	flag.Parse()

	if *version{
		fmt.Printf("Version %s\n",flag.Lookup("version").Usage)
	}
}