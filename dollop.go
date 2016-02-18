package main

import (
	"flag"
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
	//Flags to use:
	init := flag.Bool("init", false, "Set this flag to print the default config.")
	path := flag.String("config", "config.json", "Use this flag to specify a config file.")
	flag.Parse()
	//Checking if init case is wanted:
	if *init {
		c := config.EmptyConfig()
		fmt.Printf("%s\n", c.ToJson())
		os.Exit(0)
	}
	//Reading config file:
	config, err := config.ReadFile(*path)
	if err != nil {
		fmt.Printf("Error reading/parsing config file '%s':\n", *path)
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Printf("Successfully parsed config:\n\t%v\n", config)
	//Starting up:
	//FIXME IMPLEMENT .)
}
