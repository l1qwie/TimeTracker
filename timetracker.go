package main

import (
	"github.com/l1qwie/TimeTracker/api/servers"
	changeclient "github.com/l1qwie/TimeTracker/tests/changeClient"
	deleteclient "github.com/l1qwie/TimeTracker/tests/deleteClient"
	getclients "github.com/l1qwie/TimeTracker/tests/getClients"
	gettimelogs "github.com/l1qwie/TimeTracker/tests/getTimeLogs"
	newclient "github.com/l1qwie/TimeTracker/tests/newClient"
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
	//
	go servers.DeleteClient()
	deleteclient.StartTestDeleteClient()
	//
	go servers.ChangeClient()
	changeclient.StartTestChangeClient()
	//
	go servers.NewClient()
	newclient.StartTestNewClient()
}
