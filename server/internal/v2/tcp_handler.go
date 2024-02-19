package v2

import (
	"fmt"
)

var RecDataSize = 4096

func (n *NabConn) processNabMessage(data []byte) {
	sdata := string(data)
	/*if sdata[0:2] == "08" {
	}*/
	fmt.Println(sdata)
}
