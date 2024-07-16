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

// Получение данных пользователей
func getClients() {
	go servers.GetClientsInfo()
	getclients.StartTestGetClients()
}

// Получение трудозатрат по пользователю за период задача-сумма часов
// и минут с сортировкой от большей затраты к меньшей
func getTimeLogs() {
	go servers.GetTimeLogs()
	gettimelogs.StartTestGetTimeLogs()
}

// Начать отсчет времени по задаче для пользователя и
// закончить отсчет времени по задаче для пользователя
func timeManager() {
	go servers.StartTimeManager()
	timemanager.StartTestTimeManager()
}

// Удаление пользователя
func deleteClient() {
	go servers.DeleteClient()
	deleteclient.StartTestDeleteClient()
}

// Изменение данных пользователя
func changeClient() {
	go servers.ChangeClient()
	changeclient.StartTestChangeClient()
}

// Добавление нового пользователя
func newClient() {
	go servers.NewClient()
	newclient.StartTestNewClient()
}

// Включить все сервера
func turnAllOn() {
	// Получение данных пользователей
	go servers.GetClientsInfo()

	// Получение трудозатрат по пользователю за период задача-сумма часов
	// и минут с сортировкой от большей затраты к меньшей
	go servers.GetTimeLogs()

	// Начать отсчет времени по задаче для пользователя и
	// закончить отсчет времени по задаче для пользователя
	go servers.StartTimeManager()

	// Удаление пользователя
	go servers.DeleteClient()

	// Изменение данных пользователя
	go servers.ChangeClient()

	// Добавление нового пользовател
	servers.NewClient()
}

// Прокрутка всех тестов
func testAll() {
	// Получение данных пользователей
	go servers.GetClientsInfo()
	getclients.StartTestGetClients()

	// Получение трудозатрат по пользователю за период задача-сумма часов
	// и минут с сортировкой от большей затраты к меньшей
	go servers.GetTimeLogs()
	gettimelogs.StartTestGetTimeLogs()

	// Начать отсчет времени по задаче для пользователя и
	// закончить отсчет времени по задаче для пользователя
	go servers.StartTimeManager()
	timemanager.StartTestTimeManager()

	// Удаление пользователя
	go servers.DeleteClient()
	deleteclient.StartTestDeleteClient()

	// Изменение данных пользователя
	go servers.ChangeClient()
	changeclient.StartTestChangeClient()

	// Добавление нового пользовател
	go servers.NewClient()
	newclient.StartTestNewClient()
}

func main() {
	testAll() //Прокрутка всех тестов
	//turnAllOn() //Включить все сервера

	//Сервера по отдельности
	//getClients()
	//getTimeLogs()
	//timeManager()
	//deleteClient()
	//changeClient()
	//newClient()
}
