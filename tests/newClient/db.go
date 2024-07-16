package newclient

import "database/sql"

type testCon struct {
	DB *sql.DB
}

// Удаляет всех клиентов из бд
func (c *testCon) deleteClients() {
	_, err := c.DB.Exec("DELETE FROM Clients")
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

// Проверяет точно ли был создан новый клиент
func (c *testCon) newClient() bool {
	var count int
	err := c.DB.QueryRow(`
		SELECT COUNT(*) FROM Clients 
		WHERE clientid = 1 AND name = 'Ivan' AND surname = 'Ivanov' AND
		passportseries = '1234' AND passportnumber = '567890' AND address = 'Tel-Aviv Yafo'`).Scan(&count)
	if err != nil {
		panic(err)
	}
	return count > 0
}
