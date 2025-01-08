package telegram

import (
    "fmt"
	"io"
	"log"
	"net/http"
    "os"
)

type TelegramHandler struct {}

func GetBasicBotInfo(w http.ResponseWriter, r *http.Request) {
	telegramApiPath := os.Getenv("TELEGRAM_API")
	telegramGetMeEndpoint := os.Getenv("TELEGRAM_GET_ME_ENDPOINT")

	response, err := http.Get(fmt.Sprintf("%s%s", telegramApiPath, telegramGetMeEndpoint))
	if err != nil {
		log.Fatalf("Failed to make HTTP request: %v", err)
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal("Failed to read response body: %v: ", err)
	}

	w.Write(body)
}