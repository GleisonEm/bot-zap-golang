package services

import (
	"context"
	"fmt"
	"github.com/GleisonEm/bot-claudio-zap-golang/domains/message"
	domainMessage "github.com/GleisonEm/bot-claudio-zap-golang/domains/message"
	"github.com/GleisonEm/bot-claudio-zap-golang/pkg/whatsapp"
	"github.com/GleisonEm/bot-claudio-zap-golang/validations"
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

func (service serviceMessage) ReactMessage(ctx context.Context, request message.ReactionRequest) (response message.ReactionResponse, err error) {
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


