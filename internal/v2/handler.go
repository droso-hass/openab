package v2

import "fmt"

func (n *NabConn) processNabMessage(data []byte) {
	fmt.Println(string(data))
}
