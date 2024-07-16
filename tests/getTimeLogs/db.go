package gettimelogs

import (
	"database/sql"

	"github.com/l1qwie/TimeTracker/apptype"
)

type testCon struct {
	DB *sql.DB
}

// Создает 3 клиентов для тестирование запроса с задачами
func (c *testCon) createClients() {
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

// Создает по три задачи на каждого клиента (всего 3 клиента)
func (c *testCon) createTasks() {
	_, err := c.DB.Exec(`
		INSERT INTO Tasks (taskid, clientid, taskname, tasktimestart, tasktimeend)
		VALUES (nextval('task_id_seq'), 1, 'Выкопать яму', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + INTERVAL '3 hours')`)
	if err != nil {
		panic(err)
	}

	_, err = c.DB.Exec(`
		INSERT INTO Tasks (taskid, clientid, taskname, tasktimestart, tasktimeend)
		VALUES (nextval('task_id_seq'), 1, 'Посадить репку', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + INTERVAL '1 hours')`)
	if err != nil {
		panic(err)
	}

	_, err = c.DB.Exec(`
		INSERT INTO Tasks (taskid, clientid, taskname, tasktimestart, tasktimeend)
		VALUES (nextval('task_id_seq'), 1, 'Выростить сына', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + INTERVAL '18 years')`)
	if err != nil {
		panic(err)
	}

	_, err = c.DB.Exec(`
		INSERT INTO Tasks (taskid, clientid, taskname, tasktimestart, tasktimeend)
		VALUES (nextval('task_id_seq'), 2, 'Найти жену', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + INTERVAL '999 years')`)
	if err != nil {
		panic(err)
	}

	_, err = c.DB.Exec(`
		INSERT INTO Tasks (taskid, clientid, taskname, tasktimestart, tasktimeend)
		VALUES (nextval('task_id_seq'), 2, 'Купить машину', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + INTERVAL '10 minutes')`)
	if err != nil {
		panic(err)
	}

	_, err = c.DB.Exec(`
		INSERT INTO Tasks (taskid, clientid, taskname, tasktimestart, tasktimeend)
		VALUES (nextval('task_id_seq'), 2, 'Покушать', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + INTERVAL '15 minutes')`)
	if err != nil {
		panic(err)
	}

	_, err = c.DB.Exec(`
		INSERT INTO Tasks (taskid, clientid, taskname, tasktimestart, tasktimeend)
		VALUES (nextval('task_id_seq'), 3, 'Выполоть грядки', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + INTERVAL '1 hours')`)
	if err != nil {
		panic(err)
	}

	_, err = c.DB.Exec(`
		INSERT INTO Tasks (taskid, clientid, taskname, tasktimestart, tasktimeend)
		VALUES (nextval('task_id_seq'), 3, 'Посадить клубнику', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + INTERVAL '2 hours')`)
	if err != nil {
		panic(err)
	}

	_, err = c.DB.Exec(`
		INSERT INTO Tasks (taskid, clientid, taskname, tasktimestart, tasktimeend)
		VALUES (nextval('task_id_seq'), 3, 'Погулять', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + INTERVAL '5 hours')`)
	if err != nil {
		panic(err)
	}
}

// Удаляет абсолютно все задачи из таблицы Tasks
func (c *testCon) deleteTasks() {
	_, err := c.DB.Exec("DELETE FROM Tasks")
	if err != nil {
		panic(err)
	}
}

// Удаляет всех клиентов из таблицы Clients
func (c *testCon) deleteCleints() {
	_, err := c.DB.Exec("DELETE FROM Clients")
	if err != nil {
		panic(err)
	}
}

// Обноляет все счетчики
func (c *testCon) deleteSeq() {
	_, err := c.DB.Exec("ALTER SEQUENCE task_id_seq RESTART WITH 1")
	if err != nil {
		panic(err)
	}
	_, err = c.DB.Exec("ALTER SEQUENCE client_id_seq RESTART WITH 1")
	if err != nil {
		panic(err)
	}
}
