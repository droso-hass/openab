package v2

import (
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

var RecDataSize = 4096

func (n *NabConn) processNabMessage(data []byte) {
	sdata := string(data)
	if sdata[0:2] == "08" {
		if len(n.recData) >= RecDataSize {
			filename := fmt.Sprintf("./rec/%s_%d.wav", strings.Replace(n.mac, ":", "", -1), time.Now().Unix())
			err := makeWav(n.recData, filename)
			if err != nil {
				fmt.Println(err)
			}
			n.recData = []byte{}
		}
		h, err := hex.DecodeString(sdata[3:])
		if err == nil {
			n.recData = append(n.recData, h...)
		} else {
			fmt.Println(err)
		}
	}
	fmt.Println(sdata)
}
