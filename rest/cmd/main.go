package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/soqet/configjson"
	"net/http"
	"rest/internal/api"
	database "rest/internal/db"
)

type Config struct {
	ApiUrl string `json:"apiUrl"`
	Port   int    `json:"port"`
}

func main() {
	config := new(Config)
	configjson.ReadConfigFile("./config.json", config)
	db := new(database.DB)
	err := db.Init("./files.db")
	if err != nil {
		panic(err)
	}
	router := mux.NewRouter()
	// s := router.Host(config.ApiUrl).Subrouter()
	api.Init(router, db)
	http.Handle("/", router)
	fmt.Println("Server is listening")
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}
