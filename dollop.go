package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/* General structure of our configuration */
type Config struct {
	HostCertMap     map[string]string
	HostEndpointMap map[string]Endpoint
}

/* Endpoint to connect to for a given domain name suffix */
type Endpoint struct {
	Addr string
	Port string
}

/* Function to produce an empty config map */
func emptyConfig() Config {
	hostEndpointMap := make(map[string]Endpoint)
	hostEndpointMap["example.com"] = Endpoint{Addr: "127.0.0.1", Port: "8080"}
	return Config{
		HostCertMap:     map[string]string{"example.com": "./this.cert"},
		HostEndpointMap: hostEndpointMap,
	}
}

/*
  main does the following things:
  - If no parameters given, it complains with usage message.
  - If 'init' parameter is given, it prints a 'config.json' file.
  - If <configFile> parameter is given, it reads the file and tries to act accordingly.
  Exit codes:
  0: All fine
  1: Wrong usage
  2: Problems reading file
  3: Problems parsing JSON
*/
func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage:\t%v [init|<configFile>]\nSource:\thttps://github.com/runjak/improved-dollop\n", os.Args[0])
		os.Exit(1)
	}
	//Checking if init case is wanted:
	if os.Args[1] == "init" {
		json, _ := json.MarshalIndent(emptyConfig(), "", "  ")
		fmt.Printf("%s\n", json)
		os.Exit(0)
	}
	//Reading config file:
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	//Parsing config file:
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	fmt.Println("Successfully parsed config.")
	//Starting up:
	//FIXME IMPLEMENT .)
}
