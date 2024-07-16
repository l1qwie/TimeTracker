package application

import (
	"database/sql"
	"fmt"

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

// Ищет клиента в таблице Clients по id
func (c *Conn) findClient(id int) (bool, error) {
	var count int
	err := c.DB.QueryRow("SELECT COUNT(*) FROM Clients WHERE clientid = $1", id).Scan(&count)
	return count > 0, err
}

// Выбирает все таски определенного клиента (id)
func (c *Conn) selectClientTasks(id int) ([]*apptype.Task, error) {
	var (
		count int
		err   error
		tasks []*apptype.Task
		rows  *sql.Rows
	)
	apptype.Debug.Printf("Запрос 1. Нужно найти сколько всего задачей у клиента")
	err = c.DB.QueryRow("SELECT COUNT(*) FROM Tasks WHERE clientid = $1", id).Scan(&count)

	if err == nil {
		apptype.Debug.Printf("Запрос 2. Задачи успешно найдены. Теперь данные из этих задач")
		tasks = make([]*apptype.Task, count)
		rows, err = c.DB.Query(`
			SELECT taskid, taskname, AGE(tasktimeend, tasktimestart) AS time_interval
			FROM Tasks
			WHERE clientid = $1
			ORDER BY time_interval`, id)

		if err == nil {
			i := 0
			for rows.Next() && err == nil {
				tasks[i] = &apptype.Task{}
				err = rows.Scan(&tasks[i].ID, &tasks[i].Name, &tasks[i].TimeSpent)
				i++
			}
		}
	}
	return tasks, err
}

// Обновляет данные для начала отсчета вреемени
func (c *Conn) updateTaskStartTime(clientid int, taskid int) error {
	_, err := c.DB.Exec("UPDATE Tasks SET tasktimestart = CURRENT_TIMESTAMP WHERE taskid = $1 AND clientid = $2", taskid, clientid)
	return err
}

// Ищет нужную таску по переданному taskid
func (c *Conn) findTask(taskid int) (bool, error) {
	var count int
	err := c.DB.QueryRow("SELECT COUNT(*) FROM Tasks WHERE taskid = $1", taskid).Scan(&count)
	return count > 0, err
}

// Обновляет данные для окончания отсчета времени
func (c *Conn) updateTaskEndTime(clientid int, taskid int) error {
	_, err := c.DB.Exec("UPDATE Tasks SET tasktimeend = CURRENT_TIMESTAMP WHERE taskid = $1 AND clientid = $2", taskid, clientid)
	return err
}

// Окончание транзакции в зависимости от значения переданной переменной
func (c *Conn) endTransacttion(err error) {
	if err != nil {
		_, err2 := c.DB.Exec("ROLLBACK")
		if err2 != nil {
			apptype.Debug.Printf("IMPOSSIBLE TO END THE TRANSACTION! :%s", err2)
		}
	} else {
		_, err2 := c.DB.Exec("COMMIT")
		if err2 != nil {
			apptype.Debug.Printf("IMPOSSIBLE TO END THE TRANSACTION! :%s", err2)
		}
	}
}

// Удаляет все таски клиента, а потом удаляет и самого клиента
func (c *Conn) deleteClientDB(clientid int) error {
	_, err := c.DB.Exec("BEGIN ISOLATION LEVEL REPEATABLE READ")
	defer c.endTransacttion(err)
	if err == nil {
		_, err = c.DB.Exec("DELETE FROM Tasks WHERE clientid = $1", clientid)
		if err == nil {
			_, err = c.DB.Exec("DELETE FROM Clients WHERE clientid = $1", clientid)
		}
	}
	return err
}

// Проверяет существует ли переданный столбец в таблице Clients
func (c *Conn) findColumn(column string) (bool, error) {
	var ok bool
	query := "SELECT EXISTS (SELECT 1 FROM information_schema.columns WHERE table_name = 'clients' AND column_name = $1)"
	err := c.DB.QueryRow(query, column).Scan(&ok)
	apptype.Debug.Println("Query:", query, "Column:", column, "Result:", ok)
	return ok, err
}

// Обновляет данные в таблице Clients только для типа int
func (c *Conn) updateClientColumnInt(clientid int, column string, value int) error {
	query := fmt.Sprintf("UPDATE Clients SET %s = $2 WHERE clientid = $1", column)
	_, err := c.DB.Exec(query, clientid, value)
	return err
}

// Обновляет данные в таблице Clients только для типа string
func (c *Conn) updateClientColumnStr(clientid int, column, value string) error {
	query := fmt.Sprintf("UPDATE Clients SET %s = $2 WHERE clientid = $1", column)
	_, err := c.DB.Exec(query, clientid, value)
	return err
}

// Создает нового клиента в бд
func (c *Conn) addClientDB(p *apptype.People, series, number string) error {
	_, err := c.DB.Exec(`
		INSERT INTO Clients 
		(clientid, name, surname, passportseries, passportnumber, address)
		VALUES (nextval('client_id_seq'), $1, $2, $3, $4, $5)`, p.Name, p.Surname, series, number, p.Address)
	return err
}
