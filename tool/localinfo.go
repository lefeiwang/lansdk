package tool

import (
	"log"
	"net"
	"strings"
)

func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		log.Println(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}

func GetHostName() string {
	cBytes, err := ReadFile("/etc/hostname")
	if err != nil {
		log.Println(err)
		return ""
	}
	return strings.Replace(string(cBytes), "\n", "", -1)
}
