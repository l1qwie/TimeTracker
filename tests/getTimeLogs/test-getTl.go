package gettimelogs

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/l1qwie/TimeTracker/apptype"
)

// Отсылает запрос на сервер, принимает ответ и расшифровывает его
func getReqInfo(clientid int) []*apptype.Task {
	var tasks []*apptype.Task
	resp, err := http.Get(fmt.Sprintf("http://localhost:8099/client/%d/time-logs", clientid))
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
	err = json.Unmarshal([]byte(respbody), &tasks)
	if err != nil {
		panic(err)
	}
	apptype.Debug.Println("Успешное декодирование ответа от сервера")
	return tasks
}

// Сравнение между ожидаемыми данными и полученными для clientid 1
func checkAnswersFirstCL(tasks []*apptype.Task) {
	if tasks[0].ID != 2 || tasks[1].ID != 1 || tasks[2].ID != 3 {
		panic(fmt.Sprintf(`Expected: asks[0].ID = 1 && tasks[1].ID = 2 && tasks[2].ID = 3. 
		Received: asks[0].ID = %d && tasks[1].ID = %d && tasks[2].ID = %d.`, tasks[0].ID, tasks[1].ID, tasks[2].ID))
	}
	if tasks[0].Name != "Посадить репку" || tasks[1].Name != "Выкопать яму" || tasks[2].Name != "Выростить сына" {
		panic(fmt.Sprintf(`Expected: tasks[0].Name = "Посадить репку" && tasks[1].Name = "Выкопать яму" && tasks[2].Name = "Выростить сына".
		Received: tasks[0].Name = "%s" && tasks[1].Name = "%s" && tasks[2].Name = "%s".`, tasks[0].Name, tasks[1].Name, tasks[2].Name))
	}
	if tasks[0].TimeSpent != "01:00:00" || tasks[1].TimeSpent != "03:00:00" || tasks[2].TimeSpent != "18 years" {
		panic(fmt.Sprintf(`Expected: tasks[0].TimeSpent = "01:00:00" && tasks[1].TimeSpent = "03:00:00" && tasks[2].TimeSpent = "18 years". 
		Received: tasks[0].TimeSpent = "%s" && tasks[1].TimeSpent = %s" && tasks[2].TimeSpent = "%s".`, tasks[0].TimeSpent, tasks[1].TimeSpent, tasks[2].TimeSpent))
	}
}

// Сравнение между ожидаемыми данными и полученными для clientid 2
func checkAnswersSecondCL(tasks []*apptype.Task) {
	if tasks[0].ID != 5 || tasks[1].ID != 6 || tasks[2].ID != 4 {
		panic(fmt.Sprintf(`Expected: asks[0].ID = 5 && tasks[1].ID = 6 && tasks[2].ID = 4. 
		Received: asks[0].ID = %d && tasks[1].ID = %d && tasks[2].ID = %d.`, tasks[0].ID, tasks[1].ID, tasks[2].ID))
	}
	if tasks[0].Name != "Купить машину" || tasks[1].Name != "Покушать" || tasks[2].Name != "Найти жену" {
		panic(fmt.Sprintf(`Expected: tasks[0].Name = "Купить машину" && tasks[1].Name = "Покушать" && tasks[2].Name = "Найти жену".
		Received: tasks[0].Name = "%s" && tasks[1].Name = "%s" && tasks[2].Name = "%s".`, tasks[0].Name, tasks[1].Name, tasks[2].Name))
	}
	if tasks[0].TimeSpent != "00:10:00" || tasks[1].TimeSpent != "00:15:00" || tasks[2].TimeSpent != "999 years" {
		panic(fmt.Sprintf(`Expected: tasks[0].TimeSpent = "00:10:00" && tasks[1].TimeSpent = "00:15:00" && tasks[2].TimeSpent = "999 years". 
		Received: tasks[0].TimeSpent = "%s" && tasks[1].TimeSpent = %s" && tasks[2].TimeSpent = "%s".`, tasks[0].TimeSpent, tasks[1].TimeSpent, tasks[2].TimeSpent))
	}
}

// Сравнение между ожидаемыми данными и полученными для clientid 3
func checkAnswersThirdCL(tasks []*apptype.Task) {
	if tasks[0].ID != 7 || tasks[1].ID != 8 || tasks[2].ID != 9 {
		panic(fmt.Sprintf(`Expected: asks[0].ID = 7 && tasks[1].ID = 8 && tasks[2].ID = 9. 
		Received: asks[0].ID = %d && tasks[1].ID = %d && tasks[2].ID = %d.`, tasks[0].ID, tasks[1].ID, tasks[2].ID))
	}
	if tasks[0].Name != "Выполоть грядки" || tasks[1].Name != "Посадить клубнику" || tasks[2].Name != "Погулять" {
		panic(fmt.Sprintf(`Expected: tasks[0].Name = "Выполоть грядки" && tasks[1].Name = "Посадить клубнику" && tasks[2].Name = "Погулять".
		Received: tasks[0].Name = "%s" && tasks[1].Name = "%s" && tasks[2].Name = "%s".`, tasks[0].Name, tasks[1].Name, tasks[2].Name))
	}
	if tasks[0].TimeSpent != "01:00:00" || tasks[1].TimeSpent != "02:00:00" || tasks[2].TimeSpent != "05:00:00" {
		panic(fmt.Sprintf(`Expected: tasks[0].TimeSpent = "01:00:00" && tasks[1].TimeSpent = "02:00:00" && tasks[2].TimeSpent = "05:00:00". 
		Received: tasks[0].TimeSpent = "%s" && tasks[1].TimeSpent = %s" && tasks[2].TimeSpent = "%s".`, tasks[0].TimeSpent, tasks[1].TimeSpent, tasks[2].TimeSpent))
	}
}

// Делает запрос от clientid 1 и проверяет данные
func checkFirstClient() {
	apptype.Info.Println("Тест FirstClient успешно начат")
	tasks := getReqInfo(1)
	checkAnswersFirstCL(tasks)
	apptype.Info.Println("Тест FirstClient успешно закончен")
}

// Делает запрос от clientid 2 и проверяет данные
func checkSecondClient() {
	apptype.Info.Println("Тест SecondClient успешно начат")
	tasks := getReqInfo(2)
	checkAnswersSecondCL(tasks)
	apptype.Info.Println("Тест SecondClient успешно закончен")
}

// Делает запрос от clientid 3 и проверяет данные
func checkThirdClient() {
	apptype.Info.Println("Тест ThirdClient успешно начат")
	tasks := getReqInfo(3)
	checkAnswersThirdCL(tasks)
	apptype.Info.Println("Тест ThirdClient успешно закончен")
}

// Тест endPoint /clients/:id/time-logs | get запрос.
// Все возможные тесты вызываются из этой функции
func StartTestGetTimeLogs() {
	apptype.Info.Println("Тесты с endPoint /clients/:id/time-logs | get успешно начаты")
	var err error
	con := new(testCon)
	con.DB, err = apptype.ConnectToDatabase()
	if err != nil {
		panic(err)
	}
	apptype.Debug.Println("Успешное подключение к базе данны (из тестов)")
	con.createClients()
	con.createTasks()
	defer con.deleteSeq()
	defer con.deleteCleints()
	defer con.deleteTasks()
	//
	checkFirstClient()
	checkSecondClient()
	checkThirdClient()
	//
	apptype.Info.Println("Тесты с endPoint /clients/:id/time-logs | get успешно завершены")
}
