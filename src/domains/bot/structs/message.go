package structs

import (
	waProto "go.mau.fi/whatsmeow/binary/proto"
)

type SendMessageParams struct {
	MentionAllUsers bool
	Message         string
	MentionUsers    []string
	IsQuotedMessage bool
	AudioMessage    waProto.AudioMessage
}
