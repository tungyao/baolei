package main

import (
	"log"
	"os"
)

const service = `
[Unit]
Description = ping server
After = network.target syslog.target
Wants = network.target

[Service]
Type = simple
ExecStart = /root/app/ping/ping -mode -path=./
WorkingDirectory= /root/app/ping
[Install]
WantedBy = multi-user.target
`

func install() {
	fs, err := os.OpenFile("/etc/systemd/system/baolei.service", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Panic(err)
	}
	fs.WriteString(service)
	fs.Close()
}
