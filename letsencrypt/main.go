package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/caarlos0/env"
	"github.com/robfig/cron"
)

//Config Object for Hosts, CertName and CaName
type Config struct {
	Hosts           string `env:"LE_HOSTS" envDefault:"/mnt/letsencrypt/hosts"`
	Email           string `env:"LE_MAIL" envDefault:"e@mail.com"`
	ReloadContainer string `env:"LE_RP" envDefault:"nginx"`
	WwwRoot         string `env:"LE_WWW" envDefault:"/var/www/letsencrypt/"`
}

func main() {

	fmt.Println("Starting letsencrypt certificate renewal tool")

	cfg := Config{}
	env.Parse(&cfg)

	renewCertificates(cfg)

	cr := cron.New()
	cr.AddFunc("@daily", func() { renewCertificates(cfg) })
	cr.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)

	go func() {
		for sig := range c {
			_ = sig
			checkForNewHosts(cfg)
		}
	}()
	for {
		time.Sleep(1000 * time.Millisecond)
	}
}

func checkForNewHosts(cfg Config) {
	fmt.Println("Received event to check for new Certificates...")

	var newHosts []string

	for _, host := range getHosts(cfg.Hosts) {
		caPath := fmt.Sprintf("/etc/letsencrypt/live/%s/fullchain.pem", host)
		certPath := fmt.Sprintf("/etc/letsencrypt/live/%s/privkey.pem", host)
		if verifyCertificate(certPath, caPath, host) {
			fmt.Printf("Certificate for %s already available.\n", host)
		} else {
			fmt.Printf("Adding %s to request\n", host)
			newHosts = append(newHosts, host)
			requestCertificate(cfg, host)
			reloadNginx(cfg.ReloadContainer)
		}
	}

	if len(newHosts) > 0 {
		fmt.Printf("Requested new certificates for %v", newHosts)
	} else {
		fmt.Println("No new hosts to add.")
	}

}

func getHosts(filename string) []string {
	if content, err := ioutil.ReadFile(filename); err == nil {
		return deleteEmpty(strings.Split(string(content), "\n"))
	}
	var s []string
	return s
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
