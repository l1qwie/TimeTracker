package changeclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/l1qwie/TimeTracker/apptype"
)

// Отсылает запрос на сервер, принимает ответ и расшифровывает его
func changeClient(body []byte) string {
	var answer string
	req, err := http.NewRequest(http.MethodPut, "http://localhost:8039/client/change", bytes.NewBuffer(body))
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

// Запускает тест изменения имени клиента
func changeName(con *testCon) {
	apptype.Info.Println("Тест изменения имени клиента успешно начат")
	con.createClient()
	defer con.resetSeq()
	defer con.deleteClients()
	request := &apptype.Change{
		ClientId: 1,
		Column:   "name",
		ValueStr: "Natan",
	}
	jsonbody, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}
	answer := changeClient(jsonbody)
	if answer != fmt.Sprintf("The client's {%d} name was changed to Natan", 1) {
		panic(fmt.Sprintf(`Expected: "The client's {%d} name was changed to Natan". Recieved: %s`, 1, answer))
	}
	if !con.checkChangedColumn("name", 1, "Natan") {
		panic("The client's name in database hasn't been changed yet")
	}
	apptype.Info.Println("Тест изменения имени клиента успешно закончен")
}

// Запускает тест изменения фамилии клиента
func changeSurname(con *testCon) {
	apptype.Info.Println("Тест изменения фамилии клиента успешно начат")
	con.createClient()
	defer con.resetSeq()
	defer con.deleteClients()
	request := &apptype.Change{
		ClientId: 1,
		Column:   "surname",
		ValueStr: "Sahar",
	}
	jsonbody, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}
	answer := changeClient(jsonbody)
	if answer != fmt.Sprintf("The client's {%d} surname was changed to Sahar", 1) {
		panic(fmt.Sprintf(`Expected: "The client's {%d} surname was changed to Sahar". Recieved: %s`, 1, answer))
	}
	if !con.checkChangedColumn("surname", 1, "Sahar") {
		panic("The client's surname in database hasn't been changed yet")
	}
	apptype.Info.Println("Тест изменения фамилии клиента успешно закончен")
}

// Запускает тест изменения отчества клиента
func changePatronymic(con *testCon) {
	apptype.Info.Println("Тест изменения отчества клиента успешно начат")
	con.createClient()
	defer con.resetSeq()
	defer con.deleteClients()
	request := &apptype.Change{
		ClientId: 1,
		Column:   "patronymic",
		ValueStr: "Antonovich",
	}
	jsonbody, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}
	answer := changeClient(jsonbody)
	if answer != fmt.Sprintf("The client's {%d} patronymic was changed to Antonovich", 1) {
		panic(fmt.Sprintf(`Expected: "The client's {%d} patronymic was changed to Antonovich". Recieved: %s`, 1, answer))
	}
	if !con.checkChangedColumn("patronymic", 1, "Antonovich") {
		panic("The client's patronymic in database hasn't been changed yet")
	}
	apptype.Info.Println("Тест изменения отчества клиента успешно закончен")
}

// Запускает тест изменения возраста клиента
func changeAge(con *testCon) {
	apptype.Info.Println("Тест изменения возраста клиента успешно начат")
	con.createClient()
	defer con.resetSeq()
	defer con.deleteClients()
	request := &apptype.Change{
		ClientId: 1,
		Column:   "age",
		ValueInt: 55,
	}
	jsonbody, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}
	answer := changeClient(jsonbody)
	if answer != fmt.Sprintf("The client's {%d} age was changed to 55", 1) {
		panic(fmt.Sprintf(`Expected: "The client's {%d} age was changed to 55". Recieved: %s`, 1, answer))
	}
	if !con.checkChangedColumn("age", 1, 55) {
		panic("The client's age in database hasn't been changed yet")
	}
	apptype.Info.Println("Тест изменения возраста клиента успешно закончен")
}

// Запускает тест изменения серии паспорта клиента
func changePassportSeries(con *testCon) {
	apptype.Info.Println("Тест изменения серии паспорта клиента успешно начат")
	con.createClient()
	defer con.resetSeq()
	defer con.deleteClients()
	request := &apptype.Change{
		ClientId: 1,
		Column:   "passportseries",
		ValueStr: "2222",
	}
	jsonbody, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}
	answer := changeClient(jsonbody)
	if answer != fmt.Sprintf("The client's {%d} passportseries was changed to 2222", 1) {
		panic(fmt.Sprintf(`Expected: "The client's {%d} passportseries was changed to 2222". Recieved: %s`, 1, answer))
	}
	if !con.checkChangedColumn("passportseries", 1, "2222") {
		panic("The client's passport series in database hasn't been changed yet")
	}
	apptype.Info.Println("Тест изменения серии паспорта клиента успешно закончен")
}

// Запускает тест изменения номера паспорта клиента
func changePassporNumber(con *testCon) {
	apptype.Info.Println("Тест изменения номера паспорта клиента успешно начат")
	con.createClient()
	defer con.resetSeq()
	defer con.deleteClients()
	request := &apptype.Change{
		ClientId: 1,
		Column:   "passportnumber",
		ValueStr: "1235367",
	}
	jsonbody, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}
	answer := changeClient(jsonbody)
	if answer != fmt.Sprintf("The client's {%d} passportnumber was changed to 1235367", 1) {
		panic(fmt.Sprintf(`Expected: "The client's {%d} passportnumber was changed to 1235367". Recieved: %s`, 1, answer))
	}
	if !con.checkChangedColumn("passportnumber", 1, "1235367") {
		panic("The client's passoprt number in database hasn't been changed yet")
	}
	apptype.Info.Println("Тест изменения номера паспорта клиента успешно закончен")
}

// Запускает тест изменения адреса клиента
func changeAddress(con *testCon) {
	apptype.Info.Println("Тест изменения адреса клиента успешно начат")
	con.createClient()
	defer con.resetSeq()
	defer con.deleteClients()
	request := &apptype.Change{
		ClientId: 1,
		Column:   "address",
		ValueStr: "Am Yisrael Chai",
	}
	jsonbody, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}
	answer := changeClient(jsonbody)
	if answer != fmt.Sprintf("The client's {%d} address was changed to Am Yisrael Chai", 1) {
		panic(fmt.Sprintf(`Expected: "The client's {%d} address was changed to Am Yisrael Chai". Recieved: %s`, 1, answer))
	}
	if !con.checkChangedColumn("address", 1, "Am Yisrael Chai") {
		panic("The client's address in database hasn't been changed yet")
	}
	apptype.Info.Println("Тест изменения адреса клиента успешно закончен")
}

// Запускает множество тестов для изменения данных клиента
func StartTestChangeClient() {
	apptype.Info.Println("Тесты с endPoint /client/change | put успешно начаты")
	var err error
	con := new(testCon)
	con.DB, err = apptype.ConnectToDatabase()
	if err != nil {
		panic(err)
	}
	apptype.Debug.Println("Успешное подключение к базе данны (из тестов)")

	changeName(con)
	changeSurname(con)
	changePatronymic(con)
	changeAge(con)
	changePassportSeries(con)
	changePassporNumber(con)
	changeAddress(con)

	apptype.Info.Println("Тесты с endPoint /client/change | put успешно закончены")
}
