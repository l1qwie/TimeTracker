package newclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/l1qwie/TimeTracker/apptype"
)

// Контактирует с сервером /client/new, отсылает запрос, принимет ответ и расшифровывает его
func newClient(body []byte) *apptype.People {
	people := new(apptype.People)
	resp, err := http.Post("http://localhost:8019/client/new", "application/json", bytes.NewBuffer(body))
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
	err = json.Unmarshal([]byte(respbody), &people)
	if err != nil {
		panic(err)
	}
	apptype.Debug.Println("Успешное декодирование ответа от сервера")
	return people
}

// Формирует запрос, отправляет его в функцию, которая контактирует с сервером и
// проверяет на коректность данные полученный от сервера
func add(con *testCon) {
	request := &apptype.NewClient{
		Passport: "1234 567890",
	}
	jsonbody, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}
	people := newClient(jsonbody)
	if people.Name != "Ivan" {
		panic(fmt.Sprintf("The new client's name isn't Ivan. It is %s", people.Name))
	}
	if people.Surname != "Ivanov" {
		panic(fmt.Sprintf("The new client's surname isn't Ivanov. It is: %s", people.Surname))
	}
	if people.Address != "Tel-Aviv Yafo" {
		panic(fmt.Sprintf("The new client's address isn't Tel-Aviv Yafo. It is: %s", people.Address))
	}
	if !con.newClient() {
		panic("There's no new client in the database")
	}
}

// Запуск всех тестов для добавления нового клиента
func StartTestNewClient() {
	apptype.Info.Println("Тесты с endPoint /client/new | post успешно начаты")
	var err error
	con := new(testCon)
	con.DB, err = apptype.ConnectToDatabase()
	if err != nil {
		panic(err)
	}
	apptype.Debug.Println("Успешное подключение к базе данны (из тестов)")

	defer con.deleteClients()
	defer con.resetSeq()

	add(con)
	apptype.Info.Println("Тесты с endPoint /client/new | post успешно закончены")
}
