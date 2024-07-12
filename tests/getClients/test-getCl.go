package getclients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/l1qwie/TimeTracker/apptype"
)

// ?name=John&surname=Doe&patronymic=Andry&age=22
//panic(fmt.Sprintf(`Expected: . Received: .`))

// Проверяет общие данные для этого запроса (создано по 3 клиента в бд для каждого варианта запроса)
func checkForAll(cl []*apptype.Client) {
	if len(cl) != 3 {
		panic(fmt.Sprintf("Expected: len([]*apptype.Client) = 3. Received: len([]*apptype.Client) = %d.", len(cl)))
	}
	if cl[0].Address != "Turkey Istanbul" {
		panic(fmt.Sprintf(`Expected: cl[0].Address = "Turkey Istanbul". Received: cl[0].Address = "%s".`, cl[0].Address))
	}
	if cl[1].Address != "Israel Tel-Aviv" {
		panic(fmt.Sprintf(`Expected: cl[1].Address = "Israel Tel-Aviv". Received: cl[1].Address = "%s".`, cl[1].Address))
	}
	if cl[2].Address != "USA Dallas" {
		panic(fmt.Sprintf(`Expected: cl[2].Address = "USA Dallas". Received: cl[2].Address = "%s".`, cl[2].Address))
	}
	if cl[0].PassportNumber != "123456" {
		panic(fmt.Sprintf(`Expected: cl[0].PassportNumber = "123456". Received: cl[0].PassportNumber = "%s".`, cl[0].PassportNumber))
	}
	if cl[1].PassportNumber != "358782" {
		panic(fmt.Sprintf(`Expected: cl[1].PassportNumber = "358782". Received: cl[1].PassportNumber = "%s".`, cl[1].PassportNumber))
	}
	if cl[2].PassportNumber != "008923" {
		panic(fmt.Sprintf(`Expected: cl[2].PassportNumber = "008923". Received: cl[2].PassportNumber = "%s".`, cl[2].PassportNumber))
	}
	if cl[0].PassportSeries != "1111" {
		panic(fmt.Sprintf(`Expected: cl[0].PassportSeries = "1111". Received: cl[0].PassportSeries = "%s".`, cl[0].PassportSeries))
	}
	if cl[1].PassportSeries != "4349" {
		panic(fmt.Sprintf(`Expected: cl[1].PassportSeries = "4349". Received: cl[1].PassportSeries = "%s" .`, cl[1].PassportSeries))
	}
	if cl[2].PassportSeries != "3221" {
		panic(fmt.Sprintf(`Expected: cl[2].PassportSeries = "3221". Received: cl[2].PassportSeries = "%s".`, cl[2].PassportSeries))
	}
}

// Отсылает запрос на сервер, принимает ответ и расшифровывает его
func getReqInfo(query string) []*apptype.Client {
	var clients []*apptype.Client
	resp, err := http.Get(fmt.Sprintf("http://localhost:8088/client%s", query))
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
	err = json.Unmarshal([]byte(respbody), &clients)
	if err != nil {
		panic(err)
	}
	apptype.Debug.Println("Успешное декодирование ответа от сервера")
	return clients
}

func checkAnswersNSPA(cl []*apptype.Client) {
	checkForAll(cl)
	if cl[0].Name != "John" && cl[1].Name != "John" && cl[2].Name != "John" {
		panic(fmt.Sprintf(`Expected: cl[0].Name = "John" && cl[1].Name = "John" && cl[2].Name = "John". 
		Received: cl[0].Name = "%s" && cl[1].Name = "%s" && cl[2].Name = "%s".`, cl[0].Name, cl[1].Name, cl[2].Name))
	}
	if cl[0].Surname != "Doe" && cl[1].Surname != "Doe" && cl[2].Surname != "Doe" {
		panic(fmt.Sprintf(`Expected: cl[0].Surname = "Doe" && cl[1].Surname = "Doe" && cl[2].Surname = "Doe". 
		Received: cl[0].Surname = "%s" && cl[1].Surname = "%s" && cl[2].Surname = "%s".`, cl[0].Surname, cl[1].Surname, cl[2].Surname))
	}
	if cl[0].Patronymic != "Andry" && cl[1].Patronymic != "Andry" && cl[2].Patronymic != "Andry" {
		panic(fmt.Sprintf(`Expected: cl[0].Patronymic = "Andry" && cl[1].Patronymic = "Andry" && cl[2].Patronymic = "Andry". 
		Received: cl[0].Patronymic = "%s" && cl[1].Patronymic = "%s" && cl[2].Patronymic = "%s".`, cl[0].Patronymic, cl[1].Patronymic, cl[2].Patronymic))
	}
	if cl[0].Age != 22 && cl[1].Age != 22 && cl[2].Age != 22 {
		panic(fmt.Sprintf(`Expected: cl[0].Age = 22 && cl[1].Age = 22 && cl[2].Age = 22. 
		Received: cl[0].Age = %d && cl[1].Age = %d && cl[2].Age = %d.`, cl[0].Age, cl[1].Age, cl[2].Age))
	}
}

