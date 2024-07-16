package timemanager

import "database/sql"

type testCon struct {
	DB *sql.DB
}

// Создает таску в бд
func (c *testCon) createATask() {
	_, err := c.DB.Exec(`
		INSERT INTO Tasks (taskid, clientid, taskname)
		VALUES (nextval('task_id_seq'), 1, 'Выкопать яму')`)
	if err != nil {
		panic(err)
	}
}

// Удаляет все таски из базы данных
func (c *testCon) deleteTasks() {
	_, err := c.DB.Exec("DELETE FROM Tasks")
	if err != nil {
		panic(err)
	}
}

// Создает одного клиента
func (c *testCon) createClient() {
	_, err := c.DB.Exec(`
		INSERT INTO 
		Clients (clientid, name, surname, patronymic, age, passportnumber, passportseries, address) 
		VALUES (nextval('client_id_seq'), 'John', 'Do', 'Mathew', 12, '123456', '1111', 'Turkey Istanbul')`)
	if err != nil {
		panic(err)
	}
}

// Удаляет всех клиентов из базы данных
func (c *testCon) deleteClients() {
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

// Проверяет было ли изменено столбец tasktimestart после передачи данных в endPoint
func (c *testCon) checkStartTimeDB(clientid, taskid int) bool {
	var count int
	err := c.DB.QueryRow("SELECT COUNT(*) FROM Tasks WHERE clientid = $1 AND taskid = $2 AND tasktimestart = CURRENT_TIMESTAMP", clientid, taskid).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}

// Проверяет было ли изменено столбец tasktimeend после передачи данных в endPoint
func (c *testCon) checkEndTimeDB(clientid, taskid int) bool {
	var count int
	err := c.DB.QueryRow("SELECT COUNT(*) FROM Tasks WHERE clientid = $1 AND taskid = $2 AND tasktimeend = CURRENT_TIMESTAMP", clientid, taskid).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}
