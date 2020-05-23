package lebot

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/prometheus/alertmanager/notify/webhook"
	"log"
)

type Lebot struct {
	token string
	chatId int64
	bot *tgbotapi.BotAPI
}

func NewLebot(token string, chatId int64) *Lebot {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)
	bot.Debug = true
	return &Lebot{
		token: token,
		chatId: chatId,
		bot: bot,
	}
}

func (t *Lebot) Say(message string) error {
	m := tgbotapi.NewMessage(t.chatId, message)
	m.ParseMode = "Markdown"
	_, err := t.bot.Send(m)
	return err
}

func getAlertIcon(status string) string {
	if status == "firing" {
		return "\xF0\x9F\x94\xA5"
	} else {
		return "\xF0\x9F\x92\x9A"
	}
}

func (t *Lebot) Alert(message webhook.Message) error  {
	log.Printf("Send message %v", message)
	for _, alert := range message.Alerts {
		log.Printf("Alert: status=%s,Labels=%v,Annotations=%v", alert.Status, alert.Labels, alert.Annotations)

		message := fmt.Sprintf("%s *%s* \nStatus: _%s_, Severity: _%s_ \n%s",
			getAlertIcon(alert.Status),
			alert.Annotations["summary"],
			alert.Status,
			alert.Labels["severity"],
			alert.Annotations["description"],
		)
		t.Say(message)
	}
	return nil
}