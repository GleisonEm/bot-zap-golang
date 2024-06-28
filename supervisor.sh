[program:bot-zap-golang]
directory=/home/bot-zap-golang
command=/home/bot-zap-golang/main
autostart=true
autorestart=true
stderr_logfile=/var/log/bot-zap-golang.err.log
stdout_logfile=/var/log/bot-zap-golang.out.log
