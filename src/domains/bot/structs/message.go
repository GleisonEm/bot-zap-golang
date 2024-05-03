package structs

import (
	waProto "go.mau.fi/whatsmeow/binary/proto"
)

type SendMessageParams struct {
	MentionAllUsers bool
	Message         string
	MentionUsers    []string
	IsQuotedMessage bool
	AudioMessage    *waProto.AudioMessage
	IsQuotedType    string
}

type SendMessageStickerParams struct {
	MentionAllUsers bool
	Message         string
	MentionUsers    []string
	IsQuotedMessage bool
	ImageMessage   *waProto.ImageMessage
}
