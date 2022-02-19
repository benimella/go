package tcp

import (
	"fmt"
	"log"
	"net"
	"time"
)

func Server()  {
	l, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Received message %s => %s", conn.RemoteAddr(), conn.LocalAddr())

		go func(connection net.Conn) {
			for {
				buf := make([]byte, 1024)
				_, err := conn.Read(buf)
				if err != nil {
					log.Println("conn read err", err)
					break
				}
				_, err = conn.Write([]byte("ok"))
				if err != nil {
					log.Println("conn write err", err)
					break
				}
			}
		}(conn)
	}
}

func Client()  {
	conn, err := net.Dial("tcp", "192.168.0.33:8080")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		_, err = conn.Write([]byte("hello"))
		if err != nil {
			log.Fatalln(err)
		}
		time.Sleep(time.Second * 2)
	}
}
