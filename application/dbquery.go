package application

import (
	"database/sql"

	"github.com/l1qwie/TimeTracker/apptype"
	_ "github.com/lib/pq"
)

type Conn struct {
	DB *sql.DB
}

// Подготавливает массив для ответа сервера
func (c *Conn) prepareArray(rows *sql.Rows, clients []*apptype.Client) ([]*apptype.Client, error) {
	var err error
	i := 0
	for rows.Next() && err == nil {
		clients[i] = &apptype.Client{}
		err = rows.Scan(&clients[i].Name, &clients[i].Surname, &clients[i].Patronymic, &clients[i].Age,
			&clients[i].PassportNumber, &clients[i].PassportSeries, &clients[i].Address)
		i++
	}
	return clients, err
}

// Выниммет из бд все данные, кроме id, орентируясь по имемни, фамилии, отчеству и возрасту
func (c *Conn) selectNameSurPatAge(name, surname, patr string, age int) ([]*apptype.Client, error) {
	var (
		count   int
		err     error
		rows    *sql.Rows
		clients []*apptype.Client
	)

	apptype.Info.Println("Запрос 1. Для выяснении длинны массива")
	err = c.DB.QueryRow(`
		SELECT COUNT(*) 
		FROM Clients 
		WHERE name = $1 AND surname = $2 AND 
		patronymic = $3 AND age = $4`, name, surname, patr, age).Scan(&count)

	if err == nil {
		apptype.Info.Println("Запрос 2. Вынимаем все данные (по условию)")
		clients = make([]*apptype.Client, count)
		rows, err = c.DB.Query(`
			SELECT name, surname, patronymic, age, passportnumber, passportseries, address 
			FROM Clients 
			WHERE name = $1 AND surname = $2 AND patronymic = $3 AND age = $4
			ORDER BY (clientid)`, name, surname, patr, age)

		if err == nil {
			clients, err = c.prepareArray(rows, clients)
		}
	}
	return clients, err
}

// Выниммет из бд все данные, кроме id, орентируясь по имемни, фамилии и отчеству
func (c *Conn) selectNameSurPat(name, surname, patr string) ([]*apptype.Client, error) {
	var (
		count   int
		err     error
		rows    *sql.Rows
		clients []*apptype.Client
	)
	apptype.Info.Println("Запрос 1. Для выяснении длинны массива")
	err = c.DB.QueryRow(`
		SELECT COUNT(*) 
		FROM Clients 
		WHERE name = $1 AND surname = $2 AND 
		patronymic = $3`, name, surname, patr).Scan(&count)

	if err == nil {
		apptype.Info.Println("Запрос 2. Вынимаем все данные (по условию)")
		clients = make([]*apptype.Client, count)
		rows, err = c.DB.Query(`
			SELECT name, surname, patronymic, age, passportnumber, passportseries, address 
			FROM Clients 
			WHERE name = $1 AND surname = $2 AND patronymic = $3
			ORDER BY (clientid)`, name, surname, patr)

		if err == nil {
			clients, err = c.prepareArray(rows, clients)
		}
	}
	return clients, err
}

// Выниммет из бд все данные, кроме id, орентируясь по имемни и фамилии
func (c *Conn) selectNameSur(name, surname string) ([]*apptype.Client, error) {
	var (
		count   int
		err     error
		rows    *sql.Rows
		clients []*apptype.Client
	)
	apptype.Info.Println("Запрос 1. Для выяснении длинны массива")
	err = c.DB.QueryRow(`
		SELECT COUNT(*) 
		FROM Clients 
		WHERE name = $1 AND surname = $2`, name, surname).Scan(&count)

	if err == nil {
		apptype.Info.Println("Запрос 2. Вынимаем все данные (по условию)")
		clients = make([]*apptype.Client, count)
		rows, err = c.DB.Query(`
			SELECT name, surname, patronymic, age, passportnumber, passportseries, address 
			FROM Clients 
			WHERE name = $1 AND surname = $2
			ORDER BY (clientid)`, name, surname)

		if err == nil {
			clients, err = c.prepareArray(rows, clients)
		}
	}
	return clients, err
}

// Выниммет из бд все данные, кроме id, орентируясь по имемни
func (c *Conn) selectName(name string) ([]*apptype.Client, error) {
	var (
		count   int
		err     error
		rows    *sql.Rows
		clients []*apptype.Client
	)
	apptype.Info.Println("Запрос 1. Для выяснении длинны массива")
	err = c.DB.QueryRow(`
		SELECT COUNT(*) 
		FROM Clients 
		WHERE name = $1`, name).Scan(&count)

	if err == nil {
		apptype.Info.Println("Запрос 2. Вынимаем все данные (по условию)")
		clients = make([]*apptype.Client, count)
		rows, err = c.DB.Query(`
			SELECT name, surname, patronymic, age, passportnumber, passportseries, address 
			FROM Clients 
			WHERE name = $1
			ORDER BY (clientid)`, name)

		if err == nil {
			clients, err = c.prepareArray(rows, clients)
		}
	}
	return clients, err
}
