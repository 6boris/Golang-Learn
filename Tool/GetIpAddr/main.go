package main

import (
	"fmt"
	"net"
)

func main() {
	localIP()
}

//	获取本机IP
func localIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	localIP := localAddr.IP.String()
	fmt.Println(localIP)
	//log.Info("Local IP : ", localIP)
	return localIP, nil
}
