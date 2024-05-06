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
	MentionAdmin    bool
}

type SendMessageStickerParams struct {
	MentionAllUsers bool
	Message         string
	MentionUsers    []string
	IsQuotedMessage bool
	ImageMessage    *waProto.ImageMessage
}
