package main

import (
	"flag"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

var (
	clientIp     = ""
	writePath    = ""
	isServerMode = false
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "install" {
		install()
		log.Println("写入完成")
		os.Exit(0)
	}

	flag.StringVar(&writePath, "path", "", "ip_path")
	flag.BoolVar(&isServerMode, "mode", false, "is_server_mode")
	flag.StringVar(&clientIp, "ip", "", "client ip")
	flag.Parse()
	if isServerMode {
		l, err := net.Listen("tcp", ":6868")
		if err != nil {
			log.Panic(err)
		}

		for {
			conn, err := l.Accept()
			if err != nil {
				log.Println(err)
				continue
			}
			conn.Close()
			log.Println("get ip", conn.RemoteAddr().String())
			// 写到某一个位置
			fs, err := os.OpenFile(strings.TrimSuffix(writePath, "/")+"/ip", os.O_CREATE|os.O_WRONLY, 755)
			if err != nil {
				log.Println(err)
			}
			fs.Write([]byte(conn.RemoteAddr().String()))
			fs.Close()
		}
	} else {
		for {
			conn, _ := net.Dial("tcp", clientIp)
			conn.Close()
			time.Sleep(time.Minute)
		}
	}

}