func checkAnswersNSP(cl []*apptype.Client) {
	checkForAll(cl)
	if cl[0].Name != "John" && cl[1].Name != "John" && cl[2].Name != "John" {
		panic(fmt.Sprintf(`Expected: cl[0].Name = "John" && cl[1].Name = "John" && cl[2].Name = "John". 
		Received: cl[0].Name = "%s" && cl[1].Name = "%s" && cl[2].Name = "%s".`, cl[0].Name, cl[1].Name, cl[2].Name))
	}
	if cl[0].Surname != "Doe" && cl[1].Surname != "Doe" && cl[2].Surname != "Doe" {
		panic(fmt.Sprintf(`Expected: cl[0].Surname = "Doe" && cl[1].Surname = "Doe" && cl[2].Surname = "Doe". 
		Received: cl[0].Surname = "%s" && cl[1].Surname = "%s" && cl[2].Surname = "%s".`, cl[0].Surname, cl[1].Surname, cl[2].Surname))
	}
	if cl[0].Patronymic != "Andry" && cl[1].Patronymic != "Andry" && cl[2].Patronymic != "Andry" {
		panic(fmt.Sprintf(`Expected: cl[0].Patronymic = "Andry" && cl[1].Patronymic = "Andry" && cl[2].Patronymic = "Andry". 
		Received: cl[0].Patronymic = "%s" && cl[1].Patronymic = "%s" && cl[2].Patronymic = "%s".`, cl[0].Patronymic, cl[1].Patronymic, cl[2].Patronymic))
	}
	if cl[0].Age != 22 && cl[1].Age != 23 && cl[2].Age != 66 {
		panic(fmt.Sprintf(`Expected: cl[0].Age = 22 && cl[1].Age = 23 && cl[2].Age = 66. 
		Received: cl[0].Age = %d && cl[1].Age = %d && cl[2].Age = %d.`, cl[0].Age, cl[1].Age, cl[2].Age))
	}
}

func checkAnswersNS(cl []*apptype.Client) {
	checkForAll(cl)
	if cl[0].Name != "John" && cl[1].Name != "John" && cl[2].Name != "John" {
		panic(fmt.Sprintf(`Expected: cl[0].Name = "John" && cl[1].Name = "John" && cl[2].Name = "John". 
		Received: cl[0].Name = "%s" && cl[1].Name = "%s" && cl[2].Name = "%s".`, cl[0].Name, cl[1].Name, cl[2].Name))
	}
	if cl[0].Surname != "Doe" && cl[1].Surname != "Doe" && cl[2].Surname != "Doe" {
		panic(fmt.Sprintf(`Expected: cl[0].Surname = "Doe" && cl[1].Surname = "Doe" && cl[2].Surname = "Doe". 
		Received: cl[0].Surname = "%s" && cl[1].Surname = "%s" && cl[2].Surname = "%s".`, cl[0].Surname, cl[1].Surname, cl[2].Surname))
	}
	if cl[0].Patronymic != "Mathew" && cl[1].Patronymic != "Ivanov" && cl[2].Patronymic != "Moisha" {
		panic(fmt.Sprintf(`Expected: cl[0].Patronymic = "Andry" && cl[1].Patronymic = "Andry" && cl[2].Patronymic = "Andry". 
		Received: cl[0].Patronymic = "%s" && cl[1].Patronymic = "%s" && cl[2].Patronymic = "%s".`, cl[0].Patronymic, cl[1].Patronymic, cl[2].Patronymic))
	}
	if cl[0].Age != 12 && cl[1].Age != 83 && cl[2].Age != 68 {
		panic(fmt.Sprintf(`Expected: cl[0].Age = 12 && cl[1].Age = 83 && cl[2].Age = 68. 
		Received: cl[0].Age = %d && cl[1].Age = %d && cl[2].Age = %d.`, cl[0].Age, cl[1].Age, cl[2].Age))
	}
}

