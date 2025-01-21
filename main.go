package main

import (
	"log"
	"lol-statistics/api/telegram"
	"net/http"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env.local")

	if err != nil {
		log.Fatal("Failed to load env files")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello main!"))
	})

	mux.HandleFunc("/bot/info", telegram.GetBasicBotInfo)
	mux.HandleFunc("/bot/send-info", telegram.GetBasicBotInfo)

	mux.HandleFunc("/lol/stats", telegram.GetBasicBotInfo)
	mux.HandleFunc("/lol/ranking", telegram.GetBasicBotInfo)

	http.ListenAndServe(":8080", mux)
}