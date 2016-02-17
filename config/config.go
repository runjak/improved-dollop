package config

import (
	"encoding/json"
	"io/ioutil"
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
func EmptyConfig() *Config {
	hostEndpointMap := make(map[string]Endpoint)
	hostEndpointMap["example.com"] = Endpoint{Addr: "127.0.0.1", Port: "8080"}
	return &Config{
		HostCertMap:     map[string]string{"example.com": "./this.cert"},
		HostEndpointMap: hostEndpointMap,
	}
}

/* Serialize a Config to a []byte as json */
func (c *Config) ToJson() []byte {
	data, _ := json.MarshalIndent(c, "", "  ")
	return data
}

/* Parse a Config from a []byte as json */
func FromJson(data []byte) (*Config, error) {
	var config Config
	err := json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

/* Write a Config to a File */
func (c *Config) WriteFile(path string) {
	ioutil.WriteFile(path, c.ToJson(), 0111)
}

/* Read a Config from a File */
func ReadFile(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return FromJson(data)
}
