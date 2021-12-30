package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func splitHostPort(in string) (string, string) {
	host, port, err := net.SplitHostPort(in)
	if err != nil && !strings.Contains(err.Error(), "missing port in address") {
		log.Fatalf("%+v\n", err)
	} else if err != nil {
		host = in
	}

	return host, port
}

func main() {
	toResolve := os.Args[1]

	host, port := splitHostPort(toResolve)
	ip := net.ParseIP(host)

	// if this is an IP...
	if ip != nil {
		fmt.Printf("resolved address is: %s:%s\n", host, port)
		os.Exit(0)
	}

	// this is a domain, then
	if port == "" {
		// let's use the GET method
		resp, err := http.Get(fmt.Sprintf("https://%s/_harmony/server", host))
		if resp == nil {
			log.Fatalf("error GET-ing https://%s/_harmony/server %+v\n", host, err)
		}

		if resp.StatusCode == 404 {
			goto assume
		}

		var returned struct {
			Server string `json:"h.server"`
		}

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("error reading body: %+v\n", err)
		}

		err = json.Unmarshal(data, &returned)
		if err != nil {
			log.Fatalf("error reading JSON: %+v\n", err)
		}

		host, port = splitHostPort(toResolve)
		if port == "" {
			port = "2289"
		}
		ip = net.ParseIP(host)

		// if this is an IP...
		if ip != nil {
			fmt.Printf("resolved address is: %s:%s\n", host, port)
			os.Exit(0)
		}

		goto assume
	}

assume:
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatalf("failed to resolve IP: %+v\n", err)
	}

	if len(ips) < 1 {
		log.Fatalf("no IP addresses found\n")
	}

	ip = ips[0]

	fmt.Printf("resolved address is: %s:%s\n", ip, port)
	os.Exit(0)
}
