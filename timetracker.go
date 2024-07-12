package main

import (
	"github.com/l1qwie/TimeTracker/api/servers"
	getclients "github.com/l1qwie/TimeTracker/tests/getClients"
)

func main() {
	go servers.GetClientsInfo()
	getclients.StartTestGetClients()
}
