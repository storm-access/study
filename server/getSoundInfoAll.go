package server //issue4

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetSoundInfoAll(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		type soundInfo struct { //格納する構造体
			SoundID   int    `json:"soundID"`
			SoundName string `json:"soundName"`
		}
		var responseData []soundInfo
		var jsonData soundInfo

		rows, err := db.Query("SELECT sound_ID, sound_name FROM sound") //問い合わせ
		if err != nil {
			fmt.Println("エラー") //エラー処理
		}
		defer rows.Close()

		for rows.Next() { //rowsの中のデータをresponseDataに追加していく
			err := rows.Scan(&jsonData.SoundID, &jsonData.SoundName)
			if err != nil {
				fmt.Println("エラー") //エラー処理
			}
			responseData = append(responseData, jsonData)
		}

		err = rows.Err()
		if err != nil {
			fmt.Println("エラー") //エラー処理
		}

		responseMessage, err := json.Marshal(responseData) //json形式にパース
		if err != nil {
			fmt.Println("エラー") //エラー処理
		}

		log.Println(string(responseMessage))
		w.Header().Set("Content-Type", "application/json;charset=utf-8") //ヘッダ，アプリ,utf-8を明示
		w.Write(responseMessage)                                         //レスポンスするJSON
	}
}
