package changeclient

import (
	"database/sql"
	"fmt"

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

// Полностью сбрасывает счетчик клиентов в бд
func (c *testCon) resetSeq() {
	_, err := c.DB.Exec("ALTER SEQUENCE client_id_seq RESTART WITH 1")
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

// Проверяет любой столбец из бд
func (c *testCon) checkChangedColumn(column string, id int, value any) bool {
	var count int
	query := fmt.Sprintf("SELECT COUNT(*) FROM Clients WHERE clientid = $1 AND %s = $2 ", column)
	err := c.DB.QueryRow(query, id, value).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}
