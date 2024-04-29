## WhatsApp API Multi Device Version

![release version](https://img.shields.io/github/v/release/aldinokemal/go-whatsapp-web-multidevice)
<br>
![Build Image](https://github.com/aldinokemal/go-whatsapp-web-multidevice/actions/workflows/build-docker-image.yaml/badge.svg)
<br>
![release windows](https://github.com/aldinokemal/go-whatsapp-web-multidevice/actions/workflows/release-windows.yml/badge.svg)
![release linux](https://github.com/aldinokemal/go-whatsapp-web-multidevice/actions/workflows/release-linux.yml/badge.svg)
![release macos](https://github.com/aldinokemal/go-whatsapp-web-multidevice/actions/workflows/release-mac.yml/badge.svg)

### Support `ARM` Architecture

Now that we support ARM64 for Linux:

- [Release](https://github.com/aldinokemal/go-whatsapp-web-multidevice/releases/latest) for ARM64
- [Docker Image](https://hub.docker.com/r/aldinokemal2104/go-whatsapp-web-multidevice/tags) for ARM64.

### Feature

- Send WhatsApp message via http API, [docs/openapi.yml](./docs/openapi.yaml) for more details
- Compress image before send
- Compress video before send
- Change OS name become your app (it's the device name when connect via mobile)
    - `--os=Chrome` or `--os=MyApplication`
- Basic Auth (able to add multi credentials)
    - `--basic-auth=kemal:secret,toni:password,userName:secretPassword`, or you can simplify
    - `-b=kemal:secret,toni:password,userName:secretPassword`
- Customizable port and debug mode
    - `--port 8000`
    - `--debug true`
- Auto reply message
    - `--autoreply="Don't reply this message"`
- Webhook for received message
    - `--webhook="http://yourwebhook.site/handler"`, or you can simplify
    - `-w="http://yourwebhook.site/handler"`
- For more command `./main --help`

### Required (without docker)

- Mac OS:
    - `brew install vips`
    - `brew install ffmpeg`
    - `export CGO_CFLAGS_ALLOW="-Xpreprocessor"`
- Linux:
    - `sudo apt update`
    - `sudo apt install libvips-dev`
    - `sudo apt install ffmpeg`
- Windows (not recomended, prefer using [WSL](https://docs.microsoft.com/en-us/windows/wsl/install)):
    - install vips library, or you can check here https://www.libvips.org/install.html
    - install ffmpeg, download [here](https://www.ffmpeg.org/download.html#build-windows)
    - add to vips & ffmpg to [environment variable](https://www.google.com/search?q=windows+add+to+environment+path)

### How to use

#### Basic

1. Clone this repo: `git clone https://github.com/aldinokemal/go-whatsapp-web-multidevice`
2. Open the folder that was cloned via cmd/terminal.
3. run `cd src`
4. run `go run main.go`
5. Open `http://localhost:3000`

#### Docker (you don't need to install in required)

1. Clone this repo: `git clone https://github.com/aldinokemal/go-whatsapp-web-multidevice`
2. Open the folder that was cloned via cmd/terminal.
3. run `docker-compose up -d --build`
4. open `http://localhost:3000`

#### Build your own binary

1. Clone this repo `git clone https://github.com/aldinokemal/go-whatsapp-web-multidevice`
2. Open the folder that was cloned via cmd/terminal.
3. run `cd src`
4. run
    1. Linux & MacOS: `go build -o whatsapp`
    2. Windows (CMD / PowerShell): `go build -o whatsapp.exe`
5. run
    1. Linux & MacOS: `./whatsapp`
        1. run `./whatsapp --help` for more detail flags
    2. Windows: `.\whatsapp.exe` or you can double-click it
        1. run `.\whatsapp.exe --help` for more detail flags
6. open `http://localhost:3000` in browser

### Production Mode (docker)

```
docker run --detach --publish=3000:3000 --name=whatsapp --restart=always --volume=$(docker volume create --name=whatsapp):/app/storages aldinokemal2104/go-whatsapp-web-multidevice --autoreply="Dont't reply this message please"
```

### Production Mode (binary)

- download binary from [release](https://github.com/aldinokemal/go-whatsapp-web-multidevice/releases)

You can fork or edit this source code !

### Current API

You can check [docs/openapi.yml](./docs/openapi.yaml) for detail API, furthermore you can generate HTTP Client from this
API using [openapi-generator](https://openapi-generator.tech/#try)

| Feature | Menu                           | Method | URL                         |
|---------|--------------------------------|--------|-----------------------------|
| ✅       | Login                          | GET    | /app/login                  |
| ✅       | Logout                         | GET    | /app/logout                 |
| ✅       | Reconnect                      | GET    | /app/reconnect              |
| ✅       | User Info                      | GET    | /user/info                  |
| ✅       | User Avatar                    | GET    | /user/avatar                |
| ✅       | User My Group List             | GET    | /user/my/groups             |
| ✅       | User My Privacy Setting        | GET    | /user/my/privacy            |
| ✅       | Send Message                   | POST   | /send/message               |
| ✅       | Send Image                     | POST   | /send/image                 |
| ✅       | Send Audio                     | POST   | /send/audio                 |
| ✅       | Send File                      | POST   | /send/file                  |
| ✅       | Send Video                     | POST   | /send/video                 |
| ✅       | Send Contact                   | POST   | /send/contact               |
| ✅       | Send Link                      | POST   | /send/link                  |
| ✅       | Send Location                  | POST   | /send/location              |
| ✅       | Send Poll / Vote               | POST   | /send/poll                  |
| ✅       | Revoke Message                 | POST   | /message/:message_id/revoke |
| ✅       | React Message                  | POST   | /message/:message_id/react  |
| ✅       | Edit Message                   | POST   | /message/:message_id/update |
| ✅       | Join Group With Link           | POST   | /group/join-with-link       |
| ✅       | Leave Group                    | POST   | /group/leave                |
| ✅       | Create Group                   | POST   | /group                      |
| ❌       | Add More Participants in Group | POST   |                             |
| ❌       | Remove Participant in Group    | POST   |                             |
| ❌       | Promote Participant in Group   | POST   |                             |

```
✅ = Available
❌ = Not Available Yet
```

### App User Interface

1. Homepage ![Homepage](https://i.ibb.co/TBNcFT0/homepage.png)
2. Login ![Login](https://i.ibb.co/jkcB15R/login.png?v=1)
3. Send Message ![Send Message](https://i.ibb.co/rc3NXMX/send-message.png?v1)
4. Send Image ![Send Image](https://i.ibb.co/BcFL3SD/send-image.png?v1)
5. Send File ![Send File](https://i.ibb.co/f4yxjpp/send-file.png)
6. Send Video ![Send Video](https://i.ibb.co/PrD3P51/send-video.png)
7. Send Contact ![Send Contact](https://i.ibb.co/4810H7N/send-contact.png)
8. Send Location ![Send Location](https://i.ibb.co/TWsy09G/send-location.png)
9. Send Audio ![Send Location](https://i.ibb.co/p1wL4wh/Send-Audio.png)
10. Send Poll ![Send Poll](https://i.ibb.co/mq2fGHz/send-poll.png)
11. Revoke Message ![Revoke Message](https://i.ibb.co/yswhvQY/revoke.png?v1)
12. Reaction Message ![Revoke Message](https://i.ibb.co/BfHgSHG/react-message.png)
13. Edit Message ![Edit Message](https://i.ibb.co/kXfpqJw/update-message.png)
14. User Info ![User Info](https://i.ibb.co/3zjX6Cz/user-info.png?v=1)
15. User Avatar ![User Avatar](https://i.ibb.co/ZmJZ4ZW/search-avatar.png?v=1)
16. My Privacy ![My Privacy](https://i.ibb.co/Cw1sMQz/my-privacy.png)
17. My Group ![My Group](https://i.ibb.co/WB268Xy/list-group.png)
18. Auto Reply ![Auto Reply](https://i.ibb.co/D4rTytX/IMG-20220517-162500.jpg)
19. Basic Auth Prompt ![Basic Auth](https://i.ibb.co/PDjQ92W/Screenshot-2022-11-06-at-14-06-29.png)

### Mac OS NOTE

- Please do this if you have an error (invalid flag in pkg-config --cflags: -Xpreprocessor)
  `export CGO_CFLAGS_ALLOW="-Xpreprocessor"`



			// replyMessage := &waProto.Message{
			// 	ExtendedTextMessage: &waProto.ExtendedTextMessage{
			// 		Text: proto.String("kkkkkkkk"),
			// 		ContextInfo: &waProto.ContextInfo{
			// 			QuotedMessage: &waProto.Message{
			// 				Conversation: proto.String("replyText"),
			// 			},
			// 			StanzaId:    &stanzaID,
			// 			Participant: proto.String(dataWaRecipient.String()), // O participante é quem enviou a mensagem original
			// 		},
			// 	},
			// }








                err := os.Remove(path.MediaPath)
    if err != nil {
        log.Errorf("Failed to remove file %s: %v", path.MediaPath, err)
    }
    err = os.Remove(filePathAudio)
    if err != nil {
        log.Errorf("Failed to remove file %s: %v", filePathAudio, err)
    }


    	audioMessage := evt.Message.GetAudioMessage()
	if audioMessage != nil {
		path, err := ExtractMedia(config.PathStorages, audioMessage)
		if err != nil {
			log.Errorf("Failed to download audio: %v", err)
		} else {
			log.Infof("audio downloaded to %s", path)
			filePathAudio := "storages/audio.wav"
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


package whatsapp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/GleisonEm/bot-claudio-zap-golang/config"
	"github.com/GleisonEm/bot-claudio-zap-golang/internal/bot/services"
	"github.com/GleisonEm/bot-claudio-zap-golang/internal/websocket"
	pkgError "github.com/GleisonEm/bot-claudio-zap-golang/pkg/error"
	"github.com/GleisonEm/bot-claudio-zap-golang/pkg/utils"

	// "github.com/GleisonEm/bot-claudio-zap-golang/internal/rest/helpers"
	"mime"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/appstate"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/protobuf/proto"
)

var (
	cli           *whatsmeow.Client
	log           waLog.Logger
	historySyncID int32
	startupTime   = time.Now().Unix()
)

type ExtractedMedia struct {
	MediaPath string `json:"media_path"`
	MimeType  string `json:"mime_type"`
	Caption   string `json:"caption"`
}

type evtReaction struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

type evtMessage struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	RepliedId string `json:"replied_id"`
}

func SanitizePhone(phone *string) {
	if phone != nil && len(*phone) > 0 && !strings.Contains(*phone, "@") {
		if len(*phone) <= 15 {
			*phone = fmt.Sprintf("%s%s", *phone, config.WhatsappTypeUser)
		} else {
			*phone = fmt.Sprintf("%s%s", *phone, config.WhatsappTypeGroup)
		}
	}
}

func GetPlatformName(deviceID int) string {
	switch deviceID {
	case 0:
		return "UNKNOWN"
	case 1:
		return "CHROME"
	case 2:
		return "FIREFOX"
	case 3:
		return "IE"
	case 4:
		return "OPERA"
	case 5:
		return "SAFARI"
	case 6:
		return "EDGE"
	case 7:
		return "DESKTOP"
	case 8:
		return "IPAD"
	case 9:
		return "ANDROID_TABLET"
	case 10:
		return "OHANA"
	case 11:
		return "ALOHA"
	case 12:
		return "CATALINA"
	case 13:
		return "TCL_TV"
	default:
		return "UNKNOWN"
	}
}

func ParseJID(arg string) (types.JID, error) {
	if arg[0] == '+' {
		arg = arg[1:]
	}
	if !strings.ContainsRune(arg, '@') {
		return types.NewJID(arg, types.DefaultUserServer), nil
	} else {
		recipient, err := types.ParseJID(arg)
		if err != nil {
			fmt.Printf("invalid JID %s: %v", arg, err)
			return recipient, pkgError.ErrInvalidJID
		} else if recipient.User == "" {
			fmt.Printf("invalid JID %v: no server specified", arg)
			return recipient, pkgError.ErrInvalidJID
		}
		return recipient, nil
	}
}

func IsOnWhatsapp(waCli *whatsmeow.Client, jid string) bool {
	// only check if the jid a user with @s.whatsapp.net
	if strings.Contains(jid, "@s.whatsapp.net") {
		data, err := waCli.IsOnWhatsApp([]string{jid})
		if err != nil {
			panic(pkgError.InvalidJID(err.Error()))
		}

		for _, v := range data {
			if !v.IsIn {
				return false
			}
		}
	}

	return true
}

func ValidateJidWithLogin(waCli *whatsmeow.Client, jid string) (types.JID, error) {
	MustLogin(waCli)

	if !IsOnWhatsapp(waCli, jid) {
		return types.JID{}, pkgError.InvalidJID(fmt.Sprintf("Phone %s is not on whatsapp", jid))
	}

	return ParseJID(jid)
}

func InitWaDB() *sqlstore.Container {
	// Running Whatsapp
	log = waLog.Stdout("Main", config.WhatsappLogLevel, true)
	dbLog := waLog.Stdout("Database", config.WhatsappLogLevel, true)
	storeContainer, err := sqlstore.New("sqlite3", fmt.Sprintf("file:%s/%s?_foreign_keys=off", config.PathStorages, config.DBName), dbLog)
	if err != nil {
		log.Errorf("Failed to connect to database: %v", err)
		panic(pkgError.InternalServerError(fmt.Sprintf("Failed to connect to database: %v", err)))

	}
	return storeContainer
}

func InitWaCLI(storeContainer *sqlstore.Container) *whatsmeow.Client {
	device, err := storeContainer.GetFirstDevice()
	if err != nil {
		log.Errorf("Failed to get device: %v", err)
		panic(err)
	}

	osName := fmt.Sprintf("%s %s", config.AppOs, config.AppVersion)
	store.DeviceProps.PlatformType = &config.AppPlatform
	store.DeviceProps.Os = &osName
	cli = whatsmeow.NewClient(device, waLog.Stdout("Client", config.WhatsappLogLevel, true))
	cli.EnableAutoReconnect = true
	cli.AutoTrustIdentity = true
	cli.AddEventHandler(handler)

	return cli
}

func MustLogin(waCli *whatsmeow.Client) {
	if waCli == nil {
		panic(pkgError.InternalServerError("Whatsapp client is not initialized"))
	}
	if !waCli.IsConnected() {
		panic(pkgError.ErrNotConnected)
	} else if !waCli.IsLoggedIn() {
		panic(pkgError.ErrNotLoggedIn)
	}
}

func handler(rawEvt interface{}) {
	switch evt := rawEvt.(type) {
	case *events.AppStateSyncComplete:
		if len(cli.Store.PushName) > 0 && evt.Name == appstate.WAPatchCriticalBlock {
			err := cli.SendPresence(types.PresenceAvailable)
			if err != nil {
				log.Warnf("Failed to send available presence: %v", err)
			} else {
				log.Infof("Marked self as available")
			}
		}
	case *events.PairSuccess:
		websocket.Broadcast <- websocket.BroadcastMessage{
			Code:    "LOGIN_SUCCESS",
			Message: fmt.Sprintf("Successfully pair with %s", evt.ID.String()),
		}
	case *events.LoggedOut:
		websocket.Broadcast <- websocket.BroadcastMessage{
			Code:   "LIST_DEVICES",
			Result: nil,
		}
	case *events.Connected, *events.PushNameSetting:
		if len(cli.Store.PushName) == 0 {
			return
		}

		// Send presence available when connecting and when the pushname is changed.
		// This makes sure that outgoing messages always have the right pushname.
		err := cli.SendPresence(types.PresenceAvailable)
		if err != nil {
			log.Warnf("Failed to send available presence: %v", err)
		} else {
			log.Infof("Marked self as available")
		}
	case *events.StreamReplaced:
		os.Exit(0)
	case *events.Message:

		metaParts := []string{fmt.Sprintf("pushname: %s", evt.Info.PushName), fmt.Sprintf("timestamp: %s", evt.Info.Timestamp)}
		if evt.Info.Type != "" {
			metaParts = append(metaParts, fmt.Sprintf("type: %s", evt.Info.Type))
		}
		if evt.Info.Category != "" {
			metaParts = append(metaParts, fmt.Sprintf("category: %s", evt.Info.Category))
		}
		if evt.IsViewOnce {
			metaParts = append(metaParts, "view once")
		}
		if evt.IsViewOnce {
			metaParts = append(metaParts, "ephemeral")
		}
		command := utils.ProcessCommand(evt.Message.ExtendedTextMessage.GetText())

		fmt.Println("Received message ", string(evt.Info.ID), evt.Info.SourceString(), strings.Join(metaParts, ", "), evt.Message.ExtendedTextMessage.GetText(), command)

		if command == "!audio" {
			argument := utils.GetArgument(evt.Message.ExtendedTextMessage.GetText())

			fmt.Println("argument", argument)
			formattedJid, _ := utils.GetLastJid(evt.Info.SourceString())
			dataWaRecipient, _ := ValidateJidWithLogin(cli, formattedJid)
			fmt.Println("mandando")
			audioDownloaded, errAudioDownloaded := services.SearchAudioFunnyReturnFile(argument)

			if errAudioDownloaded != nil {
				msg := &waProto.Message{
					Conversation: proto.String("error ao baixar o audio"),
				}
				s, err2 := cli.SendMessage(context.Background(), dataWaRecipient, msg)
				fmt.Println("send message error download audio in errAudioDownloaded", s, err2)
			}

			audioMimeType := http.DetectContentType(audioDownloaded)

			audioUploaded, err := cli.Upload(context.Background(), audioDownloaded, whatsmeow.MediaAudio)
			if err != nil {
				fmt.Sprintf("Failed to upload audio: %v", err)
			}

			stanzaID := string(evt.Info.ID)
			duration := 120 // 2 minutos

			// Converta a duração para uint32 e pegue o endereço
			seconds := uint32(duration)
			secondsPtr := &seconds
			fmt.Println("object audio uploaded", audioUploaded)

			msg := &waProto.Message{
				AudioMessage: &waProto.AudioMessage{
					Url:           proto.String(audioUploaded.URL),
					DirectPath:    proto.String(audioUploaded.DirectPath),
					Mimetype:      proto.String(audioMimeType),
					FileLength:    proto.Uint64(audioUploaded.FileLength),
					FileSha256:    audioUploaded.FileSHA256,
					FileEncSha256: audioUploaded.FileEncSHA256,
					Seconds:       secondsPtr,
					MediaKey:      audioUploaded.MediaKey,
					ContextInfo: &waProto.ContextInfo{
						QuotedMessage: &waProto.Message{
							Conversation: proto.String(evt.Message.ExtendedTextMessage.GetText()),
						},
						StanzaId:    &stanzaID,
						Participant: proto.String(dataWaRecipient.String()), // O participante é quem enviou a mensagem original
					},
				},
			}

			// replyMessage := &waProto.Message{
			// 	ExtendedTextMessage: &waProto.ExtendedTextMessage{
			// 		Text: proto.String("kkkkkkkk"),
			// 		ContextInfo: &waProto.ContextInfo{
			// 			QuotedMessage: &waProto.Message{
			// 				Conversation: proto.String("replyText"),
			// 			},
			// 			StanzaId:    &stanzaID,
			// 			Participant: proto.String(dataWaRecipient.String()), // O participante é quem enviou a mensagem original
			// 		},
			// 	},
			// }
			s, err2 := cli.SendMessage(context.Background(), dataWaRecipient, msg)
			fmt.Println("mandado", s, err2)
		}

		img := evt.Message.GetImageMessage()
		if img != nil {
			path, err := ExtractMedia(config.PathStorages, img)
			if err != nil {
				log.Errorf("Failed to download image: %v", err)
			} else {
				log.Infof("Image downloaded to %s", path)
			}
		}

		if config.WhatsappAutoReplyMessage != "" &&
			!isGroupJid(evt.Info.Chat.String()) &&
			!strings.Contains(evt.Info.SourceString(), "broadcast") {
			_, _ = cli.SendMessage(context.Background(), evt.Info.Sender, &waProto.Message{Conversation: proto.String(config.WhatsappAutoReplyMessage)})
		}

		if config.WhatsappWebhook != "" &&
			!strings.Contains(evt.Info.SourceString(), "broadcast") &&
			!isFromMySelf(evt.Info.SourceString()) {
			if err := forwardToWebhook(evt); err != nil {
				logrus.Error("Failed forward to webhook", err)
			}
		}

	case *events.Receipt:
		if evt.Type == types.ReceiptTypeRead || evt.Type == types.ReceiptTypeReadSelf {
			log.Infof("%v was read by %s at %s", evt.MessageIDs, evt.SourceString(), evt.Timestamp)
		} else if evt.Type == types.ReceiptTypeDelivered {
			log.Infof("%s was delivered to %s at %s", evt.MessageIDs[0], evt.SourceString(), evt.Timestamp)
		}
	case *events.Presence:
		if evt.Unavailable {
			if evt.LastSeen.IsZero() {
				log.Infof("%s is now offline", evt.From)
			} else {
				log.Infof("%s is now offline (last seen: %s)", evt.From, evt.LastSeen)
			}
		} else {
			log.Infof("%s is now online", evt.From)
		}
	case *events.HistorySync:
		id := atomic.AddInt32(&historySyncID, 1)
		fileName := fmt.Sprintf("%s/history-%d-%s-%d-%s.json", config.PathStorages, startupTime, cli.Store.ID.String(), id, evt.Data.SyncType.String())
		file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			log.Errorf("Failed to open file to write history sync: %v", err)
			return
		}
		enc := json.NewEncoder(file)
		enc.SetIndent("", "  ")
		err = enc.Encode(evt.Data)
		if err != nil {
			log.Errorf("Failed to write history sync: %v", err)
			return
		}
		log.Infof("Wrote history sync to %s", fileName)
		_ = file.Close()
	case *events.AppState:
		log.Debugf("App state event: %+v / %+v", evt.Index, evt.SyncActionValue)
	}
}

// forwardToWebhook is a helper function to forward event to webhook url
func forwardToWebhook(evt *events.Message) error {
	logrus.Info("Forwarding event to webhook:", config.WhatsappWebhook)
	client := &http.Client{Timeout: 10 * time.Second}
	imageMedia := evt.Message.GetImageMessage()
	stickerMedia := evt.Message.GetStickerMessage()
	videoMedia := evt.Message.GetVideoMessage()
	audioMedia := evt.Message.GetAudioMessage()
	documentMedia := evt.Message.GetDocumentMessage()

	var message evtMessage
	message.Text = evt.Message.GetConversation()
	message.ID = evt.Info.ID
	if extendedMessage := evt.Message.ExtendedTextMessage.GetText(); extendedMessage != "" {
		message.Text = extendedMessage
		message.RepliedId = evt.Message.ExtendedTextMessage.ContextInfo.GetStanzaId()
	}

	var quotedmessage any
	if evt.Message.ExtendedTextMessage != nil && evt.Message.ExtendedTextMessage.ContextInfo != nil {
		if conversation := evt.Message.ExtendedTextMessage.ContextInfo.QuotedMessage.GetConversation(); conversation != "" {
			quotedmessage = conversation
		}
	}

	var forwarded bool
	if evt.Message.ExtendedTextMessage != nil && evt.Message.ExtendedTextMessage.ContextInfo != nil {
		forwarded = evt.Message.ExtendedTextMessage.ContextInfo.GetIsForwarded()
	}

	var waReaction evtReaction
	if reactionMessage := evt.Message.ReactionMessage; reactionMessage != nil {
		waReaction.Message = reactionMessage.GetText()
		waReaction.ID = reactionMessage.GetKey().GetId()
	}

	body := map[string]interface{}{
		"audio":          audioMedia,
		"contact":        evt.Message.GetContactMessage(),
		"document":       documentMedia,
		"forwarded":      forwarded,
		"from":           evt.Info.SourceString(),
		"image":          imageMedia,
		"list":           evt.Message.GetListMessage(),
		"live_location":  evt.Message.GetLiveLocationMessage(),
		"location":       evt.Message.GetLocationMessage(),
		"message":        message,
		"order":          evt.Message.GetOrderMessage(),
		"pushname":       evt.Info.PushName,
		"quoted_message": quotedmessage,
		"reaction":       waReaction,
		"sticker":        stickerMedia,
		"video":          videoMedia,
		"view_once":      evt.Message.GetViewOnceMessage(),
	}

	if imageMedia != nil {
		path, err := ExtractMedia(config.PathMedia, imageMedia)
		if err != nil {
			return pkgError.WebhookError(fmt.Sprintf("Failed to download image: %v", err))
		}
		body["image"] = path
	}
	if stickerMedia != nil {
		path, err := ExtractMedia(config.PathMedia, stickerMedia)
		if err != nil {
			return pkgError.WebhookError(fmt.Sprintf("Failed to download sticker: %v", err))
		}
		body["sticker"] = path
	}
	if videoMedia != nil {
		path, err := ExtractMedia(config.PathMedia, videoMedia)
		if err != nil {
			return pkgError.WebhookError(fmt.Sprintf("Failed to download video: %v", err))
		}
		body["video"] = path
	}
	if audioMedia != nil {
		path, err := ExtractMedia(config.PathMedia, audioMedia)
		if err != nil {
			return pkgError.WebhookError(fmt.Sprintf("Failed to download audio: %v", err))
		}
		body["audio"] = path
	}
	if documentMedia != nil {
		path, err := ExtractMedia(config.PathMedia, documentMedia)
		if err != nil {
			return pkgError.WebhookError(fmt.Sprintf("Failed to download document: %v", err))
		}
		body["document"] = path
	}

	postBody, err := json.Marshal(body)
	if err != nil {
		return pkgError.WebhookError(fmt.Sprintf("Failed to marshal body: %v", err))
	}

	req, err := http.NewRequest(http.MethodPost, config.WhatsappWebhook, bytes.NewBuffer(postBody))
	if err != nil {
		return pkgError.WebhookError(fmt.Sprintf("error when create http object %v", err))
	}
	req.Header.Set("Content-Type", "application/json")
	if _, err = client.Do(req); err != nil {
		return pkgError.WebhookError(fmt.Sprintf("error when submit webhook %v", err))
	}
	return nil
}

// isGroupJid is a helper function to check if the message is from group
func isGroupJid(jid string) bool {
	return strings.Contains(jid, "@g.us")
}

// isFromMySelf is a helper function to check if the message is from my self (logged in account)
func isFromMySelf(jid string) bool {
	return extractPhoneNumber(jid) == extractPhoneNumber(cli.Store.ID.String())
}

// extractPhoneNumber is a helper function to extract the phone number from a JID
func extractPhoneNumber(jid string) string {
	regex := regexp.MustCompile(`\d+`)
	// Find all matches of the pattern in the JID
	matches := regex.FindAllString(jid, -1)
	// The first match should be the phone number
	if len(matches) > 0 {
		return matches[0]
	}
	// If no matches are found, return an empty string
	return ""
}

// ExtractMedia is a helper function to extract media from whatsapp
func ExtractMedia(storageLocation string, mediaFile whatsmeow.DownloadableMessage) (extractedMedia ExtractedMedia, err error) {
	if mediaFile == nil {
		logrus.Info("Skip download because data is nil")
		return extractedMedia, nil
	}

	data, err := cli.Download(mediaFile)
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


