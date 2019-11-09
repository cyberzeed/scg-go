package line

import (
	"log"
	"net/http"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

// Bot struct
type Bot struct {
	client *linebot.Client
}

// NewBot is a factory function which create new Bot object
func NewBot(accessToken string, secret string) *Bot {
	client, err := linebot.New(secret, accessToken)
	if err != nil {
		log.Fatalf("Cannot connect LINE Messaging API, %v", err)
	}
	return &Bot{client}
}

// BroadcastMessage send messages to friend
func (bot *Bot) BroadcastMessage(messages ...string) (*linebot.BasicResponse, error) {
	sendingMessages := createLineSendingMessages(messages...)
	return bot.client.BroadcastMessage(sendingMessages...).Do()
}

// ReplyTextMessage will reply text message to friend
func (bot *Bot) ReplyTextMessage(
	event *linebot.Event,
	messages ...string,
) (*linebot.BasicResponse, error) {
	sendingMessages := createLineSendingMessages(messages...)
	return bot.client.ReplyMessage(event.ReplyToken, sendingMessages...).Do()
}

// ReplyStickerMessage will reply sticker message to friend
func (bot *Bot) ReplyStickerMessage(
	event *linebot.Event,
	packageID string,
	stickerID string,
) (*linebot.BasicResponse, error) {
	sendingMessage := linebot.NewStickerMessage(packageID, stickerID)
	return bot.client.ReplyMessage(event.ReplyToken, sendingMessage).Do()
}

// ReplyAudioMessage will reply audio message to friend
func (bot *Bot) ReplyAudioMessage(
	event *linebot.Event,
	contentURL string,
	duration int,
) (*linebot.BasicResponse, error) {
	sendingMessage := linebot.NewAudioMessage(contentURL, duration)
	return bot.client.ReplyMessage(event.ReplyToken, sendingMessage).Do()
}

// ReplyImageMessage will reply image message to friend
func (bot *Bot) ReplyImageMessage(
	event *linebot.Event,
	contentURL string,
	previewImageURL string,
) (*linebot.BasicResponse, error) {
	sendingMessage := linebot.NewImageMessage(contentURL, previewImageURL)
	return bot.client.ReplyMessage(event.ReplyToken, sendingMessage).Do()
}

// SendGreetingMessage will reply greeting message when LINE official account is added as a friend
func (bot *Bot) SendGreetingMessage(event *linebot.Event) (*linebot.BasicResponse, error) {
	greeting := []string{
		greetingMessage,
		strings.Join(usageMessage, "\n"),
	}
	return bot.ReplyTextMessage(event, greeting...)
}

// ParseRequest is function which handle webhook request
func (bot *Bot) ParseRequest(req *http.Request) error {
	// extract events from request
	events, err := bot.client.ParseRequest(req)
	if err != nil {
		return err
	}

	// execute event handlers
	for _, event := range events {
		if err = bot.routeEventHandler(event); err != nil {
			return err
		}
	}

	return err
}

func (bot *Bot) routeEventHandler(event *linebot.Event) error {
	var err error
	log.Printf("EventType: %v", event.Type)

	// select bot function for event.
	switch event.Type {
	case linebot.EventTypeMessage:
		err = bot.routeMessageType(event)
	case linebot.EventTypeFollow:
		_, err = bot.SendGreetingMessage(event)
	}

	return err
}

func (bot *Bot) routeMessageType(event *linebot.Event) error {
	var err error

	// select case that handle message event
	switch msg := event.Message.(type) {
	case *linebot.TextMessage:
		// check execute command in message before send message
		messages, err := executeCommand(msg.Text)
		if err != nil {
			break
		}
		_, err = bot.ReplyTextMessage(event, messages...)

	case *linebot.StickerMessage:
		_, err = bot.ReplyStickerMessage(event, msg.PackageID, msg.StickerID)

	case *linebot.AudioMessage:
		_, err = bot.ReplyAudioMessage(event, msg.OriginalContentURL, msg.Duration)

	case *linebot.ImageMessage:
		_, err = bot.ReplyImageMessage(event, msg.OriginalContentURL, msg.PreviewImageURL)
	}

	return err
}

func createLineSendingMessages(messages ...string) []linebot.SendingMessage {
	var sendingMessages []linebot.SendingMessage
	for _, msg := range messages {
		message := linebot.NewTextMessage(msg)
		sendingMessages = append(sendingMessages, message)
	}

	return sendingMessages
}
