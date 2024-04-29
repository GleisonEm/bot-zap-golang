package contexts

import (
	domainApp "github.com/GleisonEm/bot-claudio-zap-golang/domains/app"
	domainGroup "github.com/GleisonEm/bot-claudio-zap-golang/domains/group"
	domainMessage "github.com/GleisonEm/bot-claudio-zap-golang/domains/message"
	domainSend "github.com/GleisonEm/bot-claudio-zap-golang/domains/send"
	domainUser "github.com/GleisonEm/bot-claudio-zap-golang/domains/user"
)

type AppContext struct {
	AppService     domainApp.IAppService
	SendService    domainSend.ISendService
	UserService    domainUser.IUserService
	MessageService domainMessage.IMessageService
	GroupService   domainGroup.IGroupService
}

var Context *AppContext

func InitServiceAppContext(
	appService domainApp.IAppService,
	sendService domainSend.ISendService,
	userService domainUser.IUserService,
	messageService domainMessage.IMessageService,
	groupService domainGroup.IGroupService,
) {
	Context = &AppContext{
		AppService:     appService,
		SendService:    sendService,
		UserService:    userService,
		MessageService: messageService,
		GroupService:   groupService,
	}
}
