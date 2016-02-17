package main

import (
	"fmt"
	"github.com/runjak/improved-dollop/config"
	"os"
)

/*
  main does the following things:
  - If no parameters given, it complains with usage message.
  - If 'init' parameter is given, it prints a 'config.json' file.
  - If <configFile> parameter is given, it reads the file and tries to act accordingly.
  Exit codes:
  0: All fine
  1: Wrong usage
  2: Problems reading/parsing file
*/
func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage:\t%v [init|<configFile>]\nSource:\thttps://github.com/runjak/improved-dollop\n", os.Args[0])
		os.Exit(1)
	}
	//Checking if init case is wanted:
	if os.Args[1] == "init" {
		c := config.EmptyConfig()
		fmt.Printf("%s\n", c.ToJson())
		os.Exit(0)
	}
	//Reading config file:
	config, err := config.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Printf("Successfully parsed config:\n\t%v\n", config)
	//Starting up:
	//FIXME IMPLEMENT .)
}
