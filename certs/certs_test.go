package certs

import (
	"fmt"
	"github.com/runjak/improved-dollop/config"
	"github.com/runjak/improved-dollop/util"
	"os"
	"testing"
)

func TestCheckConfig(t *testing.T) {
	eConfig := config.EmptyConfig()
	eConfig.CreateMissingCerts = true
	eConfig.HostCertMap = make(map[string]config.CertPaths)
	//Testing empty HostCertMap:
	if CheckConfig(eConfig) != nil {
		t.Errorf("Empty HostCertMap should always pass CheckConfig.")
	}
	//Adding some stuff to HostCertMap:
	rand := util.RandString(10)
	cps := config.CertPaths{
		Certfile: "/tmp/" + rand + ".pem",
		Keyfile:  "/tmp/" + rand + ".key.pem"}
	eConfig.HostCertMap["test"] = cps
	//Testing creation of certs:
	err := CheckConfig(eConfig)
	if err != nil {
		t.Errorf("CheckConfig complains when it should have created some certs:\n\t%s\n", err)
	}
	if !fExists(cps.Certfile) {
		t.Errorf("CheckConfig didn't create expected file: %s\n", cps.Certfile)
	}
	if !fExists(cps.Keyfile) {
		t.Errorf("CheckConfig didn't create expected file: %s\n", cps.Keyfile)
	}
	fmt.Printf("Removing files:\n\t%s\n\t%s\n", cps.Certfile, cps.Keyfile)
	os.Remove(cps.Certfile)
	os.Remove(cps.Keyfile)
	//Testing missing HostCertFiles:
	eConfig.CreateMissingCerts = false
	if CheckConfig(eConfig) == nil {
		t.Errorf("CheckConfig didn't complain when it could not create missing certs.")
	}
	if fExists(cps.Certfile) {
		t.Errorf("CheckConfig created unexpected file: %s\n", cps.Certfile)
	}
	if fExists(cps.Keyfile) {
		t.Errorf("CheckConfig created unexpected file: %s\n", cps.Keyfile)
	}
}
