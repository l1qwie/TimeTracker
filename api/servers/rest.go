package servers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/l1qwie/TimeTracker/application"
	"github.com/l1qwie/TimeTracker/apptype"
)

// @Summary Информация о пользователях
// @Description Возвращает данные а найденных пользователях
// @Accept  json
// @Produce  json
// @Param   name     	query    	string     	false  "Name"
// @Param   surname  	query    	string     	false  "Surname"
// @Param   age      	query    	string     	false  "Age"
// @Param   page     	query    	int        	false  "Page"
// @Param   pagesize    query   	int        	false  "PageSize"
// @Success 200 {object} []Client
// @Failure 400 {object} string
// @Router /users [get]
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
			response = gin.H{"error": err.Error()}

		} else {
			apptype.Debug.Println("Данные успешно вынуты из запроса")
			apptype.Info.Printf("Данные из запроса: %v", req)
			con := new(application.Conn)
			con.DB, err = apptype.ConnectToDatabase()

			if err != nil {
				apptype.Debug.Printf("Не удалось подключиться к базе данных: %s", err)
				statreq = http.StatusBadRequest
				response = gin.H{"error": err.Error()}

			} else {
				apptype.Debug.Println("Успешное подключение к базе данных")
				clients, err := application.PrepareQueryToDb(con, req)

				if err != nil {
					apptype.Debug.Printf("Ошибка на пути бизнес логики: %s", err)
					statreq = http.StatusBadRequest
					response = gin.H{"error": err.Error()}

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
