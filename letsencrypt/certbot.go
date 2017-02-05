package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func renewCertificates(cfg Config) {
	fmt.Println("Renewing certificates")

	cmd := "certbot"
	args := []string{"renew"}
	if out, err := exec.Command(cmd, args...).Output(); err != nil {
		fmt.Println(err)
		fmt.Println("Couldn't renew certificates")
	} else {
		fmt.Println(string(out))
		fmt.Println("Successfully renewed certificates")
	}
}

func requestCertificates(cfg Config, hosts []string) {

	fmt.Printf("Using Certbot to generate new certificates for %v \n", hosts)

	//certonly --webroot -w /var/www/example/ --agree-tos -n -m a@b.com --expand -d a.domain.co
	cmd := "certbot"
	emailArg := fmt.Sprintf("-m %s", cfg.Email)
	wwwArg := fmt.Sprintf("-w %s", cfg.WwwRoot)
	args := []string{"certonly", "--webroot", "--agree-tos", "--non-interactive", "--expand", wwwArg, emailArg}

	for _, host := range hosts {
		args = append(args, fmt.Sprintf("-d %s", host))
	}

	command := exec.Command(cmd, args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	command.Stdout = &out
	command.Stderr = &stderr
	err := command.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		fmt.Println("Requesting new certificates failed")
		return
	}
	fmt.Println("Result: " + out.String())
	fmt.Println("Successfully requested new certificates")

}
