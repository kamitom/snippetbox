# sample webapp (golang)

##### service content
> sudo cat /lib/systemd/system/helenfitweb2.service
> 

```
[Unit]
Description=helenfitweb2
ConditionPathExists=/home/ubuntu/go/src/snippetbox
After=network.target

[Service]
Type=simple
User=ubuntu
Group=ubuntu
Restart=always
RestartSec=5s
WorkingDirectory=/home/ubuntu/go/src/snippetbox
ExecStart=/home/ubuntu/go/src/snippetbox/web

[Install]
WantedBy=multi-user.target
```