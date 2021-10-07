package main

import (
	"database/sql"
	"log"
	"net/http"
	"otokake/server"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//DBに接続
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/otokake")
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
	defer db.Close()

	//以下ハンドラ
	http.HandleFunc("/getStepSoundInfo", server.GetSoundInfoAll(db))
	http.ListenAndServe(":8080", nil)
}
