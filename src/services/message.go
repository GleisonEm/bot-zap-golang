package services

import (
	"context"
	"fmt"
	"mime"
	"os"
	"time"

	"github.com/GleisonEm/bot-claudio-zap-golang/config"
	ServiceAppContext "github.com/GleisonEm/bot-claudio-zap-golang/contexts"
	DomainBot "github.com/GleisonEm/bot-claudio-zap-golang/domains/bot/structs"
	"github.com/GleisonEm/bot-claudio-zap-golang/domains/message"
	domainMessage "github.com/GleisonEm/bot-claudio-zap-golang/domains/message"
	ServicesBot "github.com/GleisonEm/bot-claudio-zap-golang/internal/bot/services"
	"github.com/GleisonEm/bot-claudio-zap-golang/pkg/utils"
	"github.com/GleisonEm/bot-claudio-zap-golang/pkg/whatsapp"
	"github.com/GleisonEm/bot-claudio-zap-golang/validations"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
	// "time"
)

type serviceMessage struct {
	WaCli *whatsmeow.Client
}

func NewMessageService(waCli *whatsmeow.Client) domainMessage.IMessageService {
	return &serviceMessage{
		WaCli: waCli,
	}
}

func (service serviceMessage) ReactMessage(ctx context.Context, request domainMessage.ReactionRequest) (response domainMessage.ReactionResponse, err error) {
	if err = validations.ValidateReactMessage(ctx, request); err != nil {
		return response, err
	}
	dataWaRecipient, err := whatsapp.ValidateJidWithLogin(service.WaCli, request.Phone)
	if err != nil {
		return response, err
	}
	fmt.Println("ReactMessage request.MessageID", request.MessageID, request.Phone)
	// msg := &waProto.Message{
	// 	ReactionMessage: &waProto.ReactionMessage{
	// 		Key: &waProto.MessageKey{
	// 			FromMe:    proto.Bool(true),
	// 			Id:        proto.String(request.MessageID),
	// 			RemoteJid: proto.String(dataWaRecipient.String()),
	// 		},
	// 		Text:              proto.String(request.Emoji),
	// 		SenderTimestampMs: proto.Int64(time.Now().UnixMilli()),
	// 	},
	// }
	ts, err := service.WaCli.SendMessage(ctx, dataWaRecipient, service.WaCli.BuildReaction(dataWaRecipient, types.JID{
		User:   request.Phone, //Agus
		Server: types.DefaultUserServer,
	}, request.MessageID, request.Emoji))
	if err != nil {
		return response, err
	}

	response.MessageID = ts.ID
	response.Status = fmt.Sprintf("Reaction sent to %s (server timestamp: %s)", request.Phone, ts.Timestamp)
	return response, nil
}

func (service serviceMessage) RevokeMessage(ctx context.Context, request domainMessage.RevokeRequest) (response domainMessage.RevokeResponse, err error) {
	if err = validations.ValidateRevokeMessage(ctx, request); err != nil {
		return response, err
	}
	dataWaRecipient, err := whatsapp.ValidateJidWithLogin(service.WaCli, request.Phone)
	if err != nil {
		return response, err
	}
	fmt.Println("fazendo revoke")
	fmt.Println("ReactMessage request.MessageID", request.MessageID)
	ts, err := service.WaCli.SendMessage(context.Background(), dataWaRecipient, service.WaCli.BuildRevoke(dataWaRecipient, types.JID{
		User:   request.Phone, //Agus
		Server: types.DefaultUserServer,
	}, request.MessageID))

	fmt.Println("terminei revoke", ts)
	if err != nil {
		return response, err
	}

	response.MessageID = ts.ID
	response.Status = fmt.Sprintf("Revoke success %s (server timestamp: %s)", request.Phone, ts.Timestamp)
	return response, nil
}

func (service serviceMessage) UpdateMessage(ctx context.Context, request domainMessage.UpdateMessageRequest) (response domainMessage.UpdateMessageResponse, err error) {
	if err = validations.ValidateUpdateMessage(ctx, request); err != nil {
		return response, err
	}

	dataWaRecipient, err := whatsapp.ValidateJidWithLogin(service.WaCli, request.Phone)
	if err != nil {
		return response, err
	}

	msg := &waProto.Message{Conversation: proto.String(request.Message)}
	ts, err := service.WaCli.SendMessage(context.Background(), dataWaRecipient, service.WaCli.BuildEdit(dataWaRecipient, request.MessageID, msg))
	if err != nil {
		return response, err
	}

	response.MessageID = ts.ID
	response.Status = fmt.Sprintf("Update message success %s (server timestamp: %s)", request.Phone, ts.Timestamp)
	return response, nil
}

