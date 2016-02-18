package certs

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/runjak/improved-dollop/config"
	"math/big"
	"os"
	"time"
)

/* Function to test if a file exists */
func fExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

/*
	This function shall check a config by performing the following actions:
	* Iterate the HostCertMap, and for each CertPaths in it:
		* If one of the cert files doesn't exist:
		  * If Config.CreateMissingCerts:
			  * Try to create the missing cert files or report error
		  * Else report error that cert files are missing
*/
func CheckConfig(config *config.Config) error {
	for host, cPath := range config.HostCertMap {
		hasCert, hasKey := fExists(cPath.Certfile), fExists(cPath.Keyfile)
		if !hasCert || !hasKey {
			if config.CreateMissingCerts {
				err := CreateCert(host, &cPath)
				if err != nil {
					return err
				}
			} else {
				return errors.New(fmt.Sprintf("Could not find cert files: {%s,%s}\n", cPath.Certfile, cPath.Keyfile))
			}
		}
	}
	return nil
}

/* Function to create (missing) cert files */
func CreateCert(host string, cPath *config.CertPaths) error {
	//Data to work with:
	const validFor = 365 * 24 * time.Hour
	const isCA = false
	const rsaBits = 4096
	notBefore := time.Now()
	notAfter := notBefore.Add(validFor)
	//Generating a key:
	priv, err := rsa.GenerateKey(rand.Reader, rsaBits)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to generate private key: %s\n", err))
	}

	template := x509.Certificate{
		SerialNumber: new(big.Int).SetInt64(0),
		Subject: pkix.Name{
			Country:    []string{"DE"},
			Province:   []string{"Saxony"},
			CommonName: string(host),
		},
		NotBefore: notBefore,
		NotAfter:  notAfter,

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}
	template.DNSNames = append(template.DNSNames, host)
	template.IsCA = true
	template.KeyUsage |= x509.KeyUsageCertSign

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to create certificate: %s", err))
	}

	certOut, err := os.Create(cPath.Certfile)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to open %s for writing: %s", cPath.Certfile, err))
	}
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()
	fmt.Printf("Wrote %s\n", cPath.Certfile)

	keyOut, err := os.OpenFile(cPath.Keyfile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to open %s for writing: %s", cPath.Keyfile, err))
	}
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
	keyOut.Close()
	fmt.Printf("Wrote %s\n", cPath.Keyfile)
	//Finish:
	return nil
}
