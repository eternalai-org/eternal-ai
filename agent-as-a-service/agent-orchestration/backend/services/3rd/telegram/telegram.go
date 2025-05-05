package telegram

import (
	"context"
	"fmt"
	"log"

	"github.com/gotd/td/telegram"
	"github.com/gotd/td/tg"
)

type Client struct {
	apiID   int
	apiHash string
	client  *telegram.Client
}

func NewClient(apiID int, apiHash string) *Client {
	return &Client{
		apiID:   apiID,
		apiHash: apiHash,
		client:  telegram.NewClient(apiID, apiHash, telegram.Options{}),
	}
}

func (c *Client) GetMessages(ctx context.Context, channelURL string, limit int) error {
	// Create a new API client
	api := c.client.API()

	peer := &tg.InputPeerChannel{
		ChannelID:  -2439704383,
		AccessHash: 0,
	}

	// Resolve the channel username from URL
	// channelUsername := channelURL[len("https://t.me/"):]

	// // Get channel info
	// resolved, err := api.ContactsResolveUsername(ctx, &tg.ContactsResolveUsernameRequest{
	// 	Username: channelUsername,
	// })
	// if err != nil {
	// 	return fmt.Errorf("failed to resolve channel: %v", err)
	// }

	// var peer tg.InputPeerClass
	// for _, chat := range resolved.Chats {
	// 	if channel, ok := chat.(*tg.Channel); ok {
	// 		peer = &tg.InputPeerChannel{
	// 			ChannelID:  channel.ID,
	// 			AccessHash: channel.AccessHash,
	// 		}
	// 		fmt.Printf("Found group: %s, Chat ID: -100%d\n", channel.Title, channel.ID)
	// 		break
	// 	}
	// }
	// Get channel messages
	messages, err := api.MessagesGetHistory(ctx, &tg.MessagesGetHistoryRequest{
		Peer:  peer,
		Limit: limit,
	})
	if err != nil {
		return fmt.Errorf("failed to get messages: %v", err)
	}

	// Process messages
	for _, msg := range messages.(*tg.MessagesChannelMessages).Messages {
		if m, ok := msg.(*tg.Message); ok {
			log.Printf("Message ID: %d, Text: %s\n", m.ID, m.Message)
		}
	}

	return nil
}

func (c *Client) Run(ctx context.Context) error {
	return c.client.Run(ctx, func(ctx context.Context) error {
		// Here you can add authentication if needed
		return nil
	})
}
