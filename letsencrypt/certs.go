package main

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
)

func verifyCertificate(certPath string, caPath string, domain string) bool {
	if ca, err := initRootCA(caPath); err == nil {
		opts := x509.VerifyOptions{
			DNSName: domain,
			Roots:   ca,
		}
		if cert, err := readCertificate(certPath); err == nil {
			if _, err := cert.Verify(opts); err == nil {
				return true
			}
		}
	}
	return false
}

func initRootCA(path string) (cert *x509.CertPool, err error) {
	if chain, err := ioutil.ReadFile(path); err == nil {

		roots := x509.NewCertPool()
		ok := roots.AppendCertsFromPEM([]byte(chain))
		if !ok {
			return nil, errors.New("Failed to parse root CA")
		}
		return roots, nil
	}
	fmt.Println("CA file does not exist!")
	return nil, errors.New("CA_NOT_EXIST")
}

func readCertificate(path string) (cert *x509.Certificate, err error) {
	if file, err := ioutil.ReadFile(path); err == nil {
		block, _ := pem.Decode([]byte(file))
		if block == nil {
			panic("failed to parse certificate PEM")
		}

		if cert, err := x509.ParseCertificate(block.Bytes); err == nil {
			return nil, errors.New("CERT_NOT_EXIST")
		} else {
			return cert, nil
		}
	}
	fmt.Println("Certificate file does not exist!")
	return nil, errors.New("CERT_NOT_EXIST")
}
