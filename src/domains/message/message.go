package message

import (
	"context"

	DomainBot "github.com/GleisonEm/bot-claudio-zap-golang/domains/bot/structs"
	"go.mau.fi/whatsmeow"
)

type IMessageService interface {
	ReactMessage(ctx context.Context, request ReactionRequest) (response ReactionResponse, err error)
	RevokeMessage(ctx context.Context, request RevokeRequest) (response RevokeResponse, err error)
	UpdateMessage(ctx context.Context, request UpdateMessageRequest) (response UpdateMessageResponse, err error)
	ConvertMessageAudioToText(ctx context.Context, fromChat string, sender string, name string, stanzaID string, messageText string, sendMessageParams DomainBot.SendMessageParams)
	ExtractMedia(ctx context.Context, storageLocation string, mediaFile whatsmeow.DownloadableMessage) (extractedMedia ExtractedMessageMedia, err error)
}