func checkAnswersN(cl []*apptype.Client) {
	checkForAll(cl)
	if cl[0].Name != "John" && cl[1].Name != "John" && cl[2].Name != "John" {
		panic(fmt.Sprintf(`Expected: cl[0].Name = "John" && cl[1].Name = "John" && cl[2].Name = "John". 
		Received: cl[0].Name = "%s" && cl[1].Name = "%s" && cl[2].Name = "%s".`, cl[0].Name, cl[1].Name, cl[2].Name))
	}
	if cl[0].Surname != "Do" && cl[1].Surname != "Black" && cl[2].Surname != "Green" {
		panic(fmt.Sprintf(`Expected: cl[0].Surname = "Do" && cl[1].Surname = "Black" && cl[2].Surname = "Green". 
		Received: cl[0].Surname = "%s" && cl[1].Surname = "%s" && cl[2].Surname = "%s".`, cl[0].Surname, cl[1].Surname, cl[2].Surname))
	}
	if cl[0].Patronymic != "Mathew" && cl[1].Patronymic != "Ivanov" && cl[2].Patronymic != "Moisha" {
		panic(fmt.Sprintf(`Expected: cl[0].Patronymic = "Mathew" && cl[1].Patronymic = "Ivanov" && cl[2].Patronymic = "Moisha". 
		Received: cl[0].Patronymic = "%s" && cl[1].Patronymic = "%s" && cl[2].Patronymic = "%s".`, cl[0].Patronymic, cl[1].Patronymic, cl[2].Patronymic))
	}
	if cl[0].Age != 12 && cl[1].Age != 83 && cl[2].Age != 68 {
		panic(fmt.Sprintf(`Expected: cl[0].Age = 12 && cl[1].Age = 83 && cl[2].Age = 68. 
		Received: cl[0].Age = %d && cl[1].Age = %d && cl[2].Age = %d.`, cl[0].Age, cl[1].Age, cl[2].Age))
	}
}

// Тестирует запрос и ответ от сервера, если отправить
// имя, фамилию, отчество и возраст клиента
func testNameSurPatrAgeRequest(con *testCon) {
	apptype.Info.Println("Тест NSPA успешно начат")
	con.createClientsNSPA()
	defer con.deleteClients()
	clients := getReqInfo("?name=John&surname=Doe&patronymic=Andry&age=22")
	checkAnswersNSPA(clients)
	apptype.Info.Println("Тест NSPA успешно завершен")
}

func testNameSurPartRequest(con *testCon) {
	apptype.Info.Println("Тест NSP успешно начат")
	con.createClientsNSP()
	defer con.deleteClients()
	clients := getReqInfo("?name=John&surname=Doe&patronymic=Andry")
	checkAnswersNSP(clients)
	apptype.Info.Println("Тест NSP успешно завершен")
}

func testNameSurRequest(con *testCon) {
	apptype.Info.Println("Тест NS успешно начат")
	con.createClientsNS()
	defer con.deleteClients()
	clients := getReqInfo("?name=John&surname=Doe")
	checkAnswersNS(clients)
	apptype.Info.Println("Тест NS успешно завершен")
}

func testNameRequest(con *testCon) {
	apptype.Info.Println("Тест NS успешно начат")
	con.createClientsN()
	defer con.deleteClients()
	clients := getReqInfo("?name=John")
	checkAnswersN(clients)
	apptype.Info.Println("Тест NS успешно завершен")
}

// Тест endPoint /clients | get запрос.
// Все возможные тесты вызываются из этой функции
func StartTestGetClients() {
	apptype.Info.Println("Тесты с endPoint /clients | get успешно начаты")
	var err error
	con := new(testCon)
	con.DB, err = apptype.ConnectToDatabase()
	if err != nil {
		panic(err)
	}
	apptype.Debug.Println("Успешное подключение к базе данны (из тестов)")

	testNameSurPatrAgeRequest(con)
	testNameSurPartRequest(con)
	testNameSurRequest(con)
	testNameRequest(con)
	apptype.Info.Println("Тесты с endPoint /clients | get успешно завершены")
}
