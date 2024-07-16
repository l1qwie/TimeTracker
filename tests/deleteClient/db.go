package deleteclient

import (
	"database/sql"

	"github.com/l1qwie/TimeTracker/apptype"
)

type testCon struct {
	DB *sql.DB
}

// Создает клиента для тестов
func (c *testCon) createClient() {
	_, err := c.DB.Exec(`
		INSERT INTO 
		Clients (clientid, name, surname, patronymic, age, passportnumber, passportseries, address) 
		VALUES (nextval('client_id_seq'), 'John', 'Do', 'Mathew', 12, '123456', '1111', 'Turkey Istanbul')`)
	if err != nil {
		panic(err)
	}
}

// Создает таску для тестов
func (c *testCon) createATask() {
	_, err := c.DB.Exec(`
	INSERT INTO Tasks (taskid, clientid, taskname, tasktimestart, tasktimeend)
	VALUES (nextval('task_id_seq'), 1, 'Выкопать яму', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + INTERVAL '3 hours')`)
	if err != nil {
		panic(err)
	}
}

// Полностью сбрасывает все счетчики в бд
func (c *testCon) resetSeq() {
	_, err := c.DB.Exec("ALTER SEQUENCE task_id_seq RESTART WITH 1")
	if err != nil {
		panic(err)
	}
	_, err = c.DB.Exec("ALTER SEQUENCE client_id_seq RESTART WITH 1")
	if err != nil {
		panic(err)
	}
}

// Удаляет всех клиентов из бд
func (c *testCon) deleteClients() {
	_, err := c.DB.Exec("DELETE FROM Clients")
	if err != nil {
		panic(err)
	}
	apptype.Debug.Println("Тестовые данные успешно удалены из бд")
}

// Удаляет все таски из бд
func (c *testCon) deleteTasks() {
	_, err := c.DB.Exec("DELETE FROM Tasks")
	if err != nil {
		panic(err)
	}
	apptype.Debug.Println("Тестовые данные успешно удалены из бд")
}

// Проверяет на наличие "удаленной" строчки в бд и если не находит возвращает true
func (c *testCon) checkDeletedClient(clientid int) bool {
	var count int
	err := c.DB.QueryRow("SELECT COUNT(*) FROM Clients WHERE clientid = $1", clientid).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count == 0
}

// Проверяет на наличие "удаленные" строчек в бд и если не находит возвращает true
func (c *testCon) checkDeletedTask(clientid int) bool {
	var count int
	err := c.DB.QueryRow("SELECT COUNT(*) FROM Tasks WHERE clientid = $1", clientid).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count == 0
}
