package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/lealhugui/vice/data"

	"github.com/lealhugui/vice/cmd/worker/config"

	mainRoutes "github.com/lealhugui/vice/routeManager"
)

func getIPAddr() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	var ip net.IP
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {

			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			break
		}
	}
	if ip == nil {
		return "", nil
	} else {
		return ip.String(), nil
	}
}

func RegisterOnInit() {
	ip, err := getIPAddr()
	if err != nil {
		log.Fatalln(err)
	}

	body, err := json.Marshal(mainRoutes.RegisterWorkerRequest{
		Host: ip,
	})

	resp, err := http.Post(fmt.Sprintf("%s/registerWorker", config.AppConfig.MasterHost),
		"application/json", bytes.NewBuffer(body))

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var respBody data.Worker
	allBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(allBytes, &respBody)
	if err != nil {
		log.Fatalln(err)
	}

	log.Print(respBody)
}
