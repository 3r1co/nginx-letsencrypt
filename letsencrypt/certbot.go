package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
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

func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func requestCertificate(cfg Config, host string) {

	fmt.Printf("Using Certbot to generate new certificate for %s \n", host)

	command := exec.Command("/le_wrapper.sh", cfg.Email, host)

	printCommand(command)

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
