package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/* General structure of our configuration */
type Config struct {
	CertFile string
	ListenOn []string
	HostMap  map[string]string
}

/* Function to produce an empty config map */
func emptyConfig() Config {
	hostMap := make(map[string]string)
	hostMap["127.0.0.1:80"] = "127.0.0.1:8080"
	return Config{
		CertFile: "./that.cert",
		ListenOn: []string{"0.0.0.0:80", "0.0.0.0:443"},
		HostMap:  hostMap}
}

/*
  main does the following things:
  - If no parameters given, it complains with usage message.
  - If 'init' parameter is given, it prints a 'config.json' file.
  - If <configFile> parameter is given, it reads the file and tries to act accordingly.
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
	}
}
