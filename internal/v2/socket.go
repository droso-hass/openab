package v2

import (
	"fmt"
	"log/slog"
	"net"
	"strconv"
	"time"
)

var conns = map[string]nab{}

func connect(req string) error {
	// if a connection is already open, close it
	c, ok := conns[req]
	if ok {
		c.conn.Close()
	}

	slog.Info("connecting to " + req)
	conn, err := net.Dial("tcp", req)
	if err != nil {
		return err
	}

	// reading loop
	go func() {
		for {
			fmt.Println("reading ...")
			buf := make([]byte, 1024)
			conn.Read(buf)
			fmt.Println(string(buf))

		}
	}()

	// test writing loop
	go func() {
		i := 0
		for {
			time.Sleep(2 * time.Second)
			_, err = conn.Write([]byte("ping " + strconv.Itoa(i)))
			if err != nil {
				fmt.Println(err)
				return
			}
			i++
		}
	}()

	return nil
}
