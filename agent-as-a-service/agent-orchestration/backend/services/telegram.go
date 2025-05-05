package services

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	tl "github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/telegram"
	"github.com/mymmrac/telego"
)

var teleVideoActivitiesAlert *telego.Bot
var teleVideoActivitiesChatId int64

var teleMagicVideoActivitiesAlert *telego.Bot
var teleMagicVideoActivitiesChatId int64

func InitTeleMagicVideoActivitiesAlert(token, chatId string) {
	if len(token) == 0 || len(chatId) == 0 {
		return
	}
	teleMagicVideoActivitiesChatId, _ = strconv.ParseInt(chatId, 10, 64)
	teleMagicVideoActivitiesAlert, _ = telego.NewBot(token, telego.WithDefaultDebugLogger())
}
func (s *Service) SendTeleMagicVideoActivitiesAlert(msg string) {
	if teleMagicVideoActivitiesAlert != nil {
		go func() {
			teleMagicVideoActivitiesAlert.SendMessage(
				&telego.SendMessageParams{
					ChatID: telego.ChatID{
						ID: teleMagicVideoActivitiesChatId,
					},
					Text: strings.TrimSpace(msg),
				},
			)
		}()
	}
}

func InitTeleVideoActivitiesAlert(token, chatId string) {
	if len(token) == 0 || len(chatId) == 0 {
		return
	}
	teleVideoActivitiesChatId, _ = strconv.ParseInt(chatId, 10, 64)
	teleVideoActivitiesAlert, _ = telego.NewBot(token, telego.WithDefaultDebugLogger())
}

func (s *Service) SendTeleVideoActivitiesAlert(msg string, chatId ...string) {
	if teleVideoActivitiesAlert != nil {
		go func() {
			chatID := teleVideoActivitiesChatId
			var err error
			if len(chatId) > 0 {
				chatID, err = strconv.ParseInt(chatId[0], 10, 64)
				if err != nil {
					return
				}
			}
			teleVideoActivitiesAlert.SendMessage(
				&telego.SendMessageParams{
					ChatID: telego.ChatID{
						ID: chatID,
					},
					Text: strings.TrimSpace(msg),
				},
			)
		}()
	}
}

func (s *Service) GetTelegramMessage(ctx context.Context) {
	client := tl.NewClient(20670342, "94dc0a1407d2c786a802de64b048707a")

	// Create a context that will be canceled on SIGINT
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle SIGINT
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		cancel()
	}()

	// Run the client
	if err := client.Run(ctx); err != nil {
		fmt.Printf("Failed to run client: %v", err)
	}

	// Get messages from the channel
	if err := client.GetMessages(ctx, "https://t.me/fttrenches_sol", 5); err != nil {
		fmt.Printf("Failed to get messages: %v", err)
	}
}
