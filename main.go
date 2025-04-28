package main

import (
	"fmt"
	"go_final_project/pkg/api"
	"go_final_project/pkg/db"
	"log"
	"net/http"
)

func main() {

	// подключаем БД
	if err := db.Init("pkg/db/scheduler.db"); err != nil {
		fmt.Println("Ошибка в подключении БД")
		return
	}

	//Вычисляем следующую дату повторения
	api.Init()

	//Запускаем веб-сервер
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	log.Println("Сервер запущен на :7540")
	log.Fatal(http.ListenAndServe(":7540", nil))

}
