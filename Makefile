run:
	cd src && go run main.go

build:
	cd src && go build main.go

restart:
	sudo supervisorctl restart bot-zap-golang

status:
	sudo supervisorctl status bot-zap-golang

stop:
	sudo supervisorctl stop bot-zap-golang

start:
	sudo supervisorctl start bot-zap-golang
