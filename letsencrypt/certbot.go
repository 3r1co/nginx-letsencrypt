package main

import (
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
	//cmd := "certbot"
	cmd := "echo"
	emailArg := fmt.Sprintf("-m %s", cfg.Email)
	wwwArg := fmt.Sprintf("-w %s", cfg.WwwRoot)
	args := []string{"-n", "\"", "certonly", "--webroot", "--agree-tos", "--non-interactive", "--expand", wwwArg, emailArg}

	for _, host := range hosts {
		args = append(args, fmt.Sprintf("-d %s", host))
	}
	args = append(args, "\"")

	if out, err := exec.Command(cmd, args...).Output(); err != nil {
		fmt.Println(err)
		fmt.Println("Requesting new certificates failed")
	} else {
		fmt.Println(string(out))
		fmt.Println("Successfully requested new certificates")
	}

}
