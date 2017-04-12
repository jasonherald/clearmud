package handlers

import (
	"bufio"
	"bytes"
	"io"
	"net"
	"reflect"
	"strings"
	"unsafe"

	objects "../objects"
)

func RequestHandler(conn net.Conn, out chan string, user objects.User) {
	defer close(out)
	for {
		line, err := bufio.NewReader(conn).ReadBytes('\n')
		if err != nil {
			return
		}

		// look up the command which for now is just a movement
		str := BytesToString(line)
		str = strings.Trim(str, "\n")
		str = strings.Trim(str, "\r")
		switch str {
		case "n":
			user.Move(0)
		case "e":
			user.Move(1)
		case "w":
			user.Move(2)
		case "s":
			user.Move(3)
		}
		out <- "Current Position: " + user.ToString() + "\n"
	}
}

func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

func SendData(conn net.Conn, in <-chan string) {
	defer conn.Close()
	for {
		message := <-in
		if message != "" {
			io.Copy(conn, bytes.NewBufferString(message))
		}
	}
}
