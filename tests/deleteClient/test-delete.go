package deleteclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/l1qwie/TimeTracker/apptype"
)

// Отсылает запрос на сервер, принимает ответ и расшифровывает его
func deleteClient(id int) string {
	var answer string
	url := fmt.Sprintf("http://localhost:8059/client/%d/delete", id)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	apptype.Debug.Println("Успешный запрос на сервер")
	defer resp.Body.Close()
	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	apptype.Debug.Printf("Данные из ответа: %s", string(respbody))
	err = json.Unmarshal([]byte(respbody), &answer)
	if err != nil {
		panic(err)
	}
	apptype.Debug.Println("Успешное декодирование ответа от сервера")
	return answer
}

// Вызывает функцию вызова сервера и проверяет ответы
func delete(con *testCon) {
	answer := deleteClient(1)
	if answer != fmt.Sprintf("The client %d has been successfuly deleted", 1) {
		panic(fmt.Sprintf("Expected: The client {id} has been successfuly deleted. Recieved: %s", answer))
	}
	if !con.checkDeletedTask(1) {
		panic("The client's tasks weren't deleted")
	}
	if !con.checkDeletedClient(1) {
		panic("The clients wasn't deleted")
	}
}

// Запускает тесты для /client/:id/delete | delete
func StartTestDeleteClient() {
	apptype.Info.Println("Тесты с endPoint /client/:id/delete | delete успешно начаты")
	var err error
	con := new(testCon)
	con.DB, err = apptype.ConnectToDatabase()
	if err != nil {
		panic(err)
	}
	apptype.Debug.Println("Успешное подключение к базе данны (из тестов)")

	con.createClient()
	con.createATask()
	defer con.resetSeq()
	defer con.deleteClients()
	defer con.deleteTasks()

	delete(con)

	apptype.Info.Println("Тесты с endPoint /client/:id/delete | delete успешно закончены")
}
