package getclients

import (
	"database/sql"

	"github.com/l1qwie/TimeTracker/apptype"
)

type testCon struct {
	DB *sql.DB
}

// Создает 3 клментов для тестирование запроса (имя, фаммилия, отчество, возвраст)
func (c *testCon) createClientsNSPA() {
	_, err := c.DB.Exec(`
		INSERT INTO 
		Clients (clientid, name, surname, patronymic, age, passportnumber, passportseries, address) 
		VALUES (nextval('client_id_seq'), 'John', 'Doe', 'Andry', 22, '123456', '1111', 'Turkey Istanbul')`)
	if err != nil {
		panic(err)
	}

	_, err = c.DB.Exec(`
		INSERT INTO 
		Clients (clientid, name, surname, patronymic, age, passportnumber, passportseries, address) 
		VALUES (nextval('client_id_seq'), 'John', 'Doe', 'Andry', 22, '358782', '4349', 'Israel Tel-Aviv')`)
	if err != nil {
		panic(err)
	}

	_, err = c.DB.Exec(`
		INSERT INTO 
		Clients (clientid, name, surname, patronymic, age, passportnumber, passportseries, address) 
		VALUES (nextval('client_id_seq'), 'John', 'Doe', 'Andry', 22, '008923', '3221', 'USA Dallas')`)
	if err != nil {
		panic(err)
	}
	apptype.Debug.Println("Тестовые данные успешно созданы в бд")
}

// Создает 3 клментов для тестирование запроса (имя, фаммилия и отчество)
func (c *testCon) createClientsNSP() {
	_, err := c.DB.Exec(`
		INSERT INTO 
		Clients (clientid, name, surname, patronymic, age, passportnumber, passportseries, address) 
		VALUES (nextval('client_id_seq'), 'John', 'Doe', 'Andry', 22, '123456', '1111', 'Turkey Istanbul')`)
	if err != nil {
		panic(err)
	}

	_, err = c.DB.Exec(`
		INSERT INTO 
		Clients (clientid, name, surname, patronymic, age, passportnumber, passportseries, address) 
		VALUES (nextval('client_id_seq'), 'John', 'Doe', 'Andry', 23, '358782', '4349', 'Israel Tel-Aviv')`)
	if err != nil {
		panic(err)
	}

	_, err = c.DB.Exec(`
		INSERT INTO 
		Clients (clientid, name, surname, patronymic, age, passportnumber, passportseries, address) 
		VALUES (nextval('client_id_seq'), 'John', 'Doe', 'Andry', 66, '008923', '3221', 'USA Dallas')`)
	if err != nil {
		panic(err)
	}
	apptype.Debug.Println("Тестовые данные успешно созданы в бд")
}

// Создает 3 клментов для тестирование запроса (имя и фаммилия )
func (c *testCon) createClientsNS() {
	_, err := c.DB.Exec(`
		INSERT INTO 
		Clients (clientid, name, surname, patronymic, age, passportnumber, passportseries, address) 
		VALUES (nextval('client_id_seq'), 'John', 'Doe', 'Mathew', 12, '123456', '1111', 'Turkey Istanbul')`)
	if err != nil {
		panic(err)
	}

	_, err = c.DB.Exec(`
		INSERT INTO 
		Clients (clientid, name, surname, patronymic, age, passportnumber, passportseries, address) 
		VALUES (nextval('client_id_seq'), 'John', 'Doe', 'Ivanov', 83, '358782', '4349', 'Israel Tel-Aviv')`)
	if err != nil {
		panic(err)
	}

	_, err = c.DB.Exec(`
		INSERT INTO 
		Clients (clientid, name, surname, patronymic, age, passportnumber, passportseries, address) 
		VALUES (nextval('client_id_seq'), 'John', 'Doe', 'Moisha', 68, '008923', '3221', 'USA Dallas')`)
	if err != nil {
		panic(err)
	}
	apptype.Debug.Println("Тестовые данные успешно созданы в бд")
}

// Создает 3 клментов для тестирование запроса (имя и фаммилия )
func (c *testCon) createClientsN() {
	_, err := c.DB.Exec(`
		INSERT INTO 
		Clients (clientid, name, surname, patronymic, age, passportnumber, passportseries, address) 
		VALUES (nextval('client_id_seq'), 'John', 'Do', 'Mathew', 12, '123456', '1111', 'Turkey Istanbul')`)
	if err != nil {
		panic(err)
	}

	_, err = c.DB.Exec(`
		INSERT INTO 
		Clients (clientid, name, surname, patronymic, age, passportnumber, passportseries, address) 
		VALUES (nextval('client_id_seq'), 'John', 'Black', 'Ivanov', 83, '358782', '4349', 'Israel Tel-Aviv')`)
	if err != nil {
		panic(err)
	}

	_, err = c.DB.Exec(`
		INSERT INTO 
		Clients (clientid, name, surname, patronymic, age, passportnumber, passportseries, address) 
		VALUES (nextval('client_id_seq'), 'John', 'Green', 'Moisha', 68, '008923', '3221', 'USA Dallas')`)
	if err != nil {
		panic(err)
	}
	apptype.Debug.Println("Тестовые данные успешно созданы в бд")
}

// Удаляет все клиентов из таблицы Clients
func (c *testCon) deleteClients() {
	_, err := c.DB.Exec("DELETE FROM Clients")
	if err != nil {
		panic(err)
	}
	apptype.Debug.Println("Тестовые данные успешно удалены из бд")
}
