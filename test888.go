package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func conn() {
	con, err := net.Dial("tcp", "192.168.0.33:9090")
	if err != nil {
		log.Println(err)
		return
	}
	defer con.Close()
	fmt.Println("客户端：", con.LocalAddr())
	time.Sleep(time.Second * 4)
}

func main() {
	for {
		go conn()
		time.Sleep(time.Second * 1)
	}
}