func (service serviceMessage) ConvertMessageAudioToText(
	ctx context.Context, fromChat string, sender string, name string, stanzaID string, messageText string, sendMessageParams DomainBot.SendMessageParams,
) {
	dataWaRecipient, _ := whatsapp.ValidateJidWithLogin(service.WaCli, fromChat)
	dataWaRecipientSender, _ := whatsapp.ValidateJidWithLogin(service.WaCli, sender)
	audioMessage := &sendMessageParams.AudioMessage

	reactLoading, errReactLoading := service.WaCli.SendMessage(context.Background(), dataWaRecipient, service.WaCli.BuildReaction(dataWaRecipient, dataWaRecipientSender, stanzaID, "⏳"))
	fmt.Println("mandado react", reactLoading, errReactLoading)

	path, err := ServiceAppContext.Context.MessageService.ExtractMedia(context.Background(), config.PathStorages, *audioMessage)
	if err != nil {
		log.Errorf("Failed to download audio: %v", err)
		s, err2 := service.WaCli.SendMessage(
			context.Background(), dataWaRecipient, service.WaCli.BuildReaction(dataWaRecipient, dataWaRecipientSender, stanzaID, "❌"),
		)
		fmt.Println("mandado react", s, err2)
	} else {
		log.Infof("audio downloaded to %s", path)
		filePathAudio := fmt.Sprintf("%s/%d-%s%s", config.PathStorages, time.Now().Unix(), uuid.NewString(), ".wav")
		errConvertOgaToWav := utils.ConvertOgaToWav(path.MediaPath, filePathAudio)

		if errConvertOgaToWav != nil {
			s, err2 := service.WaCli.SendMessage(
				context.Background(), dataWaRecipient, service.WaCli.BuildReaction(dataWaRecipient, dataWaRecipientSender, stanzaID, "❌"),
			)
			fmt.Println("mandado react", s, err2)
			log.Errorf("Failed to convert audio: %v", errConvertOgaToWav)
		}
		response, errSendToRecogntionApiAudioFile := ServicesBot.SendToRecogntionApiAudioFile(filePathAudio)

		if errSendToRecogntionApiAudioFile != nil {
			s, err2 := service.WaCli.SendMessage(
				context.Background(), dataWaRecipient, service.WaCli.BuildReaction(dataWaRecipient, dataWaRecipientSender, stanzaID, "❌"),
			)
			fmt.Println("mandado react", s, err2)
			log.Errorf("Failed to convert audio to text: %v", errSendToRecogntionApiAudioFile)
		}

		if response.Text == "" {
			log.Errorf("Failed to convert audio to text: %v", "empty response")

			s, err2 := service.WaCli.SendMessage(
				context.Background(), dataWaRecipient, service.WaCli.BuildReaction(dataWaRecipient, dataWaRecipientSender, stanzaID, "❌"),
			)
			fmt.Println("mandado react", s, err2)
		}

		go service.WaCli.SendMessage(context.Background(), dataWaRecipient, service.WaCli.BuildReaction(dataWaRecipient, dataWaRecipientSender, stanzaID, "✅"))

		ServiceAppContext.Context.SendService.SendMessage(context.Background(), fromChat, sender, name, stanzaID, messageText, DomainBot.SendMessageParams{
			Message:         response.Text,
			AudioMessage:    *audioMessage,
			IsQuotedMessage: true,
		})
	}
}

func (service serviceMessage) ExtractMedia(ctx context.Context, storageLocation string, mediaFile whatsmeow.DownloadableMessage) (extractedMedia message.ExtractedMessageMedia, err error) {
	if mediaFile == nil {
		logrus.Info("Skip download because data is nil")
		return extractedMedia, nil
	}

	data, err := service.WaCli.Download(mediaFile)
	if err != nil {
		logrus.Info("err download", err)
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
		logrus.Info("err write", err)
		return extractedMedia, err
	}
	return extractedMedia, nil
}
