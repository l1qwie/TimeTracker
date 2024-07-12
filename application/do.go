package application

import (
	"github.com/l1qwie/TimeTracker/apptype"
)

// Функция служит перенаправителе к разным функциям которые делают запросы в базу даннных
func PrepareQueryToDb(con *Conn, req *apptype.ReqToIfo) ([]*apptype.Client, error) {
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
