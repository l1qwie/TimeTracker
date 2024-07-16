package application

import (
	"fmt"
	"strings"

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

// Проверяет существует ли клиент и его таска по переданному clientid и taskid
// и обновляет нужный столбец в бд в зависимости от запроса (либо начало, либо конец отсчета)
func TaskTime(con *Conn, timeManager *apptype.Time) (string, error) {
	var (
		err                  error
		clientidok, taskidok bool
		answer               string
	)
	apptype.Info.Println("Начало бизнес логики")
	clientidok, err = con.findClient(timeManager.ClientId)
	if clientidok {
		taskidok, err = con.findTask(timeManager.TaskId)
		if taskidok {
			if timeManager.StartTime {
				err = con.updateTaskStartTime(timeManager.ClientId, timeManager.TaskId)
				if err == nil {
					answer = "The countdown has begun"
				}
			} else {
				err = con.updateTaskEndTime(timeManager.ClientId, timeManager.TaskId)
				if err == nil {
					answer = "The countdown has ended"
				}
			}
		} else {
			apptype.Debug.Printf("Не смог найти таску по переданному taskid")
			if err == nil {
				err = fmt.Errorf("couldn't find taskid. Try to send a diffrent one")
			}
		}
	} else {
		apptype.Debug.Printf("Не смог найти клиента по переданному clientid")
		if err == nil {
			err = fmt.Errorf("couldn't find clientid. Try to send a diffrent one")
		}
	}
	apptype.Info.Println("Конец бизнес логики")
	return answer, err
}

// Проверяет сущестует такой клиент, который пришел и если да, то удаляет его
func PrepareQueryToDeleteClient(con *Conn, clientid int) (string, error) {
	var answer string
	apptype.Info.Println("Начало бизнес логики")
	ok, err := con.findClient(clientid)
	if ok {
		err = con.deleteClientDB(clientid)
		if err == nil {
			answer = fmt.Sprintf("The client %d has been successfuly deleted", clientid)
		}
	} else {
		apptype.Debug.Printf("Не смог найти клиента по переданному clientid")
		if err == nil {
			err = fmt.Errorf("couldn't find clientid. Try to send a diffrent one")
		}
	}
	apptype.Info.Println("Конец бизнес логики")
	return answer, err
}

// Осуществляет некоторые проверки перед началом изменений данны клиента, а именно:
// проверяет существует ли клиент по переданному clientid, а так же существует ли
// столбец в таблице Clients по переданному значению column
func ChangeClient(con *Conn, ch *apptype.Change) (string, error) {
	var (
		answer             string
		clientok, columnok bool
		err                error
	)
	apptype.Info.Println("Начало бизнес логики")
	clientok, err = con.findClient(ch.ClientId)
	if clientok {
		columnok, err = con.findColumn(ch.Column)
		if columnok {
			if ch.Column == "age" {
				err = con.updateClientColumnInt(ch.ClientId, ch.Column, ch.ValueInt)
				if err == nil {
					answer = fmt.Sprintf("The client's {%d} %s was changed to %d", ch.ClientId, ch.Column, ch.ValueInt)
				}
			} else {
				err = con.updateClientColumnStr(ch.ClientId, ch.Column, ch.ValueStr)
				if err == nil {
					answer = fmt.Sprintf("The client's {%d} %s was changed to %s", ch.ClientId, ch.Column, ch.ValueStr)
				}
			}
		} else {
			apptype.Debug.Printf("Не смог найти переданое название столбца в таблице бд")
			if err == nil {
				err = fmt.Errorf("couldn't find column_name. Try to send a diffrent one")
			}
		}
	} else {
		apptype.Debug.Printf("Не смог найти клиента по переданному clientid")
		if err == nil {
			err = fmt.Errorf("couldn't find clientid. Try to send a diffrent one")
		}
	}
	apptype.Info.Println("Конец бизнес логики")
	return answer, err
}

// Проверяет на коректность данные, которые были переданы и если все верно, создает нового клиента
func AddClient(con *Conn, newcl *apptype.NewClient) (*apptype.People, error) {
	var err error
	apptype.Info.Println("Начало бизнес логики")
	parts := strings.Split(newcl.Passport, " ")
	people := new(apptype.People)
	if len(parts) == 2 {
		people.Name = "Ivan"
		people.Surname = "Ivanov"
		people.Address = "Tel-Aviv Yafo"
		err = con.addClientDB(people, parts[0], parts[1])
	} else {
		err = fmt.Errorf(`couldn't divide the data of passport. Try to send a new one like this "1234 567890"`)
	}
	apptype.Info.Println("Конец бизнес логики")
	return people, err
}
