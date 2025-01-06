package main

import (
	"fmt"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env.local")

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello main!"))
	})

	http.HandleFunc("/bot", func (w http.ResponseWriter, r *http.Request) {
		fullPath := os.Getenv("TELEGRAM_API")

		resp, _ := http.Get(fmt.Sprintf("%s%s", fullPath, "/getMe"))

		w.Write([]byte(fmt.Sprintf("client: status code: %d\n", resp.StatusCode)))
	})

	http.ListenAndServe(":8080", nil)
}