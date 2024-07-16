package servers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/l1qwie/TimeTracker/application"
	"github.com/l1qwie/TimeTracker/apptype"
)

type Error struct {
	Err string `json:"error"`
}

// @Summary Информация о клиентахх
// @Description Возвращает данные о найденных клиентах
// @Accept  json
// @Produce json
// @Param   name     	query    	string     	false  	"Name"
// @Param   surname  	query    	string     	false  	"Surname"
// @Param 	patronymic	query		string		false 	"Patronymic"
// @Param   age      	query    	integer     false  	"Age"
// @Success 200 {object} []*apptype.Client
// @Failure 400 {object} *Error
// @Router /client [get]
func GetClientsInfo() {
	router := gin.Default()
	router.GET("/client", func(c *gin.Context) {
		apptype.Info.Println("Сервер /client | get запрос - запущен")
		var (
			statreq  int
			response any
		)

		req := new(apptype.ReqToIfo)
		if err := c.ShouldBindQuery(req); err != nil {
			apptype.Debug.Printf("Ошибка при попытке вынуть данные из запроса: %s", err)
			statreq = http.StatusBadRequest
			response = &Error{Err: err.Error()}

		} else {
			apptype.Debug.Println("Данные успешно получены из запроса")
			apptype.Info.Printf("Данные из запроса: %v", req)
			con := new(application.Conn)
			con.DB, err = apptype.ConnectToDatabase()

			if err != nil {
				apptype.Debug.Printf("Не удалось подключиться к базе данных: %s", err)
				statreq = http.StatusBadRequest
				response = &Error{Err: err.Error()}

			} else {
				apptype.Debug.Println("Успешное подключение к базе данных")
				clients, err := application.PrepareQueryToDbGetInfo(con, req)

				if err != nil {
					apptype.Debug.Printf("Ошибка на пути бизнес логики: %s", err)
					statreq = http.StatusBadRequest
					response = &Error{Err: err.Error()}

				} else {
					apptype.Debug.Println("Бизнес логика завершена успешно")
					statreq = http.StatusOK
					response = clients
				}
			}
		}
		apptype.Debug.Printf("Данные для ответа: %v", response)
		apptype.Info.Println("Отправлен ответ из сервера /client | get запрос")
		c.JSON(statreq, response)
	})
	router.Run(":8088")
}

// @Summary Трудозатраты клиентов
// @Description Возвращает трудозатраты по пользователю за период
// @Accept  json
// @Produce json
// @Param  id   path     string     true     "ID"
// @Success 200 {object} []*apptype.Tasks
// @Failure 400 {object} *Error
// @Router /client/{id}/time-logs [get]
func GetTimeLogs() {
	router := gin.Default()
	router.GET("/client/:id/time-logs", func(c *gin.Context) {
		apptype.Info.Println("Сервер /client/:id/time-logs | get запрос - запущен")
		var (
			statreq  int
			response any
			err      error
		)
		id := c.Param("id")
		if id == "" {
			statreq = http.StatusBadRequest
			response = &Error{Err: "there's no id in your request"}
		} else {
			apptype.Info.Printf("Полученный Id: %s", id)
			apptype.Debug.Println("Данные успешно получены из запроса")
			con := new(application.Conn)
			con.DB, err = apptype.ConnectToDatabase()

			if err != nil {
				apptype.Debug.Printf("Не удалось подключиться к базе данных: %s", err)
				statreq = http.StatusBadRequest
				response = &Error{Err: err.Error()}

			} else {
				apptype.Debug.Println("Успешное подключение к базе данных")
				clientid, err := strconv.Atoi(id)

				if err != nil {
					apptype.Debug.Println("Не удалось форматировать string к int")
					statreq = http.StatusBadRequest
					response = &Error{Err: err.Error()}

				} else {
					tasks, err := application.PrepareQueryToDbGetTimeLogs(con, clientid)

					if err != nil {
						apptype.Debug.Println("Произошла ошибка в бизнес логике")
						statreq = http.StatusBadRequest
						response = &Error{Err: err.Error()}

					} else {
						apptype.Debug.Println("Бизнес логика закончила свою работу без ошибок")
						statreq = http.StatusOK
						response = tasks
					}
				}
			}
		}
		apptype.Debug.Printf("Данные для ответа: %v", response)
		apptype.Info.Println("Отправлен ответ из сервера /client/:id/time-logs | get запрос")
		c.JSON(statreq, response)
	})
	router.Run(":8099")
}

