package v2

import (
	"fmt"
)

var RecDataSize = 4096

func (n *NabConn) processNabMessage(data []byte) {
	sdata := string(data)
	fmt.Println(sdata)
	/*if sdata[0:2] == "07" {

	}*/
}
