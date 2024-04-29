package events

import (
	"context"
	"fmt"
	"mime"
	"os"
	"time"

	"github.com/GleisonEm/bot-claudio-zap-golang/config"
	ServiceAppContext "github.com/GleisonEm/bot-claudio-zap-golang/contexts"
	DomainBot "github.com/GleisonEm/bot-claudio-zap-golang/domains/bot/structs"
	ServicesBot "github.com/GleisonEm/bot-claudio-zap-golang/internal/bot/services"
	"github.com/GleisonEm/bot-claudio-zap-golang/pkg/utils"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types/events"
)

type ExtractedMedia struct {
	MediaPath string `json:"media_path"`
	MimeType  string `json:"mime_type"`
	Caption   string `json:"caption"`
}

func OnMessage(evt *events.Message) {
	// fmt.Println("RawMessage", evt.RawMessage)

	messageText := evt.Message.GetExtendedTextMessage().GetText()
	argument := utils.GetArgument(messageText)
	command := utils.ProcessCommand(messageText)
	formattedJid, _ := utils.GetLastJid(evt.Info.SourceString())
	stanzaID := evt.Info.ID

	fmt.Println(argument, "\t", formattedJid)
	if command == "!audio" {
		ServiceAppContext.Context.SendService.SendAudioFunny(context.Background(), formattedJid, argument, stanzaID, messageText)
	}

	if command == "@todes" {
		ServiceAppContext.Context.SendService.SendMessage(context.Background(), formattedJid, argument, stanzaID, messageText, DomainBot.SendMessageParams{
			MentionAllUsers: true,
		})
	}

	if command == "!transcrever" {
		audioMessage := evt.RawMessage.ExtendedTextMessage.ContextInfo.QuotedMessage.GetAudioMessage()
		if audioMessage != nil {
			path, err := ExtractMedia(config.PathStorages, audioMessage)
			if err != nil {
				log.Errorf("Failed to download audio: %v", err)
			} else {
				log.Infof("audio downloaded to %s", path)
				filePathAudio := fmt.Sprintf("%s/%d-%s%s", config.PathStorages, time.Now().Unix(), uuid.NewString(), ".wav")
				errConvertOgaToWav := utils.ConvertOgaToWav(path.MediaPath, filePathAudio)

				if errConvertOgaToWav != nil {
					log.Errorf("Failed to convert audio: %v", errConvertOgaToWav)
				}
				response, errSendToRecogntionApiAudioFile := ServicesBot.SendToRecogntionApiAudioFile(filePathAudio)

				if errSendToRecogntionApiAudioFile != nil {
					log.Errorf("Failed to convert audio to text: %v", errSendToRecogntionApiAudioFile)
				}

				ServiceAppContext.Context.SendService.SendMessage(context.Background(), formattedJid, argument, stanzaID, messageText, DomainBot.SendMessageParams{
					Message:         response.Text,
					AudioMessage:    *audioMessage,
					IsQuotedMessage: true,
				})
			}
		}
	}

	audioMessage := evt.Message.GetAudioMessage()
	if audioMessage != nil {
		path, err := ExtractMedia(config.PathStorages, audioMessage)
		if err != nil {
			log.Errorf("Failed to download audio: %v", err)
		} else {
			log.Infof("audio downloaded to %s", path)
			filePathAudio := fmt.Sprintf("%s/%d-%s%s", config.PathStorages, time.Now().Unix(), uuid.NewString(), ".wav")
			errConvertOgaToWav := utils.ConvertOgaToWav(path.MediaPath, filePathAudio)

			if errConvertOgaToWav != nil {
				log.Errorf("Failed to convert audio: %v", errConvertOgaToWav)
			}
			response, errSendToRecogntionApiAudioFile := ServicesBot.SendToRecogntionApiAudioFile(filePathAudio)

			if errSendToRecogntionApiAudioFile != nil {
				log.Errorf("Failed to convert audio to text: %v", errSendToRecogntionApiAudioFile)
			}

			ServiceAppContext.Context.SendService.SendMessage(context.Background(), formattedJid, argument, stanzaID, messageText, DomainBot.SendMessageParams{
				Message:         response.Text,
				AudioMessage:    *audioMessage,
				IsQuotedMessage: true,
			})
		}
	}
}

func ExtractMedia(storageLocation string, mediaFile whatsmeow.DownloadableMessage) (extractedMedia ExtractedMedia, err error) {
	if mediaFile == nil {
		logrus.Info("Skip download because data is nil")
		return extractedMedia, nil
	}

	data, err := (ServiceAppContext.Context.AppService.GetWaCli(context.Background())).Download(mediaFile)
	if err != nil {
		return extractedMedia, err
	}

	switch media := mediaFile.(type) {
	case *waProto.ImageMessage:
		extractedMedia.MimeType = media.GetMimetype()
		extractedMedia.Caption = media.GetCaption()
	case *waProto.AudioMessage:
		extractedMedia.MimeType = media.GetMimetype()
	case *waProto.VideoMessage:
		extractedMedia.MimeType = media.GetMimetype()
		extractedMedia.Caption = media.GetCaption()
	case *waProto.StickerMessage:
		extractedMedia.MimeType = media.GetMimetype()
	case *waProto.DocumentMessage:
		extractedMedia.MimeType = media.GetMimetype()
		extractedMedia.Caption = media.GetCaption()
	}

	extensions, _ := mime.ExtensionsByType(extractedMedia.MimeType)
	extractedMedia.MediaPath = fmt.Sprintf("%s/%d-%s%s", storageLocation, time.Now().Unix(), uuid.NewString(), extensions[0])
	err = os.WriteFile(extractedMedia.MediaPath, data, 0600)
	if err != nil {
		return extractedMedia, err
	}
	return extractedMedia, nil
}
