package apptype

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var Debug = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
var Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

const (
	docHost     = "postgresql"
	docPort     = "5432"
	docUsername = "postgres"
	docPass     = "postgres"
	docDbname   = "postgres"
	docSslmode  = "disable"
)

type ReqToIfo struct {
	Name       string `form:"name"`
	Surname    string `form:"surname"`
	Patronymic string `form:"patronymic"`
	Age        int    `form:"age"`
}

type Client struct {
	ID             int    `json:"id"`
	PassportSeries string `json:"passport_series"`
	PassportNumber string `json:"passport_number"`
	Surname        string `json:"surname"`
	Name           string `json:"name"`
	Patronymic     string `json:"patronymic"`
	Age            int    `json:"age"`
	Address        string `json:"address"`
}

type Task struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	TimeSpent string `json:"time_spent"`
}

type Time struct {
	ClientId  int  `json:"clientid"`
	TaskId    int  `json:"taskid"`
	StartTime bool `json:"start_time"`
}

type Change struct {
	ClientId int    `json:"clientid"`
	Column   string `json:"column"`
	ValueStr string `json:"value_str"`
	ValueInt int    `json:"int"`
}

func docConnect() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		docHost,
		docPort,
		docUsername,
		docPass,
		docDbname,
		docSslmode)
}

func ConnectToDatabase() (*sql.DB, error) {
	db, err := sql.Open("postgres", docConnect())
	if err != nil {
		log.Print(err)
	}
	err = db.Ping()
	return db, err
}
