package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/prometheus/alertmanager/notify/webhook"
	"gopkg.in/alecthomas/kingpin.v2"
	lebot "lebotic/pkg"
	"log"
	"net/http"
	"os"
	"time"
)

func initHandler(token string, chatId int64) func(w http.ResponseWriter, r *http.Request) {

	log.Printf("Authorized %s %d", token, chatId)

	var bot = lebot.NewLebot(token, chatId)
	bot.Say("I have been restarted!")

	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Got alert!")

		if r.Method != http.MethodPost {
			log.Printf("Method is not allowed")
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		if r.Body == nil {
			log.Printf("Body is empty")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var alert webhook.Message

		err := json.NewDecoder(r.Body).Decode(&alert)
		log.Printf(" %v", r.Body)
		if err != nil {
			log.Printf("Bad json")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = bot.Alert(alert)
		if err != nil {
			log.Printf("Can not alert")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func main() {
	godotenv.Load()

	config := struct {
		listenAddr     string
		telegramToken  string
		telegramChatId int64
	}{}

	log.Printf("LEBOTIC_TELEGRAM_TOKEN %s", os.Getenv("LEBOTIC_TELEGRAM_TOKEN"))
	log.Printf("LEBOTIC_TELEGRAM_CHAT_ID %s", os.Getenv("LEBOTIC_TELEGRAM_CHAT_ID"))

	a := kingpin.New("lebotic", "Simplest bot for Prometheus' Alertmanager")

	a.Flag("listen.addr", "The address the lebotic listens on for incoming webhooks").
		Envar("LEBOTIC_LISTEN_ADDR").
		Default("0.0.0.0:8619").
		StringVar(&config.listenAddr)

	a.Flag("telegram.token", "Telegram token").
		Envar("LEBOTIC_TELEGRAM_TOKEN").
		Required().
		StringVar(&config.telegramToken)

	a.Flag("telegram.chat.id", "The chat id which will receive alerts").
		Envar("LEBOTIC_TELEGRAM_CHAT_ID").
		Required().
		Int64Var(&config.telegramChatId)

	_, err := a.Parse(os.Args[1:])
	if err != nil {
		fmt.Printf("error parsing commandline arguments: %v\n", err)
		a.Usage(os.Args[1:])
		os.Exit(2)
	}

	mtx := mux.NewRouter()
	alertHandler := initHandler(config.telegramToken, config.telegramChatId)
	mtx.HandleFunc("/alert", alertHandler)
	srv := &http.Server{
		Handler:      mtx,
		Addr:         config.listenAddr,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}