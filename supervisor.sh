[program:bot-zap-golang]
directory=/home/ubuntu/bot-zap-golang/src
command=/home/ubuntu/bot-zap-golang/src/main
autostart=true
autorestart=true
stderr_logfile=/var/log/bot-zap-golang.err.log
stdout_logfile=/var/log/bot-zap-golang.out.log
