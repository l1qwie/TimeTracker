package main

import (
	"github.com/l1qwie/TimeTracker/api/servers"
	getclients "github.com/l1qwie/TimeTracker/tests/getClients"
	gettimelogs "github.com/l1qwie/TimeTracker/tests/getTimeLogs"
	timemanager "github.com/l1qwie/TimeTracker/tests/timeManager"
)

func main() {
	go servers.GetClientsInfo()
	getclients.StartTestGetClients()
	//
	go servers.GetTimeLogs()
	gettimelogs.StartTestGetTimeLogs()
	//
	go servers.StartTimeManager()
	timemanager.StartTestTimeManager()
}
