package services

// import (
// 	"context"
// 	"fmt"

// 	"github.com/GleisonEm/bot-claudio-zap-golang/pkg/whatsapp"
// 	"go.mau.fi/whatsmeow"
// 	"go.mau.fi/whatsmeow/types"
// 	// "time"
// )

// func ReactMessage(ctx context.Context, waCli *whatsmeow.Client) error {

// 	dataWaRecipient, err := whatsapp.ValidateJidWithLogin(waCli, request.Phone)
// 	if err != nil {
// 		return response, err
// 	}
// 	fmt.Println("ReactMessage request.MessageID", request.MessageID, request.Phone)
// 	// msg := &waProto.Message{
// 	// 	ReactionMessage: &waProto.ReactionMessage{
// 	// 		Key: &waProto.MessageKey{
// 	// 			FromMe:    proto.Bool(true),
// 	// 			Id:        proto.String(request.MessageID),
// 	// 			RemoteJid: proto.String(dataWaRecipient.String()),
// 	// 		},
// 	// 		Text:              proto.String(request.Emoji),
// 	// 		SenderTimestampMs: proto.Int64(time.Now().UnixMilli()),
// 	// 	},
// 	// }
// 	ts, err := service.WaCli.SendMessage(
// 		ctx,
// 		dataWaRecipient,
// 		service.WaCli.BuildReaction(dataWaRecipient, types.JID{
// 			User:   request.Phone, //Agus
// 			Server: types.DefaultUserServer,
// 		}, request.MessageID, request.Emoji),
// 	)
// 	if err != nil {
// 		return response, err
// 	}

// 	response.MessageID = ts.ID
// 	response.Status = fmt.Sprintf("Reaction sent to %s (server timestamp: %s)", request.Phone, ts.Timestamp)
// 	return nil
// }
