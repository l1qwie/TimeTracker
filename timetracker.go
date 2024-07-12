package main

import (
	"github.com/l1qwie/TimeTracker/api/servers"
	getclients "github.com/l1qwie/TimeTracker/tests/getClients"
	gettimelogs "github.com/l1qwie/TimeTracker/tests/getTimeLogs"
)

func main() {
	go servers.GetClientsInfo()
	getclients.StartTestGetClients()
	//
	go servers.GetTimeLogs()
	gettimelogs.StartTestGetTimeLogs()
}
