package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//var db *sql.DB
//var err error
var tpl *template.Template

type RegisterData struct {
	Id   int
	Name string `json:"name,omitempty"`
}
type Response struct {
	Status  string
	Code    int
	Message interface{}
}

var db, err = sql.Open("mysql", "root:root@/test")
var timezone, _ = time.LoadLocation("America/New_York")
var zone, _ = time.Now().In(timezone).Zone()

func getDBConnectivity(db *sql.DB) (bool, error) {
	err = db.Ping()
	if err != nil {
		log.Println("db connection: ", err.Error())
		return false, err
	}
	return true, nil
}

func main() {
	var js RegisterData
	//js.Id = 0
	js.Name = "jdhhsdhjsd"

	resp, err := json.Marshal(js)
	if err != nil {
		fmt.Println("error:", err)
	}
	err = Request_insert_json(string(resp))
	if err != nil {
		fmt.Println("Error while connecting to server:", err)
	}

}

func Request_insert_json(status string) error {
	//data:=status+"|"+strconv.Itoa(Jobid)
	url := fmt.Sprintf("http://localhost:3033/index?json=%s", url.QueryEscape(status))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("NewRequest: ", err.Error)
		return err
	}

	//fmt.Println(url)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Do: ", err)
		return err
	}

	//fmt.Println(string(resp.Body))
	defer resp.Body.Close()
	//fmt.Println("**************************")
	return nil
}