// @Summary Начало отсчета времени
// @Description Начало отсчета времени для задачи клиента
// @Accept  json
// @Produce json
// @Param  clientid   query     string     true     "ClientId"
// @Param  taskid 	  query	    string	   true	    "TaskId"
// @Success 200 {object} string
// @Failure 400 {object} *Error
// @Router /client/tasks/start [post]
func StartTimeManager() {
	router := gin.Default()
	router.POST("/client/tasks/timeManager", func(c *gin.Context) {
		apptype.Info.Println("Сервер /client/tasks/timeManager | post запрос - запущен")
		var (
			statreq  int
			response any
			err      error
		)
		timemanager := new(apptype.Time)
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			statreq = http.StatusBadRequest
			response = &Error{Err: err.Error()}
		} else {
			apptype.Debug.Println("Данные успешно получены из запроса")
			err = json.Unmarshal(body, &timemanager)
			if err != nil {
				apptype.Debug.Println("Не удалось успешно расшифровать данные из запроса")
				statreq = http.StatusBadRequest
				response = &Error{Err: err.Error()}
			} else {
				apptype.Debug.Println("Данные успешно расшифрованны из запроса")
				con := new(application.Conn)
				con.DB, err = apptype.ConnectToDatabase()

				if err != nil {
					apptype.Debug.Printf("Не удалось подключиться к базе данных: %s", err)
					statreq = http.StatusBadRequest
					response = &Error{Err: err.Error()}

				} else {
					apptype.Debug.Println("Успешное подключение к базе данных")
					completed, err := application.TaskTime(con, timemanager)

					if err != nil {
						apptype.Debug.Println("Произошла ошибка в бизнес логике")
						statreq = http.StatusBadRequest
						response = &Error{Err: err.Error()}

					} else {
						apptype.Debug.Println("Бизнес логика закончила свою работу без ошибок")
						statreq = http.StatusOK
						response = completed
					}
				}
			}
		}
		apptype.Debug.Printf("Данные для ответа: %v", response)
		apptype.Info.Println("Отправлен ответ из сервера /client/tasks/timeManager | post запрос")
		c.JSON(statreq, response)
	})
	router.Run(":8079")
}

// @Summary Удаление клиента из базы данных
// @Description Полное удаление задач клиента и самого клиента
// @Accept  json
// @Produce json
// @Param  clientid   path     string     true     "Id"
// @Success 200 {object} string
// @Failure 400 {object} *Error
// @Router /client/{id}/delete [delete]
func DeleteClient() {
	router := gin.Default()
	router.DELETE("/client/:id/delete", func(c *gin.Context) {
		apptype.Info.Println("Сервер /client/:id/delete | delete запрос - запущен")
		var (
			statreq  int
			response any
			err      error
		)
		id := c.Param("id")
		if id == "" {
			statreq = http.StatusBadRequest
			response = &Error{Err: "there's no id in your request"}

		} else {
			apptype.Info.Printf("Полученный Id: %s", id)
			apptype.Debug.Println("Данные успешно получены из запроса")
			con := new(application.Conn)
			con.DB, err = apptype.ConnectToDatabase()
			if err != nil {
				apptype.Debug.Printf("Не удалось подключиться к базе данных: %s", err)
				statreq = http.StatusBadRequest
				response = &Error{Err: err.Error()}

			} else {
				apptype.Debug.Println("Успешное подключение к базе данных")
				clientid, err := strconv.Atoi(id)

				if err != nil {
					apptype.Debug.Println("Не удалось форматировать string к int")
					statreq = http.StatusBadRequest
					response = &Error{Err: err.Error()}

				} else {
					answer, err := application.PrepareQueryToDeleteClient(con, clientid)

					if err != nil {
						apptype.Debug.Println("Ошибка во время выполнения бизнес логики")
						statreq = http.StatusBadRequest
						response = &Error{Err: err.Error()}

					} else {
						statreq = http.StatusOK
						response = answer
					}
				}
			}
		}
		apptype.Debug.Printf("Данные для ответа: %v", response)
		apptype.Info.Println("Отправлен ответ из сервера /client/:id/delete | delete запрос")
		c.JSON(statreq, response)
	})
	router.Run(":8059")
}

// @Summary Изменение данныз клиента
// @Description Полное удаление задач клиента и самого клиента
// @Accept  json
// @Produce json
// @Param  clientid   query     int        true     "ClientId"
// @Param  column 	  query	    string	   true	    "Column"
// @Param  valueint   query 	int 	   true 	"ValueInt"
// @Param  valuestr   query 	string 	   true 	"ValueStr"
// @Success 200 {object} string
// @Failure 400 {object} *Error
// @Router /client/change [put]
func ChangeClient() {
	router := gin.Default()
	router.PUT("/client/change", func(c *gin.Context) {
		apptype.Info.Println("Сервер /client/change | put запрос - запущен")
		var (
			statreq  int
			response any
			err      error
		)
		changes := new(apptype.Change)
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			statreq = http.StatusBadRequest
			response = &Error{Err: err.Error()}

		} else {
			apptype.Debug.Println("Данные успешно получены из запроса")
			err = json.Unmarshal(body, &changes)

			if err != nil {
				apptype.Debug.Println("Не удалось успешно расшифровать данные из запроса")
				statreq = http.StatusBadRequest
				response = &Error{Err: err.Error()}
			} else {
				apptype.Debug.Println("Данные успешно расшифрованны из запроса")
				con := new(application.Conn)
				con.DB, err = apptype.ConnectToDatabase()

				if err != nil {
					apptype.Debug.Printf("Не удалось подключиться к базе данных: %s", err)
					statreq = http.StatusBadRequest
					response = &Error{Err: err.Error()}

				} else {
					apptype.Debug.Println("Успешное подключение к базе данных")
					completed, err := application.ChangeClient(con, changes)

					if err != nil {
						apptype.Debug.Println("Произошла ошибка в бизнес логике")
						statreq = http.StatusBadRequest
						response = &Error{Err: err.Error()}

					} else {
						apptype.Debug.Println("Бизнес логика закончила свою работу без ошибок")
						statreq = http.StatusOK
						response = completed
					}
				}

			}
		}
		apptype.Debug.Printf("Данные для ответа: %v", response)
		apptype.Info.Println("Отправлен ответ из сервера /client/change | put запрос")
		c.JSON(statreq, response)
	})
	router.Run(":8039")
}
