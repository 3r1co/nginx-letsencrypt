package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"
)

func reloadNginx(container string) {

	httpc := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", "/var/run/docker.sock")
			},
		},
	}

	connection := fmt.Sprintf("http://unix/containers/%s/kill?signal=HUP", container)

	if _, err := httpc.Post(connection, "", strings.NewReader("")); err != nil {
		fmt.Println("Couldn't connect to docker socket, please check you mounted /var/run/docker.sock")
	} else {
		fmt.Printf("Sent HUP Signal to %s", container)
	}
}
