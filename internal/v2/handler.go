package v2

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func (n *NabConn) processNabMessage(data []byte) {
	sdata := string(data)
	if sdata[0:2] == "008" {
		if n.recFile == nil {
			f, err := os.OpenFile("./rec.wav", os.O_WRONLY|os.O_CREATE, 0777)
			if err != nil {
				log.Fatal(err)
			}
			n.recFile = f
		}
		h, err := hex.DecodeString(sdata[3:])
		if err != nil {
			log.Fatal(err)
		}
		_, err = n.recFile.Write(h)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(sdata)
}
