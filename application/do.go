package application

import (
	"fmt"

	"github.com/l1qwie/TimeTracker/apptype"
)

// Функция служит перенаправителе к разным функциям которые делают запросы в базу даннных
func PrepareQueryToDbGetInfo(con *Conn, req *apptype.ReqToIfo) ([]*apptype.Client, error) {
	var (
		clients []*apptype.Client
		err     error
	)
	apptype.Info.Println("Начало бизнес логики")
	if req.Name != "" && req.Surname != "" && req.Patronymic != "" && req.Age != 0 {
		apptype.Debug.Println("Выбран путь - Запрос в бд с именем, фамилией, отчеством и возрастом")
		clients, err = con.selectNameSurPatAge(req.Name, req.Surname, req.Patronymic, req.Age)

	} else if req.Name != "" && req.Surname != "" && req.Patronymic != "" && req.Age == 0 {
		apptype.Debug.Println("Выбран путь - Запрос в бд с именем, фамилией и отчеством")
		clients, err = con.selectNameSurPat(req.Name, req.Surname, req.Patronymic)

	} else if req.Name != "" && req.Surname != "" && req.Patronymic == "" && req.Age == 0 {
		apptype.Debug.Println("Выбран путь - Запрос в бд с именем и фамилией")
		clients, err = con.selectNameSur(req.Name, req.Surname)

	} else if req.Name != "" && req.Surname == "" && req.Patronymic == "" && req.Age == 0 {
		apptype.Debug.Println("Выбран путь - Запрос в бд с именем")
		clients, err = con.selectName(req.Name)

	}
	apptype.Info.Println("Конец бизнес логики")
	return clients, err
}

// Проверяет существует ли такой id который передали и перенаправляетт в функцию, которая делаю запрос в бд
func PrepareQueryToDbGetTimeLogs(con *Conn, id int) ([]*apptype.Task, error) {
	var (
		tasks []*apptype.Task
		err   error
	)
	apptype.Info.Println("Начало бизнес логики")
	ok, err := con.findClient(id)
	if ok {
		tasks, err = con.selectClientTasks(id)
	} else {
		apptype.Debug.Printf("Не смог найти клиента по переданному id")
		if err == nil {
			err = fmt.Errorf("couldn't find clientid. Try to send a diffrent one")
		}
	}
	apptype.Info.Println("Конец бизнес логики")
	return tasks, err
}
