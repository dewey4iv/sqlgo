package main

import (
	"log"
	"net"
)

func main() {
	log.Println("Starting...")
	cname, addrs, err := net.LookupSRV("", "", "consul.service.iad.node.iad.consul")
	log.Println("Got Request")
	if err != nil {
		panic(err)
	}

	log.Printf(cname)

	for i := range addrs {
		log.Println(addrs[i])
	}
}
