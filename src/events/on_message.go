package events

import (
	"context"
	"fmt"
	"strings"

	"github.com/GleisonEm/bot-claudio-zap-golang/config"
	ServiceAppContext "github.com/GleisonEm/bot-claudio-zap-golang/contexts"
	DomainBot "github.com/GleisonEm/bot-claudio-zap-golang/domains/bot/structs"
	"github.com/GleisonEm/bot-claudio-zap-golang/pkg/utils"
	"go.mau.fi/whatsmeow/types/events"
)

type ExtractedMedia struct {
	MediaPath string `json:"media_path"`
	MimeType  string `json:"mime_type"`
	Caption   string `json:"caption"`
}

func OnMessage(evt *events.Message) {

	// fmt.Println("RawMessage", evt.RawMessage.String())
	messageText := ""
	if evt.Message.GetExtendedTextMessage().GetText() != "" {
		messageText = evt.Message.GetExtendedTextMessage().GetText()
	} else if evt.Message.GetConversation() != "" {
		messageText = evt.Message.GetConversation()
	}
	argument := utils.GetArgument(messageText)
	command := utils.ProcessCommand(messageText)
	sender := evt.Info.Sender.User + "@" + evt.Info.Sender.Server
	fromChat := evt.Info.Chat.String()
	stanzaID := evt.Info.ID
	// if fromChat == "558796485300-1461896371@g.us" {
	// 	return
	// }
	fmt.Println("Received message ", string(evt.Info.ID), evt.Info.SourceString(), "is group:", evt.Info.IsGroup)
	// , "is user", evt.Info.Chat.IsUser(), "is broadcast", evt.Info.Chat.IsBroadcast(), "is server", evt.Info.Chat.IsServer(), "is status", evt.Info.Chat.IsStatus(), "is group", evt.Info.Chat.IsGroup(), "is user", evt.Info.Chat.IsUser(), "is broadcast", evt.Info.Chat.IsBroadcast(), "is server", evt.Info.Chat.IsServer(), "is status", evt.Info.Chat.IsStatus()
	// fmt.Println(argument, "\t", sender, "\t", evt.Info.Sender.Server, "\t", evt.Info.Chat)

	fmt.Println("command", command, "argument", argument, "sender", sender, "fromChat", fromChat, "stanzaID", stanzaID, "messageText", messageText, evt.Message.GetConversation(), evt.Message.String())

	if !strings.Contains(config.ChatsDevEnabled, fromChat) {
		return
	}

	if command == "!sticker" {
		go ServiceAppContext.Context.SendService.SendSticker(context.Background(), fromChat, sender, argument, stanzaID, messageText, DomainBot.SendMessageStickerParams{
			ImageMessage: evt.RawMessage.ExtendedTextMessage.ContextInfo.QuotedMessage.GetImageMessage(),
		})
	}

	if command == "!audio" {
		go ServiceAppContext.Context.SendService.SendAudioFunny(context.Background(), fromChat, sender, argument, stanzaID, messageText)
	}

	if command == "@todes" {
		go handleTodes(context.Background(), fromChat, sender, argument, stanzaID, messageText, DomainBot.SendMessageParams{
			MentionAllUsers: true,
		})
	}

	if command == "@supremacy" {
		go handleTodes(context.Background(), fromChat, sender, argument, stanzaID, messageText, DomainBot.SendMessageParams{
			MentionAdmin: true,
		})
	}

	if command == "!balinha" {
		go ServiceAppContext.Context.SendService.SendMessage(context.Background(), fromChat, sender, argument, stanzaID, messageText, DomainBot.SendMessageParams{
			Message: "Aí é com o famoso 😉",
		})
	}

	if command == "!documentacao" {
		messageDoc := "Verificação documental completa."
		if argument != "h1234" {
			messageDoc = "Nenhum documento encontrado."
		}

		ServiceAppContext.Context.SendService.SendMessage(context.Background(), fromChat, sender, argument, stanzaID, messageText, DomainBot.SendMessageParams{
			Message: messageDoc,
		})
	}

	if command == "!localizacao" {
		messageLoc := "TE"
		if argument != "h1234" {
			messageLoc = "Nenhuma localização encontrada para esse documento."
		}

		ServiceAppContext.Context.SendService.SendMessage(context.Background(), fromChat, sender, argument, stanzaID, messageText, DomainBot.SendMessageParams{
			Message: messageLoc,
		})
	}

	if command == "!transcrever" {
		audioMessage := evt.RawMessage.ExtendedTextMessage.ContextInfo.QuotedMessage.GetAudioMessage()
		// fmt.Println("audio message transcrever", audioMessage)
		if audioMessage != nil {
			go ServiceAppContext.Context.MessageService.ConvertMessageAudioToText(context.Background(), fromChat, sender, argument, stanzaID, messageText, DomainBot.SendMessageParams{
				AudioMessage: audioMessage,
			})
		}
	}

	if !evt.Info.IsGroup {
		audioMessage := evt.Message.GetAudioMessage()
		// fmt.Println("audio message direto", audioMessage)
		if audioMessage != nil {
			go ServiceAppContext.Context.MessageService.ConvertMessageAudioToText(context.Background(), fromChat, sender, argument, stanzaID, messageText, DomainBot.SendMessageParams{
				AudioMessage: audioMessage,
			})
		}
	}
}

func handleTodes(ctx context.Context, fromChat string, sender string, name string, stanzaID string, messageText string, sendMessageParams DomainBot.SendMessageParams) {
	ServiceAppContext.Context.SendService.SendMessage(ctx, fromChat, sender, name, stanzaID, messageText, sendMessageParams)
}
