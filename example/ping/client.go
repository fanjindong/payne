package main

import (
	"bufio"
	"fmt"
	"github.com/fanjindong/payne/codec"
	"github.com/fanjindong/payne/msg"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	codec := codec.LvCodec{}
	go func() {
		for {
			m, err := codec.Decode(conn)
			if err != nil {
				fmt.Sprintln("codec.Decode err:", err)
				break
			}
			fmt.Println("reply:", string(m.GetData()))
		}
	}()
	for {
		text := readByStd()
		data, err := codec.Encode(msg.NewMsg([]byte(text)))
		if err != nil {
			panic(err)
		}
		if _, err = conn.Write(data); err != nil {
			panic(err)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func readByStd() string {
	fmt.Print("Send text: ")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")
	return input
}
