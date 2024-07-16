package timemanager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/l1qwie/TimeTracker/apptype"
)

// Отсылает запрос на сервер, принимает ответ и расшифровывает его
func postStartTime(body []byte) string {
	var answer string
	resp, err := http.Post("http://localhost:8079/client/tasks/timeManager", "application/json", bytes.NewBuffer(body))
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

// Подготавливает запрос, и проверяет на точность ответ для Start Time
func testStartTime(con *testCon) {
	request := &apptype.Time{
		ClientId:  1,
		TaskId:    1,
		StartTime: true,
	}
	jsondata, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}
	answer := postStartTime(jsondata)
	if answer != "The countdown has begun" {
		panic(fmt.Sprintf(`Expected: "The countdown has begun". Received: %s`, answer))
	}
	if con.checkStartTimeDB(1, 1) {
		panic("The start time wasn't written into database")
	}
}

// Подготавливает запрос, и проверяет на точность ответ для End Time
func testEndTime(con *testCon) {
	request := &apptype.Time{
		ClientId:  1,
		TaskId:    1,
		StartTime: false,
	}
	jsondata, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}
	answer := postStartTime(jsondata)
	if answer != "The countdown has ended" {
		panic(fmt.Sprintf(`Expected: "The countdown has ended". Received: %s`, answer))
	}
	if con.checkEndTimeDB(1, 1) {
		panic("The end time wasn't written into database")
	}
}

// Запускает несколько предподготовительны функций, для инициализации окружения теста для Start Time
func startTime() {
	apptype.Info.Println("Тесты с endPoint /client/tasks/timeManager | post успешно начаты")
	var err error
	con := new(testCon)
	con.DB, err = apptype.ConnectToDatabase()
	if err != nil {
		panic(err)
	}
	apptype.Debug.Println("Успешное подключение к базе данны (из тестов)")

	con.createClient()
	con.createATask()
	defer con.deleteSeq()
	defer con.deleteClients()
	defer con.deleteTasks()

	testStartTime(con)
	apptype.Info.Println("Тесты с endPoint /client/tasks/timeManager | post успешно завершены")
}

// Запускает несколько предподготовительны функций, для инициализации окружения теста для End Time
func endTime() {
	apptype.Info.Println("Тесты с endPoint /client/tasks/timeManager | post успешно начаты")
	var err error
	con := new(testCon)
	con.DB, err = apptype.ConnectToDatabase()
	if err != nil {
		panic(err)
	}
	apptype.Debug.Println("Успешное подключение к базе данны (из тестов)")

	con.createClient()
	con.createATask()
	defer con.deleteSeq()
	defer con.deleteClients()
	defer con.deleteTasks()

	testEndTime(con)
	apptype.Info.Println("Тесты с endPoint /client/tasks/timeManager | post успешно завершены")
}

// Тестирует функционал начала отсчета времени и окончания отсчета времени
func StartTestTimeManager() {
	apptype.Info.Println("Тесты с endPoint /client/tasks/timeManager | post успешно начаты")
	startTime()
	endTime()
	apptype.Info.Println("Тесты с endPoint /client/tasks/timeManager | post успешно завершены")
}
